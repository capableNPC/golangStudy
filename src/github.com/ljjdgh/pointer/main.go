package main

import "fmt"

/*
*指针：
	区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
	它的指针地址、指针类型和指针取值这三种性质和C/C++中的指针类似
	Go语言中new和make是内建的两个函数，主要用来分配内存。
	new:
		它的函数签名：func new(Type) *Type
	maek:
		make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
		而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
		因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
		它的函数签名：func make(t Type, size ...IntegerType) Type
*/
func main() {
	var p1 *int
	n1 := 123
	p1 = &n1
	*p1 = 321
	fmt.Println(*p1)
	p2 := new(int)
	fmt.Println(*p2)
	s1 := make([]int, 3, 3)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
}
