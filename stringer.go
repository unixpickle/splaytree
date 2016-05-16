package splaytree

import (
	"fmt"
	"strings"
)

func (t *Tree) String() string {
	return t.Root.String()
}

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
