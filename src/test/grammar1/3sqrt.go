package main

import (
	"fmt"
	"math"
)

/*
导出名
在 Go 中，首字母大写的名称是被导出的。

在导入包之后，你只能访问包所导出的名字，任何未导出的名字是不能被包外的代码访问的。

Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。

执行代码，注意编译器报的错误。

然后将 math.pi 改名为 math.Pi 再试着执行一下。
 */

func main() {
	fmt.Println("Now u have %g problems.",math.Sqrt(7))
}