package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

/*
*map:
	Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。
	map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	map的定义语法如下：
		map[KeyType]ValueType
		KeyType:表示键的类型。
		ValueType:表示键对应的值的类型。
	map的初始类型为nil，需要使用make()函数来分配内存:
		make(map[KeyType]ValueType, [cap])
		其中cap表示map的容量，该参数虽然不是必须的，
		但是我们应该在初始化map的时候就为其指定一个合适的容量。
	使用delete()函数删除键值对：delete(map, key)
	遍历：
		遍历map时的元素顺序与添加键值对的顺序无关
*/

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["张三"] = 10
	m1["lisi"] = 20
	fmt.Println(m1)
	fmt.Println(m1["lisi"])
	fmt.Println(m1["abc"]) //如果key不存在，得到的数据为对应值类型的零值
	v, ok := m1["abc"]
	if !ok {
		fmt.Println("key不存在")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	delete(m1, "张三")
	fmt.Println(m1)
	delete(m1, "abc") //如果key不存在不做任何反应
	fmt.Println(m1)

	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	//元素为map的切片
	var s1 = make([]map[string]int, 3, 10)
	fmt.Println(s1)
	s1[0] = make(map[string]int, 3)
	s1[0]["张三"] = 11
	s1[0]["李四"] = 22
	fmt.Println(s1)

	//值为切片的map
	var m2 = make(map[string][]int, 10)
	m2["as"] = []int{1, 2, 3, 4}
	fmt.Println(m2)

	//统计句子中的单词出现次数
	str := "As the New Year has passed away, I think the cold weather will pass away and gets warm soon, but I am wrong. This morning, it snows again and I have to wear a lot of clothes. I guess this is the last snow and the summer is coming soon. I miss summer. I can swim and play with my friends in summer."
	str = strings.ToLower(str)
	str = strings.Join(strings.Split(str, ","), " ")
	str = strings.Join(strings.Split(str, "."), " ")
	strSlice := strings.Fields(str)
	var strMap = make(map[string]int, 200)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)

	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	s = append(s, 4)
	ptr0 := &s[0]
	ptr1 := &s[1]
	ptr2 := &s[2]
	ptr3 := &s[3]
	fmt.Println(">???>>>>>  ", *ptr0, *ptr1, *ptr2, *ptr3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[3:]...)

	fmt.Println(">???>>>>>  ", *ptr0, *ptr1, *ptr2, *ptr3)
	fmt.Printf("%+v\n", s)
	fmt.Println(len(s), cap(s))
	fmt.Printf("%+v\n", m["q1mi"])
	fmt.Println(len(m["q1mi"]), cap(m["q1mi"]))
}
