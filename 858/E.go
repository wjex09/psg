package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)


var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

const N = 100000
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

//var stor [][][] int
var cnt [][] int
var tab [N][MAXLG] int
var dis [] int

var par [] int

/*
	if x == y  then sum of weight on the path from x to 1
	else there should exist some p when x = y
*/


func lower_bound(array []int, target int) int {
	low, high, mid := 0, len(array), 0
	for low < high {
		mid = low + (high - low) / 2
		if array[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
		if low < len(array) && array[low] < target {
			low++
		}
	}
	return low
}
const MAXLG = 20
func get_idx(d,x int) int {  // x-th at d depth
	return lower_bound(cnt[d],x);
}

func dfs(no, d int) {
	dis[no] = d
	cnt[dis[no]] = append(cnt[dis[no]],no)
	tab[no][0] = par[no]
	for i := 1 ; i < MAXLG ; i++{
		tab[no][i] = tab[tab[no][i-1]][i-1]
	}
	for _,nxt := range tree[no] {
		dfs(nxt,d+1)
	}
}


const rootN = 333

func anc(x, k int)  int{   // return k-th ancestor ok node
	for i := 0 ; i < MAXLG ; i++ {
		if k&(1<<uint8(i)) == 1 {
			x = tab[x][i];
		}
	}
	return x
}

func solve(){
	var n,q int
	fmt.Fscan(in,&n,&q)
 	tree = make([][]int, n+5)
 	cnt = make([][]int,n+5)
	wt := make([]int, n+5)
	dis = make([]int, n+5)
	par = make([]int, n+5)
	mark := make([]bool,n+5)
	end := make([]int, n+5)
	//stor := make([][][]int, n+5,N)
	stor := [N+5][20][20] int {}
	for i := 1 ; i <= n ; i++{
		fmt.Fscan(in,&wt[i])
	}
	for i := 2 ; i <= n ; i++{
		fmt.Fscan(in,&par[i])
		tree[par[i]] = append(tree[par[i]],i)
	}


	//fmt.Println("POG")

	dfs(1,0)

	//fmt.Println(cnt)


	// for all heights upto n precompute sum in stor

	last := -1
	for i:=0 ; i <= n ; i++{
		nocount := len(cnt[i]) // nodes at i-th dis
		if nocount == 0 || nocount > 333 { //  0 < #cnt < root(N)
			end[i] = last
			last = i
			continue
		}
		sort.Ints(cnt[i])
		mark[i] = true
		end[i] = last
		//stor[i] = make([][]int,nocount+10)


		for j := 0 ; j<nocount ; j++{
			for k := j ; k<nocount ; k++{
				//stor[i][j] = make([]int,nocount+10)

				u,v := cnt[i][j] , cnt[i][k]
				//fmt.Println(u,v)
				//os.Exit(0)
				if i > 0 && mark[i-1] == true {
					//fmt.Println(i-1,par[u],par[v])
					x , y := get_idx(i-1,par[u]) , get_idx(i-1,par[v])

					if x > y  {
						x,y = y,x
					}
					//fmt.Println(stor[i-1])

					stor[i][j][k] = stor[i-1][x][y]

					if stor[i-1][x][y] == 0 {
					//	fmt.Println(i)
					}
				}
				//fmt.Println(stor[i])
				stor[i][j][k] += wt[u] * wt[v]

			}
		}
	}

	//fmt.Println("Good til here")

	//os.Exit(0)

	for i := 0 ; i < q ; i++{
		var x, y int
		fmt.Fscan(in,&x,&y)
		ans , d := 0 , dis[x]
		for d >= 0{
			if mark[d] == true {
				if x > y {
					x,y = y,x
				}
				ans += stor[d][get_idx(d,x)][get_idx(d,y)]
				//fmt.Println(d,x,y,ans)
				x,y = anc(x,d-end[d]) , anc(y,d-end[d])
			//	fmt.Println(x,y)
				d = end[d]
			} else {
					d = d - 1
					ans += wt[x] * wt[y]
					x,y = par[x] , par[y]
			}
		}
		fmt.Println(ans)
	}
}


func main() {
	defer out.Flush()
	t := 1
  //fmt.Fscan(in,&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
