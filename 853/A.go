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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if gcd(a[i], a[j]) <= 2 {
				fmt.Println("YES")
				return
			}
		}
	}

	fmt.Println("NO")

}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
