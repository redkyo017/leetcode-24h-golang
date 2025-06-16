package main

// 53. Maximum Subarray
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	total := 0
	for _, num := range nums {
		total += num
		maxSum = max(maxSum, total)
		if total < 0 {
			// the sub array end here
			total = 0
		}
	}
	return maxSum
}
