package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/atotto/clipboard"
	"github.com/spf13/pflag"
)

func spongebob(s string) string {
	var b strings.Builder
	upper := true
	for _, r := range s {
		if unicode.IsLetter(r) {
			if upper {
				r = unicode.ToUpper(r)
			} else {
				r = unicode.ToLower(r)
			}
			upper = !upper
		}
		b.WriteRune(r)
	}
	return b.String()
}

func die(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg, err)
		os.Exit(1)
	}
}

func main() {
	textArg := pflag.StringP("text", "t", "", "the text to convert")
	fileArg := pflag.StringP("file", "f", "", "the path to a file containing the text to convert")
	clipboardArg := pflag.BoolP("clipboard", "c", false, "copy the output to the clipboard")
	overwriteArg := pflag.BoolP("overwrite", "o", false, "overwrite the input file")
	pflag.Parse()

	text := *textArg
	if *fileArg != "" {
		b, err := os.ReadFile(*fileArg)
		die(err, "Failed to read file:")
		text = string(b)
	}

	out := spongebob(text)
	fmt.Println(out)

	if *fileArg != "" && *overwriteArg {
		die(os.WriteFile(*fileArg, []byte(out), 0644), "Failed to write to file:")
	}
	if *clipboardArg {
		die(clipboard.WriteAll(out), "Failed to copy text to clipboard:")
	}
}

