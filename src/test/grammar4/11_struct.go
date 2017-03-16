package main

import "fmt"

type User struct {
	username string
	password string
	age int32
}

func (this User)Set() {
	this.username="set name"
}

func (this *User)Set2() {
	this.username="set name"
}
// 一定要看！ todo:http://studygolang.com/topics/1259
//这两者的区别在于：它们是不同的类型实现的方法
func main() {
	user:=&User{username:"a",password:"123456",age:25}
	user.Set()
	fmt.Println(user)

	user2:=&User{username:"a",password:"123456",age:25}
	user2.Set2()
	fmt.Println(user2)
}

