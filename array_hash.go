package main

import (
	"fmt"
	"slices"
)

// 217. Contains Duplicate
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}

// 242. Valid Anagram
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	length := len(s)
	for i := 0; i < length; i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, count := range m {
		if count < 0 {
			return false
		}
	}
	return true
}

// 1. Two Sum
func TwoSum(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{}
	}
	res := []int{}
	m := make(map[int]int)
	for index, num := range nums {
		complement := target - num
		if i, ok := m[complement]; ok {
			return []int{index, i}
		}
		m[num] = index
	}

	return res
}

func IsValidSudoku(board [][]byte) bool {
	if len(board) < 9 || len(board[0]) < 9 {
		return false
	}
	row_map := make(map[int][]byte)
	col_map := make(map[int][]byte)
	box_map := make(map[string][]byte)

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}

			boxId := fmt.Sprintf("%d-%d", r/3, c/3)
			if slices.Contains(row_map[r], val) || slices.Contains(col_map[c], val) || slices.Contains(box_map[boxId], val) {
				return false
			}
			row_map[r] = append(row_map[r], val)
			col_map[c] = append(col_map[c], val)
			box_map[boxId] = append(box_map[boxId], val)
		}
	}
	return true
}

// 238. Product of Array Except Self

func ProductExceptSelf(nums []int) []int {
	// PREFIX SUM
	if len(nums) == 0 {
		return []int{}
	}
	res := make([]int, len(nums))
	current := 1
	for i := 0; i < len(nums); i++ {
		res[i] = current
		current *= nums[i]
	}
	current = 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] = res[j] * current
		current *= nums[j]
	}

	return res
}

// 287. Find the Duplicate Number
func FindDuplicate(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = true
	}
	return -1
}

// 128. Longest Consecutive Sequence
func LongestConsecutive(nums []int) int {
	numSet := make(map[int]bool, 0)
	for _, num := range nums {
		numSet[num] = true
	}
	res := 0
	for num := range numSet {
		if !numSet[num-1] {
			count := 1
			currentNum := num
			for numSet[currentNum+1] {
				currentNum++
				count++
			}
			res = max(res, count)
		}
	}
	return res
}

// 49. Group Anagrams
func GroupAnagrams(strs []string) [][]string {
	res := [][]string{}
	group := make(map[string][]string, len(strs))
	for _, str := range strs {
		countChar := [26]int{}
		for i := 'a' - 'a'; i <= 'z'-'z'; i++ {
			countChar[i] = 0
		}
		for _, c := range str {
			countChar[c-'a']++
		}
		key := ""
		for char, count := range countChar {
			if count != 0 {
				key += string(char) + string(count)
			}
		}
		if _, ok := group[key]; ok {
			group[key] = append(group[key], str)

		} else {
			group[key] = []string{str}
		}
	}
	for _, set := range group {
		res = append(res, set)
	}
	return res
}

// // 347. Top K Frequent Elements
// func TopKFrequent(nums []int, k int) []int {

// }
