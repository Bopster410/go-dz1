package uniq

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"
)

type Options struct {
	Count      bool // -c
	Repeated   bool // -d
	Unique     bool // -u
	SkipFields int  // -f
	SkipChars  int  // -s
	IgnoreCase bool // -i
}

// Formats given line in accordance with options
func processLine(line string, Counter int, options Options) (processed string, err error) {
	// If -d flag or -u flag
	if (options.Repeated && Counter > 1 || !options.Repeated) && (options.Unique && Counter == 1 || !options.Unique) {
		processed = line
		// If -c flag
		if options.Count {
			processed = fmt.Sprintf("%d %v", Counter, line)
		}
	} else {
		err = fmt.Errorf("current line doesn't meet the requirements")
	}

	return
}

// Returns part which then will be used to compare lines (in accordance with options)
func getPartToCompare(currentLine string, options Options) (partToCompare string) {
	skippedFields := 0
	wasSkipped := false
	indToSlice := 0
	for ind, symb := range currentLine {
		if symb == ' ' && !wasSkipped {
			wasSkipped = true
			skippedFields++
			indToSlice = ind
		} else if symb != ' ' && wasSkipped {
			wasSkipped = false
		}

		if skippedFields == options.SkipFields {
			break
		}
	}

	partToCompare = currentLine[indToSlice:]
	if options.SkipChars >= utf8.RuneCountInString(partToCompare) {
		partToCompare = ""
	} else {
		partToCompare = string([]rune(partToCompare)[options.SkipChars:])
		if options.IgnoreCase {
			partToCompare = strings.ToLower(partToCompare)
		}
	}

	return
}

// Unig main function
func Uniq(input []string, options Options) (string, error) {
	// checks options
	if options.Unique && options.Repeated || options.Unique && options.Count || options.Repeated && options.Unique {
		return "", fmt.Errorf("-u, -d, -c flags can't be used simultaneously")
	}

	// initial values
	var outputSlice []string
	var prevLine string = input[0]
	var prevPartToCompare = prevLine
	var counter int = 0

	// iterate through all lines
	for i, currentLine := range input {
		// get part to use to compare lines
		var partToCompare string = currentLine
		if options.SkipFields > 0 || options.SkipChars > 0 || options.IgnoreCase {
			partToCompare = getPartToCompare(currentLine, options)
			if i == 0 {
				prevPartToCompare = partToCompare
			}
		}

		// line is unique
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

// Separates lines
func parseString(input string) (output []string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		currentLine := scanner.Text()
		output = append(output, currentLine)
	}
	return
}
