package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const cf = true

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

type SqrtInt interface {
	SignedSqrt | UnsignedSqrt
}

type SignedSqrt interface {
	~int | ~int32 | ~int64
}

type UnsignedSqrt interface {
	~uint | ~uint32 | ~uint64 | ~uintptr
}

func sqrt[T SqrtInt](x T) T {
	l, r := T(0), T(2000000123)
	for l < r {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func reverseStr(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Ordered interface {
	Integer | Float | ~string
}

func abs[T Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func min[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func max[T Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

var (
	in  = getReader()
	out = getWriter()
)

func getReader() *bufio.Reader {
	if cf {
		return bufio.NewReader(os.Stdin)
	} else {
		in, _ := os.Open("input.txt")
		return bufio.NewReader(in)
	}
}

func getWriter() *bufio.Writer {
	if cf {
		return bufio.NewWriter(os.Stdout)
	} else {
		out, _ := os.Create("output.txt")
		return bufio.NewWriter(out)
	}
}

func program() {
	arr := []int{0, 0, 0}
	fmt.Fscan(in, &arr[0], &arr[1], &arr[2])
	sort.Ints(arr)
	fmt.Fprintln(out, arr[1])
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		program()
	}
}
