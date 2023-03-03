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

const N = 200050
const MOD = 998244353

func pow(a,x,mod int) int {
	mod = MOD
	ret:=1
	for x>0{
		if x&1==1{
			ret = ret%mod*a%mod
		}
		a = a%mod * a%mod
		x/=2
	}
	return ret
}

func min(a, b int) int{
	if a > b { return b }
	return a
}

func max(a, b int) int{
	if a < b { return b }
	return a
}



func solve(){
	var n , k , x int
	fmt.Fscan(in,&n,&k,&x)
	a := make([]int,n+5)
	mn := make([]int,n+5)
	pref := make([]int, n+5)

	for i:=1 ; i <= n ; i++{
		fmt.Fscan(in,&a[i])
		a[i]-=x
	}
	for i:=1 ; i <= n ; i++{
		pref[i] =  pref[i-1] + a[i]
	}
	for i:=1 ; i <= n ; i++{
		mn[i] = min(mn[i-1],pref[i])
	}

	ans , acc := -MOD , 0

	// pos x
	if x>=0 {
		for i:=1 ; i<=n ; i++{
			for j:=i ; j>= max(1,i-k+1); j--{
				acc = pref[i] - pref[j-1] + 2 * (i-j+1) * x
				ans = max(ans,acc)
			}
			if i > k {
				acc =  pref[i] - mn[i-k-1] + 2 * k * x
				ans = max(ans , acc)
			}
		}
		ans = max(0,ans)
		fmt.Println(ans)
	} else { // negs
			for i:=1 ; i <=n ; i++{
				if i < n - k{
					ans = max(ans,pref[i] - mn[i-1] )
				} else {
						for j:=i ; j>0 ; j--{
							acc = pref[i]	- pref[j-1]  + 2 * max(0,k-(n-(i-j+1))) * x
							ans = max(ans, acc)
						}
				}
			}
		ans = max(0,ans)
		fmt.Println(ans)
	}
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in,&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
