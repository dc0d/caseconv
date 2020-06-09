package caseconv

import "strings"

func ToPascal(str string) string {
	chunks := chunk(str)
	for i, c := range chunks {
		chunks[i] = strings.Title(c)
	}
	return strings.Join(chunks, "")
}
