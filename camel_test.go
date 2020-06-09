package caseconv_test

import (
	"testing"

	"github.com/dc0d/caseconv"

	assert "github.com/stretchr/testify/require"
)

func Test_ToCamel(t *testing.T) {
	var (
		assert = assert.New(t)
	)

	var (
		testCases []testCase
	)

	{
		testCases = []testCase{
			{"", ""},
			{"test", "test"},
			{"test string", "testString"},
			{"Test String", "testString"},
			{"TestV2", "testV2"},
			{"_foo_bar_", "fooBar"},
			{"version 1.2.10", "version1210"},
			{"version 1.21.0", "version1210"},
			{"version 1.2.10", "version1210"},
		}
	}

	for _, tc := range testCases {
		assert.Equal(tc.expectedOutput, caseconv.ToCamel(tc.input))
	}
}
