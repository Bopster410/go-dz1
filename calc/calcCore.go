package calc

import (
	"regexp"
	"strconv"
)

type ExprPart interface {
	CalcExpr() int
}

// ----Num----
type Num struct {
	value int
}

func (num Num) CalcExpr() int {
	return num.value
}

// ----Expr----
type Expr struct {
	left   ExprPart
	right  ExprPart
	action int
}

// Action flags
const (
	NONE = iota
	ADD  = iota
	SUB  = iota
	MUL  = iota
	DIV  = iota
)

var ACTION = map[string]int{
	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": DIV,
}

func ParseExpr(expr string) (exprStruct Expr) {
	re := regexp.MustCompile(`([1-9]|\(.*\))\s*([+\-\/*])\s*([1-9]|\(.*\))`)
	exprStruct.action = NONE
	for _, item := range re.FindStringSubmatch(expr)[1:] {
		_, isAction := ACTION[item]
		if isAction {
			exprStruct.action = ACTION[item]
		} else if exprStruct.action == NONE {
			left, err := strconv.Atoi(item)
			if err == nil {
				exprStruct.left = Num{value: left}
			} else {
				exprStruct.left = ParseExpr(item)
			}
		} else {
			right, err := strconv.Atoi(item)
			if err == nil {
				exprStruct.right = Num{value: right}
			} else {
				exprStruct.right = ParseExpr(item)
			}
		}
	}
	return
}

func (exprStruct Expr) CalcExpr() int {
	// Regex for numerical expressions
	var answer int
	switch exprStruct.action {
	case ADD:
		answer = exprStruct.left.CalcExpr() + exprStruct.right.CalcExpr()
	case SUB:
		answer = exprStruct.left.CalcExpr() - exprStruct.right.CalcExpr()
	case MUL:
		answer = exprStruct.left.CalcExpr() * exprStruct.right.CalcExpr()
	case DIV:
		answer = exprStruct.left.CalcExpr() / exprStruct.right.CalcExpr()
	}

	return answer
}
