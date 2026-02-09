package font

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Font represents a font used for rendering text.
// It contains metadata and the character mapping.
type Font struct {
	Name       string              `json:"name"`
	Height     int                 `json:"height"`
	Characters map[string][]string `json:"characters"`
}

// Load reads a font file from the specified path and returns a Font struct.
func Load(path string) (*Font, error) {
	if path == "" {
		return nil, fmt.Errorf("font path cannot be empty")
	}

	cleanPath := filepath.Clean(path)
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %w", err)
	}

	var f Font
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("failed to parse font data: %w", err)
	}

	if err := f.validate(); err != nil {
		return nil, fmt.Errorf("invalid font: %w", err)
	}

	return &f, nil
}

// validate checks if the font structure is valid.
func (f *Font) validate() error {
	if f.Name == "" {
		return fmt.Errorf("font name is required")
	}
	if f.Height <= 0 {
		return fmt.Errorf("font height must be positive")
	}
	if len(f.Characters) == 0 {
		return fmt.Errorf("font must contain characters")
	}

	for char, lines := range f.Characters {
		if len(lines) != f.Height {
			return fmt.Errorf("character %q has height %d, expected %d", char, len(lines), f.Height)
		}
	}
	return nil
}
