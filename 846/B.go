
package main

import(
	"fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)
func gcd(a ,b int) int {
	for b!=0{
		a,b = b, a%b
	}
	return a
}

func max(a, b int) int {
	if (a < b)  {return b}
	return a
}


func solve(){
	var n int
	fmt.Fscan(in,&n)
	acc := 0
	v := make([]int,n)
	for i:=0 ; i<n ; i++{
		fmt.Fscan(in,&v[i])
		acc+=v[i]
	}
	ans := 1
	sm := 0
	for i:=0 ; i < n-1 ; i++{
		sm+=v[i]
		acc-=v[i]
		ans = max(ans,gcd(sm,acc))
	}
	fmt.Println(ans)
}


func main(){
	defer out.Flush()
	t:=0
	fmt.Fscan(in,&t)
	for i := 0 ; i < t ; i++{
		solve()
	}
}

