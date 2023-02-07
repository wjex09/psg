package main
import (
	"fmt"
)
const mod = 1000000007
func main(){
	N := 1000
	ans := 0
	p2 := make([]int,N+5)
	pr := make([]bool,N+5)
	p2[0] = 1
	for i:=1 ; i <=N ; i++{
		p2[i] = (2 * p2[i-1])%mod
		pr[i] =  true
	}
	for i:=2 ; i*i<N ;i++{
		if(!pr[i]) {
			continue
		}
		for j := 2*i ; j < N  ; j+=i{
			pr[j] = false
		}
	}
	//fmt.Println(p2,ans)

	cnt := 0
	var primeans[] int
	// taking care of prime n's
	for i := 2 ; i <= N ; i++{
		if(pr[i]){
			cnt = (i)/2
			primeans=append(primeans,cnt%mod*(p2[i]%mod + p2[i-1]%mod)+2)
		}
		else {

		}
	}
	fmt.Println(primeans)
	for _ , v := range primeans {
		ans = ans%mod+v%mod
	}
	fmt.Println(ans)
}
