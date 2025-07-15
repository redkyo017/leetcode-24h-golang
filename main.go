package main

import "log"

func main() {
	log.Println("Hello Leetcode 24h for Golang")
	// s := "A man, a plan, a canal: Panama"
	// var p *TreeNode
	// q := &TreeNode{Val: 0}
	// log.Println(TopKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
	strs := []string{"neet", "code", "love", "you"}
	sol := Solution{}
	encoded := sol.Encode(strs)
	decoded := sol.Decoded(encoded)
	log.Println("con co", encoded, decoded)
}
