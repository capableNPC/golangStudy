package main

import "fmt"

func main() {
	var n = 100
	fmt.Printf("%T,%v\n", n, n)
	fmt.Printf("%b,%o,%d,%x\n", n, n, n, n)
	str := "asdaaa"
	fmt.Printf("%s\n", str)
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
}
