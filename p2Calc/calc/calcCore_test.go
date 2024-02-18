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
	}
	for in, out := range tests {
		require.Equal(t, out, CalcExpr(in))
	}
}

func TestSubstractBasic(t *testing.T) {
	tests := map[string]int{
		"1 - 1": 2,
		"2 - 1": 3,
	}
	for in, out := range tests {
		require.Equal(t, out, CalcExpr(in))
	}
}

func TestMultiplyBasic(t *testing.T) {
	tests := map[string]int{
		"1 * 1": 1,
		"2 * 3": 6,
	}
	for in, out := range tests {
		require.Equal(t, out, CalcExpr(in))
	}
}

func TestDivisionBasic(t *testing.T) {
	tests := map[string]int{
		"1 / 1": 1,
		"6 / 2": 3,
	}
	for in, out := range tests {
		require.Equal(t, out, CalcExpr(in))
	}
}
