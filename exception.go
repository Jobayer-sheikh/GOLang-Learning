package main

import "fmt"

func main() {
	//A common use of panic is to abort if a function returns an error value
	//that we don’t know how to (or want to) handle.
	//panic("Its a panic")

	// i = 0
	for i := 0; i <= 5; i++ {
		n, err := fmt.Println(1 / i)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(n)
	}

	fmt.Println("end")
}
