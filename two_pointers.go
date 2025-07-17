package main

import (
	"sort"
	"strings"
)

// 125 Valid Palindrome
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	isAlphanumeric := func(char byte) bool {
		return (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9')
	}
	l, r := 0, len(s)-1
	for l < r {
		for !isAlphanumeric(s[l]) && l < r {
			l++
		}
		for !isAlphanumeric(s[r]) && l < r {
			r--
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

// 121. Best Time to Buy and Sell Stock
func MaxProfit(prices []int) int {
	l := 0
	res := 0
	for r := range prices {
		if prices[r] < prices[l] {
			l = r
		}
		res = max(res, prices[r]-prices[l])
	}
	return res
}

// 11. Container With Most Water
func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l <= r {
		res = max(res, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

// 167. Two Sum II - Input Array Is Sorted
func TwoSumII(numbers []int, target int) []int {
	res := []int{}
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			res = append(res, l+1, r+1)
			return res
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return res
}

// 15. 3Sum
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := num + nums[l] + nums[r]
			if threeSum > 0 {
				r -= 1
			} else if threeSum < 0 {
				l += 1
			} else {
				res = append(res, []int{num, nums[l], nums[r]})
				l += 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
		}
	}

	return res
}

// 42. Trapping Rain Water
func Trap(height []int) int {
	res := 0
	if len(height) == 0 {
		return res
	}
	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	for l < r {
		if leftMax < rightMax {
			l += 1
			leftMax = max(leftMax, height[l])
			res += leftMax - height[l]
			// here dont need to check if (leftmax - height[l]) is negative or not cause we already set the leftmax before canculating, worse case is leftmax = height[i] and the calculating result will be 0, no impact on the res, or we can change the order of above lines of code as
			// if leftMax - height[l] > 0 {
			// res += leftMax - height[l]
			// }
			// leftMax = max(leftMax, height[l])
			// same for the right case below
		} else {
			r -= 1
			rightMax = max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}
	return res
}
