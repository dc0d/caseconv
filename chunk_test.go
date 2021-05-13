package caseconv

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_chunk(t *testing.T) {
	testCases := []testCase{
		{"", ""},

		// single words
		{"test", "test"},
		{"TEST", "test"},
		{"PIPPI LÅNGSTRUMP", "pippi långstrump"},

		// camel case
		{"testString", "test string"},
		{"testString123", "test string123"},
		{"testString_1_2_3", "test string 1 2 3"},
		{"x_256", "x 256"},
		{"anHTMLTag", "an html tag"},
		{"ID123String", "id123 string"},
		{"Id123String", "id123 string"},
		{"foo bar123", "foo bar123"},
		{"a1bStar", "a1b star"},

		// constant case
		{"CONSTANT_CASE ", "constant case"},
		{"CONST123_FOO", "const123 foo"},

		// random cases
		{"FOO_bar", "foo bar"},
		{"XMLHttpRequest", "xml http request"},
		{"IQueryAArgs", "i query a args"},

		// non-alphanumeric separators
		{"dot.case", "dot case"},
		{"path/case", "path case"},
		{"snake_case", "snake case"},
		{"snake_case123", "snake case123"},
		{"snake_case_123", "snake case 123"},

		// Punctuation.
		{`"quotes"`, "quotes"},

		// space between number parts
		{"version 0.45.0", "version 0 45 0"},
		{"version 0..78..9", "version 0 78 9"},
		{"version 4_99/4", "version 4 99 4"},

		// whitespace.
		{"  test  ", "test"},

		// Number string input.
		{"something_2014_other", "something 2014 other"},

		// https://github.com/blakeembrey/change-case/issues/21
		{"amazon s3 data", "amazon s3 data"},
		{"foo_13_bar", "foo 13 bar"},
		{"Pippil ÅNGSTRUMP", "pippil ångstrump"},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedOutput, strings.Join(chunk(tc.input), " "))
	}
}

type testCase struct {
	input          string
	expectedOutput string
}
