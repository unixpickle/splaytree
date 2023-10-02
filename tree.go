// Package splaytree implements the splay tree
// data structure, a self-balancing binary
// search tree with amortized logarithmic time
// operations.
package splaytree

// Value is an abstract comparable type of value.
type Value[T any] interface {
	// Compare returns 1 if the callee
	// is greater than v2, -1 if it is
	// less than v2, or 0 if they are equal.
	Compare(v2 T) int
}

// Node is a node in a tree.
type Node[T Value[T]] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// Tree is a splay tree, which consists of
// a single root node.
type Tree[T Value[T]] struct {
	Root *Node[T]
}

// Insert inserts the value into the tree.
// It is possible to have multiple copies
// of the same value in a tree at once.
func (t *Tree[T]) Insert(v T) {
	if t.Root == nil {
		t.Root = &Node[T]{Value: v}
		return
	}

	splay(&t.Root, v)

	comparison := v.Compare(t.Root.Value)
	newNode := &Node[T]{Value: v}
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
func (t *Tree[T]) Delete(v Value[T]) {
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
		splay(&t.Root.Left, greatestValue[T]{})
		t.Root.Left.Right = t.Root.Right
		t.Root = t.Root.Left
	}
}

// Height returns the height of the tree,
// where 0 indicates that the tree is empty.
func (t *Tree[T]) Height() int {
	return t.Root.height()
}

// Min gets the leftmost value, or the zero value of T if
// the tree is empty.
func (t *Tree[T]) Min() T {
	if t.Root == nil {
		var zeroVal T
		return zeroVal
	}
	n := t.Root
	for n.Left != nil {
		n = n.Left
	}
	return n.Value
}

// Max gets the rightmost value, or the zero value of T if
// the tree is empty.
func (t *Tree[T]) Max() T {
	if t.Root == nil {
		var zeroVal T
		return zeroVal
	}
	n := t.Root
	for n.Right != nil {
		n = n.Right
	}
	return n.Value
}

// Len gets the number of values stored in the tree.
func (t *Tree[T]) Len() int {
	return t.Root.Len()
}

// Iterate iterates over elements in the tree from least to greatest.
// Calls f for each element, breaking the loop if f returns false.
//
// It is not safe to modify the tree during iteration.
func (t *Tree[T]) Iterate(f func(T) bool) {
	t.Root.Iterate(f)
}

func (n *Node[T]) height() int {
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

func (n *Node[T]) Len() int {
	if n == nil {
		return 0
	} else {
		return 1 + n.Left.Len() + n.Right.Len()
	}
}

func (n *Node[T]) Iterate(f func(T) bool) bool {
	if n == nil {
		return true
	}
	if n.Left != nil {
		if !n.Left.Iterate(f) {
			return false
		}
	}
	if !f(n.Value) {
		return false
	}
	if n.Right != nil {
		return n.Right.Iterate(f)
	}
	return true
}

// splay searches for a value v and
// splays the tree as it does so.
// If the value is found, the root
// will be set to a node with that
// value.
func splay[T Value[T], V Value[T]](root **Node[T], v V) {
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

type greatestValue[T Value[T]] struct{}

func (_ greatestValue[T]) Compare(v T) int {
	return 1
}
