package utilpq

import (
	//"fmt"
	"github.com/priya23/finalpq"
	"github.com/priya23/finalpq/implementheap"
)

func CreateHeap() *implementheap.PriorityQueue {
	v1 := implementheap.CreateHeap()
	return v1
}

func CreateNewNode(val int, prior int) *implementheap.Item {
	k := implementheap.CreateNew(val, prior)
	return k
}

func (l *finalpq.PQ) Push(val string, prior int) {
	item := implementheap.CreateNew(val, prior)
	l.Push(item)
}

/*func (k *finalpq.PQ) NewCreate() {

}*/
