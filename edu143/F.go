package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

const N = 255000
const MOD = 998244353

func pow(a, x, mod int) int {
	mod = MOD
	ret := 1
	for x > 0 {
		if x&1 == 1 {
			ret = ret % mod * a % mod
		}
		a = a % mod * a % mod
		x /= 2
	}
	return ret
}

func inv(a int) int {
	return pow(a, MOD-2, MOD)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

var tree [][]int
var dep []int
var black []int
var cnt []int
var tmp []int
var arr []int
var req []int
var ok bool
var k int

func dfs(no, par int) {
	dep[no] = 0
	cnt[no] = black[no]
	tmp[no] = req[no]
	for _, nxt := range tree[no] {
		//fmt.Println(no, nxt, par, tree[no])
		if nxt == par {
			continue
		}
		dfs(nxt, no)
		if tmp[nxt] > 0 {
			if cnt[no] > 0 {
				ok = false
				return
			}
			cnt[no] = 1
			tmp[no] = tmp[nxt] - 1
		} else if cnt[nxt] == 0 {
			dep[no] = max(dep[no], dep[nxt]+1)
		}
	}
	if dep[no] >= tmp[no] {
		tmp[no] = 0
	}
}

func check(x int) bool {
	for i := 1; i <= k; i++ {
		if i <= x {
			req[arr[i]] = (x-i)/k + 1
		} else {
			req[arr[i]] = 0
		}
	}
	ok = true
	dfs(1, 0)
	if tmp[1] > 0 {
		ok = false
	}
	return ok
}

func solve() {
	var n int
	fmt.Fscan(in, &n)
	tree = make([][]int, N)
	dep = make([]int, n+5)
	req = make([]int, n+5)
	black = make([]int, n+5)
	tmp = make([]int, n+5)
	cnt = make([]int, n+5)
	for i := 0; i <= n; i++ {
		dep[i] = 0
		black[i] = 0
		req[i] = 0
		for j := 0; j < len(tree[i]); j++ {
			tree[i][j] = 0
		}
	}
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		tree[x] = append(tree[x], y)
		tree[y] = append(tree[y], x)
	}

	//fmt.Println(tree)
	fmt.Fscan(in, &k)
	arr = make([]int, k+1)
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &arr[i])
		black[arr[i]] = 1
	}

	L, R, ans := 1, n-k, 0

	for L <= R {
		mid := (L + R) / 2
		if check(mid) {
			L = mid + 1
			ans = mid
		} else {
			R = mid - 1
		}
	}

	fmt.Println(ans)
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
