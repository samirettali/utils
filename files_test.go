package utils

import (
	"bytes"
	"testing"
)

func compare(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestReadlines(t *testing.T) {
	type test struct {
		input    string
		expected []string
	}

	tests := []test{
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    "\n",
			expected: []string{"", ""},
		},
		{
			input:    "line1",
			expected: []string{"line1"},
		},
		{
			input:    "1\n2\n3",
			expected: []string{"1", "2", "3"},
		},
		{
			input:    "1\n2\n3\n ",
			expected: []string{"1", "2", "3", " "},
		},
	}

	for ti, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			var buffer bytes.Buffer
			buffer.WriteString(tc.input)
			actual, err := readlines(&buffer)
			if err != nil || !compare(actual, tc.expected) {
				t.Errorf("test %d: expected %+v of length %d, got %+v of length %d", ti, tc.expected, len(tc.expected), actual, len(actual))
			}
		})
	}

}
