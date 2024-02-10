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

func uniq(input string, options Options) string {
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
