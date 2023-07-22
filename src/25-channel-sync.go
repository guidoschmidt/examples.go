package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Work it, make it, do it…")
	time.Sleep(time.Second)
	fmt.Println("\ndone.")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
