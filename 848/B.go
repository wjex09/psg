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
	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)
	p := make([]int, n+1)
	pos := make([]int, n+1)
	a := make([]int, m)  
	ans := 1000000007
	for i := 1; i <= n; i++ { 
		fmt.Fscan(in, &p[i])  
		pos[p[i]] = i 
	} 

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i]) 
		if(i>0){
			ans = min(ans,pos[a[i]]-pos[a[i-1]]) 
			if(d < n-1){
				ans = min(ans, pos[a[i-1]] - pos[a[i]] + d + 1)
			}
		}
	} 
	fmt.Println(max(0,ans))
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
} 

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Ordered interface {
	Integer | ~string | Float
} 


func min[T Ordered](x T, value ...T) T{
	ret:=x 
	for _,val:= range value{
		if ret > val { ret = val }
	} 
	return ret
} 

func max[T Ordered](x T, value ...T) T{
	ret:=x 
	for _,val:= range value{
		if ret < val { ret = val }
	} 
	return ret
}