package main

import "fmt"

//常量声明
const pi = 3.1415926

//批量常量声明
const (
	statusOk = 200
	notFound = 404
)

//批量声明常量时，如果某一行没有赋值，默认和上一行一致
const (
	number1 = 100
	number2
	number3
)

/*	iota是go语言的常量计数器，在const关键字出现时会被重置为0，
*	const中每新增一行常量声明iota自加一，即相当于const语句块的索引
*	它会使得const语句块相当于枚举
 */
const (
	n1 = iota // 0
	n2        // 1
	n3        // 2
)
const (
	a1 = iota // 0
	a2 = iota // 1
	_  = iota // 2
	a3 = iota // 3
)
const (
	b1 = iota // 0
	b2 = 100  // 100
	b3 = iota // 2
)
const (
	c1, c2 = iota, iota + 1 // 0,1
	c3     = iota           // 1
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

func main() {
	fmt.Println("常量声明")
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
