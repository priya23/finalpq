package implementheap

import (
	"container/heap"
	"github.com/priya23/finalpq"
	//"fmt"
)

type item struct {
	val      interface{}
	priority int
	index    int
}

type priorityQueue []*item

//for sort implementation
func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	//fmt.Println("i and j are", i, j)
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	//fmt.Println("in swapi and j are", i, j, pq[i].index, pq[j].index)
	pq[i], pq[j] = pq[j], pq[i]
}

//implementing default interface
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) Update(k interface{}, val interface{}, prior int) {
	i := k.(*item)
	i.val = val
	i.priority = prior
	heap.Fix(pq, i.index)
}

func CreateNew(val interface{}, priority int) interface{} {
	itt := item{val: val, priority: priority}
	return &itt
}

func (pq *priorityQueue) Insert(value string, prior int) {
	i := item{val: value, priority: prior}
	pq.Push(&i)
}

func (pq *priorityQueue) Remove() int {
	returnval := pq.Pop()
	rval := returnval.(*item)
	return rval.priority
}

func CreateHeap() finalpq.PQ {
	p := make(priorityQueue, 0)
	heap.Init(&p)
	return &p
}
