package calc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrongInput(t *testing.T) {
	tests := map[string]error{
		"(1 + 1 * 2": fmt.Errorf("wrong expression syntax"),
		"2 * 1 + 1)": fmt.Errorf("wrong expression syntax"),
		"aboba15":    fmt.Errorf("wrong expression syntax"),
	}
	for in, outError := range tests {
		in := in
		outError := outError
		t.Run(fmt.Sprintf("Test %q", in), func(t *testing.T) {
			t.Parallel()
			expr, err := ParseExpr(in)
			require.Equal(t, outError, err)
			require.Nil(t, expr)
		})
	}
}

func TestCorrectInput(t *testing.T) {
	tests := []string{
		"(1 + 1) * 2",
		"(1+1)*2",
		"2 * (1 + 1)",
		"2*(1+1)",
		"1 + 2 + 3",
		"1+2+3",
	}
	for _, in := range tests {
		in := in
		t.Run(fmt.Sprintf("Test %q", in), func(t *testing.T) {
			t.Parallel()
			expr, err := ParseExpr(in)
			require.Nil(t, err)
			require.NotNil(t, expr)
		})
	}
}

func TestSumBasic(t *testing.T) {
	tests := map[string]float64{
		"1 + 1":     2,
		"1 + 2":     3,
		"1    +  2": 3,
		"1+2":       3,
	}
	for in, out := range tests {
		in := in
		out := out
		t.Run(fmt.Sprintf("Test %q", in), func(t *testing.T) {
			t.Parallel()
			expr, _ := ParseExpr(in)
			val, err := expr.CalcExpr()
			require.Equal(t, out, val)
			require.Nil(t, err)
		})
	}
}

func TestSubstractBasic(t *testing.T) {
	tests := map[string]float64{
		"1 - 1":   0,
		"2 - 1":   1,
		"2-1":     1,
		"2   - 1": 1,
	}
	for in, out := range tests {
		in := in
		out := out
		t.Run(fmt.Sprintf("Test %q", in), func(t *testing.T) {
			t.Parallel()
			expr, _ := ParseExpr(in)
			val, err := expr.CalcExpr()
			require.Equal(t, out, val)
			require.Nil(t, err)
		})
	}
}

func TestMultiplyBasic(t *testing.T) {
	tests := map[string]float64{
		"1 * 1":    1,
		"2 * 3":    6,
		"2*3":      6,
		"2   *  3": 6,
	}
	for in, out := range tests {
		in := in
		out := out
		t.Run(fmt.Sprintf("Test %q", in), func(t *testing.T) {
			t.Parallel()
			expr, _ := ParseExpr(in)
			val, err := expr.CalcExpr()
			require.Equal(t, out, val)
			require.Nil(t, err)
		})
	}
}

func TestDivisionBasic(t *testing.T) {
	tests := map[string]float64{
		"1 / 1":    1,
		"6 / 2":    3,
		"6/2":      3,
		"6   /  2": 3,
	}
	for in, out := range tests {
		in := in
		out := out
		t.Run(fmt.Sprintf("Test %q", in), func(t *testing.T) {
			t.Parallel()
			expr, _ := ParseExpr(in)
			val, err := expr.CalcExpr()
			require.Equal(t, out, val)
			require.Nil(t, err)
		})
	}
}

func TestComplex(t *testing.T) {
	tests := map[string]float64{
		"(1 + 1) * 2":     4,
		"2 * (1 + 1)":     4,
		"1 + 3 + 8 - 4":   8,
		"1 + 3 + 8 * 4":   36,
		"(1 + 3) * 8 - 4": 28,
	}
	for in, out := range tests {
		in := in
		out := out
		t.Run(fmt.Sprintf("Test %q", in), func(t *testing.T) {
			t.Parallel()
			expr, _ := ParseExpr(in)
			val, err := expr.CalcExpr()
			require.Equal(t, out, val)
			require.Nil(t, err)
		})
	}
}

func TestZeroDivision(t *testing.T) {
	t.Parallel()
	in := "1 / 0"
	expr, _ := ParseExpr(in)
	_, calcErr := expr.CalcExpr()
	require.Equal(t, fmt.Errorf("zero division occurred"), calcErr)
}
