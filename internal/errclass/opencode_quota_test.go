package errclass

import "testing"

func TestOpenCodeQuotaLimitBoundaries(t *testing.T) {
	for _, tc := range []struct {
		message string
		status  int
		limited bool
		reset   int
	}{
		{"OpenCode Go usage limit reached (weekly, resets in 4d 16h).", 429, true, 403200},
		{"OpenCode Go usage limit reached (rolling, resets in 2h 3m 4s).", 429, true, 7384},
		{"OpenCode Go usage limit reached (monthly).", 429, true, 0},
		{"OpenCode Go usage limit reached (weekly, resets in 999999999999999999999d).", 429, true, 0},
		{"OpenCode Go usage limit reached (weekly, resets in tomorrow).", 429, true, 0},
		{"OpenCode Go usage limit reached (weekly, resets in 500d).", 429, true, 0},
		{"GoUsageLimitError", 429, false, 0},
		{"rate limit exceeded", 429, false, 0},
		{"quota exhausted", 429, false, 0},
		{"OpenCode Go usage limit reached (weekly).", 503, false, 0},
	} {
		got := FromStatus(tc.status, tc.message)
		if got.CredentialScoped != tc.limited || got.RetryAfterSeconds != tc.reset {
			t.Errorf("%q: scoped=%t reset=%d", tc.message, got.CredentialScoped, got.RetryAfterSeconds)
		}
	}
}
