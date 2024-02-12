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

func processLine(line string, counter int, options Options) (processed string, err error) {
	// If -d flag or -u flag
	if (options.repeated && counter > 1 || !options.repeated) && (options.unique && counter == 1 || !options.unique) {
		processed = line
		// If -c flag
		if options.count {
			processed = fmt.Sprintf("%d %v", counter, line)
		}
	} else {
		err = fmt.Errorf("current line doesn't meet the requirements")
	}

	return
}

func uniq(input []string, options Options) string {
	var outputSlice []string
	var prevLine string = input[0]
	var counter int = 0

	for _, currentLine := range input {
		if currentLine != prevLine {
			processed, err := processLine(prevLine, counter, options)
			if err == nil {
				outputSlice = append(outputSlice, processed)
			}

			prevLine = currentLine
			counter = 1
		} else {
			counter++
		}
	}

	processed, err := processLine(prevLine, counter, options)
	if err == nil {
		outputSlice = append(outputSlice, processed)
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
