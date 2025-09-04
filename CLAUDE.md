# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Build and Test
- Build: `go build -v ./...`
- Test: `go test -v ./...`
- Install: `go install github.com/mpr1255/mhtml-to-html@latest`

### Usage
- Run the tool: `go run main.go *.mht` or `go run main.go *.mhtml`
- Built binary usage: `mhtml-to-html *.mht`
- Flags: `--verbose` for verbose output, `--about` for project info

## Architecture

This is a Go CLI tool that converts MHTML files to HTML files with extracted resources.

### Core Components

- **main.go**: Entry point that initializes and runs the MHTMLToHTML command
- **cmd/cmd.go**: Main command logic, file processing orchestration, and CLI interface using Kong
- **cmd/parse.go**: MIME parsing logic for multipart MHTML files with support for various content encodings

### Processing Flow

1. **File Discovery**: Automatically finds *.mht and *.mhtml files if none specified
2. **MIME Parsing**: Parses MHTML as multipart MIME with recursive handling of nested parts
3. **Content Extraction**: Separates HTML content from embedded resources (images, CSS, JS)
4. **Resource Handling**: Saves resources to `{filename}_files/` directory organized by MIME type
5. **Reference Updating**: Updates HTML references to point to extracted local files using goquery for DOM manipulation

### Key Dependencies

- **github.com/alecthomas/kong**: CLI argument parsing and help generation
- **github.com/PuerkitoBio/goquery**: HTML parsing and DOM manipulation for reference updates
- **Standard library**: MIME multipart parsing, base64/quoted-printable decoding

### File Structure

The tool creates output files as:
- `{input}.html`: Main HTML file with updated references
- `{input}_files/`: Directory containing extracted resources organized by MIME type