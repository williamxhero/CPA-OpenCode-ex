package plugin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginabi"
	"github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginapi"
)

const goWeeklyLimitMessage = "OpenCode Go usage limit reached (weekly, resets in 4d 16h). To keep going now, turn on Extra Usage or upgrade to Go Plus at https://opencode.ai/console/org_synthetic/go."

func TestExecuteGoUsageLimitIsCredentialScoped(t *testing.T) {
	for _, stream := range []bool{false, true} {
		for _, model := range []string{"opencode-go/glm-5.3", "opencode-go/minimax-m3", "opencode-go/gpt-5.6-luna"} {
			t.Run(fmt.Sprintf("stream=%t/model=%s", stream, model), func(t *testing.T) {
				m, f := newExecManager(t)
				f.responder = wrapWithCatalog(multiRouteCatalog, func(method string, payload []byte) ([]byte, error) {
					switch method {
					case pluginabi.MethodHostHTTPDo:
						return hostOK(pluginapi.HTTPResponse{StatusCode: http.StatusTooManyRequests, Body: []byte(goWeeklyLimitMessage)}), nil
					case pluginabi.MethodHostHTTPDoStream:
						return hostOK(hostStreamStartResp{StatusCode: http.StatusTooManyRequests, StreamID: "quota-upstream"}), nil
					case pluginabi.MethodHostHTTPStreamRead:
						return hostOK(hostStreamReadResp{Payload: []byte(goWeeklyLimitMessage), Done: true}), nil
					default:
						return hostOK(map[string]any{}), nil
					}
				})
				method := "executor.execute"
				request := execReqBody(model, "openai", []byte(ccRequestBody), false)
				if stream {
					method = "executor.execute_stream"
					request = execStreamReqBody(model, "openai", []byte(ccRequestBody), "quota-downstream")
				}
				raw, err := m.HandleCall(method, request)
				if err != nil {
					t.Fatal(err)
				}
				env := decodeEnv(t, raw)
				if env.OK || env.Error == nil || env.Error.HTTPStatus != http.StatusTooManyRequests {
					t.Fatalf("expected upstream quota error: %+v", env)
				}
				var detail struct {
					Error struct {
						CredentialScoped  bool `json:"credential_scoped"`
						RetryAfterSeconds int  `json:"retry_after_seconds"`
					} `json:"error"`
				}
				if err := json.Unmarshal(raw, &detail); err != nil {
					t.Fatal(err)
				}
				if !detail.Error.CredentialScoped || env.Error.Code != "credential_quota" {
					t.Fatalf("first Go usage limit must exhaust the credential, got code=%q credential_scoped=%t", env.Error.Code, detail.Error.CredentialScoped)
				}
				if detail.Error.RetryAfterSeconds != 403200 {
					t.Fatalf("reset hint = %d, want 403200", detail.Error.RetryAfterSeconds)
				}
			})
		}
	}
}
