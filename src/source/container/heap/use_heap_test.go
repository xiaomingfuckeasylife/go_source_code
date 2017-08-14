package heap

import (
	"testing"
	"container/heap"
	"fmt"
)


// heap sort using a tree structure to create a min-heap

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
	h := IntHeap{1,2,5}
	heap.Init(&h)
	heap.Push(&h,3)
	heap.Push(&h,100)
	for h.Len() > 0 {
		fmt.Printf(" %v " ,heap.Pop(&h))
	}
}


// Priority Queue.

type Item struct {
	value string	// The value of the Item ; arbitrary .
	priority int 	// The priority of the Item in the queue
	// The index is needed by update and is maintained by the heap.Interface methods
	index int	// the index of the Item in the heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq)}

// We want Pop to give us the higest , not lowest priority so we use greater than here.
func (pq PriorityQueue) Less(i , j int) bool {return pq[i].priority > pq[j].priority}

// Swap value between Queue
func (pq PriorityQueue) Swap(i , j int) {
	pq[i] , pq[j] = pq[j] , pq[j]
	pq[i].priority = i
	pq[j].priority = j
}

//
func (pq *PriorityQueue) Push(x interface{}){
	n :=pq.Len()
	item := x.(*Item)
	item.index = n
	*pq = append(*pq , item)
}

func (pq *PriorityQueue) Pop() interface{}{
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(Item *Item , value string , priority int) {
	Item.value = value
	Item.priority = priority
	heap.Fix(pq,Item.index)
}

func TestPriorityHeap(t *testing.T){
	// some Items and their priorities
	Items := map[string]int{
		"banana":3,"apple":2,"pear":4,
	}
	//
	pq := make(PriorityQueue,len(Items))
	i := 0
	for value , priority := range Items {
		pq[i] = &Item{
			value:value,
			priority:priority,
			index:i,
		}
		i++
	}
	heap.Init(&pq)
	//
	item := &Item{
		value:"orange",
		priority:1,
	}
	heap.Push(&pq,item)
	pq.update(item,item.value,5)
	//
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s",item.priority,item.value)
	}
}





