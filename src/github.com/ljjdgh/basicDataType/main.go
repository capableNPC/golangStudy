package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

//Go 语言的基本类型和其他语言大同小异
/*
*	整型：
		整型分为以下两个大类：
		按长度分为：int8、int16、int32、int64
		对应的无符号整型：uint8、uint16、uint32、uint64
		特殊的有：uint、int、uintptr
		其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，
		int64对应C语言中的long型。
		uint和int在不同操作系统上会根据操作系统的位数改变，32位系统为int32,64位系统为int64
	浮点数：
		Go语言支持两种浮点型数：float32和float64。
		float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
		float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。
		打印浮点数时，可以使用fmt包配合动词%f
	复数：
		复数分为complex64和complex128
		复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
	布尔值：
		Go语言中以bool类型进行声明布尔型数据。
		布尔型数据只有true（真）和false（假）两个值。
	字符串：
		Go语言中的字符串以原生数据类型出现，使用字符串就像使用
		其他原生数据类型（int、bool、float32、float64 等）一样。
		Go 语言里的字符串的内部实现使用UTF-8编码。
		字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符。
		字符串转义符：\r回车符、\n换行符、\t制表符、\'单引号、\"双引号\、反斜杠
		多行字符串：
			Go语言中要定义一个多行字符串时，就必须使用反引号字符。
			反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
		字符串的常用操作
						 方法							介绍
						len(str)					   求长度
					+或fmt.Sprintf					  拼接字符串
					strings.Split						分割
					strings.contains				 判断是否包含
			strings.HasPrefix,strings.HasSuffix		前缀/后缀判断
			strings.Index(),strings.LastIndex()		子串出现的位置
			strings.Join(a[]string, sep string)		   join操作
		byte和rune类型：
			组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。
			字符用单引号（'）包裹起来。
		字符:
			Go 语言的字符有两种：
				uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
				rune类型，代表一个 UTF-8字符。rune类型实际是一个int32
		Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，
		也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
		字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
		字符串是不能修改的,字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
		rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
		修改字符串:
			要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
			无论哪种转换，都会重新分配内存，并复制字节数组。
		类型转换:
			Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
*/

func main() {
	fmt.Println("基本数据类型")
	//布尔值
	var b1 bool //布尔值的初始值为false
	b2 := true
	//不能将整型强制转换为布尔型，且布尔型无法参与数值运算，也无法与其他类型进行转换
	//b3 := bool(1)
	fmt.Printf("%T,%v\n", b1, b1)
	fmt.Printf("%T,%v\n", b2, b2)

	//数字字面量语法
	n1 := 99   //十进制
	n2 := 077  //八进制
	n3 := 0xff //十六进制

	f1 := 1.23
	f2 := math.MaxFloat32
	fmt.Println(f1, f2)
	//go语言中小数默认为float64
	fmt.Printf("%T\n", f1) //查看变量的类型
	f3 := float32(1.2)
	fmt.Printf("%T\n", f3) //查看变量的类型
	//go语言中float32类型的值无法直接赋值给float64
	//f1 = f3

	fmt.Println(n1)
	fmt.Printf("%o\n", n1) //十进制转八进制输出
	fmt.Printf("%x\n", n1) //十进制转十六进制输出
	fmt.Println(n2)
	fmt.Printf("%T\n", n2) //查看变量的类型
	fmt.Println(n3)
	n4 := int8(10)
	fmt.Printf("%T\n", n4) //查看变量的类型

	//字符串
	str1 := "asd这是一个字符串！"
	str2 := `1
		ABCD
		EFGH
    IJKL
		 MNOP
Q
	`
	str3 := `c:\asdwwa\aa`
	c1 := 'a'
	c2 := '啊'
	fmt.Println(str1)
	fmt.Printf("%#v\n", str2)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Printf("%c[%T],%c[%T]\n", c1, c1, c2, c2) //查看变量的类型

	fmt.Println(len(str1))
	fmt.Println(str1 + str3)
	fmt.Println(fmt.Sprintf("%s%s", str1, str3))
	fmt.Println(strings.Split(str3, "\\"))
	fmt.Println(strings.Contains(str3, ":"))
	fmt.Println(strings.HasPrefix(str3, "c:"))
	fmt.Println(strings.HasSuffix(str3, "as"))
	fmt.Println(strings.Index(str3, "a"))
	fmt.Println(strings.LastIndex(str3, "a"))
	fmt.Println(strings.Join(strings.Split(str3, "\\"), "+"))

	for i := 0; i < len(str1); i++ {
		fmt.Println(str1[i])        //只会输出对应的整数值
		fmt.Printf("%c\n", str1[i]) //非ASCII码字符乱码
	}
	for _, c := range str1 {
		fmt.Printf("%c\n", c)
	}
	strE1 := "西瓜"
	strE2 := []rune(strE1)
	strE2[0] = '冬'
	fmt.Println(string(strE2))

	count := 0
	for _, c := range str1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Printf("str1:%#v中的汉字个数为：%d ", str1, count)
}
