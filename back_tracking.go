package main

import "sort"

// 78. Subsets
func Subsets(nums []int) [][]int {
	res := [][]int{}
	current := []int{}

	var dfs func(current []int, i int)
	dfs = func(current []int, i int) {
		if i == len(nums) {
			subset := make([]int, len(current))
			copy(subset, current)
			res = append(res, subset)
			return
		}
		// skip
		dfs(current, i+1)

		// include
		current = append(current, nums[i])
		dfs(current, i+1)

	}
	dfs(current, 0)
	return res
}

// 39. Combination Sum
func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) <= 0 {
		return [][]int{}
	}
	res := [][]int{}
	cur := []int{}

	var dfs func(current []int, sum int, start int)
	dfs = func(current []int, sum int, start int) {
		if sum == target {
			subset := make([]int, len(current))
			copy(subset, current)
			res = append(res, subset)
			return
		}
		if sum > target || start == len(candidates) {
			return
		}

		// for i := start; i < len(candidates); i++ {
		// 	current = append(current, candidates[i])
		// 	dfs(current, sum+candidates[i], i)
		// 	current = current[:len(current)-1]
		// }

		dfs(current, sum, start+1)

		current = append(current, candidates[start])
		dfs(current, sum+candidates[start], start)
		current = current[:len(current)-1]
	}
	dfs(cur, 0, 0)
	return res
}

// 40. Combination Sum II
func CombinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)

	var dfs func(int, []int, int)
	dfs = func(i int, current []int, total int) {
		if total == target {
			temp := make([]int, len(current))
			copy(temp, current)
			res = append(res, temp)
			return
		}
		if i >= len(candidates) || total > target {
			return
		}
		current = append(current, candidates[i])
		dfs(i+1, current, total+candidates[i])
		current = current[:len(current)-1]

		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i += 1
		}
		dfs(i+1, current, total)
	}

	dfs(0, []int{}, 0)
	return res
}
