package main

import "fmt"

func main() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum) //128
}

/*
for 是 Go 的 “while”
基于此可以省略分号：C 的 while 在 Go 中叫做 for 。
 */