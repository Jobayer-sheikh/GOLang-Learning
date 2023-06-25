package main

import (
	"fmt"
)

func main() {
	str := make([]string, 100)
	fmt.Println(str)

	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(len(s))

	s[2] = "3"
	fmt.Println(s)

	c := make([]string, len(s))

	copy(c, s)
	c[0] = "1"
	c[1] = "2"
	fmt.Println(c, s)

	l := c[0:3]
	fmt.Println(l)

	l = c[:3]
	fmt.Println(l)

	l = c[1:]
	fmt.Println(l)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		var len = i + 1
		twoD[i] = make([]int, len)
		for j := 0; j < len; j++ {
			twoD[i][j] = i
		}
	}
	fmt.Println(twoD)

	mp := make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3

	value, hasValue := mp["a"]
	fmt.Println(value, hasValue)

	mpString := map[string]string{"a": "1", "b": "2", "c": "3"}
	fmt.Println(mpString)

	for i, c := range "Jobayer" {
		fmt.Println(i, (string(c)))
	}
}
