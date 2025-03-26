package rope

import (
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

func (n *RopeNode) Left() *RopeNode {
	return n.left
}

func (n *RopeNode) Right() *RopeNode {
	return n.right
}

func (n *RopeNode) Data() string {
	return n.data
}

func (n *RopeNode) Weight() int {
	weight := 0
	if n.left != nil {
		leavesOnLeftChild := n.left.CollectLeaves()
		for _,leaf := range leavesOnLeftChild {
			weight += leaf.len
		}
	}
	return weight
}

/**
	Collects all leaves starting from node.
**/
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

func (root *RopeNode) String() string {

	if root == nil {
		return ""
	}

	if root.left == nil && root.right == nil {
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

func NewRope(s string) *RopeNode {
	if s == "" {
		return nil
	}
	length := len(s)
	half := length / 2
	rope := &RopeNode{len: half}
	if half == 0 {
		rope.data = s
		rope.len = 1
		return rope
	}
	leftString := s[:half]
	rightString := s[half:]
	rope.left = NewRope(leftString)
	rope.right = NewRope(rightString)
	return rope
}

func IsValidRope(n *RopeNode) bool {
	return n.data == "" || n.isLeaf() || n.isNonLeaf()
}

/**
	Concatenates ropes "a" and "b".
**/
func Concatenate(a *RopeNode, b *RopeNode) *RopeNode {

	root := &RopeNode{left: a, right: b}

	return root
}

/**
	Splits "rope" at position "pos" into two new ropes.
**/
func Split(rope *RopeNode, pos int) (*RopeNode, *RopeNode) {
	return nil, nil
}

func (n *RopeNode) Insert(pos int, str string) {

}

/**
	Delete a substring from a specified position and length.
**/
func (n *RopeNode) Delete(start int, length int) {

}

