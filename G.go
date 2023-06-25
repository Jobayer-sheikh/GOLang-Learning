package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FastReader struct {
	index  int
	tokens []string
	*bufio.Reader
}

func (r *FastReader) UpdatesTokenIfRequired() {
	if len(r.tokens) == r.index {
		line, _ := r.ReadString('\n')
		r.tokens = strings.Fields(line)
		r.index = 0
	}
}

func (r *FastReader) NextToken() string {
	r.UpdatesTokenIfRequired()
	token := (r.tokens)[r.index]
	r.index++
	return token
}

func (r *FastReader) ReadInt() int {
	val, _ := strconv.Atoi(r.NextToken())
	return val
}

var reader = FastReader{Reader: bufio.NewReader(os.Stdin)}

func __main__() {

	n := reader.ReadInt()
	k := reader.ReadInt()

	pos := k
	neg := n - k

	for i := 1; i <= n; i++ {
		val := 0
		if i&1 == 1 {
			if neg > 0 {
				val = -i
				neg--
			} else {
				val = i
			}
		} else {
			if pos > 0 {
				val = i
				pos--
			} else {
				val = -i
			}
		}

		if i == 1 {
			fmt.Printf("%d", val)
		} else {
			fmt.Printf(" %d", val)
		}
	}
	fmt.Printf("\n")
}

var mn int = -100000

func max(a, b, c int) int {
	var mx = a
	if mx < b {
		mx = b
	}
	if mx < c {
		mx = c
	}

	return mx
}

func run(idx int, diff int, rods []int, dp [][]int) int {

	if idx == len(rods) {
		if diff == 0 {
			return 0
		}
		return mn
	}

	if dp[idx][diff+5000] != 0 {
		return dp[idx][diff+5000]
	}

	var x int = run(idx+1, diff, rods, dp)
	var a int = rods[idx] + run(idx+1, diff+rods[idx], rods, dp)
	var b int = rods[idx] + run(idx+1, diff-rods[idx], rods, dp)
	dp[idx][diff+5000] = max(x, a, b)

	return dp[idx][diff+5000]
}

func tallestBillboard(rods []int) int {
	dp := make([][]int, 21)

	for i := 0; i < 21; i++ {
		dp[i] = make([]int, 10001)
	}

	for i := 0; i < 21; i++ {
		for j := 0; j < 1000; j++ {
			dp[i][j] = -1
		}
	}

	return run(0, 0, rods, dp) >> 1
}

func main() {
	//cas := reader.ReadInt()
	//for i := 0; i < cas; i++ {
	//	__main__()
	//}

	dp := make([][]int, 21)

	for i := 0; i < 21; i++ {
		dp[i] = make([]int, 10001)
	}

	for i := 0; i < 21; i++ {
		for j := 0; j <= 10000; j++ {
			dp[i][j] = -1
		}
	}

	fmt.Println(dp)
}
