package color

import "fmt"

// Color definitions.
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// Map maps string names to color codes.
var Map = map[string]string{
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"purple":  Magenta,
	"magenta": Magenta,
	"pink":    Magenta,
	"cyan":    Cyan,
	"white":   White,
}

// Apply wraps the given text with the specified color code.
// If the color is not found or empty, it returns the text as is.
func Apply(text, colorName string) string {
	code, ok := Map[colorName]
	if !ok {
		return text
	}
	return fmt.Sprintf("%s%s%s", code, text, Reset)
}