package main

import "fmt"

func main() {
	number := 100
	fmt.Println("Even panic this defer method will be called")
	defer fmt.Println("Leaving from method, but it will print before panic")

	if number == 100 {
		panic("number is 100")
	}

	//variation()
}

func variation() {
	number := 100
	fmt.Println("Panic this defer method will not be called")

	if number == 100 {
		panic("number is 100")
	}

	defer fmt.Println("This defer will not call")
}
