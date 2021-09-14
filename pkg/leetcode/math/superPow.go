package math

// 372
// 剩余定理也可以
func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	// a的1234次方 = a的4次方 * （a的1次方的)10次方)的2次方的)10次方)的3次方)的10次方）
	// (a*b)%k = ((a%k)*(b%k))%k
	res := calPow(a, b[len(b)-1]) * calPow(superPow(a, b[:len(b)-1]), 10)
	return res % 1337
}

func calPow(a, k int) int {
	res := 1
	a %= 1337
	for i := 0; i < k; i++ {
		res *= a
		res %= 1337
	}
	return res
}

// 有机会把求模规模缩小一半
func calPow2(a, k int) int {
	if k == 0 {
		return 1
	}
	a %= 1337
	if k%2 == 1 {
		return (a * calPow2(a, k-1)) % 1337
	} else {
		return (calPow2(a, k/2) * calPow2(a, k/2)) % 1337
	}
}
