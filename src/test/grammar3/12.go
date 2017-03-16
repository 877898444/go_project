package main

import "fmt"

func main() {

	var a [] int
	var b [4] int

	fmt.Printf("%T\n", a)  // []int
	fmt.Printf("%T\n", b) // [4]int

	printSlice("a", a)

	//append works in nil slice.
	a = append(a, 0)
	printSlice("a", a)

	//the slice grows as needs
	a = append(a, 1)
	printSlice("a", a)

	//we can add more than one element at a time
	a = append(a, 2, 3, 4)
	printSlice("a", a)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)

}

/*
a len=0 cap=0 []
a len=1 cap=1 [0]
a len=2 cap=2 [0 1]
a len=5 cap=6 [0 1 2 3 4]
 */


/*
向 slice 添加元素
向 slice 的末尾添加元素是一种常见的操作，因此 Go 提供了一个内建函数 append 。 内建函数的文档对 append 有详细介绍。

func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个元素类型为 T 的 slice ，其余类型为 T 的值将会附加到该 slice 的末尾。

append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。

如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。

（了解更多关于 slice 的内容，参阅文章Go 切片：用法和本质。）
 */