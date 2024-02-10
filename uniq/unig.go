package uniq

import (
	"bufio"
	"strings"
)

func uniq(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var output string
	seen := map[string]bool{}
	for scanner.Scan() {
		currentLine := scanner.Text()
		if _, exists := seen[currentLine]; !exists {
			if len(seen) == 0 {
				output += currentLine
			} else {
				output += "\n" + currentLine
			}
			seen[currentLine] = true
		}
	}
	return output
}
