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

// convertToSpongebobCase converts the input text to SpongeBob case.
//
// The input text is the string to be converted.
// Returns the converted text as a string.
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

// copyToClipboard copies the given text to the clipboard.
//
// Parameters:
// - text: the text to be copied to the clipboard.
//
// Returns:
// - None.
func copyToClipboard(text string) {
	err := clipboard.WriteAll(text)
	checkErrorAndExit(err, "Failed to copy text to clipboard:")
}

// checkErrorAndExit checks if the given error is not nil and prints an error message to stderr along with the error.
// If the error is not nil, the program exits with a status code of 1.
//
// Parameters:
// - err: the error to check.
// - message: the message to print before the error.
//
// Returns:
// - None.
func checkErrorAndExit(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %v\n", message, err)
		os.Exit(1)
	}
}

// handleFileInput reads the contents of a file and returns it as a string.
//
// Parameters:
// - filePath: the path to the file to read.
//
// Returns:
// - string: the contents of the file as a string.
func handleFileInput(filePath string) string {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "No such file: '%s'\n", filePath)
		os.Exit(1)
	}

	bytes, err := os.ReadFile(filePath)
	checkErrorAndExit(err, "Failed to read file:")

	return string(bytes)
}

// handleFileOutput writes the given output string to the specified file path.
//
// Parameters:
// - filePath: the path to the file to write to.
// - output: the string to write to the file.
//
// Returns:
// - None.
func handleFileOutput(filePath, output string) {
	err := os.WriteFile(filePath, []byte(output), 0644)
	checkErrorAndExit(err, "Failed to write to file:")
}

func main() {
	pflag.Parse()

	var text string
	if *fileArg != "" {
		text = handleFileInput(*fileArg)
	} else {
		text = *textArg
	}

	output := convertToSpongebobCase(text)
	fmt.Println(output)

	if *fileArg != "" && *overwriteArg {
		handleFileOutput(*fileArg, output)
	}

	if *clipboardArg {
		copyToClipboard(output)
	}
}
