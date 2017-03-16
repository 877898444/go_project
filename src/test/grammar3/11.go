package main

import "fmt"

func main() {
	var z[]int
	fmt.Println(z, len(z), cap(z)) // [] 0 0

	if z == nil {
		fmt.Println("nil !") // nil !
	}
}

/*
nil slice
slice 的零值是 nil 。

一个 nil 的 slice 的长度和容量是 0。
 */