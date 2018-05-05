package segment_tree

import "fmt"

type Node struct {
	low   int
	high  int
	val   int
	left  *Node
	right *Node
}

type SegTree struct {
	root *Node
}

//        [36]
//    [10]     [26]
//  [3   7   11  15]
// [1,2,3,4,5,6,7,8]
func (s *SegTree) Build(arr []int) {
	s.root = build(0, len(arr)-1, arr)
}

func (s *SegTree) Print() {
	print(s.root)
}

func print(node *Node) {
	if node == nil {
		return
	}
	print(node.left)
	fmt.Printf("(%v, %v, %v)\n", node.low, node.high, node.val)
	print(node.right)
}
func build(low, high int, arr []int) *Node {
	node := &Node{}
	if low == high {
		node.val = arr[low]

	} else {
		mid := (low + high) / 2
		node.left = build(low, mid, arr)
		node.right = build(mid+1, high, arr)
		node.val = node.left.val + node.right.val
	}
	node.low = low
	node.high = high
	return node
}

// Query is query
func (s *SegTree) Query(node *Node, start, end int) int {
	fmt.Printf("%v %v %v %v\n", node.low, node.high, start, end)
	if node == nil {
		return 0
	}
	if end < node.low || start > node.high {
		return 0
	} else if start <= node.low && end >= node.high {
		return node.val
	} else {
		return s.Query(node.left, start, end) + s.Query(node.right, start, end)
	}
}
