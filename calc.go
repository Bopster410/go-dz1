package main

import (
	"fmt"
	"os"

	"github.com/Bopster410/go-dz1/calc"
)

func main() {
	exprStr := os.Args[1]
	expr := calc.ParseExpr(exprStr)
	fmt.Printf("%d", expr.CalcExpr())
}
