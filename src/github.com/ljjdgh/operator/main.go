package main

import "fmt"

/*
*运算符:
	Go 语言内置的运算符有:
		算术运算符:+、-、*、/、%
		关系运算符：==、!=、>、>=、<、<=
		逻辑运算符：&&、||、!
		位运算符：&、|、^、<<、>>
		赋值运算符：=、+=、-=、*=、/=、%=、<<=、>>=、&=、|=、^=
*
*/
func main() {
	//数组中的数均出现了两次，但有一个只出现一次，找出只出现一次的数字
	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 6, 7, 8, 9}
	repeatedNumber := 0
	for _, number := range array {
		repeatedNumber ^= number
	}
	fmt.Println("重复的数字是： ", repeatedNumber)
}
