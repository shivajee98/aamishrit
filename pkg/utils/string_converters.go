package utils

import "strconv"

func ParseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
