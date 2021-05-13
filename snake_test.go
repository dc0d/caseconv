package caseconv_test

import (
	"testing"

	"github.com/dc0d/caseconv"

	assert "github.com/stretchr/testify/require"
)

func Test_ToSnake(t *testing.T) {
	testCases := []testCase{
		{"", ""},
		{"test", "test"},
		{"test string", "test_string"},
		{"Test String", "test_string"},
		{"TestV2", "test_v2"},
		{"PippiLÅNGSTRUMP", "pippi_långstrump"},
		{"PippilÅNGSTRUMP", "pippil_ångstrump"},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedOutput, caseconv.ToSnake(tc.input))
	}
}
