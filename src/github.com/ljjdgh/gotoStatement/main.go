package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; ; j++ {
			if j > i {
				goto level1
			}
			fmt.Printf("%d*%d=%2d ", j, i, i*j)
		}
	level1:
		fmt.Println()
	}
}
