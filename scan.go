package main

import (
	"fmt"
	"sort"
)

var t int

func solve() {
	vecs := [3]int{}
	var a, b, c int

	fmt.Scanln(&vecs[0], &vecs[1], &vecs[2]) // 1 2 3
	fmt.Scanf("%d %d %d\n", &a, &b, &c)      // 1 2 3
	fmt.Scan(&a, &b, &c)                     // 1 2 3

	sort.Ints(vecs[:])
	fmt.Println(vecs[1])
}

func main() {
	fmt.Scanln(&t)        // 1
	fmt.Scanf("%d\n", &t) // 5
	fmt.Scan(&t)          // 5

	for i := 0; i < t; i++ {
		solve()
	}
}
