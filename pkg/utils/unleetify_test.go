package utils

import "testing"

func TestUnleetify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Maulik Test",
			input:    "v1bh475u",
			expected: "vibhatsu",
		},
		{
			name:     "Praneeth Test",
			input:    "ph03n1x",
			expected: "phoenix",
		},
		{
			name:     "Osama Test",
			input:    "0S4M4",
			expected: "oSaMa",
		},
		{
			name:     "Hashkat Test",
			input:    "h4shk4t",
			expected: "hashkat",
		},
		{
			name:     "Thunder God Test",
			input:    "7π4nd5R_G0d",
			expected: "thandsR_God",
		},
		{
			name:     "Basic leet replacements",
			input:    "h3ll0 w0rld! 7h15 15 4 73s7.",
			expected: "hello world! this is a test.",
		},
		{
			name:     "No leet characters",
			input:    "hello world!",
			expected: "hello world!",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Only leet characters",
			input:    "4@837$+",
			expected: "aabetst",
		},
		{
			name:     "Complex leet string",
			input:    "7h1$ 1$ @ c0mpl3x l33t $7r1ng.",
			expected: "this is a complex leet string.",
		},
		{
			name:     "Leet characters with spaces",
			input:    "7h1$ 1$ 4 73s7 w1th $p4c3$.",
			expected: "this is a test with spaces.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unleetify(tt.input)
			if result != tt.expected {
				t.Errorf("Test case %q failed: Unleetify(%q) = %q; expected %q", tt.name, tt.input, result, tt.expected)
			}
		})
	}
}
