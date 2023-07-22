package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 12; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	for i := 1; i < 10; i++ {
		go f("gogogo")
	} 

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
