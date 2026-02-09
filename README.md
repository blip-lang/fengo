# fengo

<p align="center">
  <img src="assets/img/logo.png" alt="fengo logo" width="300px">
</p>

<p align="center">
  <img src="https://img.shields.io/github/license/blip-lang/fengo" alt="License">
  <img src="https://img.shields.io/github/v/release/blip-lang/fengo" alt="Release Status">
  <img src="https://img.shields.io/github/go-mod/go-version/blip-lang/fengo" alt="Go Version">
  <img src="https://github.com/blip-lang/fengo/actions/workflows/ci.yml/badge.svg" alt="Build Status">
</p>

`fengo` is a lightweight, customizable ASCII art font rendering engine written in Go. It is designed to be simple, secure, and easily extensible with your own JSON-based font formats.

## Features

- **Custom JSON Fonts**: Define your own fonts in a simple JSON format.
- **Improved Resolution**: High-quality 5-line height fonts for better readability.
- **Color Support**: Built-in support for standard ANSI colors including `pink`, `cyan`, and `magenta`.
- **Case Insensitivity**: Automatically falls back to uppercase if lowercase characters are missing in the font file.
- **Multi-line Support**: Renders multi-line strings correctly.
- **Zero Dependencies**: Uses only the Go standard library for core logic.

## Installation

### From Source

Ensure you have [Go](https://go.dev/dl/) installed:

```bash
go install github.com/blip-lang/fengo/cmd/fengo@latest
```

Or clone and build manually:

```bash
git clone https://github.com/blip-lang/fengo.git
cd fengo
go build -o fengo cmd/fengo/main.go
```

### Pre-built Binaries

You can download pre-built binaries for Windows, macOS, and Linux from the [Releases](https://github.com/blip-lang/fengo/releases) page.

## Usage

```bash
fengo [flags] <text>
```

### Flags

- `--font`: Name of a font in `assets/fonts/` (e.g., `mini`, `block`, `dots`, `slant`) or a direct path to a custom `.json` font file. Default: `mini`.
- `--color`: Color to apply to the output.

### Available Fonts

- `mini` (default): Clean, sans-serif style.
- `block`: Bold, dense characters using `@`.
- `dots`: Refined style using dots.
- `slant`: Stylish slanted ASCII art.

### Available Colors

`red`, `green`, `yellow`, `blue`, `purple`, `magenta`, `pink`, `cyan`, `white`.

### Examples

```bash
# Render with default font
fengo "Hello"

# Render with slant font in cyan
fengo --font slant --color cyan "Fengo"

# Render with block font in green
fengo --font block --color green "Success"
```

## Creating Custom Fonts

Fonts are simple JSON files. Create a file (e.g., `myfont.json`) with the following structure:

```json
{
  "name": "myfont",
  "height": 5,
  "characters": {
    "A": [
      "  A  ",
      " A A ",
      "AAAAA",
      "A   A",
      "A   A"
    ],
    " ": [
      "     ",
      "     ",
      "     ",
      "     ",
      "     "
    ]
  }
}
```

Then use it with:
```bash
fengo --font ./myfont.json "AAA"
```

## License

MIT