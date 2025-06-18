package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Min() (v any) {
	return (*h)[0]
}

func (h *MinHeap) Pop() (v any) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return
}

type KthLargest struct {
	minHeap *MinHeap
	size    int
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := &MinHeap{}
	for _, num := range nums {
		heap.Push(minHeap, num)
		if len(*minHeap) > k {
			heap.Pop(minHeap)
		}
	}
	return KthLargest{minHeap, k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	if len(*this.minHeap) > this.size {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}
