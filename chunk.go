package caseconv

import (
	"strings"
)

func chunk(str string) []string {
	var (
		lastRune rune
		parts    []string
	)

	parts = chunkBy(str, func(r, next rune) (split, drop bool) {
		defer func() { lastRune = r }()

		if r == '_' || r == ' ' || r == '.' || r == '/' || r == '"' {
			return true, true
		}

		if lastRune == 0 {
			return false, false
		}

		if next != 0 {
			if 'A' <= r && r <= 'Z' &&
				'A' <= lastRune && lastRune <= 'Z' &&
				'a' <= next && next <= 'z' {
				return true, false
			}
		}

		if (('a' <= lastRune && lastRune <= 'z') || ('0' <= lastRune && lastRune <= '9')) &&
			'A' <= r && r <= 'Z' {
			return true, false
		}

		return false, false
	})

	buffer := parts
	parts = make([]string, len(buffer))
	for i, part := range buffer {

		parts[i] = strings.ToLower(part)
	}

	return parts
}

func chunkBy(str string, chunkFn func(r, next rune) (split, drop bool)) (result []string) {
	var (
		lastChunk []rune
	)

	for i, r := range str {
		var next rune
		if i < len(str)-1 {
			next = []rune(str)[i+1]
		}
		split, drop := chunkFn(r, next)
		if !split && !drop {
			lastChunk = append(lastChunk, r)
			continue
		}
		if split {
			if len(lastChunk) > 0 {
				result = append(result, string(lastChunk))
			}
			lastChunk = nil
			if !drop {
				lastChunk = append(lastChunk, r)
			}
			continue
		}
	}

	if len(lastChunk) > 0 {
		result = append(result, string(lastChunk))
	}

	return
}
