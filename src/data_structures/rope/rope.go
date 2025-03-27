package rope

import (
	"data_structures/stack"
	"fmt"
	"strings"
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
func Rebuild(leaves []*RopeNode) *RopeNode {

	length := len(leaves)

	if length <= 2 {
		var right *RopeNode
		if length == 2 {
			right = leaves[1]
		}
		return &RopeNode{left: leaves[0], right: right, len: leaves[0].len}
	}

	nodes := make([]*RopeNode, 0)

	for ind := 0; ind < length; ind = ind + 2 {
		var right *RopeNode
		if ind+1 < length {
			right = leaves[ind+1]
		}
		nodes = append(nodes, &RopeNode{left: leaves[ind], right: right, len: leaves[ind].len})
	}

	return Rebuild(nodes)
}

/*
*

	Rebalances the root node and returns a new, balanced root node

*
*/
func Rebalance(root *RopeNode) *RopeNode {

	return Rebuild(root.CollectLeaves())
}

/*
*

	Concatenates ropes "a" and "b".

*
*/
func Concatenate(a *RopeNode, b *RopeNode) *RopeNode {

	root := &RopeNode{left: a, right: b}

	return Rebuild(root.CollectLeaves())
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

func (n *RopeNode) Print() {
	fmt.Print(n.data)
	if n.left != nil {
		Print(n.left, 1)
	}
	if n.right != nil {
		Print(n.right, 1)
	}
}

func Print(n *RopeNode, depth int) {
	if n.isLeaf() {
		fmt.Printf("->%s%s\n", strings.Repeat(" ", depth), n.data)
	} else {
		fmt.Printf("<-%s\n", strings.Repeat(" ", depth))
	}
	if n.left != nil {
		Print(n.left, depth+1)
	}
	if n.right != nil {
		Print(n.right, depth+1)
	}
}
