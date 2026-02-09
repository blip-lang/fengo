package render

import (
	"strings"

	"github.com/blip-lang/fengo/pkg/color"
	"github.com/blip-lang/fengo/pkg/font"
)

// Options configuration for rendering.
type Options struct {
	Color string
}

// Render converts the input text into ASCII art using the provided font.
func Render(text string, f *font.Font, opts *Options) string {
	if f == nil {
		return ""
	}

	// Split by newline to handle multi-line input
	lines := strings.Split(text, "\n")
	var renderedLines []string

	for _, line := range lines {
		renderedLines = append(renderedLines, renderLine(line, f)...)
	}

	result := strings.Join(renderedLines, "\n")
	
	if opts != nil && opts.Color != "" {
		return color.Apply(result, opts.Color)
	}

	return result
}

func renderLine(text string, f *font.Font) []string {
	if text == "" {
		// If line is empty, we return empty lines of height to maintain vertical spacing?
		// But usually empty string in Split results means an empty line.
		// Let's return a blank vertical spacer (height * empty strings)
        // Wait, if I split "A\n\nB", I get "A", "", "B".
        // Rendering "" should probably result in a vertical gap of height?
        // Or just nothing?
        // Standard figlet: empty line -> empty line in output (height lines of spaces? No).
        // If I return empty list, it collapses.
        // Let's return f.Height empty strings.
        emptyBuffer := make([]string, f.Height)
        return emptyBuffer
	}

	// Initialize buffer for this line of text
	// buffer[0] is the top line of the output
	// buffer[1] is the second line...
	buffer := make([]string, f.Height)

	for _, char := range text {
		charStr := string(char)
		charLines, ok := f.Characters[charStr]
		if !ok {
			// Try uppercase
			upperCharStr := strings.ToUpper(charStr)
			if charLines, ok = f.Characters[upperCharStr]; !ok {
				// Fallback to "?" if available, else space, else blank
				if fallback, exists := f.Characters["?"]; exists {
					charLines = fallback
				} else if fallback, exists := f.Characters[" "]; exists {
					charLines = fallback
				} else {
					// Absolute fallback: empty space of correct width
					charLines = make([]string, f.Height)
					for i := 0; i < f.Height; i++ {
						charLines[i] = " "
					}
				}
			}
		}

		// Append character lines to buffer
		for i := 0; i < f.Height; i++ {
			// Add spacing if it's not the first character
			if len(buffer[i]) > 0 {
				buffer[i] += " "
			}
			buffer[i] += charLines[i]
		}
	}

	return buffer
}