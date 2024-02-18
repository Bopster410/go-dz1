package calc

import (
	"regexp"
	"strconv"
)

func CalcExpr(expr string) int {
	var action string
	var left, right int
	// Regex for numerical expressions
	re := regexp.MustCompile(`([1-9])\s*([+\-\/*])\s*([1-9])`)
	for _, item := range re.FindStringSubmatch(expr)[1:] {
		if item == "+" || item == "-" || item == "*" || item == "/" {
			action = item
		} else if action == "" {
			left, _ = strconv.Atoi(item)
		} else {
			right, _ = strconv.Atoi(item)
		}
	}

	var answer int
	switch action {
	case "+":
		answer = left + right
	case "-":
		answer = left - right
	case "*":
		answer = left * right
	case "/":
		answer = left / right
	}

	return answer
}
