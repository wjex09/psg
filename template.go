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

func solve() {
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
