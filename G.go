package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func dfs(root *TreeNode, sum int) int {
//	if root == nil {
//		return sum
//	}
//	ans := 0
//	if root.Left != nil && root.Right != nil {
//		x, y := dfs(root.Left, sum+1)+1, dfs(root.Right, sum+1)+1
//		if x < y {
//			ans = x
//		} else {
//			ans = y
//		}
//	} else if root.Left != nil {
//		ans = dfs(root.Left, sum+1) + 1
//	} else if root.Right != nil {
//		ans = dfs(root.Right, sum+1) + 1
//	}
//
//	return ans
//}

//func minDepth(root *TreeNode) int {
//	return dfs(root, 1)
//}

func __man__() {

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

func max3(a, b, c int) int {
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
	dp[idx][diff+5000] = max3(x, a, b)

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

func buddyStrings(s string, goal string) bool {

	if len(s) != len(goal) {
		return false
	}

	size := len(s)
	cnt := make(map[rune]int, size)

	if s == goal {
		for _, c := range s {
			cnt[c]++
			if cnt[c] == 2 {
				return true
			}
		}

		return false
	}

	index := make([]int, 2)
	for i := 0; i < size; i++ {
		if s[i] != goal[i] {
			index = append(index, i)
		}
	}

	if len(index) != 2 {
		return false
	}

	return s[index[0]] == goal[index[1]] && s[index[1]] == goal[index[0]]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(t *TreeNode, k int, vis map[int]bool, ans *[]int) int {
	if vis[t.Val] == true || k < 0 {
		return 0
	}

	vis[t.Val] = true
	if k == 0 {
		*ans = append(*ans, t.Val)
	}

	*ans = append(*ans, t.Val)

	if t.Left != nil {
		return dfs(t.Left, k-1, vis, ans)
	}

	if t.Right != nil {
		return dfs(t.Right, k-1, vis, ans)
	}

	return 0
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var ans []int
	var vis map[int]bool = make(map[int]bool)

	dfs(target, k, vis, &ans)
	return ans
}

func canFinish(n int, p [][]int) bool {

	var q = make([]int, 0)
	var m = make(map[int]int)
	var g = make(map[int][]int)

	for i := 0; i < len(p); i++ {
		m[p[i][1]]++
		if _, ok := g[p[i][0]]; ok == false {
			g[p[i][0]] = make([]int, 0)
			g[p[i][0]] = append(g[p[i][0]], p[i][1])
		} else {
			g[p[i][0]] = append(g[p[i][0]], p[i][1])
		}
	}

	var t = make(map[int]bool)
	for i := 0; i < len(p); i++ {
		if m[p[i][0]] == 0 && !t[p[i][0]] {
			t[p[i][0]] = true
			q = append(q, p[i][0])
		}
	}

	var vis = 0
	for len(q) != 0 {
		var f = q[0]

		vis++
		q = q[1:]

		if v, ok := g[f]; ok {
			for i := 0; i < len(v); i++ {
				var x = v[i]
				m[x]--
				if m[x] == 0 && !t[x] {
					t[x] = true
					q = append(q, x)
				}
			}
		}
	}

	return vis == n
}

func eraseOverlapIntervals(a [][]int) int {

	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})

	var sz, y, ans = len(a), a[0][1], 0

	for i := 1; i < sz; i++ {

		if y <= a[i][0] {
			y = a[i][1]
			continue
		}

		ans++
		if y > a[i][1] {
			y = a[i][1]
		}
	}

	return ans
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func asteroidCollision(a []int) []int {
	var ans = make([]int, 1)

	ans[0] = a[0]
	for i := 1; i < len(a); i++ {
		var sz = len(ans)
		if sz == 0 {
			ans = append(ans, a[i])
			continue
		}
		var x = ans[sz-1]
		if x >= 0 && a[i] < 0 {
			if x > absInt(a[i]) {
				continue
			} else if x < absInt(a[i]) {
				var ok bool = false

				for len(ans) > 0 {
					sz = len(ans)
					if ans[sz-1] < 0 || a[i] >= 0 {
						ok = true
						break
					}
					if absInt(ans[sz-1]) > absInt(a[i]) {
						ok = false
						break
					} else if absInt(ans[sz-1]) < absInt(a[i]) {
						ok = true
						ans = ans[0 : sz-1]
					} else {
						ok = false
						ans = ans[0 : sz-1]
						break
					}
				}

				if ok {
					ans = append(ans, a[i])
				}

			} else {
				ans = ans[0 : sz-1]
			}
		} else {
			ans = append(ans, a[i])
		}
	}

	return ans
}

func main() {
	//cas := reader.ReadInt()
	//for i := 0; i < cas; i++ {
	//	__main__()
	//}

	var e = []int{-2, -2, 1, -2}
	asteroidCollision(e)

	var r [][]int = [][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}
	eraseOverlapIntervals(r)

	var g [][]int = [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}

	canFinish(5, g)

	fmt.Println(buddyStrings("ab", "ba"))

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
