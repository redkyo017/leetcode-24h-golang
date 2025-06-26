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
func MinCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}
	return min(cost[0], cost[1])
}

// 198. House Robber
func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n1, n2 := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		n1, n2 = n2, max(n1+nums[i], n2)
	}
	return n2
}

// 213. House Robber II
func RobII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var robber func(amounts []int) int
	robber = func(amounts []int) int {
		n1, n2 := 0, amounts[0]
		for i := 1; i < len(amounts); i++ {
			n1, n2 = n2, max(n1+amounts[i], n2)
		}
		return n2
	}
	amount1 := robber(nums[:len(nums)-1])
	amount2 := robber(nums[1:])
	return max(amount1, amount2)
}

// 5. Longest Palindromic Substring
func LongestPalindrome(s string) string {
	maxChar := 0
	res := ""
	for i, _ := range s {
		// odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxChar {
				maxChar = r - l + 1
				res = s[l : r+1]
			}
			l, r = l-1, r+1
		}

		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxChar {
				maxChar = r - l + 1
				res = s[l : r+1]
			}
			l, r = l-1, r+1
		}
	}

	return res
}

// 647. Palindromic Substrings
func CountSubstrings(s string) int {
	res := 0
	for i, _ := range s {
		// odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res += 1
			l, r = l-1, r+1
		}

		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res += 1
			l, r = l-1, r+1
		}
	}

	return res
}
