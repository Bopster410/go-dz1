package main

import (
	"fmt"
	"os"

	"github.com/Bopster410/go-dz1/calc"
)

func main() {
	exprStr := os.Args[1]
	expr, parseErr := calc.ParseExpr(exprStr)
	if parseErr != nil {
		fmt.Printf("An error occurred while parsing the expression: %q", parseErr)
	} else {
		val, calcErr := expr.CalcExpr()
		if calcErr != nil {
			fmt.Printf("An error occurred during the calculation: %q", calcErr)
		} else {
			fmt.Printf("%d", val)
		}
	}

}
