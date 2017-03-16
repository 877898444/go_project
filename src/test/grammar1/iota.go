package main

import "fmt"

//const (
//	a = iota
//	b = iota
//	c = iota
//)
//
//const (
//	a = iota
//	b
//	c
//)

func iota_test() {

	const (
		a = iota        // 0
		b                // 1
		c                // 2
		d = "ha"                // 独立值，iota+=1
		e                //"ha" iota+=1
		f = 100                //iota+=1
		g                //100 iota+=1
		h = iota                //7，恢复计数
		i                //8
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
}

//一个有趣的的 iota 实例：

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func fun_example()  {
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}

func main() {
	iota_test()
	fun_example()
	//fun_example:iota表示从0开始自动加1，所以i=1<<0,j=3<<1（<<表示左移的意思），即：i=1,j=6，这没问题，关键在k和l，从输出结果看，k=3<<2，l=3<<3。
}