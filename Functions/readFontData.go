package functions


import (
	"os" // This line imports the "os" package, which provides functions for interacting with the operating system (like opening files)
	"bufio" // This line imports the "bufio" package, which provides functions for reading data from readers in a more efficient way
)

// This function reads the font data from a file and stores it in a map
func ReadFontData(filePath string) (map[int][]string, error) {
	fontData := map[int][]string{} 
	file, err := os.Open(filePath) 
	if err != nil { 
		return nil, err 
	}
	defer file.Close() 
	scanner := bufio.NewScanner(file) 
	scanner.Split(bufio.ScanLines) 

	currentChar := 31 

	for scanner.Scan() { 
		line := scanner.Text() 
		if line == "" { 
			currentChar++ 
		} else {
			fontData[currentChar] = append(fontData[currentChar], line) 
		}
	}

	return fontData, nil 
}
