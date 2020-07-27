package main

import (
	"encoding/json"
	"fmt"
)

/*
*struct:
	Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
	Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

*/

type myInt int     //自定义类型
type nikeInt = int //类型别名

type student struct {
	name  string
	age   int
	sex   string
	class string
}

func newStudent(name string, age int, sex string, class string) *student {
	return &student{
		name:  name,
		age:   age,
		sex:   sex,
		class: class,
	}
}

type address struct {
	province string
	city     string
}

type teacher struct {
	name, sex string
	age       int
	addr      address
}

type anonymity struct {
	string
	int
}

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Println("move-------")
}

type dog struct {
	voice string
	animal
}

func (d dog) wangWW() {
	fmt.Println(d.voice)
}

type studentN struct {
	Name  string
	Age   int
	Sex   string `json:"sex" db:"sex" ini:"SEX"`
	class string
}

func f1(s student) {
	s.age = 99
}

func f2(s *student) {
	(*s).age = 99
	s.sex = "未知" //语法糖，二者等价
}

func main() {
	var s1 student
	s1.name = "阿斯顿"
	s1.age = 10
	s1.sex = "男"
	fmt.Println(s1)
	fmt.Println(s1.age)

	//匿名结构体
	var book struct {
		name  string
		price float32
	}
	book.name = "阿萨德"
	book.price = 12.34
	fmt.Println(book)

	f1(s1)
	fmt.Println(s1)
	f2(&s1)
	fmt.Println(s1)

	var s2 = new(student)
	fmt.Println(s2)
	s2.age = 1234
	fmt.Println(s2)

	var s3 = student{
		name:  "qwe",
		sex:   "nan",
		age:   12,
		class: "1ban",
	}
	fmt.Printf("%T\n", s3)
	fmt.Println(s3)
	var s4 = &student{
		"qwe",
		12,
		"nan",
		"1ban",
	}
	fmt.Printf("%T\n", s4)
	fmt.Println(s4)

	s5 := newStudent("zxc", 1234, "asd", "qwe")
	fmt.Println(s5)

	a := anonymity{
		"asd",
		123,
	}
	fmt.Println(a)
	fmt.Println(a.string)

	t1 := teacher{
		"asd",
		"a",
		33,
		address{
			"sd",
			"lc",
		},
	}
	fmt.Println(t1)

	d1 := dog{voice: "wwwwwwwww"}
	fmt.Printf("%T\n", d1)
	d1.move()
	d1.wangWW()

	SN1 := studentN{
		"qwe",
		12,
		"nan",
		"1ban",
	}

	b, err := json.Marshal(SN1)
	if err != nil {
		fmt.Printf("m f ,err: %v\n", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	SNStr := `{"Name":"qwe","Age":12,"sex":"nan"}`
	var SN2 studentN
	json.Unmarshal([]byte(SNStr), &SN2)
	fmt.Println(SN2)
}
