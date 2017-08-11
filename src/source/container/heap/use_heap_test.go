package heap

import (
	"testing"
	"container/heap"
	"fmt"
)

// customer integer type
type IntHeap []int

// implement sort method . why not using pointer because we are not changing the underly data
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i , j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap (i, j int) {h[i],h[j] = h[j] , h[i]}

// implement heap method.
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length . not just its contents
	*h = append(*h,x.(int))
}
func (h *IntHeap) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TestHeap(t *testing.T){
	h := IntHeap{2,1,5}
	heap.Init(&h)
	heap.Push(&h,3)
	heap.Push(&h,100)
	for h.Len() > 0 {
		fmt.Printf(" %v " ,h.Pop())
	}
}
