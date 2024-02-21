package calc

import (
	"regexp"
	"strconv"
	"strings"
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

func trimParentheses(line string) string {
	if !strings.HasPrefix(line, "(") || !strings.HasSuffix(line, ")") {
		return line
	}

	var open, close int
	for ind, symb := range line {
		if symb == '(' {
			open++
		} else if symb == ')' {
			close++
		}

		if open == close {
			if ind != len(line)-1 {
				return line
			}
			return line[1:ind]
		}
	}
	return line
}

func ParseExpr(expr string) ExprPart {
	expr = trimParentheses(expr)
	exprParts := regexp.MustCompile(`^(.*\(.*\)[^(]*?|[^)(]+?)\s*([+-])\s*(.*)`).FindStringSubmatch(expr)
	if len(exprParts) == 0 {
		exprParts = regexp.MustCompile(`^(.*?\(.*\)[^(]*?|[^)(]+?)\s*([*/])\s*(.*)`).FindStringSubmatch(expr)
		if len(exprParts) == 0 {
			num, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(expr))
			if err == nil {
				return Num{value: num}
			}
		}
	}

	if len(exprParts) > 0 {
		var exprStruct Expr
		exprStruct.left = ParseExpr(exprParts[1])
		exprStruct.action = ACTION[exprParts[2]]
		exprStruct.right = ParseExpr(exprParts[3])
		return exprStruct
	}

	return nil
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
