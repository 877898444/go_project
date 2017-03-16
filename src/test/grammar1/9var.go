package main

import "fmt"

/*
初始化变量
变量定义可以包含初始值，每个变量对应一个。

如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
 */
var c,python,java bool

func main() {
	var i int
	fmt.Println(i,c,python,java)
}
