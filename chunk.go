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
		var (
			br    = breakRune(r)
			bnext = breakRune(next)
			blast = breakRune(lastRune)
		)

		if br.isSplitter() {
			return true, true
		}

		if lastRune == 0 {
			return false, false
		}

		if next != 0 {
			if br.isUppercase() &&
				blast.isUppercase() &&
				bnext.isLowercase() {
				return true, false
			}
		}

		if (blast.isLowercase() || blast.isNumber()) &&
			br.isUppercase() {
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

type breakRune rune

func (br breakRune) isLowercase() bool { return 'a' <= br && br <= 'z' }
func (br breakRune) isNumber() bool    { return '0' <= br && br <= '9' }
func (br breakRune) isUppercase() bool { return 'A' <= br && br <= 'Z' }
func (br breakRune) isSplitter() bool {
	return br == '_' || br == ' ' || br == '.' || br == '/' || br == '"'
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
