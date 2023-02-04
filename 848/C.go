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
