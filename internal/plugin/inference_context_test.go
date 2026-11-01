package plugin

import (
	"encoding/json"
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginabi"
	"github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginapi"
)

func TestExecutorInvocationForwardsScopedHTTPContext(t *testing.T) {
	for _, stream := range []bool{false, true} {
		t.Run(map[bool]string{false: "http", true: "sse"}[stream], func(t *testing.T) {
			m, f, _, _ := newIntegrationManager(t)
			original, _ := json.Marshal(map[string]any{"model": "opencode-go/glm-5.2", "messages": []any{map[string]any{"role": "user", "content": "synthetic"}}, "stream": stream})
			wire, _ := json.Marshal(struct {
				pluginapi.ExecutorRequest
				HostCallbackID string `json:"host_callback_id"`
				StreamID       string `json:"stream_id"`
			}{pluginapi.ExecutorRequest{Model: "opencode-go/glm-5.2", AuthID: "synthetic", AuthProvider: ProviderID, AuthAttributes: map[string]string{"api_key": "sk-test-1"}, SourceFormat: "openai", OriginalRequest: original, Payload: original, Stream: stream}, "issued-inference-context", "downstream-test"})
			method := pluginabi.MethodExecutorExecute
			if stream {
				method = pluginabi.MethodExecutorExecuteStream
			}
			mustHandle(t, m, method, wire)
			f.mu.Lock()
			defer f.mu.Unlock()
			seen := false
			for _, call := range f.calls {
				if call.method != pluginabi.MethodHostHTTPDo && call.method != pluginabi.MethodHostHTTPDoStream {
					continue
				}
				var req hostHTTPReq
				json.Unmarshal(call.payload, &req)
				if req.Method != "POST" {
					continue
				}
				seen = true
				if req.HostCallbackID != "issued-inference-context" {
					t.Errorf("actual inference callback lost host scope stream=%v", stream)
				}
			}
			if !seen {
				t.Fatal("no actual inference callback")
			}
		})
	}
}
