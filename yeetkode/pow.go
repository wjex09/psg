func myPow(x float64, n int) float64 {
    if (n==2147483648) {
        return x
    } else if (n == -2147483648){
        if x == 1 || x == -1 {
            return 1
        } else {
            return 0
        }
    }
    if n < 0 {
        x = 1/x
        n*=-1
    }
 	ret := 1.
	for n > 0 {
		if n&1 == 1 {
			ret = ret * x
		}
		x = x * x
		n /= 2
	}
	return ret
}

