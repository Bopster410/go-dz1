package calc

import (
	"regexp"
	"strconv"
)

type Expr struct {
	left   int
	right  int
	action rune
}

func CalcExpr(expr string) int {
	var exprStruct Expr
	// Regex for numerical expressions
	re := regexp.MustCompile(`([1-9])\s*([+\-\/*])\s*([1-9])`)
	for _, item := range re.FindStringSubmatch(expr)[1:] {
		if item == "+" || item == "-" || item == "*" || item == "/" {
			exprStruct.action = rune(item[0])
		} else if exprStruct.action == '\x00' {
			exprStruct.left, _ = strconv.Atoi(item)
		} else {
			exprStruct.right, _ = strconv.Atoi(item)
		}
	}

	var answer int
	switch exprStruct.action {
	case '+':
		answer = exprStruct.left + exprStruct.right
	case '-':
		answer = exprStruct.left - exprStruct.right
	case '*':
		answer = exprStruct.left * exprStruct.right
	case '/':
		answer = exprStruct.left / exprStruct.right
	}

	return answer
}
