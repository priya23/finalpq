package binomialheap

import (
	"fmt"
)

type BiNode struct {
	key     int
	degree  int
	child   *BiNode
	sibling *BiNode
	parent  *BiNode
}
type Blist struct {
	head *BiNode
	size int
}

func CreateNewHeap() *Blist {
	return &Blist{head: nil, size: 0}
}

func CreateNewNode(keyval int) *BiNode {
	return &BiNode{key: keyval, degree: 0, child: nil, sibling: nil, parent: nil}
}

func (bh *Blist) Insert(newnode *BiNode) {
	bh.size += 1
	bh.addinto(newnode)
}

func (bh *Blist) Give(val string, prior int) {
	newn := CreateNewNode(prior)
	bh.Insert(newn)
}
func (bh *Blist) Take() int {
	return bh.Pop()
}
func (bh *Blist) Pop() int {
	bh.size -= 1
	track := bh.head
	minival := bh.head
	if minival == nil {
		return minival.key
	}
	for track != nil {
		if minival.key > track.key {
			minival = track
		}
		track = track.sibling
	}
	removeFromHead(&bh.head, minival)
	for _, children := range iterator(minival.child) {
		removeFromHead(&minival.child, children)
		bh.addinto(children)
	}
	return minival.key
}
func (bf *Blist) addinto(newnode *BiNode) {
	//fmt.Println("inside add")
	samedegrenode := checkForSameDegree(bf.head, newnode.degree)
	fmt.Println("\nsame degeree returned is %v", samedegrenode)
	if samedegrenode == nil {
		insertToForest(&bf.head, newnode)
	} else {
		fmt.Println("inside same degree condition")
		removeFromHead(&bf.head, samedegrenode)
		freshnode := joinnode(samedegrenode, newnode)
		bf.addinto(freshnode)
	}
	fmt.Println("/n head of list is %v", bf.head)
}

func joinnode(n1 *BiNode, n2 *BiNode) *BiNode {
	if n1.degree < n2.degree {
		n1.degree += 1
		n1.adopt(n2)
		return n1
	} else {
		n2.degree += 1
		n2.adopt(n1)
		return n2
	}
}

func (n1 *BiNode) adopt(n2 *BiNode) {
	insertToForest(&n1.sibling, n2)
	n2.parent = n1
}
func removeFromHead(bf **BiNode, samedegreenode *BiNode) {
	tracknode := getleftnode(*bf, samedegreenode)
	if tracknode == nil {
		*bf = samedegreenode.sibling
		fmt.Println("INSIDE sibling")
	} else {
		tracknode.sibling = samedegreenode.sibling
	}
	samedegreenode.sibling = nil
}

func getleftnode(head *BiNode, node *BiNode) *BiNode {
	if head == node {
		return nil
	}

	checknode := head

	for checknode.sibling != node {
		checknode = checknode.sibling
	}

	return checknode
}
func checkForSameDegree(bf *BiNode, newdegree int) *BiNode {
	cnode := bf
	fmt.Println("\n value of bf is %v", bf)
	for cnode != nil {
		if cnode.degree == newdegree {
			return cnode
		}
		fmt.Println("cnode is %v", cnode)
		cnode = cnode.sibling
	}
	return nil
}

func insertToForest(bf **BiNode, newnode *BiNode) {
	var prev *BiNode
	var next *BiNode
	prev = nil
	next = *bf
	for next != nil && newnode.degree < next.degree {
		prev = next
		next = next.sibling
	}
	if prev == nil && next == nil {
		fmt.Println("inside both nil")
		*bf = newnode
	} else if prev != nil && next == nil {
		//fmt.Println("INSIDE 1")
		prev.sibling = newnode
	} else if prev == nil && next != nil {
		newnode.sibling = *bf
		*bf = newnode
	} else if prev != nil && next != nil {
		prev.sibling = newnode
		newnode.sibling = next
	}
}

func iterator(parent *BiNode) []*BiNode {
	arr := make([]*BiNode, 0)
	track := parent
	for track != nil {
		arr = append(arr, track)
		track = track.sibling
	}
	fmt.Println("/narray is %v", arr)
	return arr
}
func (n1 *BiNode) PrintNode() {
	fmt.Printf("/n KEY: %d degree: %d sibling is %v", n1.key, n1.degree, n1.sibling)
}

func (bf *BiNode) Print_Level() {
	fmt.Printf("bf is %v", bf)
	bf.PrintNode()

}

func (bh *Blist) PrintValue() {
	if bh.head == nil {
		fmt.Print("heap is empty.")
	}
	fmt.Println("\n head is %v", bh.head)
	for _, node := range iterator(bh.head) {
		fmt.Printf("value of node is %v", node)
		node.Print_Level()
	}
}
