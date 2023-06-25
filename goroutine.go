package main

import (
	"fmt"
	"sync"
	"time"
)

func count(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		count("a")
	}()

	go func() {
		defer wg.Done()
		count("b")
	}()

	wg.Wait()
	fmt.Println("Done")
}
