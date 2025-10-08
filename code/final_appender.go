package main

import (
	"fmt"
	"os"
	"strings"
)

func getContent() string {
	var sb strings.Builder
	sb.WriteString("\n\n## [panic](https://gobyexample.com/panic)\n\n")
	sb.WriteString("A `panic` typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.\n\n")
	sb.WriteString("```go\n")
	sb.WriteString("package main\n\n")
	sb.WriteString("import \"os\"\n\n")
	sb.WriteString("func main() {\n")
	sb.WriteString("    // We'll use panic throughout this site to check for unexpected errors.\n")
	sb.WriteString("    // This is the only program on the site designed to panic.\n")
	sb.WriteString("    panic(\"a problem\")\n\n")
	sb.WriteString("    // A common use of panic is to abort if a function returns an error value\n")
	sb.WriteString("    // that we don't know how to (or want to) handle. Here's an example of\n")
	sb.WriteString("    // panicking if we get an unexpected error when creating a new file.\n")
	sb.WriteString("    _, err := os.Create(\"/tmp/file\")\n")
	sb.WriteString("    if err != nil {\n")
	sb.WriteString("        panic(err)\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")
	sb.WriteString("Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.\n\n")
	sb.WriteString("```shell\n")
	sb.WriteString("$ go run panic.go\n")
	sb.WriteString("panic: a problem\n\n")
	sb.WriteString("goroutine 1 [running]:\n")
	sb.WriteString("main.main()\n")
	sb.WriteString("    /.../panic.go:12 +0x47\n")
	sb.WriteString("...\n")
	sb.WriteString("exit status 2\n")
	sb.WriteString("```\n\n")
	sb.WriteString("Note that unlike some other languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.\n")
	return sb.String()
}

func main() {
	filePath := "/Users/mac/Documents/person/knowledges/go_learns/go_by_example/学习内容.md"
	contentToAppend := getContent()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(contentToAppend); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully appended content to %s\n", filePath)
}
