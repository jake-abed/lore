package utils

import "strings"

func SanitizeArgs(args []string) []string {
	for i, arg := range args {
		lower := strings.ToLower(arg)
		trimmed := strings.TrimSpace(lower)
		args[i] = trimmed
	}
	return args
}
