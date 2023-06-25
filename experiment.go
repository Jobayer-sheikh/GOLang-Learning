package main

import (
	"fmt"
	"time"
)

func main() {
	//c := make(chan int)
	//
	//go func() {
	//	c <- 100
	//}()
	//
	//go func() {
	//	m := <-c
	//	fmt.Println(m)
	//}()
	//
	//m := <-c
	//fmt.Println(m)
	//
	//fmt.Println(<-time.After(1))

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	//c2 := make(chan string, 1)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c2 <- "result 2"
	//}()
	//
	//select {
	//case res := <-c2:
	//	fmt.Println(res)
	//case <-time.After(3 * time.Second):
	//	fmt.Println("timeout 2")
	//}

	sd := []string{"a", "b"}
	sd = append(sd, "c")

	fmt.Println(sd)
}
