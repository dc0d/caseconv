package caseconv_test

import (
	"testing"

	"github.com/dc0d/caseconv"

	assert "github.com/stretchr/testify/require"
)

func Test_ToSnake(t *testing.T) {
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
			{"test string", "test_string"},
			{"Test String", "test_string"},
			{"TestV2", "test_v2"},
		}
	}

	for _, tc := range testCases {
		assert.Equal(tc.expectedOutput, caseconv.ToSnake(tc.input))
	}
}
