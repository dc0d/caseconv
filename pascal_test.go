package caseconv_test

import (
	"testing"

	"github.com/dc0d/caseconv"

	assert "github.com/stretchr/testify/require"
)

func Test_ToPascal(t *testing.T) {
	var (
		assert = assert.New(t)
	)

	var (
		testCases []testCase
	)

	{
		testCases = []testCase{
			{"", ""},
			{"test", "Test"},
			{"test string", "TestString"},
			{"Test String", "TestString"},
			{"TestV2", "TestV2"},
			{"version 1.2.10", "Version1210"},
			{"version 1.21.0", "Version1210"},
			{"LÅNGSTRUMP", "Långstrump"},
			{"PippiLÅNGSTRUMP", "PippiLångstrump"},
		}
	}

	for _, tc := range testCases {
		assert.Equal(tc.expectedOutput, caseconv.ToPascal(tc.input))
	}
}

type testCase struct {
	input          string
	expectedOutput string
}
