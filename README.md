# mhtml-to-html
This command line converts .mhtml to .html

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gonejack/mhtml-to-html)
![Build](https://github.com/gonejack/mhtml-to-html/actions/workflows/go.yml/badge.svg)
[![GitHub license](https://img.shields.io/github/license/gonejack/mhtml-to-html.svg?color=blue)](LICENSE)

### Install
```shell
> go get github.com/gonejack/mhtml-to-html
```

### Usage

By default, `mhtml-to-html` outputs HTML to stdout (like `pdftotext`):

```shell
# Extract HTML content to stdout
> mhtml-to-html document.mht

# Save HTML to a specific file
> mhtml-to-html document.mht -o output.html

# Process with verbose output
> mhtml-to-html document.mhtml --verbose -o converted.html
```

#### Flags
```
  -h, --help              Show context-sensitive help.
  -o, --output=STRING     Output file (default: stdout).
      --verbose           Verbose printing.
      --about             Show about.
```

#### Examples
```shell
# Extract text content to terminal
> mhtml-to-html email.mht

# Save converted HTML file
> mhtml-to-html email.mht -o email.html

# Pipe output to other tools
> mhtml-to-html document.mht | grep "important text"
```
