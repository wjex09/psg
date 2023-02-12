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
 
const mod = 998244353


func solve() { 
	var n int 
	fmt.Fscan(in,&n) 
	if n&1 != 1 {
		fmt.Printf("NO\n")
	} else {
		fmt.Printf("YES\n")
		m:=(n-1)/2
		for i:=m ; i>=1 ; i--{
			fmt.Printf("%d %d\n", i , 2*m + 2*(m-i) + 3)
		} 
		for i:=m+1 ; i <= 2*m+1 ; i++{
			fmt.Printf("%d %d\n", i , 2*2*m - 2*(i-(m+1)) + 2)
		}
	}

}

func main() {
	defer out.Flush() 
	var t int
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

