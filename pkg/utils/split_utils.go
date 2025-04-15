package utils

import "strings"

func SplitAndTrim(input string) []string {
	raw := strings.Split(input, ",")
	var out []string
	for _, item := range raw {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}
