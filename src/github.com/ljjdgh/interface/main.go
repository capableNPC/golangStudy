package main

import "fmt"

/*
*接口：接口（interface）定义了一个对象的行为规范，只定义规范不实现，
	由具体的对象来实现规范的细节。
	在Go语言中接口（interface）是一种类型，一种抽象的类型。

*/
type animal interface {
	speaker
	mover
}

type speaker interface {
	speak()
}
type mover interface {
	move()
}
type dog struct {
}

type emptyI interface{} //空接口

func (d dog) speak() {
	fmt.Println("WWW")
}
func (d dog) move() {
	fmt.Println("-----------------")
}

type cat struct {
}

func (c cat) speak() {
	fmt.Println("MMM")
}
func (c cat) move() {
	fmt.Println("================")
}
func hungry(x speaker) {
	x.speak()
}
func hunter(x mover) {
	x.move()
}
func getFood(x animal) {
	x.speak()
	x.move()
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func anyType(x interface{}) {
	fmt.Printf("%T\n", x)
}

func justifyType1(x interface{}) {
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

func justifyType2(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	d1 := dog{}
	c1 := cat{}
	hungry(d1)
	hunter(d1)
	hungry(c1)
	hunter(c1)

	getFood(d1)
	getFood(c1)

	var xyj haier
	var xyjItfc WashingMachine
	xyjItfc = xyj
	xyjItfc.wash()
	xyjItfc.dry()

	//空接口
	//interface：关键字
	//interface{}：空接口类型
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 10)
	m1["name"] = "zhazhah"
	m1["age"] = 10
	m1["hobby"] = [...]string{"twly", "bzb", "hszb"}
	fmt.Println(m1)

	anyType(m1)
	anyType(xyj)

	//类型断言
	//方法1
	justifyType1(123)
	justifyType1("123")
	//方法2
	justifyType2(123)
	justifyType2("123")
}
