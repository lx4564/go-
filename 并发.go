package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	println("hello goroutine ：" + fmt.Sprint(i))
}
func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}
func main() {
	HelloGoRoutine()
}
