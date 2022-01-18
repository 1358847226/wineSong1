package util

import "strings"

func S2a(s string) []string {
	var a []string
	if s == "" {
		return a
	} else {
		a = strings.Split(s, ",")
		return a
	}
}

func A2S(a []string) string {
	s := strings.Join(a, ",")
	return s
}
