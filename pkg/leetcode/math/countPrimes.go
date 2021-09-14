package math

// 204
// solution 1
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	res := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
		}
	}
	return res
}

// solution2
func countPrimes2(n int) int {
	notPrime := make([]bool, n)
	res := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			res++
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}
	return res
}
