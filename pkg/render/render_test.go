package render

import (
	"strings"
	"testing"

	"github.com/blip-lang/fengo/pkg/font"
)

func TestRender(t *testing.T) {
	nl := string([]byte{10})
	testFont := &font.Font{
		Name:   "test",
		Height: 2,
		Characters: map[string][]string{
			"A": {"^", "A"}, // Removed space to simplify test expectations vs implementation spacing
			"B": {"|)", "|)"},
			" ": {" ", " "},
		},
	}

	// Note: implementation adds a space between characters
	tests := []struct {
		name     string
		input    string
		opts     *Options
		expected string
	}{
		{
			name:  "Single character",
			input: "A",
			expected: strings.Join([]string{
				"^",
				"A",
			}, nl),
		},
		{
			name:  "Two characters",
			input: "AB",
			// A: ^, A
			// B: |), |)
			// Spacing: " "
			// Line 0: ^ + " " + |) -> ^ |)
			// Line 1: A + " " + |) -> A |)
			expected: strings.Join([]string{
				"^ |)",
				"A |)",
			}, nl),
		},
		{
			name:  "With space",
			input: "A B",
			// A: ^, A
			// Space: " ", " "
			// B: |), |)
			// Spacing between chars
			// ^ + " " + " " + " " + |) -> ^   |)
			// A + " " + " " + " " + |) -> A   |)
			expected: strings.Join([]string{
				"^   |)",
				"A   |)",
			}, nl),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Render(tt.input, testFont, tt.opts)
			if got != tt.expected {
				t.Errorf("Render() = %v want %v", got, tt.expected)
			}
		})
	}
}

func TestRenderMultiline(t *testing.T) {
	nl := string([]byte{10})
	testFont := &font.Font{
		Name:   "test",
		Height: 1,
		Characters: map[string][]string{
			"X": {"X"},
		},
	}

	input := "X" + nl + "X"
	expected := "X" + nl + "X"

	got := Render(input, testFont, nil)
	if got != expected {
		t.Errorf("Render() multiline = %q want %q", got, expected)
	}
}