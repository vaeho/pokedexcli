package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   foo   bar   baz   ",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    "   singleword   ",
			expected: []string{"singleword"},
		},
		{
			input:    "      ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   UPPERCASE  lowercase   MixedCase   ",
			expected: []string{"uppercase", "lowercase", "mixedcase"},
		},
		{
			input:    "   multiple   spaces   between   words   ",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("For input '%s', expected '%s' but got '%s'", c.input, c.expected[i], actual[i])
			}
		}
	}
}
