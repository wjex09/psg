package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
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

var graph [][]int
var deg []int
var vis []int
var occ map [int]int
func dfs_out(u int) -> int{
	cnt := 0
	if (vis[u] == 1) {
		return 0;
	}
	for _,i := range(graph[u]) {
		if occ[i] == 0{
			cnt+=dfs_out(i)
		}
	}
}

func dfs_central(u int) -> int {
	cnt := 0
	if (vis[u] == 1){
		return 0;
	}
	for _,i := range(graph[u]) {
		if occ[i] > 0 {
			cnt+=dfs_central(i)
		}
	}
}

func solve(){
	var n,m int
	fmt.Fscan(in,&n,&m)
	for i:= 0 ; i<=n ; i++ {
		vis[i] = 0
		deg[i] = 0
	}
	root := int(math.Sqrt(n))
	deg = make(int[],n+5)
	graph = make(int[][], n+5)
	occ = make(map[int]int)
	for i := 0 ; i < m ; i++{
		var u , v int
		fmt.Fscan(in,&u,&v)
		graph[u] = append(graph[u],v)
		graph[v] = append(graph[v],u)
		deg[u]++
		deg[v]++
	}
	if root <=2 {
		fmt.Println("NO")
		return
	}
	outside := make(int[],n)
	center := make(int[],n)
	for i:=1 ; i<=n ; i++{
		if(deg[i] == 2) outside = append(outside,i)
		if(deg[i] == 4) center = append(center,i)
	}

	node := outside[0]
	out_cnt:=dfs_out(node)
	if(out_cnt!=root) {
		fmt.Println("NO")
		return
	}
	for _,i := range(center) {
		occ[i]++
	}
	ok := true
	for _,i := range(center){
		vis[i] = 0
		tmp:=dfs_central(i)
		if(tmp!=root) ok = false
		for _,nxt := range(graph[i]){
			if deg[nxt] == 4 {
				dc++
			} else {
				do++
			}
		}
		if dc!=2 || do!=2 {
			ok = false
		}
	}
	if ok == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
/*
	tu kare jaye aalas ghatak addat
	raha nahi balak wala hai ye mijaz
	sab aadar kare hum pen se bimaar hai taakat
	hai taakat lagat shaurat ainth aur aafat
	maal aur kavita benz aur profit
	pyar aur parivaar sang
	aur saath me sa aur sare aam road pe moshpit
	hot shit high power socket
	kaha gaya dhyan mera ,shit i lost it
	khada hu mai chand pe
	i may jump off it
	chud chuki maa mere inn kandho ki
	chid chuki gand tabhi inn bando ki
	chad rahi jhaanh tabhi inn bhasdo ki
	rahi nahi maansik hai inn baccho si
	kare yaha na bhool ke bhi koi bakchodi

	wo chahte bano doctor(uhh)
	jab chota tha mai banna tha bus kakarot
	trauma leke bade hote bande bhot
	role model tha bus hip hop
	dost gaya jail dus saal kam se kam
	koi bola ek hi life hai dont fuck this up
	usne sune bus aakhri teen shabd
	tabhi karam iss gane ka sheershak
	par koi mange maafi likhit toh likhunga deez nuts
	shiv naresh ki tshirt
	se dosto me tees take ki deal tak
	abj mai officon me dere the interview
	use same 20 rupee bhi the bhot
	body count uss samay the galt
	bana ni ek raat me bada sa cult
	nalle bate kare jaye
	gane bajejaye
	muh kara sabka band
	*/
}


func main() {
	defer out.Flush()
	t := 1
  fmt.Fscan(in,&t)
	for i := 0; i < t; i++ {
		solve()
	}
}

