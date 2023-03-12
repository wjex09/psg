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

func solve() {
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	mp := make(map[int]int)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		mp[a[i]] += m + 1
	}
	for i := 1; i <= m; i++ {
		var pos, v int
		fmt.Fscan(in, &pos, &v)
		mp[a[pos]] -= m - i + 1
		a[pos] = v
		mp[a[pos]] += m - i + 1
	}
	ans := n * (m + 1) * m
	for _, v := range mp {
		ans -= v * (v - 1) / 2
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
