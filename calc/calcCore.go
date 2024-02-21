package calc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ExprPart interface {
	CalcExpr() (int, error)
}

// ----Num----
type Num struct {
	value int
}

func (num Num) CalcExpr() (int, error) {
	return num.value, nil
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

func ParseExpr(expr string) (ExprPart, error) {
	expr = trimParentheses(expr)
	exprParts := regexp.MustCompile(`^(.*\(.*\)[^(]*?|[^)(]+?)\s*([+-])\s*(.*)`).FindStringSubmatch(expr)
	if len(exprParts) == 0 {
		exprParts = regexp.MustCompile(`^(.*?\(.*\)[^(]*?|[^)(]+?)\s*([*/])\s*(.*)`).FindStringSubmatch(expr)
		if len(exprParts) == 0 {
			num, err := strconv.Atoi(expr)
			if err == nil {
				return Num{value: num}, nil
			}
		}
	}

	if len(exprParts) > 0 {
		var exprStruct Expr
		exprStruct.left, _ = ParseExpr(exprParts[1])
		exprStruct.action = ACTION[exprParts[2]]
		exprStruct.right, _ = ParseExpr(exprParts[3])
		if exprStruct.left != nil && exprStruct.right != nil {
			return exprStruct, nil
		}
	}

	return nil, fmt.Errorf("wrong expression syntax")
}

func (exprStruct Expr) CalcExpr() (int, error) {
	// Regex for numerical expressions
	var answer int
	leftVal, leftErr := exprStruct.left.CalcExpr()
	rightVal, rightErr := exprStruct.right.CalcExpr()
	if leftErr != nil {
		return 0, leftErr
	}
	if rightErr != nil {
		return 0, rightErr
	}

	switch exprStruct.action {
	case ADD:
		answer = leftVal + rightVal
	case SUB:
		answer = leftVal - rightVal
	case MUL:
		answer = leftVal * rightVal
	case DIV:
		if rightVal == 0 {
			return 0, fmt.Errorf("zero division occurred")
		}
		answer = leftVal / rightVal
	}

	return answer, nil
}
