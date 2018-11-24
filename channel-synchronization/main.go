package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("waiting...")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Done!")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
	fmt.Println("Returning.")
}
