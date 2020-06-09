package caseconv

import "strings"

func ToKebab(str string) string {
	chunks := chunk(str)
	return strings.Join(chunks, "-")
}
