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

func sum (x int)int {
		acc := 0 
		for x!=0{
			acc+=x%10
			x/=10
		} 
		return acc
} 

func solve() {
	var n int 
	fmt.Fscan(in,&n)	
	x := n/2
	y := n-x 
	for i:=1 ; i < mod ; i*=10{
		dx,dy := x/i%10, y/i%10 
		for j:=0 ; j < dx && j < 9 - dy ; j++{
			if x - y > 1{
				x -= i
				y -= i
			} else { break }
		}
	} 
	fmt.Println(x,y)
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

