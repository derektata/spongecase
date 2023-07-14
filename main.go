package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/atotto/clipboard"
	"github.com/spf13/pflag"
)

var (
	textArg      = pflag.StringP("text", "t", "", "the text to convert")
	fileArg      = pflag.StringP("file", "f", "", "the path to a file containing the text to convert")
	clipboardArg = pflag.BoolP("clipboard", "c", false, "copy the output to the clipboard")
	overwriteArg = pflag.BoolP("overwrite", "o", false, "overwrite the input file")
)

func convertToSpongebobCase(text string) string {
	var builder strings.Builder
	count := 0
	for _, c := range text {
		if unicode.IsLetter(c) {
			if count%2 == 0 {
				builder.WriteRune(unicode.ToUpper(c))
			} else {
				builder.WriteRune(unicode.ToLower(c))
			}
			count++
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func copyToClipboard(text string) {
	err := clipboard.WriteAll(text)
	checkErrorAndExit(err, "Failed to copy text to clipboard:")
}

func checkErrorAndExit(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %v\n", message, err)
		os.Exit(1)
	}
}

func main() {
	pflag.Parse()

	var text string
	if *fileArg != "" {
		if _, err := os.Stat(*fileArg); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "No such file: '%s'\n", *fileArg)
			os.Exit(1)
		}

		bytes, err := os.ReadFile(*fileArg)
		checkErrorAndExit(err, "Failed to read file:")

		text = string(bytes)
	} else {
		text = *textArg
	}

	output := convertToSpongebobCase(text)
	fmt.Println(output)

	if *fileArg != "" && *overwriteArg {
		err := os.WriteFile(*fileArg, []byte(output), 0644)
		checkErrorAndExit(err, "Failed to write to file:")
	}

	if *clipboardArg {
		copyToClipboard(output)
	}
}
