// Package splaytree implements the splay tree
// data structure, a self-balancing binary
// search tree with amortized logarithmic time
// operations.
package splaytree

// Value is any comparable type of value.
type Value interface {
	// Compare returns 1 if the callee
	// is greater than v2, -1 if it is
	// less than v2, or 0 if they are equal.
	Compare(v2 Value) int
}

// Node is a node in a tree.
type Node struct {
	Value Value
	Left  *Node
	Right *Node
}

// Tree is a splay tree, which consists of
// a single root node.
type Tree struct {
	Root *Node
}

// Insert inserts the value into the tree.
// It is possible to have multiple copies
// of the same value in a tree at once.
func (t *Tree) Insert(v Value) {
	if t.Root == nil {
		t.Root = &Node{Value: v}
		return
	}

	splay(&t.Root, v)

	comparison := v.Compare(t.Root.Value)
	newNode := &Node{Value: v}
	if comparison < 0 {
		newNode.Left = t.Root.Left
		t.Root.Left = nil
		newNode.Right = t.Root
	} else if comparison > 0 {
		newNode.Right = t.Root.Right
		t.Root.Right = nil
		newNode.Left = t.Root
	} else {
		newNode.Left = t.Root.Left
		newNode.Right = t.Root
		t.Root.Left = nil
	}
	t.Root = newNode
}

// Delete deletes one instance of v appearing
// in the tree.
// If the value v is present in the tree more
// than once, only one instance is deleted.
func (t *Tree) Delete(v Value) {
	if t.Root == nil {
		return
	}
	splay(&t.Root, v)

	if v.Compare(t.Root.Value) != 0 {
		return
	}

	if t.Root.Left == nil {
		t.Root = t.Root.Right
	} else {
		splay(&t.Root.Left, greatestValue{})
		t.Root.Left.Right = t.Root.Right
		t.Root = t.Root.Left
	}
}

// Height returns the height of the tree,
// where 0 indicates that the tree is empty.
func (t *Tree) Height() int {
	return t.Root.height()
}

// Min gets the leftmost value.
func (t *Tree) Min() Value {
	if t.Root == nil {
		return nil
	}
	n := t.Root
	for n.Left != nil {
		n = n.Left
	}
	return n.Value
}

// Max gets the rightmost value.
func (t *Tree) Max() Value {
	if t.Root == nil {
		return nil
	}
	n := t.Root
	for n.Right != nil {
		n = n.Right
	}
	return n.Value
}

func (n *Node) height() int {
	if n == nil {
		return 0
	}
	h1 := n.Left.height()
	h2 := n.Right.height()
	if h1 > h2 {
		return 1 + h1
	} else {
		return 1 + h2
	}
}

// splay searches for a value v and
// splays the tree as it does so.
// If the value is found, the root
// will be set to a node with that
// value.
func splay(root **Node, v Value) {
	if (*root) == nil {
		return
	}
	n := (*root)
	comparison := v.Compare(n.Value)
	if comparison < 0 {
		if n.Left == nil {
			return
		}
		n1 := n.Left
		comparison = v.Compare(n1.Value)
		if comparison > 0 && n1.Right != nil {
			splay(&n1.Right, v)
			n2 := n1.Right
			n.Left = n2.Right
			n2.Right = n
			n1.Right = n2.Left
			n2.Left = n1
			(*root) = n2
		} else if comparison < 0 && n1.Left != nil {
			splay(&n1.Left, v)
			n2 := n1.Left
			n.Left = n1.Right
			n1.Right = n
			n1.Left = n2.Right
			n2.Right = n1
			(*root) = n2
		} else {
			(*root) = n1
			n.Left = n1.Right
			n1.Right = n
		}
	} else if comparison > 0 {
		if n.Right == nil {
			return
		}
		n1 := n.Right
		comparison = v.Compare(n1.Value)
		if comparison < 0 && n1.Left != nil {
			splay(&n1.Left, v)
			n2 := n1.Left
			n.Right = n2.Left
			n2.Left = n
			n1.Left = n2.Right
			n2.Right = n1
			(*root) = n2
		} else if comparison > 0 && n1.Right != nil {
			splay(&n1.Right, v)
			n2 := n1.Right
			n.Right = n1.Left
			n1.Left = n
			n1.Right = n2.Left
			n2.Left = n1
			(*root) = n2
		} else {
			(*root) = n1
			n.Right = n1.Left
			n1.Left = n
		}
	}
}

type greatestValue struct{}

func (_ greatestValue) Compare(v Value) int {
	return 1
}

type smallestValue struct{}

func (_ smallestValue) Compare(v Value) int {
	return -1
}
