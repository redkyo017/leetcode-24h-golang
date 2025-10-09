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

// 235. Lowest Common Ancestor of a Binary Search Tree
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}

	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}

	return root
}

// 102. Binary Tree Level Order Traversal
func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// 199. Binary Tree Right Side View
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				res = append(res, node.Val)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}
	return res
}

// 1448. Count Good Nodes in Binary Tree
func GoodNodes(root *TreeNode) int {
	// res := 0
	// BFS
	// if root == nil {
	// 	return 0
	// }
	// queue := []struct {
	// 	node   *TreeNode
	// 	maxVal int
	// }{{root, math.MinInt32}}
	// for len(queue) > 0 {
	// 	front := queue[0]
	// 	queue = queue[1:]

	// 	node := front.node
	// 	maxVal := front.maxVal

	// 	if node.Val >= maxVal {
	// 		res++
	// 	}
	// 	newMaxVal := max(maxVal, node.Val)
	// 	if node.Left != nil {
	// 		queue = append(queue, struct {
	// 			node   *TreeNode
	// 			maxVal int
	// 		}{node: node.Left, maxVal: newMaxVal})
	// 	}
	// 	if node.Right != nil {
	// 		queue = append(queue, struct {
	// 			node   *TreeNode
	// 			maxVal int
	// 		}{node: node.Right, maxVal: newMaxVal})
	// 	}
	// }
	// return res

	// DFS
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode, maxVal int) int
	dfs = func(node *TreeNode, maxVal int) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val >= maxVal {
			res = 1
		}
		maxVal = max(maxVal, node.Val)
		res += dfs(node.Left, maxVal)
		res += dfs(node.Right, maxVal)
		return res
	}
	return dfs(root, root.Val)
}

// 230. Kth Smallest Element in a BST
func KthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	// var arr []int
	var dfs func(node *TreeNode)
	// BRUTE FORCE
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	arr = append(arr, node.Val)
	// 	dfs(node.Left)
	// 	dfs(node.Right)
	// }
	// dfs(root)
	// sort.Ints(arr)
	// return arr[k-1]

	// IN-ORDER TRAVERSE
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	dfs(node.Left)
	// 	arr = append(arr, node.Val)
	// 	dfs(node.Right)
	// }
	// dfs(root)
	// return arr[k-1]

	// RECURSIVE OPTIMAL
	counter, res := k, 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		counter--
		if counter == 0 {
			res = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
func BuildTree(preorder []int, inorder []int) *TreeNode {
	// DFS
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	mid := 0
	for i, val := range inorder {
		if val == preorder[0] {
			mid = i
		}
	}
	root.Left = BuildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = BuildTree(preorder[mid+1:], inorder[mid+1:])
	return root

	// DFS with hash
	// indices := make(map[int]int)
	// for i, val := range inorder {
	// 	indices[val] = i
	// }
	// preIdx := 0
	// var dfs func(int, int) *TreeNode
	// dfs = func(left int, right int) *TreeNode {
	// 	if left > right {
	// 		return nil
	// 	}
	// 	rootVal := preorder[preIdx]
	// 	preIdx++

	// 	root := &TreeNode{Val: rootVal}
	// 	mid := indices[rootVal]
	// 	root.Left = dfs(left, mid-1)
	// 	root.Right = dfs(mid+1, right)
	// 	return root
	// }
	// return dfs(0, len(inorder)-1)
}
