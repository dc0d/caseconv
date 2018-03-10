package caseconv

import (
	"bufio"
	"fmt"
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
		"a_b": {
			Snake:  "ab",
			Kebab:  "a_b",
			Camel:  "a_b",
			Pascal: "a_b",
		},
		"a_a_b": {
			Snake:  "aab",
			Kebab:  "a_a_b",
			Camel:  "a_a_b",
			Pascal: "a_a_b",
		},
		"aa_a1_2b": {
			Snake:  "aaa12b",
			Kebab:  "aa_a1_2b",
			Camel:  "aa_a1_2b",
			Pascal: "aa_a1_2b",
		},
		"a-b": {
			Snake:  "a-b",
			Kebab:  "ab",
			Camel:  "a-b",
			Pascal: "a-b",
		},
		"1a-2b": {
			Snake:  "1a-2b",
			Kebab:  "1a2b",
			Camel:  "1a-2b",
			Pascal: "1a-2b",
		},
		"1a-2b-jlnlxas": {
			Snake:  "1a-2b-jlnlxas",
			Kebab:  "1a2bjlnlxas",
			Camel:  "1a-2b-jlnlxas",
			Pascal: "1a-2b-jlnlxas",
		},
		"1a-2b_": {
			Snake:  "1a-2b",
			Kebab:  "1a2b_",
			Camel:  "1a-2b_",
			Pascal: "1a-2b_",
		},
		"A": {
			Snake:  "A",
			Kebab:  "A",
			Camel:  "a",
			Pascal: "A",
		},
		"a": {
			Snake:  "a",
			Kebab:  "a",
			Camel:  "a",
			Pascal: "a",
		},
		"aA": {
			Snake:  "aA",
			Kebab:  "aA",
			Camel:  "aA",
			Pascal: "aA",
		},
		"aa": {
			Snake:  "aa",
			Kebab:  "aa",
			Camel:  "aa",
			Pascal: "aa",
		},
		"aaBcDe": {
			Snake:  "aaBcDe",
			Kebab:  "aaBcDe",
			Camel:  "aaBcDe",
			Pascal: "aaBcDe",
		},
		"abcdefg": {
			Snake:  "abcdefg",
			Kebab:  "abcdefg",
			Camel:  "abcdefg",
			Pascal: "abcdefg",
		},
		"aBCDEFG": {
			Snake:  "aBCDEFG",
			Kebab:  "aBCDEFG",
			Camel:  "aBCDEFG",
			Pascal: "aBCDEFG",
		},
		"Abcdefg": {
			Snake:  "Abcdefg",
			Kebab:  "Abcdefg",
			Camel:  "abcdefg",
			Pascal: "Abcdefg",
		},
		"ABCDEFG": {
			Snake:  "ABCDEFG",
			Kebab:  "ABCDEFG",
			Camel:  "aBCDEFG",
			Pascal: "ABCDEFG",
		},
		"Aa": {
			Snake:  "Aa",
			Kebab:  "Aa",
			Camel:  "aa",
			Pascal: "Aa",
		},
		"AA": {
			Snake:  "AA",
			Kebab:  "AA",
			Camel:  "aA",
			Pascal: "AA",
		},
		"AbCd": {
			Snake:  "AbCd",
			Kebab:  "AbCd",
			Camel:  "abCd",
			Pascal: "AbCd",
		},
		"AbCDEfG": {
			Snake:  "AbCDEfG",
			Kebab:  "AbCDEfG",
			Camel:  "abCDEfG",
			Pascal: "AbCDEfG",
		},
		"": {
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
		{"aa-bb-cc", -1, -1, false, "", ErrInvalidCase},
		{"aa-bb-cc", Kebab, Snake, false, "aa_bb_cc", nil},
		{"aa-bb-cc_", Kebab, Snake, true, "", ErrCheckCase},
		{"aa-bb-cc", Kebab, Camel, false, "aaBbCc", nil},
		{"aa-bb-cc", Kebab, Pascal, false, "AaBbCc", nil},
		{"aa_bb_cc", Snake, Kebab, true, "aa-bb-cc", nil},
		{"aa_bb_cc-", Snake, Kebab, true, "", ErrCheckCase},
		{"aa_bb_cc", Snake, Camel, true, "aaBbCc", nil},
		{"aa_bb_cc", Snake, Pascal, true, "AaBbCc", nil},
		{"aaBbCc", Camel, Snake, true, "aa_Bb_Cc", nil},
		{"aaBbCc", Camel, Kebab, true, "aa-Bb-Cc", nil},
		{"aaBbCc", Camel, Pascal, true, "AaBbCc", nil},
		{"aa_BbCc", Camel, Pascal, true, "", ErrCheckCase},
		{"AaBbCc", Pascal, Snake, true, "Aa_Bb_Cc", nil},
		{"AaBbCc", Pascal, Kebab, true, "Aa-Bb-Cc", nil},
		{"AaBbCc", Pascal, Camel, true, "aaBbCc", nil},
		{"AaBb-Cc", Pascal, Camel, true, "", ErrCheckCase},

		{"AdminHandler", Pascal, Snake, true, "Admin_Handler", nil},

		{"cc", Pascal, Snake, true, "", ErrCheckCase},
		{"Cc", Pascal, Snake, true, "Cc", nil},
		{"C", Pascal, Snake, true, "C", nil},

		{"CDE", Pascal, Snake, true, "CDE", nil},
		{"cde", Camel, Snake, true, "cde", nil},
	}

	for _, v := range cases {
		s, err := Convert(v.str, v.from, v.to, v.check)
		assert.Equal(v.expected, s)
		if !assert.Equal(errors.Cause(v.expectedErr), err) {
			t.Log(v.str)
		}
	}
}

func ExampleConvert() {
	name := "SomePascalCaseName"

	s, err := Convert(name, Pascal, Snake)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	s, err = Convert(name, Pascal, Kebab)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	s, err = Convert(name, Pascal, Camel)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	// Output:
	// Some_Pascal_Case_Name
	// Some-Pascal-Case-Name
	// somePascalCaseName
}
