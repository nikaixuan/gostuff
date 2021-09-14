package dp

type pair struct {
	first  int
	second int
}

// 877
// the question is always return true
// this solution is a general solution for using dp to return the diff of the optimize case
// if the pile number of the stone is not even, can use this
func stoneGame(piles []int) bool {
	pairs := make([][]pair, len(piles))
	for i := range pairs {
		pairs[i] = make([]pair, len(piles))
		pairs[i][i].first = piles[i]
		pairs[i][i].second = 0
	}

	// how to traverse the dp array from left top to right bottom
	for l := 2; l < len(piles); l++ {
		for i := 0; i <= len(piles)-l; i++ {
			j := l + i - 1
			// take left or take right, the number depends on the last second hand + current pile of stone
			left := pairs[i+1][j].second + piles[j]
			right := pairs[i][j-1].second + piles[j]
			// take the bigger one
			if left < right {
				pairs[i][j].first = right
				// second hand number equal to the first hand of last one
				pairs[i][j].second = pairs[i][j-1].first
			} else {
				pairs[i][j].first = left
				pairs[i][j].second = pairs[i+1][j].first
			}
		}
	}
	res := pairs[0][len(piles)-1]
	return res.first > res.second
}
