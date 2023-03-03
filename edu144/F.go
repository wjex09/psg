package main

import (
	"bufio"
	"fmt"
	"sort"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

const N = 20
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

var tree [][] int
var dp [] int
var good bool

func get(no , par , length , mx int){
	if mx < 0{
		return
	}
	st2 := make([]int,0)
	for _, i := range(tree[no]) {
		if i == par {
			continue
		}
		st2 = append(st2,dp[i])
	}
	if par!=0 {
		st2 = append(st2,mx)
	}
	sort.Ints(st2)

	if length <= st2[0]+1 && (len(st2) == 1  || length <= st2[1]){
		good = false
	}
	if par == 0 && len(st2) == 1 {
		st2 = append(st2,-1)
	}

	//fmt.Println("st2: " , st2)

	for _ , i := range(tree[no]) {
		if i == par {
			continue
		}
		if st2[0] == dp[i]{
			if len(st2) < 3 || length <= st2[2] {
				get(i,no,length,st2[1]+1)
			}
		} else if dp[i] == st2[1] {
				if len(st2) < 3 || length <= st2[2] {
						get(i,no,length,st2[0]+1)
				}
		} else {
				if length <= st2[1] {
					get(i,no,length,st2[0]+1)
			}
		}
	}
}

func dfs(no,par,length int){
	st:= make([]int,0)
	dp[no] = 0

	//fmt.Printf("no:%d par:%d length:%d\n" , no ,par , length)
	for _ , i := range(tree[no]){
		if i == par {
			continue
		}
		dfs(i,no,length)
		st = append(st,i)
		if dp[i] == -N {
			dp[no] = -N
		}
	}

	//fmt.Println(dp)


	if dp[no] == -N || len(st) == 0{
		return;
	}

	//fmt.Println("st before : " , st)


	sort.SliceStable(st,func(i,j int) bool {
		return dp[i] < dp[j]
	})

	//fmt.Println("st : " , st)

	if len(st) > 1 && length > dp[st[1]] {
		dp[no] = -N
	} else {
			if length <= dp[st[0]] {
				dp[no] = dp[st[len(st)-1]] + 1
			} else {
					dp[no] = dp[st[0]] + 1
			}
	}
}


func solve(){
	var n int
	fmt.Fscan(in,&n)
	tree = make([][]int,n+5)
	dp = make([]int,n+5)


	for i:=0 ; i<n-1 ;  i++{
		var x , y int
		fmt.Fscan(in,&x,&y)
		tree[x] = append(tree[x],y)
		tree[y] = append(tree[y],x)
	}

	//fmt.Println(tree)
	L,R:=0,n
	for R-L > 1 {
		mid := (L+R)/2
		mx := mid + 15
		dfs(1,0,mid)
		good = true
		get(1,0,mid,mx)
		if good == false {
			L =  mid
		}  else {
			R = mid
		}
	}

	fmt.Println(L+1)
}

func main() {
	defer out.Flush()
	t := 1
	fmt.Fscan(in,&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
