package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scanf("%d", &n)
	var q []int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	for i, v := range q {
		fmt.Println("%d %d", i, v)
	}
}

func main() {
	t := 1
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
