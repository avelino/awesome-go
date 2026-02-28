package slug

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with spaces",
			input:    "some string with spaces",
			expected: "some-string-with-spaces",
		},
		{
			name:     "with out any non-literal chars",
			input:    "inputstring",
			expected: "inputstring",
		},
		{
			name:     "with whitespace prefix and suffix",
			input:    "   inputstring   ",
			expected: "inputstring",
		},
		{
			name:     "a mix of special characters",
			input:    "   an input string (with.special/chars,such_as:§\\?$/§&!)   ",
			expected: "an-input-string-with-specialchars-such-as",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.input); got != tt.expected {
				t.Errorf("Generate() = %v, want %v", got, tt.expected)
			}
		})
	}
}
