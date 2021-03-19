package caseconv_test

import (
	"testing"

	"github.com/dc0d/caseconv"

	assert "github.com/stretchr/testify/require"
)

func Test_ToKebab(t *testing.T) {
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
			{"test string", "test-string"},
			{"Test String", "test-string"},
			{"TestV2", "test-v2"},
			{"PippiLÅNGSTRUMP", "pippi-långstrump"},
		}
	}

	for _, tc := range testCases {
		assert.Equal(tc.expectedOutput, caseconv.ToKebab(tc.input))
	}
}
