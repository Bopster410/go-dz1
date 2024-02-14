package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Bopster410/go-dz1/uniq"
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

	var text []string
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		text = append(text, in.Text())
	}

	output, _ := uniq.Uniq(text, options)
	fmt.Println(output)
}
