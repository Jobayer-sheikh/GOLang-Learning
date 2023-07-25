package main

import (
	"bufio"
	"fmt"
	"os"
)

func absNUmber(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func maxNumber(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minNumber(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var rd = bufio.NewReader(os.Stdin)

func __main__() string {
	var n, k int
	fmt.Fscan(rd, &n, &k)

	var c []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &c[i])
	}

	var sz = len(c)
	if c[0] == c[sz-1] {
		var s int = 0
		for i := 0; i < n; i++ {
			if c[i] == c[0] {
				s++
			}
		}

		if s >= k {
			return "YES"
		}

		return "NO"
	}

	var x, y = 0, 0
	var l, r = 0, sz - 1

	for l < sz {
		if c[l] == c[0] {
			x++
		}
		if x == k {
			break
		}
		l++
	}

	for r >= 0 {
		if c[r] == c[sz-1] {
			y++
		}
		if y == k {
			break
		}
		r--
	}

	if x == k && y == k && l < r {
		return "YES"
	}

	return "NO"
}

func main() {
	var cas int
	fmt.Fscan(rd, &cas)

	for cas > 0 {
		var res = __main__()
		fmt.Println(res)
		cas--
	}
}
