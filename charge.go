package main

import (
	"strings"
	"unicode"
)

func ChargeParse(s string) string {
	tokens := strings.Split(s, " ")
	buffer := []string{}
	for _, t := range tokens {
		if len(strings.TrimSpace(t)) == 0 {
			continue
		}
		if !unicode.IsLetter(rune(t[0])) {
			break
		}
		if strings.HasPrefix(t, "HTTPS") {
			break
		}
		buffer = append(buffer, t)
	}

	return strings.Join(buffer, " ")
}
