package utils

import "strings"

func SanitizeArgs(args []string) []string {
	for i, arg := range args {
		trimmed := strings.TrimSpace(arg)
		args[i] = trimmed
	}
	return args
}

func TruncateString(s string, max int) string {
	var out string
	if len(s) > max {
		out = s[:max]
	} else {
		out = s
	}
	return out
}
