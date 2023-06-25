package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 50)
	go countChannel("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func countChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
