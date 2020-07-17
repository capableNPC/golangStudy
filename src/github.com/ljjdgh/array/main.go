package main

import "fmt"

/*
*数组：
	在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
	数组的长度是数组类型的一部分
*/

var a1 [3]int
var a2 [4]int

func main() {
	fmt.Printf("a1:type(%T)\na2:type(%T)\n", a1, a2)
	a1 = [3]int{1, 1, 0}
	fmt.Println(a1)
	fmt.Println(a2)
	a3 := [2]bool{true, false}
	a4 := [...]int{1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}
	fmt.Println(a3)
	fmt.Println(a4)
	a5 := [4]int{1, 2}
	a6 := [4]int{0: 1, 3: 2} //根据索引初始化
	fmt.Println(a5)
	fmt.Println(a6)
	letters1 := [...]string{"a", "b", "c", "d"}
	letters2 := [...]uint8{'a', 'b', 'c', 'd'}
	for i := 0; i < len(letters1); i++ {
		fmt.Print(letters1[i], " ")
	}
	fmt.Println()
	for _, v := range letters2 {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
	var a7 [3][2]int
	len1, len2 := 3, 2
	a7 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a7)
	fmt.Print("[")
	for len1, i := 3, 0; i < 3; i++ {
		fmt.Print("[")
		for len2, j := 2, 0; j < 2; j++ {
			fmt.Print(a7[i][j])
			if j < len2-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("]")
		if i != len1-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("]\n")
	fmt.Print("[")
	for i1, v1 := range a7 {
		fmt.Print("[")
		for i2, v2 := range v1 {
			fmt.Print(v2)
			if i2 < len2-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("]")
		if i1 != len1-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("]\n")
	a8 := a1
	a8[2] = 2
	fmt.Println(a1)
	fmt.Println(a8)

	fmt.Println(a1 == a8)
	fmt.Println(a1 != a8)

	var p1 *[3]int
	var p2 [3]*int
	n1, n2, n3 := 1, 2, 3
	p2[0] = &n1
	p2[1] = &n2
	p2[2] = &n3
	p1 = &a1
	p1[2] = 2
	fmt.Println(a1)
	fmt.Println(p1)
	fmt.Println(&n1, &n2, &n3)
	fmt.Println(p2)
	fmt.Println(*p2[0], *p2[1], *p2[2])
}
