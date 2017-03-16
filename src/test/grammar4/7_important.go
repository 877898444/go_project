package main

import (
	"time"
	"fmt"
)




// 深刻理解go里的继承

//个人理解：要返回error，就要实现error，要想实现error，就得继承实现Error，这时要自定义一个MyError，然后实现Error


type MyError struct {
	When time.Time
	What string
}

//todo:区别见11——struct.go!!!!
//这两者的区别在于：它们是不同的类型实现的方法
func (e *MyError)Error() string { //todo：加上* 与不加，结果相同
	return fmt.Sprintf("When: %v ,%v  IS WRONG !", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "It did not work!",}
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
/*
错误
Go 程序使用 error 值来表示错误状态。

与 fmt.Stringer 类似， error 类型是一个内建接口：

type error interface {
    Error() string
}
（与 fmt.Stringer 类似，fmt 包在输出时也会试图匹配 error。）

通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil， 来进行错误处理。

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
error 为 nil 时表示成功；非 nil 的 error 表示错误。
 */