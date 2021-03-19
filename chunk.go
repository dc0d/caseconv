package caseconv

import (
	"strings"
)

func chunk(str string) []string {
	var (
		cp    chunkerPredicate
		parts []string
	)

	parts = chunkBy(str, cp.run)

	buffer := parts
	parts = make([]string, len(buffer))
	for i, part := range buffer {
		parts[i] = strings.ToLower(part)
	}

	return parts
}

type chunkerPredicate struct{ lastRune rune }

func (cp *chunkerPredicate) run(r, next rune) (split, drop bool) {
	defer func() { cp.lastRune = r }()

	var (
		br    = breakRune(r)
		bnext = breakRune(next)
		blast = breakRune(cp.lastRune)
	)

	switch {
	case br.isSplitter():
		split, drop = true, true
	case cp.lastRune == 0:
		return
	case next != 0 && br.isUppercase() &&
		blast.isUppercase() &&
		bnext.isLowercase():
		fallthrough
	case (blast.isLowercase() || blast.isNumber()) &&
		br.isUppercase():
		split = true
	}

	return
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

	runes := []rune(str)

	for i, r := range runes {
		var next rune
		if i < len(runes)-1 {
			next = []rune(runes)[i+1]
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
