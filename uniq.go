package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Bopster410/go-dz1/uniq"
)

func main() {
	const HELP_MSG string = `Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]
	Filter adjacent matching lines from input_file (or standard input),
	writing to output_file (or standard output).

	-c		prefix lines by the number of occurrences
	-d		only print duplicate lines, one for each group
 	-f=N   	avoid comparing the first N fields
  	-i		ignore differences in case when comparing
  	-s=N    avoid comparing the first N characters
  	-u		only print unique lines
	`

	options := uniq.Options{}
	flag.BoolVar(&options.Count, "c", false, "count repeated lines")
	flag.BoolVar(&options.Repeated, "d", false, "print repeated lines")
	flag.BoolVar(&options.Unique, "u", false, "print unique lines")
	flag.IntVar(&options.SkipFields, "f", 0, "skip fields")
	flag.IntVar(&options.SkipChars, "s", 0, "print unique lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flag.Parse()

	if !uniq.CheckOptions(options) {
		fmt.Println(HELP_MSG)
		return
	}

	args := flag.Args()

	// Default input - stdin
	in := os.Stdin
	if len(args) > 0 {
		var err error
		in, err = os.Open(args[0])
		if err != nil {
			fmt.Printf("An error occurred while opening input file: %q\n", err)
			return
		}
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
			var err error
			out, err = os.OpenFile(args[1], os.O_WRONLY|os.O_CREATE, 0222)
			if err != nil {
				fmt.Printf("An error occurred while opening the output file: %q\n", err)
				return
			}
		}
	}

	// Get output from uniq function
	output, err := uniq.Uniq(text, options)
	if err != nil {
		fmt.Printf("An error occurred while uniq function work: %q\n", err)
		return
	}

	// Write output to output file (or stdout)
	out.Write([]byte(output))
	if len(args) == 2 {
		out.Close()
	}
}
