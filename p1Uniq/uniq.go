package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/Bopster410/go-dz1/p1Uniq/uniq"
)

func main() {
	options := uniq.Options{}
	flag.BoolVar(&options.Count, "c", false, "count repeated lines")
	flag.BoolVar(&options.Repeated, "d", false, "print repeated lines")
	flag.BoolVar(&options.Unique, "u", false, "print unique lines")
	flag.IntVar(&options.SkipFields, "f", 0, "skip fields")
	flag.IntVar(&options.SkipChars, "s", 0, "print unique lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flag.Parse()

	args := flag.Args()

	// Default input - stdin
	in := os.Stdin
	if len(args) > 0 {
		in, _ = os.Open(args[0])
	}
	// Scan input file (or stdin)
	inScanner := bufio.NewScanner(in)
	var text []string
	for inScanner.Scan() {
		text = append(text, inScanner.Text())
	}

	// Default output - stdout
	out := os.Stdout
	if len(args) > 0 {
		in.Close()
		if len(args) == 2 {
			out, _ = os.OpenFile(args[1], os.O_WRONLY|os.O_CREATE, 0222)
		}
	}

	// Get output from uniq function
	output, _ := uniq.Uniq(text, options)
	// Write output to output file (or stdout)
	out.Write([]byte(output))
	if len(args) == 2 {
		out.Close()
	}
}
