package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve() {
	var n int
	fmt.Fscan(in, &n)
	acc := 0
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		acc += arr[i]
	}

	ans := -1e9 + 7
	for i := 0; i < n-1; i++ {
		ans = math.Max(ans, float64(acc-2*(arr[i]+arr[i+1])))
	}
	fmt.Fprintln(out, ans)
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
