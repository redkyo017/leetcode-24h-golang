package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 226. Invert Binary Tree
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = InvertTree(root.Right)
	root.Right = InvertTree(temp)
	return root
}

// 104. Maximum Depth of Binary Tree
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + MaxDepth(root.Left)
	right := 1 + MaxDepth(root.Right)

	return max(left, right)

	// // DSF
	// maxD := 0
	// // var dfs func(node *TreeNode, count int)
	// // dfs = func(node *TreeNode, count int) {
	// // 	if node == nil {
	// // 		maxD = max(maxD, count)
	// // 		return
	// // 	}
	// // 	dfs(node.Left, count+1)
	// // 	dfs(node.Right, count+1)
	// // }
	// // dfs(root, 0)

	// // BFS
	// queue := []*TreeNode{root}
	// for len(queue) > 0 {
	// 	n := len(queue)
	// 	for i := 0; i < n; i++ {
	// 		node := queue[0]
	// 		queue = queue[1:]
	// 		if node.Left != nil {
	// 			queue = append(queue, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			queue = append(queue, node.Right)
	// 		}
	// 	}
	// 	maxD++
	// }
	// return maxD
}

// 543. Diameter of Binary Tree

func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0

	var findMaxDepth func(node *TreeNode) int
	findMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		left := 1 + findMaxDepth(node.Left)
		right := 1 + findMaxDepth(node.Right)
		res = max(res, left+right)
		return max(left, right)
	}
	findMaxDepth(root)

	return res
}

// 110. Balanced Binary Tree
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := true
	var findDepth func(node *TreeNode) int
	findDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := 1 + findDepth(node.Left)
		right := 1 + findDepth(node.Right)
		if math.Abs(float64(left-right)) > 1 {
			res = false
		}
		return max(left, right)
	}
	findDepth(root)
	return res
}

// 100. Same Tree
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// 98. Validate Binary Search Tree
func IsValidBST(root *TreeNode) bool {
	// RECURSION DIVIDE AND CONQUER
	// var findIfInvalidBST func(root *TreeNode, left *TreeNode, right *TreeNode) bool
	// findIfInvalidBST = func(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	// 	if root == nil {
	// 		return true
	// 	}
	// 	if min != nil && min.Val > root.Val {
	// 		return false
	// 	}
	// 	if min != nil && min.Val < root.Val {
	// 		return false
	// 	}
	// 	return findIfInvalidBST(root.Left, min, root) && findIfInvalidBST(root.Right, root, max)
	// }
	// return findIfInvalidBST(root, nil, nil)
	// ITERATIVE - DEPTH FIRST SEARCH
	var dfs func(root *TreeNode, left int, right int) bool
	dfs = func(root *TreeNode, left int, right int) bool {
		if root == nil {
			return true
		}
		if root.Val > left && root.Val < right {
			return dfs(root.Left, left, root.Val) && dfs(root.Right, root.Val, right)
		}
		return false
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

// 572. Subtree of Another Tree
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	var sameTree func(t1 *TreeNode, t2 *TreeNode) bool
	sameTree = func(t1 *TreeNode, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil || (t1.Val != t2.Val) {
			return false
		}
		return sameTree(t1.Left, t2.Left) && sameTree(t1.Right, t2.Right)
	}

	if sameTree(root, subRoot) {
		return true
	}
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
