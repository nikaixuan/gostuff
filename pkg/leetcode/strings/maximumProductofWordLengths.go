package strings

import (
	"math"
	"strings"
)

// 318
func maxProduct(words []string) int {
	if len(words) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if !strings.ContainsAny(words[i], words[j]) {
				max = int(math.Max(float64(max), float64(len(words[i])*len(words[j]))))
			}
		}
	}
	return max
}
