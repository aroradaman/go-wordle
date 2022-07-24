package solver

import "fmt"

type Node struct {
	val      rune
	children map[rune]*Node
}

func newNode(val rune) *Node {
	return &Node{val: val, children: make(map[rune]*Node)}
}

func (n *Node) Equals(m *Node) bool {

	if n.val != m.val || len(n.children) != len(m.children) {
		return false
	}

	for r, _ := range n.children {
		_, ok := m.children[r]
		if !ok {
			return false
		}
		if !n.children[r].Equals(m.children[r]) {
			return false
		}
	}

	return true
}

func (n Node) print(depths ...int) {
	depth := 0
	if len(depths) > 0 {
		depth = depths[0]
	}

	for i := 0; i < depth; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("%s\n", string(n.val))

	for _, child := range n.children {
		child.print(depth + 1)
	}
}
