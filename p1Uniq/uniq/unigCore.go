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

func Uniq(input []string, options Options) (string, error) {
	if options.Unique && options.Repeated || options.Unique && options.Count || options.Repeated && options.Unique {
		return "", fmt.Errorf("-u, -d, -c flags can't be used simultaneously")
	}

	var outputSlice []string
	var prevLine string = input[0]
	var prevPartToCompare = prevLine
	var Counter int = 0

	for i, currentLine := range input {
		var partToCompare string = currentLine
		if options.SkipFields > 0 || options.SkipChars > 0 || options.IgnoreCase {
			partToCompare = getPartToCompare(currentLine, options)
			if i == 0 {
				prevPartToCompare = partToCompare
			}
		}

		if partToCompare != prevPartToCompare {
			processed, err := processLine(prevLine, Counter, options)
			if err == nil {
				outputSlice = append(outputSlice, processed)
			}

			prevLine = currentLine
			prevPartToCompare = partToCompare
			Counter = 1
		} else {
			Counter++
		}
	}

	processed, err := processLine(prevLine, Counter, options)
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
