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

// 74. Search a 2D Matrix
func SearchMatrix(matrix [][]int, target int) bool {
	ROWS, COLS := len(matrix), len(matrix[0])

	top, bot := 0, ROWS-1
	for top <= bot {
		row := (top + bot) / 2
		if target > matrix[row][len(matrix[row])-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bot = row - 1
		} else {
			break
		}
	}
	if top > bot {
		return false
	}

	row := (top + bot) / 2
	l, r := 0, COLS-1
	for l <= r {
		mid := (l + r) / 2
		if target > matrix[row][mid] {
			l = mid + 1
		} else if target < matrix[row][mid] {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}

// 153. Find Minimum in Rotated Sorted Array
func FindMin(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] < nums[r] {
			res = min(res, nums[l])
			break
		}
		m := (l + r) / 2
		res = min(res, nums[m])
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}

// 33. Search in Rotated Sorted Array
func SearchInRotated(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if target < nums[l] || target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
