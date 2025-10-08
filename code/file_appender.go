package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./file_appender <file_path>")
		fmt.Println("Content to append should be piped via stdin.")
		os.Exit(1)
	}

	filePath := os.Args[1]

	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString("\n"); err != nil {
		fmt.Printf("Error writing newline to file: %v\n", err)
		os.Exit(1)
	}

	if _, err := f.Write(content); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully appended content to %s\n", filePath)
}
