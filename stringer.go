package splaytree

import (
	"fmt"
	"strings"
)

// String returns a human-readable,
// indendentation-based representation of
// the tree.
// Values are formatted using Go's fmt
// package.
func (t *Tree) String() string {
	return t.Root.String()
}

// String behaves much like Tree.String(),
// only for a node within a tree.
func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%v\n%s\n%s", n.Value,
		indentString(n.Left.String()),
		indentString(n.Right.String()))
}

func indentString(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = "  " + line
	}
	return strings.Join(lines, "\n")
}
