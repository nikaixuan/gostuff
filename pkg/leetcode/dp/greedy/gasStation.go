package greedy

// 134
func canCompleteCircuit(gas []int, cost []int) int {
	start, tank, total := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank = tank + gas[i] - cost[i]
		if tank < 0 {
			total = total + tank
			start = i + 1
			tank = 0
		}
	}
	if total+tank < 0 {
		return -1
	}
	return start
}
