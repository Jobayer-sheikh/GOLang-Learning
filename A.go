package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func replace(s *string, idx int, c rune) *string {
	if idx == 0 {
		*s = string(c) + (*s)[idx+1:]
	} else if idx+1 == len(*s) {
		*s = (*s)[0:idx] + string(c)
	} else {
		*s = (*s)[:idx] + string(c) + (*s)[idx+1:]
	}

	return s
}

func __main__() string {

	var s string
	fmt.Fscan(rd, &s)

	var sz, f, h = len(s), -1, 0

	for i := sz - 1; i >= 0; i-- {
		if int(s[i])+h >= '5' {
			f = i
			h = 1
		} else {
			h = 0
		}
	}

	if h == 1 {
		return "1" + strings.Repeat("0", sz)
	}

	if f >= 0 {
		var t = s[0:f]

		var c string = *replace(&t, f-1, rune(s[f-1])+1)
		return c + strings.Repeat("0", sz-f)
	}

	return s
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
