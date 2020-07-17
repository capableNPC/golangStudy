package main

import "fmt"

//导入语句

//函数外部只能放置标识符（变量、常量、函数、类型）的声明
/*
*	go语言中的变量名为下划线_和数字以及字母的组合，但第一个字符不能是数字
*	声明变量基本公式： var 变量名 变量类型
 */

//声明变量
var name string // ""
var age int     // 0
var sex string
var state bool // false

//批量声明
var (
	price float64
	date  string
)

func goods() (float64, string) {
	return 1.1, "2020-01-01"
}

//程序的入口函数
func main() {
	fmt.Println("变量声明")
	var number = 1 //类型推导
	//go语言的非全局变量声明后必须要使用，否则编译不能通过
	fmt.Println(number)
	str := "12.1" //简短变量声明，仅在函数内部使用
	fmt.Printf(str)
	totalPrice, _ := goods() //_为匿名变量，它不占用命名空间且不分配内存，表示忽略值
	fmt.Println("totalPrice: ", totalPrice)
	name = "user"
	age = 10
	price = 1.23
	fmt.Println("name: ", name, " age: ", age, " price: ", price)
	fmt.Print(state)
	fmt.Printf("name:%s\n", name)
	fmt.Println()
}
