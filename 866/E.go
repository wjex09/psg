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

const N = 200050
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

var tr [][]int
var root int
var st [N]int
var ans []int

func dfs(u, par int) int {
	st[u] = 1
	for _, i := range tr[u] {
		if i == par {
			continue
		}
		st[u] += dfs(i, u)
	}
	return st[u]
}

func dum(x int) {

}

func og(x int) {

}

func solve() {
	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u, v = u-1, v-1
		tr[u] = append(tr[u], v)
		tr[v] = append(tr[v], u)
	}
	st, ru := dfs(0, -1), dfs(root, -1)

}

func main() {
	defer out.Flush()
	t := 1
	//fmt.Fscan(in,&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
