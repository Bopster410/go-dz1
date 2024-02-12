package uniq

import (
	"bufio"
	"fmt"
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
	var prevLine string = input[0]
	var counter int = 0

	for _, currentLine := range input {
		if currentLine != prevLine {
			formattedLine := prevLine
			if options.count {
				formattedLine = fmt.Sprintf("%d %v", counter, prevLine)
			}
			outputSlice = append(outputSlice, formattedLine)
			prevLine = currentLine
			counter = 1
		} else {
			counter++
		}
	}

	formattedLine := prevLine
	if options.count {
		formattedLine = fmt.Sprintf("%d %v", counter, prevLine)
	}

	outputSlice = append(outputSlice, formattedLine)

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
