package main

import (
	 "asciiart/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide text to print as an argument.")
		return
	}

	fontData, err := functions.ReadFontData("fonts/standard.txt")
	if err != nil {
		fmt.Println("Error reading font data:", err)
		return
	}

	functions.PrintText(os.Args[1], fontData)
}
