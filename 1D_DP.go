package main

// 70. Climbing Stairs
func ClimbStairs(n int) int {
	n1, n2 := 0, 1
	for i := 0; i < n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}

// 746. Min Cost Climbing Stairs
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}
	return min(cost[0], cost[1])
}
