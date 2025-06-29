package main

import "strings"

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
