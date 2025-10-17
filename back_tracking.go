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

// 46. Permutations
func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	// RECURSIVE
	// res := [][]int{}
	// perm := Permute(nums[1:])
	// for _, p := range perm {
	// 	for i := 0; i <= len(p); i++ {
	// 		pCopy := append([]int{}, p...)
	// 		pCopy = append(pCopy[:i], append([]int{nums[0]}, pCopy[i:]...)...)
	// 		res = append(res, pCopy)
	// 	}
	// }
	// return res
	// ITERATIVE
	// perms := [][]int{{}}
	// for _, num := range nums {
	// 	var newPerms [][]int
	// 	for _, p := range perms {
	// 		for i := 0; i <= len(p); i++ {
	// 			pCopy := append([]int{}, p...)
	// 			pCopy = append(pCopy[:i], append([]int{num}, pCopy[i:]...)...)
	// 			newPerms = append(newPerms, pCopy)
	// 		}
	// 	}
	// 	perms = newPerms
	// }
	// return perms
	// BACKTRACKING
	var res [][]int
	var backtrack func(res *[][]int, perm []int, nums []int, pick []bool)
	backtrack = func(res *[][]int, perm, nums []int, pick []bool) {
		if len(perm) == len(nums) {
			temp := append([]int{}, perm...)
			*res = append(*res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !pick[i] {
				perm = append(perm, nums[i])
				pick[i] = true
				backtrack(res, perm, nums, pick)
				perm = perm[:len(perm)-1]
				pick[i] = false
			}
		}
	}
	backtrack(&res, []int{}, nums, make([]bool, len(nums)))
	return res
}

func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	// BRUTE FORCE
	// res := make(map[string][]int)
	// var backtrack func(int, []int)
	// backtrack = func(i int, subset []int) {
	//     if i == len(nums) {
	//         key := fmt.Sprint(subset)
	//         res[key] = append([]int{}, subset...)
	//         return
	//     }

	//     subset = append(subset, nums[i])
	//     backtrack(i+1, subset)
	//     subset = subset[:len(subset)-1]
	//     backtrack(i+1, subset)
	// }

	// backtrack(0, []int{})

	// var result [][]int
	// for _, v := range res {
	//     result = append(result, v)
	// }
	// return result
	// ITERATION
	// res := [][]int{{}}
	// prevIdx, idx := 0, 0
	// for i := 0; i < len(nums); i++ {
	// 	if i > 0 && nums[i] == nums[i-1] {
	// 		idx = prevIdx
	// 	} else {
	// 		idx = 0
	// 	}
	// 	prevIdx = len(res)
	// 	for j := idx; j < prevIdx; j++ {
	// 		tmp := append([]int{}, res[j]...)
	// 		tmp = append(tmp, nums[i])
	// 		res = append(res, tmp)
	// 	}
	// }
	// return res
	// BACKTRACKING
	var res [][]int
	var backtrack func(int, []int)
	backtrack = func(i int, subset []int) {
		if i == len(nums) {
			res = append(res, append([]int{}, subset...))
			return
		}
		subset = append(subset, nums[i])
		backtrack(i+1, subset)
		subset = subset[:len(subset)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		backtrack(i+1, subset)
	}
	backtrack(0, []int{})
	return res
}
