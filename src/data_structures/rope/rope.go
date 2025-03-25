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

func IsValidRope(n *RopeNode) bool {
	return n.data == "" || n.isLeaf() || n.isNonLeaf()
}

func NewRope(s string) *RopeNode {
	if s == "" {
		return nil
	}
	length := len(s)
	half := length / 2
	rope := &RopeNode{len: length}
	if half == 0 {
		rope.data = s
		rope.len = length
		return rope
	}
	leftString := s[:half]
	rightString := s[half:]
	rope.left = NewRope(leftString)
	rope.right = NewRope(rightString)
	return rope
}

func String(r *RopeNode) string {
	var s string
	if r == nil {
		return s
	}
	if r.data != "" {
		s = s + r.data
	}
	if r.left != nil {
		s = s + String(r.left)
	}
	if r.right != nil {
		s = s + String(r.right)
	}
	return s
}

func IterativeString(root *RopeNode) string {

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
		head := st.Peek()
		if head != nil {
			if head.left != nil {
				st.Push(head.left)
			} else {
				st.Push(head.right)
			}
		} else {
			st.Pop()
		}
	}

	return str
}
