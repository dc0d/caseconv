package caseconv

import (
	"bufio"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_sentinelErr(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(sentinelErr("unknown error"), ErrUnknown)
}

func Test_guess(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]Case{
		"a_b":           Snake,
		"a_a_b":         Snake,
		"aa_a1_2b":      Snake,
		"a-b":           Kebab,
		"1a-2b":         Kebab,
		"1a-2b-jlnlxas": Kebab,
		"1a-2b_":        None,
		"A":             Pascal,
		"a":             Camel,
		"aA":            Camel,
		"aa":            Camel,
		"aaBcDe":        Camel,
		"abcdefg":       Camel,
		"aBCDEFG":       Camel,
		"Abcdefg":       Pascal,
		"ABCDEFG":       Pascal,
		"Aa":            Pascal,
		"AA":            Pascal,
		"AbCd":          Pascal,
		"AbCDEfG":       Pascal,
	}

	for k, v := range cases {
		d := guess([]rune(k))
		if !assert.Equal(v, d) {
			t.Log(k, v, d)
		}
	}
}

func Test_splitter(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]map[Case]string{
		"a_b": map[Case]string{
			Snake:  "ab",
			Kebab:  "a_b",
			Camel:  "a_b",
			Pascal: "a_b",
		},
		"a_a_b": map[Case]string{
			Snake:  "aab",
			Kebab:  "a_a_b",
			Camel:  "a_a_b",
			Pascal: "a_a_b",
		},
		"aa_a1_2b": map[Case]string{
			Snake:  "aaa12b",
			Kebab:  "aa_a1_2b",
			Camel:  "aa_a1_2b",
			Pascal: "aa_a1_2b",
		},
		"a-b": map[Case]string{
			Snake:  "a-b",
			Kebab:  "ab",
			Camel:  "a-b",
			Pascal: "a-b",
		},
		"1a-2b": map[Case]string{
			Snake:  "1a-2b",
			Kebab:  "1a2b",
			Camel:  "1a-2b",
			Pascal: "1a-2b",
		},
		"1a-2b-jlnlxas": map[Case]string{
			Snake:  "1a-2b-jlnlxas",
			Kebab:  "1a2bjlnlxas",
			Camel:  "1a-2b-jlnlxas",
			Pascal: "1a-2b-jlnlxas",
		},
		"1a-2b_": map[Case]string{
			Snake:  "1a-2b",
			Kebab:  "1a2b_",
			Camel:  "1a-2b_",
			Pascal: "1a-2b_",
		},
		"A": map[Case]string{
			Snake:  "A",
			Kebab:  "A",
			Camel:  "a",
			Pascal: "A",
		},
		"a": map[Case]string{
			Snake:  "a",
			Kebab:  "a",
			Camel:  "a",
			Pascal: "a",
		},
		"aA": map[Case]string{
			Snake:  "aA",
			Kebab:  "aA",
			Camel:  "aA",
			Pascal: "aA",
		},
		"aa": map[Case]string{
			Snake:  "aa",
			Kebab:  "aa",
			Camel:  "aa",
			Pascal: "aa",
		},
		"aaBcDe": map[Case]string{
			Snake:  "aaBcDe",
			Kebab:  "aaBcDe",
			Camel:  "aaBcDe",
			Pascal: "aaBcDe",
		},
		"abcdefg": map[Case]string{
			Snake:  "abcdefg",
			Kebab:  "abcdefg",
			Camel:  "abcdefg",
			Pascal: "abcdefg",
		},
		"aBCDEFG": map[Case]string{
			Snake:  "aBCDEFG",
			Kebab:  "aBCDEFG",
			Camel:  "aBCDEFG",
			Pascal: "aBCDEFG",
		},
		"Abcdefg": map[Case]string{
			Snake:  "Abcdefg",
			Kebab:  "Abcdefg",
			Camel:  "abcdefg",
			Pascal: "Abcdefg",
		},
		"ABCDEFG": map[Case]string{
			Snake:  "ABCDEFG",
			Kebab:  "ABCDEFG",
			Camel:  "aBCDEFG",
			Pascal: "ABCDEFG",
		},
		"Aa": map[Case]string{
			Snake:  "Aa",
			Kebab:  "Aa",
			Camel:  "aa",
			Pascal: "Aa",
		},
		"AA": map[Case]string{
			Snake:  "AA",
			Kebab:  "AA",
			Camel:  "aA",
			Pascal: "AA",
		},
		"AbCd": map[Case]string{
			Snake:  "AbCd",
			Kebab:  "AbCd",
			Camel:  "abCd",
			Pascal: "AbCd",
		},
		"AbCDEfG": map[Case]string{
			Snake:  "AbCDEfG",
			Kebab:  "AbCDEfG",
			Camel:  "abCDEfG",
			Pascal: "AbCDEfG",
		},
		"": map[Case]string{
			Snake:  "",
			Kebab:  "",
			Camel:  "",
			Pascal: "",
		},
	}

	for k, v := range cases {
		k, v := k, v
		for vk, vv := range v {
			scn := bufio.NewScanner(strings.NewReader(k))
			scn.Split(newSplitter(vk).split())
			var parts []string
			for scn.Scan() {
				parts = append(parts, scn.Text())
			}
			s := strings.Join(parts, "")
			if vk == Camel {
				s = changeHead(s, true)
			}
			if !assert.Equal(vv, s) {
				t.Log(vv, s, vk)
			}
		}
	}
}

