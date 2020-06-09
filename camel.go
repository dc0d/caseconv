package caseconv

import "strings"

func ToCamel(str string) string {
	chunks := chunk(str)
	for i, c := range chunks {
		if i == 0 {
			chunks[i] = strings.ToLower(c)
			continue
		}
		chunks[i] = strings.Title(c)
	}
	return strings.Join(chunks, "")
}
