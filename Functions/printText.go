package functions

import (
	"fmt"
	"os"
	"strings"
)

func PrintText(text string, fontData map[int][]string) {

	// to validate the text.
	Validation(text)
	text = strings.ReplaceAll(text, "\\n", "\n") // remove newlines)

	lines := strings.Split(text, "\n")
	prev := ""
	// Printing a single line
	for u, line := range lines {
		if line == "" {
			if prev != "" || u < len(lines)-1 {
				fmt.Println() // newline after each row
			}
			continue
		}
		for height := 0; height < 8; height++ {
			for _, char := range line {
				fmt.Print(fontData[int(char)][height])
			}
			if height < 7 {
				fmt.Println() // newline after each row
			}
		}
		fmt.Println() // newline after each row
		prev = line
	}
}

func Validation(text string) {
	for _, char := range text {
		if char < 32 || char > 126 {
			fmt.Println("Error: text contains invalid characters.")
			os.Exit(0)
		}
	}
}
