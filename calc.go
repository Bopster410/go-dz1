package main

import (
	"fmt"
	"os"

	"github.com/Bopster410/go-dz1/calc"
)

func getInput() (string, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf("wrong input format")
	}
	return os.Args[1], nil
}

func main() {
	exprStr, err := getInput()
	if err != nil {
		fmt.Printf("An error occurred: %q\n", err)
		return
	}

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
