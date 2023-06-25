package main

import (
	"errors"
	"fmt"
	"math"
)

type combine struct {
	name string
	id   int
}

func main() {

	// region basic

	var a int = 100
	v := 100

	var sum int = a + v

	fmt.Println("Hello world", v, a, sum)

	if a < 200 {
		fmt.Println("a is less than 200")
	} else {
		fmt.Println("a is not greater than 200")
	}

	var arr [10]int
	fmt.Println(arr)

	brr := []int{1, 2, 4}
	crr := brr
	brr = append(brr, 200)
	fmt.Println(brr)
	fmt.Println(crr)

	// Map
	map_var := make(map[string]int)
	map_var["hello"] = 1
	map_var["hello"] = 1

	delete(map_var, "hello")

	fmt.Println(map_var["hello"])

	// For loop
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// While loop
	i := 0
	for i < 10 {
		fmt.Print(i)
		if i == 5 {
			break
		}
		i++
	}
	fmt.Println()

	str := []string{"hello", "world"}
	for i, s := range str {
		fmt.Println(i, s)
	}

	// Map range
	map_range := make(map[string]int)
	map_range["a"] = 1
	map_range["b"] = 2
	map_range["c"] = 3

	for i, s := range map_range {
		fmt.Println(i, s)
	}

	// endregion

	sm, err := square(1, 0)
	if err == nil {
		fmt.Println(sm)
	} else {
		fmt.Println(err.Error())
	}

	c := combine{name: "Jobayer", id: 2}
	fmt.Println(c.name)

	// Multiple variable at once
	var v1, v2 int
	fmt.Println(v1, v2)

	var o = 1
	fmt.Println(o)

	// Switch Case
	var sc = 0
	switch sc {
	case 1:
		{
			fmt.Println("One")
		}
	case 3:
		{
			fmt.Println("Three")
		}
	}

	who := func(dataType interface{}) interface{} {
		switch dataType.(type) {
		case int:
			{
				fmt.Println("Integer")
			}
		case float64, float32:
			{
				fmt.Println("Float64")
			}
		default:
			{
				fmt.Println("Can't define type")
			}
		}
		return dataType
	}

	fmt.Println(who(1))

	// 2D Arrays
	// brr := []int{1, 2, 4}
	var twoD = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Print(twoD)
}

func square(x int, y int) (float64, error) {
	if x == 0 || y == 0 {
		return 0, errors.New("0 not accepted")
	}
	return math.Sqrt((float64)(x + y)), nil
}
