package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoDuplicate(t *testing.T) {
	input := "hello world!"
	output := uniq(input)
	correctOutput := "hello world!"
	require.Equalf(t, correctOutput, output, "Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
}

func TestOneDuplicate(t *testing.T) {
	input := "hello world!\nhello world!"
	output := uniq(input)
	correctOutput := "hello world!"
	require.Equalf(t, correctOutput, output, "Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
}

func TestDifferentLines(t *testing.T) {
	input := "hello world!\nhello world!\nhi"
	output := uniq(input)
	correctOutput := "hello world!\nhi"
	require.Equalf(t, correctOutput, output, "Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
}

func TestDeleteIfNear(t *testing.T) {
	input := `hi
hello world!
hello world!
hi`
	output := uniq(input)
	correctOutput := `hi
hello world!
hi`
	require.Equalf(t, correctOutput, output, "Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
}

func TestEmptyLines(t *testing.T) {
	input := `hi

hello world!
hello world!


hi`
	output := uniq(input)
	correctOutput := `hi

hello world!

hi`
	require.Equalf(t, correctOutput, output, "Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
}
