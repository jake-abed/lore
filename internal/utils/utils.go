package utils

import "strings"

func SanitizeArgs(args []string) []string {
	for i, arg := range args {
		trimmed := strings.TrimSpace(arg)
		args[i] = trimmed
	}
	return args
}
