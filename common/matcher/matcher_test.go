package matcher

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}

func TestMatchPercent(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		target   string
		expected float64
	}{
		{
			name:     "Full match 100%",
			input:    "I love going to the school everyday",
			target:   "Love school",
			expected: 100.0,
		},
		{
			name:     "Partial match 50%",
			input:    "Quantum energy and galaxy exploration",
			target:   "Quantum farming medieval",
			expected: 33.33, // 1 из 3
		},
		{
			name:     "Zero match",
			input:    "Cats and dogs playing outside",
			target:   "Blockchain validator consensus",
			expected: 0.0,
		},
		{
			name:     "Empty target",
			input:    "Anything here",
			target:   "",
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MatchPercent(tt.input, tt.target)

			if !floatEquals(result, tt.expected) {
				t.Errorf("Expected %.2f, got %.2f", tt.expected, result)
			}
		})
	}
}