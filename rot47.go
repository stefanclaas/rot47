package main

import (
	"bufio"
	"fmt"
	"os"
)

func rot47(input string) string {
	rotated := ""
	for _, char := range input {
		if char >= 33 && char <= 126 {
			char = ((char - 33 + 47) % 94) + 33
		}
		rotated += string(char)
	}
	return rotated
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(rot47(text))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
}
