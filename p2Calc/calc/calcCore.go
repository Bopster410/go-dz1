package calc

import (
	"regexp"
	"strconv"
)

type ExprPart interface {
	CalcExpr() int
}

type Num struct {
	value int
}

func (num Num) CalcExpr() int {
	return num.value
}

type Expr struct {
	left   ExprPart
	right  ExprPart
	action rune
}

func ParseExpr(expr string) (exprStruct Expr) {
	re := regexp.MustCompile(`([1-9])\s*([+\-\/*])\s*([1-9])`)
	for _, item := range re.FindStringSubmatch(expr)[1:] {
		if item == "+" || item == "-" || item == "*" || item == "/" {
			exprStruct.action = rune(item[0])
		} else if exprStruct.action == '\x00' {
			left, _ := strconv.Atoi(item)
			exprStruct.left = Num{value: left}
		} else {
			right, _ := strconv.Atoi(item)
			exprStruct.right = Num{value: right}
		}
	}
	return
}

func (exprStruct Expr) CalcExpr() int {
	// Regex for numerical expressions
	var answer int
	switch exprStruct.action {
	case '+':
		answer = exprStruct.left.CalcExpr() + exprStruct.right.CalcExpr()
	case '-':
		answer = exprStruct.left.CalcExpr() - exprStruct.right.CalcExpr()
	case '*':
		answer = exprStruct.left.CalcExpr() * exprStruct.right.CalcExpr()
	case '/':
		answer = exprStruct.left.CalcExpr() / exprStruct.right.CalcExpr()
	}

	return answer
}
