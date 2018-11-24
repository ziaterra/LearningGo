package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {

			time.Sleep(time.Millisecond * 500)
			c1 <- "hey"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			c2 <- "gal"

		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
