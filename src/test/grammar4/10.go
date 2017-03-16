package main

import (
	"net/http"
	"fmt"
	"time"
	"log"
)

type Hello struct {
	When time.Time
	name string
}

//todo:加上*会报错 (见11.go)
func (h Hello)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"hello:%v,the time is : ",h.name,h.When)
}

func main() {
	var h Hello
	err:=http.ListenAndServe("localhost:4000",h)
	if err != nil {
		log.Fatal(err)
	}
}