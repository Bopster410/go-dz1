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

func uniq(input []string, options Options) (string, error) {
	if options.unique && options.repeated || options.unique && options.count || options.repeated && options.unique {
		return "", fmt.Errorf("-u, -d, -c flags can't be used simultaneously")
	}

	var outputSlice []string
	var prevLine string = input[0]
	var prevPartToCompare = prevLine
	var counter int = 0

	for i, currentLine := range input {
		var partToCompare string = currentLine
		if options.skipFields > 0 {
			skippedFields := 0
			wasSkipped := false
			indToSlice := 0
			for ind, symb := range currentLine {
				if symb == ' ' && !wasSkipped {
					wasSkipped = true
				} else if symb != ' ' && wasSkipped {
					skippedFields++
					wasSkipped = false
					indToSlice = ind
				}

				if skippedFields == options.skipFields {
					break
				}
			}
			partToCompare = currentLine[indToSlice:]
			if i == 0 {
				prevPartToCompare = partToCompare
			}
		}

		if partToCompare != prevPartToCompare {
			processed, err := processLine(prevLine, counter, options)
			if err == nil {
				outputSlice = append(outputSlice, processed)
			}

			prevLine = currentLine
			prevPartToCompare = partToCompare
			counter = 1
		} else {
			counter++
		}
	}

	processed, err := processLine(prevLine, counter, options)
	if err == nil {
		outputSlice = append(outputSlice, processed)
	}

	return strings.Join(outputSlice, "\n"), nil
}

func parseString(input string) (output []string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		currentLine := scanner.Text()
		output = append(output, currentLine)
	}
	return
}
