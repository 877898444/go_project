package main

import (
	"fmt"
	"math/rand"
)

/*
导入
这个代码用圆括号组合了导入，这是“打包”导入语句。

同样可以编写多个导入语句，例如：

import "fmt"
import "math"
不过使用打包的导入语句是更好的形式
 */

func main() {
	fmt.Println("My Favorite number is ",rand.Intn(10))
}
