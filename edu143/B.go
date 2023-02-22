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

func solve() {
	var n, k int
	fmt.Fscan(in, &n, &k)
	l, r := false, false
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if x == k {
			l = true
		}
		if y == k {
			r = true
		}
	}
	if l && r {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
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
