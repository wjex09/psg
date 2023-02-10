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
/*
const MAXSZ = 200005
var tree[MAXSZ][]  int
var st[MAXSZ] int
var en[MAXSZ] int
var timer int
var vertex [MAXSZ] int

func dfs(node,par int) {
	st[node] = timer
	timer+=1
	ver[st[node]] = node
	for _,next := range tree[node]{
		if(node == par) {continue}
		dfs(next,node)
	}
	en[node] = timer
}

const TSIZE = 1000

var tape [TSIZE] int
func copy() {

}
func repeat()
func repeatcopy()
*/


func solve(){
	var n int
	fmt.Fscan(in,&n)
	var od []int
	var ev []int
	o:=0
	e:=0
	for i := 0 ; i < n ; i++{
		var x int
		fmt.Fscan(in,&x)
		if(x%2!=0){
			o+=1
			od = append(od,i+1)
		} else {
			e+=1
			ev = append(ev,i+1)
		}
	}

	if o>2 {
		fmt.Printf("YES\n%d %d %d\n",od[0],od[1],od[2])
	} else if o > 0 && e > 1 {
		fmt.Printf("YES\n%d %d %d\n",od[0] , ev[0] ,ev [1] )
	} else{
		fmt.Printf("NO\n")
	}
}

func main(){
	defer out.Flush()
	t:=0
	fmt.Fscan(in,&t)
	for i := 0 ; i < t ; i++{
		solve()
	}
}
