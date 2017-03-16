package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1,2})
}

/*
结构体
一个结构体（ struct ）就是一个字段的集合。

（而 type 的含义跟其字面意思相符。）
 */