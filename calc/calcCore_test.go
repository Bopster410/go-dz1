package calc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumBasic(t *testing.T) {
	tests := map[string]int{
		"1 + 1":     2,
		"1 + 2":     3,
		"1    +  2": 3,
		"1+2":       3,
	}
	for in, out := range tests {
		expr := ParseExpr(in)
		require.Equal(t, out, expr.CalcExpr())
	}
}

func TestSubstractBasic(t *testing.T) {
	tests := map[string]int{
		"1 - 1":   0,
		"2 - 1":   1,
		"2-1":     1,
		"2   - 1": 1,
	}
	for in, out := range tests {
		expr := ParseExpr(in)
		require.Equal(t, out, expr.CalcExpr())
	}
}

func TestMultiplyBasic(t *testing.T) {
	tests := map[string]int{
		"1 * 1":    1,
		"2 * 3":    6,
		"2*3":      6,
		"2   *  3": 6,
	}
	for in, out := range tests {
		expr := ParseExpr(in)
		require.Equal(t, out, expr.CalcExpr())
	}
}

func TestDivisionBasic(t *testing.T) {
	tests := map[string]int{
		"1 / 1":    1,
		"6 / 2":    3,
		"6/2":      3,
		"6   /  2": 3,
	}
	for in, out := range tests {
		expr := ParseExpr(in)
		require.Equal(t, out, expr.CalcExpr())
	}
}

func TestComplex(t *testing.T) {
	tests := map[string]int{
		"(1 + 1) * 2": 4,
		"2 * (1 + 1)": 4,
	}
	for in, out := range tests {
		expr := ParseExpr(in)
		require.Equal(t, out, expr.CalcExpr())
	}
}
