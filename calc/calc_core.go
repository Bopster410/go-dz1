package calc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Represents structures that can be placed into the math expression
type ExprPart interface {
	CalcExpr() (float64, error)
}

// ----Num----
// represents a single number, e.g. 1, 10, 32323
type Num struct {
	value float64
}

func (num Num) CalcExpr() (float64, error) {
	return num.value, nil
}

// ----Expr----
// represents a math expression, e.g. (12 + 33) * 2 - 10
type Expr struct {
	left   ExprPart
	right  ExprPart
	action int
}

const (
	NONE = iota // action is not assigned
	ADD  = iota // '+'
	SUB  = iota // '-'
	MUL  = iota // '*'
	DIV  = iota // '/'
)

// map for convenience
var ACTION = map[string]int{
	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": DIV,
}

func (exprStruct Expr) CalcExpr() (float64, error) {
	var answer float64
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

// ----Other functoins----
// Removes outer parentheses if needed
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

// Parses input string into Expr structure
func ParseExpr(expr string) (ExprPart, error) {
	expr = trimParentheses(expr)
	// search for sum (substraction) parts
	exprParts := regexp.MustCompile(`^(.*\(.*\)[^(]*?|[^)(]+?)\s*([+-])\s*(.*)`).FindStringSubmatch(expr)
	if len(exprParts) == 0 {
		// if nothing was found, search for multiplication (division) parts
		exprParts = regexp.MustCompile(`^(.*?\(.*\)[^(]*?|[^)(]+?)\s*([*/])\s*(.*)`).FindStringSubmatch(expr)
		if len(exprParts) == 0 {
			// if nothing was found convert to int and return Num
			num, err := strconv.ParseFloat(expr, 32)
			if err == nil {
				return Num{value: num}, nil
			}
		}
	}

	// parse found parts into the Expr
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
