package caseconv

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

//-----------------------------------------------------------------------------

// Convert converts names between cases, it does not change the upper/lower cases
// except for the first rune in camel and pascal cases.
func Convert(name string, from, to Case, check ...bool) (string, error) {
	if !from.isValid() || !to.isValid() {
		return "", errors.Cause(errors.WithMessage(ErrInvalidCase, "from or to is invalid"))
	}
	if len(check) > 0 && check[0] {
		if c := guess([]rune(name)); c != from || c == None {
			return "", errors.Cause(errors.WithMessage(ErrCheckCase, "check failed"))
		}
	}
	scn := bufio.NewScanner(strings.NewReader(name))
	scn.Split(newSplitter(from).split())
	var parts []string
	for scn.Scan() {
		parts = append(parts, scn.Text())
	}
	if len(parts) > 0 {
		switch from {
		case Camel:
			parts[0] = changeHead(parts[0], true)
		case Pascal:
			parts[0] = changeHead(parts[0])
		}
	}

	switch to {
	case Snake:
		return strings.Join(parts, _underscore), nil
	case Kebab:
		return strings.Join(parts, _dash), nil
	case Camel:
		for k := range parts {
			if k == 0 {
				parts[k] = changeHead(parts[k], true)
				continue
			}
			parts[k] = changeHead(parts[k])
		}
		return strings.Join(parts, ""), nil
	case Pascal:
		for k := range parts {
			parts[k] = changeHead(parts[k])
		}
		return strings.Join(parts, ""), nil
	}
	return "", errors.Cause(ErrUnknown)
}

//-----------------------------------------------------------------------------

type splitter struct {
	_case Case
}

func newSplitter(_case Case) *splitter { return &splitter{_case: _case} }

func (s *splitter) split() bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF {
			if len(data) > 0 {
				return 0, data[:], nil
			}
			return 0, nil, nil
		}
		switch s._case {
		case Snake:
			found := bytes.IndexRune(data, _runeUnderscore)
			if found >= 0 {
				return found + 1, data[:found], nil
			}
		case Kebab:
			found := bytes.IndexRune(data, _runeDash)
			if found >= 0 {
				return found + 1, data[:found], nil
			}
		case Camel:
			runes := bytes.Runes(data)
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
				data = []byte(string(runes))
			}
			s._case = Pascal
			fallthrough
		case Pascal:
			var (
				lastcase bool
				index    int
				changes  int
			)
			seg := bytes.IndexFunc(data, func(r rune) bool {
				defer func() {
					lastcase = unicode.IsUpper(r)
					index++
				}()
				if index == 0 {
					return false
				}
				if unicode.IsUpper(r) != lastcase {
					changes++
				}
				return changes > 1
			})
			if seg > -1 {
				return seg, data[:seg], nil
			}

		}
		return len(data), data[:], nil
	}
}

func guess(name []rune) Case {
	if len(name) == 0 {
		return None
	}
	var (
		isSnake  = false
		isKebab  = false
		isCamel  = false
		isPascal = false
	)
	if unicode.IsLower(name[0]) {
		isCamel = true
	}
	var (
		prev     rune
		anyLower bool
	)
	for _, vr := range name {
		if vr == _runeDash && len(name) >= 3 {
			isKebab = true
		}
		if vr == _runeUnderscore && len(name) >= 3 {
			isSnake = true
		}
		isLower := unicode.IsLower(prev)
		if unicode.IsUpper(vr) && (prev == 0 || isLower) {
			isPascal = true
		}
		anyLower = anyLower || isLower
		prev = vr
	}
	if !anyLower {
		isPascal = true
	}
	if !isSnake && !isKebab {
		if isCamel {
			return Camel
		}
		if isPascal {
			return Pascal
		}
	}
	if isSnake && !isKebab {
		return Snake
	}
	if isKebab && !isSnake {
		return Kebab
	}
	return None
}

func changeHead(s string, lower ...bool) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return s
	}
	if len(lower) > 0 && lower[0] {
		runes[0] = unicode.ToLower(runes[0])
	} else {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

//-----------------------------------------------------------------------------

// Case represents the X-case (snake, kebab, camel, pascal)
type Case int

func (c Case) String() string {
	switch c {
	case Snake:
		return "snake case"
	case Kebab:
		return "kebab case"
	case Camel:
		return "camel case"
	case Pascal:
		return "pascal case"
	}
	return "none"
}

func (c Case) isValid() bool {
	return Snake <= c && c <= Pascal
}

// Case valid values
const (
	None Case = iota
	Snake
	Kebab
	Camel
	Pascal
)

//-----------------------------------------------------------------------------

const (
	_runeDash       = rune('-')
	_runeUnderscore = rune('_')
	_dash           = "-"
	_underscore     = "_"
)

//-----------------------------------------------------------------------------

// errors
var (
	ErrUnknown     error = sentinelErr("unknown error")
	ErrInvalidCase error = sentinelErr("invalid case value")
	ErrCheckCase   error = sentinelErr("check case error")
)

type sentinelErr string

func (v sentinelErr) Error() string { return string(v) }

//-----------------------------------------------------------------------------
