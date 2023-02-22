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

func solve() {
	var n int
	fmt.Fscan(in, &n)
	num, den := 1, 1
	mod := MOD
	for i := 1; i <= n/3; i++ {
		num = num % mod * i % mod
	}
	for i := 1; i <= n/6; i++ {
		den = den % mod * i % mod
	}

	ans := num % mod * inv(pow(den, 2, mod)) % mod

	for i := 0; i < n/3; i++ {
		count := make(map[int]int)
		mi := MOD
		for j := 0; j < 3; j++ {
			var x int
			fmt.Fscan(in, &x)
			count[x] += 1
			mi = min(mi, x)
		}
		ans = ans % MOD * count[mi] % MOD
	}

	fmt.Println(ans)

}

func main() {
	defer out.Flush()
	t := 1
	//fmt.Fscan(in,&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
