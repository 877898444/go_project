package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-c:
			fmt.Println("quit")
			return
		}
	}
}
//todo: understand
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

//与上面等价
/*func main() {
	c := make(chan int)
	quit := make(chan int)
	go func(j int) {
		for i := 0; i < j; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}(10)
	fibonacci(c, quit)
}*/
/*
select
select 语句使得一个 goroutine 在多个通讯操作上等待。

select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。
 */