package main

import (
	"fmt"
	"os"
	"sort"
)

type byLength []string

func (b byLength) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {

	s := []string{"v", "c", "a"}

	sort.Strings(s)
	fmt.Println(s)

	b := []string{"nb", "ac", "ii", "abc"}

	sort.Sort(byLength(b))
	fmt.Println(b)

	// region Swap

	sw := []string{"a", "b"}
	fmt.Println(sw)

	//sw[0], sw[1] = sw[1], sw[0]

	func(a, b *string) {
		*a, *b = *b, *a
		fmt.Println(*a, *b)
	}(&sw[0], &sw[1])

	fmt.Println(sw)

	// endregion

	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.
	panic("a problem")

	// A common use of panic is to abort if a function
	// returns an error value that we don't know how to
	// (or want to) handle. Here's an example of
	// `panic`king if we get an unexpected error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
