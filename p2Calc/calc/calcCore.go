package calc

import (
	"bufio"
	"strconv"
	"strings"
)

func CalcExpr(expr string) int {
	scanner := bufio.NewScanner(strings.NewReader(expr))
	scanner.Split(bufio.ScanWords)
	var action string
	var left, right int
	for scanner.Scan() {
		item := scanner.Text()
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
