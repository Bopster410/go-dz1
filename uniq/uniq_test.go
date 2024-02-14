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
		out, err := Uniq(parseString(in), Options{})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
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
		out, err := Uniq(parseString(in), Options{count: true})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
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
		out, err := Uniq(parseString(in), Options{repeated: true})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
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
		out, err := Uniq(parseString(in), Options{unique: true})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}

func TestNoParallelFlags(t *testing.T) {
	tests := []string{
		"hello world!",
		"hello world!\nhello world!",
		"hello world!\nhello world!\nhi",
		"hi\nhello world!\nhello world!\nhi",
		"hi\n\nhello world!\nhello world!\n\n\nhi",
	}

	for _, in := range tests {
		out, err := Uniq(parseString(in), Options{unique: true, repeated: true, count: true})
		require.Equalf(t, "", out, "Strings don't match!\ntest string - '%v'", in)
		require.NotEqualf(t, err, nil, "Must be an error!\ntest string - '%v'", in)
	}
}

func TestSkipFields(t *testing.T) {
	tests := map[string]struct {
		out     string
		options Options
	}{
		"hello world!": {
			out:     "hello world!",
			options: Options{skipFields: 1},
		},
		"hello world!\nhello world!": {
			out:     "hello world!",
			options: Options{skipFields: 1},
		},
		"hello world!\nhello world!\nhi": {
			out:     "hello world!\nhi",
			options: Options{skipFields: 1},
		},
		"hello world!\n\n\nhello world!": {
			out:     "hello world!\n\nhello world!",
			options: Options{skipFields: 1},
		},
		"hi world!\nhello world!\nhello world!": {
			out:     "hi world!",
			options: Options{skipFields: 1},
		},
		"hi world!\nhi\nhello world!\nhello world!": {
			out:     "hi world!\nhi\nhello world!",
			options: Options{skipFields: 1},
		},
		"hi world!\nhello epic world!\nhello world!": {
			out:     "hi world!",
			options: Options{skipFields: 2},
		},
		"hi hello George\namazing work George\nhello world!\nbye bye George": {
			out:     "hi hello George\nhello world!\nbye bye George",
			options: Options{skipFields: 2},
		},
		"Привет мир!\nО мир!": {
			out:     "Привет мир!",
			options: Options{skipFields: 1},
		},
	}

	for in, testParams := range tests {
		out, err := Uniq(parseString(in), testParams.options)
		require.Equalf(t, testParams.out, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}

func TestSkipFieldsCount(t *testing.T) {
	tests := map[string]struct {
		out     string
		options Options
	}{
		"hello world!": {
			out:     "1 hello world!",
			options: Options{count: true, skipFields: 1},
		},
		"hello world!\nhello world!": {
			out:     "2 hello world!",
			options: Options{count: true, skipFields: 1},
		},
		"hello world!\nhello world!\nhi": {
			out:     "2 hello world!\n1 hi",
			options: Options{count: true, skipFields: 1},
		},
		"hello world!\n\n\nhello world!": {
			out:     "1 hello world!\n2 \n1 hello world!",
			options: Options{count: true, skipFields: 1},
		},
		"hi world!\nhello world!\nhello world!": {
			out:     "3 hi world!",
			options: Options{count: true, skipFields: 1},
		},
		"hi world!\nhi\nhello world!\nhello world!": {
			out:     "1 hi world!\n1 hi\n2 hello world!",
			options: Options{count: true, skipFields: 1},
		},
		"hi world!\nhello epic world!\nhello world!": {
			out:     "3 hi world!",
			options: Options{count: true, skipFields: 2},
		},
		"hi hello George\namazing work George\nhello world!\nbye bye George": {
			out:     "2 hi hello George\n1 hello world!\n1 bye bye George",
			options: Options{count: true, skipFields: 2},
		},
		"Привет мир!\nО мир!": {
			out:     "2 Привет мир!",
			options: Options{count: true, skipFields: 1},
		},
	}

	for in, testParams := range tests {
		out, err := Uniq(parseString(in), testParams.options)
		require.Equalf(t, testParams.out, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}

func TestSkipFieldsUnique(t *testing.T) {
	tests := map[string]struct {
		out     string
		options Options
	}{
		"hello world!": {
			out:     "hello world!",
			options: Options{unique: true, skipFields: 1},
		},
		"hello world!\nhello world!": {
			out:     "",
			options: Options{unique: true, skipFields: 1},
		},
		"hello world!\nhello world!\nhi": {
			out:     "hi",
			options: Options{unique: true, skipFields: 1},
		},
		"hello world!\n\n\nhello world!": {
			out:     "hello world!\nhello world!",
			options: Options{unique: true, skipFields: 1},
		},
		"hi world!\nhello world!\nhello world!": {
			out:     "",
			options: Options{unique: true, skipFields: 1},
		},
		"hi world!\nhi\nhello world!\nhello world!": {
			out:     "hi world!\nhi",
			options: Options{unique: true, skipFields: 1},
		},
		"hi world!\nhello epic world!\nhello world!": {
			out:     "",
			options: Options{unique: true, skipFields: 2},
		},
		"hi hello George\namazing work George\nhello world!\nbye bye George": {
			out:     "hello world!\nbye bye George",
			options: Options{unique: true, skipFields: 2},
		},
		"Привет мир!\nО мир!": {
			out:     "",
			options: Options{unique: true, skipFields: 1},
		},
	}

	for in, testParams := range tests {
		out, err := Uniq(parseString(in), testParams.options)
		require.Equalf(t, testParams.out, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}

func TestSkipFieldsRepeated(t *testing.T) {
	tests := map[string]struct {
		out     string
		options Options
	}{
		"hello world!": {
			out:     "",
			options: Options{repeated: true, skipFields: 1},
		},
		"hello world!\nhello world!": {
			out:     "hello world!",
			options: Options{repeated: true, skipFields: 1},
		},
		"hello world!\nhello world!\nhi": {
			out:     "hello world!",
			options: Options{repeated: true, skipFields: 1},
		},
		"hello world!\n\n\nhello world!": {
			out:     "",
			options: Options{repeated: true, skipFields: 1},
		},
		"hi world!\nhello world!\nhello world!": {
			out:     "hi world!",
			options: Options{repeated: true, skipFields: 1},
		},
		"hi world!\nhi\nhello world!\nhello world!": {
			out:     "hello world!",
			options: Options{repeated: true, skipFields: 1},
		},
		"hi world!\nhello epic world!\nhello world!": {
			out:     "hi world!",
			options: Options{repeated: true, skipFields: 2},
		},
		"hi hello George\namazing work George\nhello world!\nbye bye George": {
			out:     "hi hello George",
			options: Options{repeated: true, skipFields: 2},
		},
		"Привет мир!\nО мир!": {
			out:     "Привет мир!",
			options: Options{repeated: true, skipFields: 1},
		},
	}

	for in, testParams := range tests {
		out, err := Uniq(parseString(in), testParams.options)
		require.Equalf(t, testParams.out, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}

func TestSkipChars(t *testing.T) {
	tests := map[string]struct {
		out     string
		options Options
	}{
		"hello world!": {
			out:     "hello world!",
			options: Options{skipChars: 1},
		},
		"hello world!\nhello world!": {
			out:     "hello world!",
			options: Options{skipChars: 1},
		},
		"hello world!\nhello world!\nhi": {
			out:     "hello world!\nhi",
			options: Options{skipChars: 5},
		},
		"hello world!\n\n\nhello world!": {
			out:     "hello world!\n\nhello world!",
			options: Options{skipChars: 3},
		},
		"kingsman\nwingsman": {
			out:     "kingsman",
			options: Options{skipChars: 5},
		},
		"kingsman\nman": {
			out:     "kingsman\nman",
			options: Options{skipChars: 5},
		},
		"kingsman\nhi\nwingsman\nkingsman": {
			out:     "kingsman\nhi\nwingsman",
			options: Options{skipChars: 5},
		},
		"Привет мир!\nО нет, мир!": {
			out:     "Привет мир!",
			options: Options{skipChars: 6},
		},
		"Привет мир!\nмир!": {
			out:     "Привет мир!\nмир!",
			options: Options{skipChars: 6},
		},
	}

	for in, testParams := range tests {
		out, err := Uniq(parseString(in), testParams.options)
		require.Equalf(t, testParams.out, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}

func TestSkipFieldsChars(t *testing.T) {
	tests := map[string]struct {
		out     string
		options Options
	}{
		"hello world!": {
			out:     "hello world!",
			options: Options{skipFields: 1, skipChars: 1},
		},
		"hello world!\nhello world!": {
			out:     "hello world!",
			options: Options{skipFields: 1, skipChars: 2},
		},
		"hello world!\nhello world!\nhi": {
			out:     "hello world!\nhi",
			options: Options{skipFields: 1, skipChars: 5},
		},
		"hello world!\n\n\nhello world!": {
			out:     "hello world!\n\nhello world!",
			options: Options{skipFields: 1, skipChars: 3},
		},
		"hello world!\n\nhello world!": {
			out:     "hello world!",
			options: Options{skipFields: 1, skipChars: 10},
		},
		"hello epicworld!\nhello damnworld!": {
			out:     "hello epicworld!",
			options: Options{skipFields: 1, skipChars: 5},
		},
		"Привет мир!\nмир!": {
			out:     "Привет мир!\nмир!",
			options: Options{skipFields: 1, skipChars: 2},
		},
	}

	for in, testParams := range tests {
		out, err := Uniq(parseString(in), testParams.options)
		require.Equalf(t, testParams.out, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}

func TestIgnoreCase(t *testing.T) {
	tests := map[string]string{
		"hello world!": "hello world!",
		"hello world!\nhelLo WorLd!\nhello world!": "hello world!",
		"hello world!\nhEllO woRlD!\nhi":           "hello world!\nhi",
		"hi\n\n\nHI":                               "hi\n\nHI",
	}

	for in, correctOut := range tests {
		out, err := Uniq(parseString(in), Options{ignoreCase: true})
		require.Equalf(t, correctOut, out, "Strings don't match!\ntest string - '%v'", in)
		require.Equalf(t, err, nil, "Error occurred: %v\ntest string - '%v'", err, in)
	}
}
