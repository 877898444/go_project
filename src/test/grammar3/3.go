package main

import "fmt"

type Vertex struct {
	X int  //X是大写
	y int //注意：y是小写
}

func main() {
	v:=Vertex{1,2}
	v.X=4
	v.y=6
	fmt.Println(v.X)
	fmt.Println(v.y)
}

/*
结构体字段
结构体字段使用点号来访问。
 */