func Test_Convert(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		str         string
		from, to    Case
		check       bool
		expected    string
		expectedErr error
	}

	cases := []testCase{
		testCase{"aa-bb-cc", -1, -1, false, "", ErrInvalidCase},
		testCase{"aa-bb-cc", Kebab, Snake, false, "aa_bb_cc", nil},
		testCase{"aa-bb-cc_", Kebab, Snake, true, "", ErrCheckCase},
		testCase{"aa-bb-cc", Kebab, Camel, false, "aaBbCc", nil},
		testCase{"aa-bb-cc", Kebab, Pascal, false, "AaBbCc", nil},
		testCase{"aa_bb_cc", Snake, Kebab, true, "aa-bb-cc", nil},
		testCase{"aa_bb_cc-", Snake, Kebab, true, "", ErrCheckCase},
		testCase{"aa_bb_cc", Snake, Camel, true, "aaBbCc", nil},
		testCase{"aa_bb_cc", Snake, Pascal, true, "AaBbCc", nil},
		testCase{"aaBbCc", Camel, Snake, true, "aa_Bb_Cc", nil},
		testCase{"aaBbCc", Camel, Kebab, true, "aa-Bb-Cc", nil},
		testCase{"aaBbCc", Camel, Pascal, true, "AaBbCc", nil},
		testCase{"aa_BbCc", Camel, Pascal, true, "", ErrCheckCase},
		testCase{"AaBbCc", Pascal, Snake, true, "Aa_Bb_Cc", nil},
		testCase{"AaBbCc", Pascal, Kebab, true, "Aa-Bb-Cc", nil},
		testCase{"AaBbCc", Pascal, Camel, true, "aaBbCc", nil},
		testCase{"AaBb-Cc", Pascal, Camel, true, "", ErrCheckCase},

		testCase{"AdminHandler", Pascal, Snake, true, "Admin_Handler", nil},

		testCase{"cc", Pascal, Snake, true, "", ErrCheckCase},
		testCase{"Cc", Pascal, Snake, true, "Cc", nil},
		testCase{"C", Pascal, Snake, true, "C", nil},

		testCase{"CDE", Pascal, Snake, true, "CDE", nil},
		testCase{"cde", Camel, Snake, true, "cde", nil},
	}

	for _, v := range cases {
		s, err := Convert(v.str, v.from, v.to, v.check)
		assert.Equal(v.expected, s)
		if !assert.Equal(errors.Cause(v.expectedErr), err) {
			t.Log(v.str)
		}
	}
}
