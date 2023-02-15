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

const N = 255000

var G [][][]int
var vis [N]int
var p [N]int
var tree [][][]int
var ans [N]int
var good bool
var val [N] int

func compute(node, wt int) {
	val[node] = wt
	for i := 0; i < len(G[node]); i++ {
		nxt := G[node][i][0]
		w := G[node][i][1]
		if p[nxt] == -1 {
			p[nxt] = p[node] ^ w
			compute(nxt, wt)
		} else {
			if p[nxt] != w^p[node] {
				good = true
			}
		}
	}
}

func dfs(node, par int) {
	for i := 0; i < len(tree[node]); i++ {
		nxt := tree[node][i][0]
		idx := tree[node][i][1]
		if nxt == par {
			continue
		}
		ans[idx] = p[node] ^ p[nxt]
		dfs(nxt, node)
	}
}

func solve() {
	var n, m int
	fmt.Fscan(in, &n, &m)
	tree = make([][][]int, N)
	G = make([][][]int, N)

	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		tree[x] = append(tree[x], []int{y, i})
		tree[y] = append(tree[y], []int{x, i})
	}

	for i := 0; i < m; i++ {
		var u, v, x int
		fmt.Fscan(in, &u, &v, &x)
		G[u] = append(G[u], []int{v, x})
		G[v] = append(G[v], []int{u, x})
	}

	for i := 0; i < N; i++ {
		p[i] = -1
	}
	for i:=1 ; i<=n ; i++{
		if p[i] == -1{
			p[i] = 0
			compute(i,i)
		}
	}

	if good == true{
		fmt.Println("NO")
		os.Exit(0)
	}


	count := make(map[int] int)
	ok := 0
	for i:=1 ; i<=n; i++  {
		if len(tree[i])&1==1{
			ok^=p[i]
			if val[i] != 1{
				count[val[i]]+=1
			}
		}
	}

	for k,v := range count{
		if v&1==1{
			for i:=1 ; i<=n; i++{
				if val[i] == k {
					p[i]^=ok
				}
			}
			break
		}
	}
	dfs(1,1)
	fmt.Println("YES")
	for i:=0 ; i<n-1  ;i++{
		fmt.Printf("%d ", ans[i])
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
