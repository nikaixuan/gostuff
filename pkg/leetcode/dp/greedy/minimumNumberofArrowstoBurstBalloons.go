package greedy

import "sort"

// 452
func findMinArrowShots(points [][]int) int {
	stack := make([][]int, 0)

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	stack = append(stack, points[0])
	i := 1
	for i < len(points) {
		for len(stack) > 0 && stack[len(stack)-1][1] >= points[i][0] {
			if stack[len(stack)-1][1] <= points[i][1] {
				newPoint := []int{points[i][0], stack[len(stack)-1][1]}
				stack = stack[:len(stack)-1]
				stack = append(stack, newPoint)
			} else {
				stack = stack[:len(stack)-1]
				stack = append(stack, points[i])
			}
			i++
		}
		if stack[len(stack)-1][1] < points[i][0] {
			stack = append(stack, points[i])
			i++
		}
	}
	return len(stack)
}

func findMinArrowShotsNoStack(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	pos := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		if pos >= points[i][0] {
			continue
		}
		count++
		pos = points[i][1]
	}
	return count
}
