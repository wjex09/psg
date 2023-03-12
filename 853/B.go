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
	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	cnt, pre := 0, 0
	for i := 0; i < n-i-1; i++ {
		if pre == 0 && s[i] != s[n-i-1] {
			cnt++
		}
		if s[i] == s[n-i-1] {
			pre = 0
		} else {
			pre++
		}
	}
	if cnt >= 2 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
