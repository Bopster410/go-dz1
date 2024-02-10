package uniq

import (
	"bufio"
	"strings"
)

func uniq(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var outputSlice []string
	var workingLine string = ""
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine != workingLine {
			outputSlice = append(outputSlice, currentLine)
			workingLine = currentLine
		}
	}
	return strings.Join(outputSlice, "\n")
}
