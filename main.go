package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	mouthStart = 32
	ascii_File = "rajni_ascii"
)

func rajnisay(message string) {

	asciiArt, err := loadAsciiArtFromFile(ascii_File)
	if err != nil {
		fmt.Println("error occured while getting ascii art")
	}

	// Prepare the speech bubble content
	lines := strings.Split(strings.TrimSpace(message), "\n")
	maxWidth := 24 // Based on placeholder bubble "<                        >"
	for i, line := range lines {
		lines[i] = fmt.Sprintf("           < %-*s >", maxWidth, line)
	}
	for len(lines) < 3 {
		lines = append(lines, fmt.Sprintf("           < %-*s >", maxWidth, ""))
	}

	// Replace the bubble in the ASCII art
	var output []string
	artLines := strings.Split(asciiArt, "\n")
	replaced := 0

	for _, line := range artLines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "<") && replaced < 3 {
			suffix, _ := strings.CutPrefix(trimmed, "<                        >")
			output = append(output, lines[replaced]+suffix)
			replaced++
		} else {
			output = append(output, line)
		}
	}

	// Print the final result
	fmt.Println(strings.Join(output, "\n"))
}

func loadAsciiArtFromFile(filePath string) (string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil

}

func main() {
	if len(os.Args) > 1 {
		rajnisay(os.Args[1])
	} else {
		rajnisay("Chumma adhurudhilla")
	}
}
