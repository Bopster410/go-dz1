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

func ParseExpr(expr string) (exprStruct Expr) {
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
	return
}

func (exprStruct Expr) CalcExpr(expr string) int {
	// Regex for numerical expressions
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
