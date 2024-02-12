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

func TestCount(t *testing.T) {
	tests := map[string]string{
		"hello world!":                             "1 hello world!",
		"hello world!\nhello world!":               "2 hello world!",
		"hello world!\nhello world!\nhi":           "2 hello world!\n1 hi",
		"hi\nhello world!\nhello world!\nhi":       "1 hi\n2 hello world!\n1 hi",
		"hi\n\nhello world!\nhello world!\n\n\nhi": "1 hi\n1 \n2 hello world!\n2 \n1 hi",
	}

	for in, correctOut := range tests {
		out := uniq(parseString(in), Options{count: true})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
	}
}

func TestRepeated(t *testing.T) {
	tests := map[string]string{
		"hello world!":                             "",
		"hello world!\nhello world!":               "hello world!",
		"hello world!\nhello world!\nhi":           "hello world!",
		"hi\nhello world!\nhello world!\nhi":       "hello world!",
		"hi\n\nhello world!\nhello world!\n\n\nhi": "hello world!\n",
	}

	for in, correctOut := range tests {
		out := uniq(parseString(in), Options{repeated: true})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
	}
}

func TestUnique(t *testing.T) {
	tests := map[string]string{
		"hello world!":                             "hello world!",
		"hello world!\nhello world!":               "",
		"hello world!\nhello world!\nhi":           "hi",
		"hi\nhello world!\nhello world!\nhi":       "hi\nhi",
		"hi\n\nhello world!\nhello world!\n\n\nhi": "hi\n\nhi",
	}

	for in, correctOut := range tests {
		out := uniq(parseString(in), Options{unique: true})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
	}
}

// func TestCount(t *testing.T) {
// 	tests := map[string]struct {
// 		out		string
// 		options	Options
// 	} {
// 		"hello world!\nhello world!\nhi": {
// 			out: "hello world!",
// 			options: Options{count: true},
// 		},
// 		"hello world!\nhello world!\nhi":           "hello world!\nhi",
// 		"hi\nhello world!\nhello world!\nhi":       "hi\nhello world!\nhi",
// 		"hi\n\nhello world!\nhello world!\n\n\nhi": "hi\n\nhello world!\n\nhi",
// 	}

// 	for in, correctOut := range tests {
// 		out := uniq(parseString(in), Options{skip})
// 		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
// 	}
// }
