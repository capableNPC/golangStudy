package main

import (
	"fmt"

	mypkg "github.com/ljjdgh/pkg" //包的别名
	//_ "github.com/ljjdgh/pkg"//匿名包
)

func main() {
	fmt.Println("------------package------------")
	mypkg.ReadAndChge()
	mypkg.ReadAndChge()
}
