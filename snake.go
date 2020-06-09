package caseconv

import "strings"

func ToSnake(str string) string {
	chunks := chunk(str)
	return strings.Join(chunks, "_")
}
