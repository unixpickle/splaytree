package splaytree

import (
	"math/rand"
	"testing"
)

func TestInsertions(t *testing.T) {
	tree := &Tree{}
	allValues := map[int]int{}
	for i := 0; i < 2000; i++ {
		n := rand.Intn(1000)
		allValues[n]++
		tree.Insert(NumValue(n))
		if !properlySorted(tree.Root) {
			t.Fatal("tree not properly sorted")
		}
		if !containsValues(tree, allValues) {
			t.Fatal("tree does not contain correct values")
		}
	}
}

func TestDeletions(t *testing.T) {
	tree := &Tree{}
	var valueCount int
	allValues := map[int]int{}

	for i := 0; i < 2000; i++ {
		valueCount++
		n := rand.Intn(1000)
		allValues[n]++
		tree.Insert(NumValue(n))
	}

	for _, x := range []int{1001, -50, 10002} {
		tree.Delete(NumValue(x))
		if !properlySorted(tree.Root) {
			t.Fatal("tree not properly sorted")
		}
		if !containsValues(tree, allValues) {
			t.Fatal("tree does not contain correct values")
		}
	}

	for valueCount > 0 {
		for k, c := range allValues {
			if c == 0 {
				continue
			}
			valueCount--
			allValues[k]--
			tree.Delete(NumValue(k))
			if !properlySorted(tree.Root) {
				t.Fatal("tree not properly sorted")
			}
			if !containsValues(tree, allValues) {
				t.Fatal("tree does not contain correct values")
			}
		}
	}

	tree.Insert(NumValue(87))
	tree.Insert(NumValue(800))
	tree.Insert(NumValue(900))
	tree.Delete(NumValue(87))
	exp := map[int]int{800: 1, 900: 1}
	if !properlySorted(tree.Root) || !containsValues(tree, exp) {
		t.Fatal("tree does not handle left-most deletions")
	}
}

func TestMinMax(t *testing.T) {
	tree := &Tree{}
	var min, max int
	for i := 0; i < 100; i++ {
		n := rand.Intn(10000) + ((i*17 + 29) % 13)
		tree.Insert(NumValue(n))
		if i == 0 || n < min {
			min = n
		}
		if i == 0 || n > max {
			max = n
		}
	}

	actualMin := int(tree.Min().(NumValue))
	actualMax := int(tree.Max().(NumValue))
	if actualMin != min {
		t.Errorf("min should be %d but got %d", min, actualMin)
	}
	if actualMax != max {
		t.Errorf("max should be %d but got %d", max, actualMax)
	}
}

func BenchmarkInsertions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree Tree
		for i := 0; i < 1000; i++ {
			tree.Insert(NumValue((17 * i) % 337))
		}
	}
}

func properlySorted(t *Node) bool {
	if t == nil {
		return true
	}
	if !properlySorted(t.Left) || !properlySorted(t.Right) {
		return false
	}
	if t.Left != nil && t.Left.Value.Compare(t.Value) > 0 {
		return false
	}
	if t.Right != nil && t.Right.Value.Compare(t.Value) < 0 {
		return false
	}
	return true
}

func containsValues(t *Tree, vals map[int]int) bool {
	subMap := map[int]int{}
	for x, y := range vals {
		subMap[x] = y
	}
	subtractNodeValues(t.Root, subMap)
	for _, v := range subMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func subtractNodeValues(n *Node, vals map[int]int) {
	if n == nil {
		return
	}
	vals[int(n.Value.(NumValue))]--
	subtractNodeValues(n.Left, vals)
	subtractNodeValues(n.Right, vals)
}

type NumValue int

func (n NumValue) Compare(x Value) int {
	xv := x.(NumValue)
	if n > xv {
		return 1
	} else if n < xv {
		return -1
	} else {
		return 0
	}
}
