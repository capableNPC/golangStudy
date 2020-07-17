package main

import (
	"fmt"
	"strings"
)

/*
*函数：
	go语言函数中没有默认参数的概念
	defer语句:Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
		在defer归属的函数即将返回时，
		将延迟处理的语句按defer定义的逆序进行执行，类似于对语句进行了入栈和出栈的操作
		由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。
		比如：资源清理、文件关闭、解锁及记录时间等。
		  在Go语言的函数中return语句在底层并不是原子操作，
		  它分为给返回值赋值和RET指令两步。
		  而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
	高阶函数：高阶函数分为函数作为参数和函数作为返回值两部分。
	匿名函数:匿名函数就是没有函数名的函数，匿名函数的定义格式如下：
		func(参数)(返回值){
    		函数体
		}
		匿名函数没有函数名，没办法像普通函数那样调用，
		所以匿名函数需要保存到某个变量或者作为立即执行函数
		立即执行函数：匿名函数定义完加(参数)直接执行
			func(参数)(返回值){
				函数体
			}(参数)
		匿名函数多用于实现回调函数和闭包。
	闭包:
		闭包指的是一个函数和与其相关的引用环境组合而成的实体。
		简单来说，闭包=函数+引用环境。



*/

//函数的定义
func add(x int, y int) (result int) { //命名返回值相当于在函数中声明一个变量
	return x + y //a
}
func f1(x int, y int) {
	fmt.Println("f1")
}
func f2() {
	fmt.Println("f2")
}
func f3() int {
	fmt.Println("f3")
	return 0
}
func f4(x int, y int) (rx int, ry int) {
	rx = x
	ry = y
	fmt.Println("f4")
	return
}
func f5(x int, y int) (int, int) {
	fmt.Println("f5")
	return x, y
}
func f6(x, y, z int) (int, int) {
	fmt.Println("f6")
	return x, y
}
func f7(x, y int, s1, s2 string) (int, int) {
	fmt.Println("f7")
	return x, y
}
func f8(x bool, y ...int) {
	fmt.Println("f8")
	fmt.Printf("%T,%v\n", y, y) //切片
}

//定义函数类型
type funcType func(int, bool) int

func f9(x int, y bool) int {
	fmt.Println("f9")
	return 9
}
func f10(x int, y bool) int {
	fmt.Println("f10")
	return 10
}

//高阶函数
//作为参数
func calc(x, y int, argF func(int, int) int) int {
	return argF(x, y)
}

//作为返回值
func f11() func(int, int) int {
	return add
}

//defer
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
}

//闭包
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func calcp(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func dispatchCoin(coins int, nameMap map[string]int) int {
	for k := range nameMap {
		for _, r := range k {
			switch r {
			case 'e', 'E':
				nameMap[k]++
				coins--
			case 'i', 'I':
				nameMap[k] += 2
				coins -= 2
			case 'o', 'O':
				nameMap[k] += 3
				coins -= 3
			case 'u', 'U':
				nameMap[k] += 4
				coins -= 4
			default:
				continue
			}
		}
	}
	return coins
}

func main() {
	f8(true)
	f8(true, 1)
	f8(true, 1, 2)
	var funccc1 funcType //函数类型变量
	funccc1 = f9
	fmt.Println(funccc1(1, true))
	funccc2 := f10
	fmt.Println(funccc2(1, true))

	fmt.Println(calc(0, 1, add))
	fmt.Println(f11()(1, 2))

	deferDemo()

	add := func(x, y int) {
		fmt.Println(x + y)
	} // 将匿名函数保存到变量
	add(10, 10)
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 10) //自执行函数

	//闭包
	var pf1 = adder()
	fmt.Println(pf1(10)) //10
	fmt.Println(pf1(20)) //30
	fmt.Println(pf1(30)) //60

	pf2 := adder()
	fmt.Println(pf2(40)) //40
	fmt.Println(pf2(50)) //90

	var pf3 = adder2(10)
	fmt.Println(pf3(10)) //20
	fmt.Println(pf3(20)) //40
	fmt.Println(pf3(30)) //70

	fp4 := adder2(20)
	fmt.Println(fp4(40)) //60
	fmt.Println(fp4(50)) //110

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	fp5, fp6 := calcp(10)
	fmt.Println(fp5(1), fp6(2)) //11 9
	fmt.Println(fp5(3), fp6(4)) //12 8
	fmt.Println(fp5(5), fp6(6)) //13 7

	/*
	   你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
	   分配规则如下：
	   a. 名字中每包含1个'e'或'E'分1枚金币
	   b. 名字中每包含1个'i'或'I'分2枚金币
	   c. 名字中每包含1个'o'或'O'分3枚金币
	   d: 名字中每包含1个'u'或'U'分4枚金币
	   写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
	   程序结构如下，请实现 ‘dispatchCoin’ 函数
	*/
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int, len(users))
	)
	for _, v := range users {
		distribution[v] = 0
	}
	left := dispatchCoin(coins, distribution)
	fmt.Println(distribution)
	fmt.Println("剩下：", left)
}
