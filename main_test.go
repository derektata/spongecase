package main

import (
	"testing"
)

func TestConvertToSpongebobCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "lowercase string",
			input:    "hello",
			expected: "HeLlO",
		},
		{
			name:     "uppercase string",
			input:    "WORLD",
			expected: "WoRlD",
		},
		{
			name:     "string with spaces",
			input:    "hello world",
			expected: "HeLlO wOrLd",
		},
		{
			name:     "string with punctuation",
			input:    "hello, world!",
			expected: "HeLlO, wOrLd!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToSpongebobCase(tt.input)
			if result != tt.expected {
				t.Errorf("convertToSpongebobCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
