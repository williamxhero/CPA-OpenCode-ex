package errclass

import (
	"regexp"
	"strconv"
	"strings"
)

// Only OpenCode's explicit subscription rejection is credential-wide. Generic
// 429s and messages merely mentioning quota must retain normal retry semantics.
var goUsageLimit = regexp.MustCompile(`(?i)^OpenCode Go usage limit reached \((rolling|weekly|monthly)(?:, resets in ([^)]+))?\)\.`)
var goResetPart = regexp.MustCompile(`^(\d+)([dhms])$`)

func openCodeQuotaLimit(message string) (bool, int) {
	match := goUsageLimit.FindStringSubmatch(strings.TrimSpace(message))
	if match == nil {
		return false, 0
	}
	var seconds int64
	parts := strings.Fields(strings.ToLower(match[2]))
	if len(parts) == 0 || len(parts) > 4 {
		return true, 0
	}
	for _, part := range parts {
		m := goResetPart.FindStringSubmatch(part)
		if m == nil {
			return true, 0
		}
		n, err := strconv.ParseInt(m[1], 10, 64)
		unit := map[string]int64{"d": 86400, "h": 3600, "m": 60, "s": 1}[m[2]]
		// One year is a safety bound, not an invented fallback reset time.
		if err != nil || n > 366*86400/unit || seconds > 366*86400-n*unit {
			return true, 0
		}
		seconds += n * unit
	}
	return true, int(seconds)
}
