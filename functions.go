package main

import "fmt"

func sum(val ...int) int {
	sum := 0
	// _ means skip
	for _, v := range val {
		sum += v
	}
	return sum
}

// Closure, return function anonymous
func intSet() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr...))

	nextInt := intSet()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
