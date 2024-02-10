package uniq

import (
	"bufio"
	"strings"
)

type Options struct {
	count      bool // -c
	repeated   bool // -d
	unique     bool // -u
	skipFields int  // -f
	skipChars  int  // -s
	ignoreCase bool // -i
}

func uniq(input []string, options Options) string {
	var outputSlice []string
	var workingLine string = ""
	for _, currentLine := range input {
		if currentLine != workingLine {
			outputSlice = append(outputSlice, currentLine)
			workingLine = currentLine
		}
	}
	return strings.Join(outputSlice, "\n")
}

func parseString(input string) (output []string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		currentLine := scanner.Text()
		output = append(output, currentLine)
	}
	return
}
