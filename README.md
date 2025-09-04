# mhtml-to-html
This command line converts .mhtml to .html

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mpr1255/mhtml-to-html)
![Build](https://github.com/mpr1255/mhtml-to-html/actions/workflows/go.yml/badge.svg)
[![GitHub license](https://img.shields.io/github/license/mpr1255/mhtml-to-html.svg?color=blue)](LICENSE)

### Install
```shell
> go install github.com/mpr1255/mhtml-to-html@latest
```

### Usage

By default, `mhtml-to-html` outputs HTML to stdout (like `pdftotext`).

**Output Behavior:**
- **Stdout mode** (no `-o` flag): Images are stripped for clean text processing in pipelines
- **File mode** (`-o` flag): Images and resources are preserved
- **Extract files mode** (`--extract-files`): Creates `{filename}_files/` directory with external resources

```shell
# Extract text content to stdout (images removed)
> mhtml-to-html document.mht

# Save HTML to file (images preserved)
> mhtml-to-html document.mht -o output.html

# Extract with external resource files
> mhtml-to-html document.mht -o output.html --extract-files

# Process with verbose output
> mhtml-to-html document.mhtml --verbose -o converted.html
```

#### Flags
```
  -h, --help              Show context-sensitive help.
      --verbose           Verbose printing.
      --version           Show version.
      --extract-files     Extract resources to external files (default: embed as data URIs).
  -o, --output=STRING     Output file (default: stdout).
```

#### Examples
```shell
# Extract text content to terminal (images stripped)
> mhtml-to-html email.mht

# Save converted HTML file (images preserved)
> mhtml-to-html email.mht -o email.html

# Save with external resource files
> mhtml-to-html email.mht -o email.html --extract-files

# Pipe output to other tools (text processing)
> mhtml-to-html document.mht | grep "important text"
```
