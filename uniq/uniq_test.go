package uniq

import (
	"testing"
)

func TestNoDuplicate(t *testing.T) {
	input := "hello world!"
	output := uniq(input)
	correctOutput := "hello world!"
	if output != correctOutput {
		t.Fatalf("Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
	}
}

func TestOneDuplicate(t *testing.T) {
	input := "hello world!\nhello world!"
	output := uniq(input)
	correctOutput := "hello world!"
	if output != correctOutput {
		t.Fatalf("Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
	}
}

func TestDifferentLines(t *testing.T) {
	input := "hello world!\nhello world!\nhi"
	output := uniq(input)
	correctOutput := "hello world!\nhi"
	if output != correctOutput {
		t.Fatalf("Strings don't match!\n(correct is '%v', yours is '%v'", correctOutput, output)
	}
}
