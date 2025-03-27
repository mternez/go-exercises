package rope

import (
	"data_structures/queue"
	"data_structures/stack"
)

type RopeNode struct {
	len   int
	left  *RopeNode
	right *RopeNode
	data  string
}

func (n *RopeNode) isLeaf() bool {
	return n.len > 0 && n.len == len(n.data) && n.left == nil && n.right == nil
}

func (n *RopeNode) isNonLeaf() bool {
	return n.len == (n.right.len+n.left.len) && n.left != nil && n.right != nil && IsValidRope(n.left) && IsValidRope(n.right)
}

func (n *RopeNode) Data() string {
	return n.data
}

func (n *RopeNode) weight() int {
	weight := 0
	if n.left != nil {
		leavesOnLeftChild := n.left.CollectLeaves()
		for _, leaf := range leavesOnLeftChild {
			weight += leaf.len
		}
	}
	return weight
}

/*
*

	Collects all leaves starting from node.

*
*/
func (node *RopeNode) CollectLeaves() []*RopeNode {

	leaves := make([]*RopeNode, 0)

	st := &stack.Stack[*RopeNode]{}

	st.Push(node)

	for !st.Empty() {
		head := *st.Pop()
		if head != nil {
			if head.isLeaf() {
				leaves = append(leaves, head)
			}
			if head.right != nil {
				st.Push(head.right)
			}
			if head.left != nil {
				st.Push(head.left)
			}
		}
	}
	return leaves
}

func (root *RopeNode) ToString() string {

	if root == nil || (root.left == nil && root.right == nil) {
		return ""
	}

	var str string
	st := &stack.Stack[*RopeNode]{}
	st.Push(root)
	for !st.Empty() {
		head := *st.Pop()
		if head != nil {
			if head.isLeaf() {
				str += head.data
			} else {
				if head.right != nil {
					st.Push(head.right)
				}
				if head.left != nil {
					st.Push(head.left)
				}
			}
		}
	}
	return str
}

func NewRope(s string, minNodeSize int) *RopeNode {

	if s == "" {
		return nil
	}

	length := len(s)
	half := length / 2

	// Is a Leaf
	if length == minNodeSize || length/minNodeSize == 0 || half < minNodeSize {
		return &RopeNode{data: s, len: length}
	}

	// Is not a Leaf
	return &RopeNode{left: NewRope(s[:half], minNodeSize), right: NewRope(s[half:], minNodeSize)}
}

func IsValidRope(n *RopeNode) bool {
	return n.data == "" || n.isLeaf() || n.isNonLeaf()
}

/*
*

	Rebuilds a Rope from an ordered collection of leaves

*
*/
func Rebalance(leaves []*RopeNode) *RopeNode {

	length := len(leaves)

	q := &queue.Queue[*RopeNode]{}

	for ind := 0; ind < length; ind++ {

		node := &RopeNode{left: leaves[ind], len: leaves[ind].len}

		if ind+1 < length {
			node.right = leaves[ind+1]
		}

		q.Push(node)
	}

	for q.Size() > 2 {

		left := *q.Pop()

		node := &RopeNode{left: left, len: left.weight()}

		if q.Peek() != nil {
			node.right = *q.Pop()
		}

		q.Push(node)
	}

	root := &RopeNode{left: *q.Pop()}
	if !q.Empty() {
		root.right = *q.Pop()
	}

	root.len = root.left.len

	return root
}

/*
*

	Concatenates ropes "a" and "b".

*
*/
func Concatenate(a *RopeNode, b *RopeNode) *RopeNode {

	root := &RopeNode{left: a, right: b}

	return root
}

/*
*

	Splits "rope" at position "pos" into two new ropes.

*
*/
func Split(rope *RopeNode, pos int) (*RopeNode, *RopeNode) {
	return nil, nil
}

func (n *RopeNode) Insert(pos int, str string) {

}

/*
*

	Delete a substring from a specified position and length.

*
*/
func (n *RopeNode) Delete(start int, length int) {

}
