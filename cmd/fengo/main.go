package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blip-lang/fengo/pkg/color"
	"github.com/blip-lang/fengo/pkg/font"
	"github.com/blip-lang/fengo/pkg/render"
)

func main() {
	fontName := flag.String("font", "mini", "Font name or path to font file")
	colorName := flag.String("color", "", "Color name (red, green, blue, etc.)")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: fengo [flags] <text>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	text := strings.Join(args, " ")
	text = strings.ReplaceAll(text, "\\n", "\n")
	fontPath := resolveFontPath(*fontName)

	f, err := font.Load(fontPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading font:", err)
		os.Exit(1)
	}

	opts := &render.Options{
		Color: strings.ToLower(*colorName),
	}

	if opts.Color != "" {
		if _, ok := color.Map[opts.Color]; !ok {
			fmt.Fprintf(os.Stderr, "Warning: unknown color %q. Available colors: %s\n", *colorName, getAvailableColors())
		}
	}

	output := render.Render(text, f, opts)
	fmt.Println(output)
}

func getAvailableColors() string {
	var colors []string
	for k := range color.Map {
		colors = append(colors, k)
	}
	return strings.Join(colors, ", ")
}

func resolveFontPath(name string) string {
	// If it's a file that exists, return it
	if _, err := os.Stat(name); err == nil {
		return name
	}

	// If it doesn't end with .json, append it
	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	// Check in assets/fonts/
	candidate := filepath.Join("assets", "fonts", name)
	if _, err := os.Stat(candidate); err == nil {
		return candidate
	}

	// Return original if nothing found, to let Load return the error
	return name
}