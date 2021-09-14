package math

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n%2 == 0 {
		return myPow(x, n/2) * x * x
	}
	return myPow(x, n/2) * x * x * x
}
