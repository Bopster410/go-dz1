package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Bopster410/go-dz1/uniq"
)

func main() {
	var text []string

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		text = append(text, in.Text())
	}
	output, _ := uniq.Uniq(text, uniq.Options{})
	fmt.Println(output)
}
