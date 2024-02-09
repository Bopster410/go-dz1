package main

import "fmt"

func f(s *[]string) {
	*s = append(*s, "o.O")
}

func main() {
	var a int = 10
	fmt.Printf("a is %d\n", a)

	b := 14
	fmt.Printf("b is %d\n", b)

	arr := [...]int{1, 2, 3}
	fmt.Printf("arr is %v\n", arr)

	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6}
	slice := append(slice1, slice2...)
	fmt.Printf("slice is %#v\n", slice)

	l, c := len(slice), cap(slice)
	fmt.Println(l, c)

	strSlice := []string{"hellooo", "friend"}
	fmt.Println(strSlice)
	f(&strSlice)
	fmt.Println(strSlice)
}
