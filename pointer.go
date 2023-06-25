package main

import "fmt"

func inc(val *int) {
	*val++
}

func main() {

	i := 10
	fmt.Println(i)

	inc(&i)
	fmt.Println(i)
}
