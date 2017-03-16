package main

import "fmt"

/*
数值常量
数值常量是高精度的 值 。

一个未指定类型的常量由上下文来决定其类型。

也尝试一下输出 needInt(Big) 吧。

（int 可以存放最大64位的整数，根据平台不同有时会更少。）
 */
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//fmt.Println(Big,Small)
	fmt.Println(needInt(Small)) //21
	fmt.Println(needFloat(Small)) //0.2
	fmt.Println(needFloat(Big)) //1.2676506002282295e+29

}
