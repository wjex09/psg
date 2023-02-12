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
	arr := make([]int, n) 
	o,t,ans:=0,0,0
	for i:=0 ; i < n ; i++{ 
		fmt.Fscan(in,&arr[i])   
		if arr[i] == 2 {
			t++
		}  else {
			o++
		}
	}   
	cnt:=0
	if t&1==1{
		ans =-1
	} else if t==0 {
		ans = 1
	} else {
		for i := 0 ; i<n ; i++{ 
			//fmt.Println(i,cnt,arr[i])
			if arr[i] == 2 {
				cnt++
			} 
			if cnt == t/2 {
				ans = i + 1  
				break
			}
		}
	}
	fmt.Println(ans)
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

