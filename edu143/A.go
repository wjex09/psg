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
	var n, m int
	var A, B string
	fmt.Fscan(in, &n, &m)
	fmt.Fscan(in, &A, &B)
	var S []rune
	chars := []rune(B)
	for i := len(chars) - 1; i >= 0; i-- {
		S = append(S, chars[i])
	}

	str := A + string(S)
	cnt := 0
	for i := 0; i < n+m-1; i++ {
		if str[i] == str[i+1] {
			cnt++
		}
	}
	if cnt <= 1 {
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
