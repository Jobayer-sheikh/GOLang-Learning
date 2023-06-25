package main

import (
	"fmt"
	"time"
)

func main() {

	// Channel has its own blocking nature
	// We don't have to use WaitGroup explicitly
	// Sending and receiving message are blocking operation

	// region Normal channel

	c := make(chan int)

	go func(n int) {
		summation(n, c)
	}(10)

	msg := <-c
	fmt.Println(msg)

	// endregion

	// region Multiple message

	mc := make(chan int)
	go multipleReceive(mc)

	/*
		for {
			msg, open := <-mc
			if !open {
				break
			}
			fmt.Println(msg)
		}
	*/

	for msg := range mc {
		fmt.Println(msg)
	}

	// endregion

	// region Buffer Channel

	// Send wil be block until something will ready to receive
	// But buffer channel make it possible
	bc := make(chan string, 2)
	bc <- "hello"
	bc <- "world"

	ms := <-bc
	println(ms)

	ms = <-bc
	println(ms)

	// endregion

	// region Channel using Select

	one := make(chan string)
	two := make(chan string)

	go func() {
		for {
			one <- "Slow channel"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			two <- "Fast channel"
			time.Sleep(time.Microsecond * 500)
		}
	}()

	/*
		for {
			x := <-one
			fmt.Println(x)
			y := <-two
			fmt.Println(y)
		}
	*/

	for {
		select {
		case x := <-one:
			{
				fmt.Println(x)
			}
		case y := <-two:
			{
				fmt.Println(y)
			}
		}
	}

	// endregion

	fmt.Println("Done")
}

func summation(n int, c chan int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	time.Sleep(time.Second * 1)
	c <- sum
}

func multipleReceive(c chan int) {
	for i := 0; i < 5; i++ {
		//time.Sleep(time.Second * 2)
		c <- i
	}
	close(c)
}
