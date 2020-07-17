package main

import "fmt"

/*
*错误处理：
	Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
	panic可以在任何地方引发，但recover只有在defer调用的函数中有效。
	recover()必须搭配defer使用。
	defer一定要在可能引发panic的语句之前定义。
*/

func fA() {
	fmt.Println("A")
}
func fB() {
	defer func() {
		err := recover()
		fmt.Println("in B defer  ", err)
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
	//fmt.Println("end of B")//不会运行至此
}
func fC() {

	fmt.Println("C")
}

func main() {
	fA()
	fB()
	fC()
}
