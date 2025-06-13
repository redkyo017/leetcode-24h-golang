package main

import "math"

// 704. Binary Search
func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 875. Koko Eating Bananas
func MinEatingSpeed(piles []int, h int) int {
	l, r := 1, 1000000000
	res := r
	var canEat func(piles []int, k int, h int) bool
	canEat = func(piles []int, k int, h int) bool {
		for _, pile := range piles {
			h -= int(math.Ceil(float64(pile) / float64(k)))
		}
		return h >= 0
	}
	for l <= r {
		k := (l + r) / 2
		if canEat(piles, k, h) {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}
