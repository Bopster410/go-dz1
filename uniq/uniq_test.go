package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoFlags(t *testing.T) {
	tests := map[string]string{
		"hello world!":                             "hello world!",
		"hello world!\nhello world!":               "hello world!",
		"hello world!\nhello world!\nhi":           "hello world!\nhi",
		"hi\nhello world!\nhello world!\nhi":       "hi\nhello world!\nhi",
		"hi\n\nhello world!\nhello world!\n\n\nhi": "hi\n\nhello world!\n\nhi",
	}

	for in, correctOut := range tests {
		out := uniq(parseString(in), Options{})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
	}
}
