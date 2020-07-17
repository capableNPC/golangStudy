package main

import (
	"fmt"
	"sort"
)

/*
*切片：
	切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。
	它非常灵活，支持自动扩容。
	切片是一个引用类型，它的内部结构包含地址、长度和容量。
	切片是引用类型，不支持直接比较，只能和nil比较
	切片一般用于快速地操作一块数据集合。
	切片的长度和容量：
		切片拥有自己的长度和容量。
		可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
		长度：它的元素个数
		容量：它的底层数组从切片的第一个元素到最后一个元素的数量
	切片再切片：切片可以在切片的基础上进行二次切片，但是底层数组不变
	make()函数构造切片：
		动态的创建一个切片需要使用内置的make()函数
		make([]T, size, cap)
	append()添加元素：
		Go语言的内建函数append()可以为切片动态添加元素。
		可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。
		切片的容量扩容:
			首先判断，如果新申请容量大于2倍的旧容量，最终容量就是新申请的容量
			否则判断，如果旧切片的长度小于1024，则最终容量就是旧容量的两倍
			否则判断，如果旧切片长度大于等于1024，则最终容量从旧容量开始循环增加原来的1/4，即直到最终容量大于等于新申请的容量
			如果最终容量计算值溢出，则最终容量就是新申请容量
	copy()复制切片:copy(destSlice, srcSlice []T)
	从切片中删除元素:Go语言中并没有删除切片元素的专用方法.



*/

func main() {
	var s1 []int                     //声明一个存放int类型的切片
	s2 := []int{}                    //声明一个存放int类型的切片并初始化
	var s3 = []int{0, 1, 2, 3, 4, 5} //声明一个存放int类型的切片并初始化
	fmt.Println("s1: ", s1 == nil, s2 == nil, s3 == nil)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))
	fmt.Println(s1, s2, s3)
	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 = a1[0:4] //基于一个数组切割，范围为数组的[0,4),即左闭右开区间
	fmt.Println("s1: ", s1 == nil, len(s1), cap(s1), s1)
	s4 := a1[:4] //[0:4]
	s5 := a1[3:] //[3:len(a1)]
	s6 := a1[:]  //[0:len(a1)]

	fmt.Println("s4: ", s4 == nil, len(s4), cap(s4), s4)
	fmt.Println("s5: ", s5 == nil, len(s5), cap(s5), s5)
	fmt.Println("s6: ", s6 == nil, len(s6), cap(s6), s6)

	s7 := a1[3:6:6] //[3:6],容量为6-3=3
	s8 := a1[3:6]   //[3:6],容量为len(a1)-3=x
	fmt.Println("s7: ", s7 == nil, len(s7), cap(s7), s7)
	fmt.Println("s8: ", s8 == nil, len(s8), cap(s8), s8)
	s9 := s5[3:6]
	fmt.Println("s9: ", s9 == nil, len(s9), cap(s9), s9)
	a1[7] = 77
	fmt.Println("s9: ", s9 == nil, len(s9), cap(s9), s9) //切片只是一个引用
	s9[1] = 7
	fmt.Println(a1)

	s10 := make([]int, 5, 10)
	fmt.Println("s10: ", s10 == nil, len(s10), cap(s10), s10)

	for i := 0; i < len(s3); i++ {
		fmt.Print(s3[i], " ")
	}
	fmt.Println()
	for _, v := range s3 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	//append()方法添加元素
	var s11 []int
	s11 = append(s11, 1) //[1]
	fmt.Println("s11: ", s11 == nil, len(s11), cap(s11), s11)
	s11 = append(s11, 2, 3, 4) //[1,2,3,4]
	fmt.Println("s11: ", s11 == nil, len(s11), cap(s11), s11)
	s12 := []int{5, 6, 7}
	s11 = append(s11, s12...)
	fmt.Println("s11: ", s11 == nil, len(s11), cap(s11), s11) //len=7,cap=8

	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	//copy()复制切片
	s13 := []int{1, 2, 3, 4}
	s14 := s13
	fmt.Println("s13: ", s13 == nil, len(s13), cap(s13), s13)
	fmt.Println("s14: ", s14 == nil, len(s14), cap(s14), s14)
	s14[1] = 0
	fmt.Println("s13: ", s13 == nil, len(s13), cap(s13), s13)
	fmt.Println("s14: ", s14 == nil, len(s14), cap(s14), s14)
	s15 := make([]int, 4, 4)
	copy(s15, s13)
	s15[1] = 2
	fmt.Println("s13: ", s13 == nil, len(s13), cap(s13), s13)
	fmt.Println("s15: ", s15 == nil, len(s15), cap(s15), s15)

	//从切片中删除元素
	s15 = append(s15[:1], s15[2:]...)
	fmt.Println("s15: ", s15 == nil, len(s15), cap(s15), s15)

	//数组的排序
	var a = [...]int{3, 7, 8, 9, 1}
	fmt.Println(a)
	b := a[:]
	sort.Ints(b) //sort.Sort(IntSlice(b))
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(a)
}
