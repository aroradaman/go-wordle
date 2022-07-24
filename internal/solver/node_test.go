package solver

import (
	"testing"
)

func TestNode_Equals(t *testing.T) {
	var n, m, p *Node
	n = newNode('A')
	m = newNode('A')
	p = newNode('B')

	if !n.Equals(m) {
		t.Errorf("Node(A) & Node(A) are equal")
	}

	if n.Equals(p) || m.Equals(p) {
		t.Errorf("Node(A) & Node(B) are not equal")
	}

	// Node for word: AWAIT & AWAKE
	n = newNode('A')
	n.children['W'] = newNode('W')
	n.children['W'].children['A'] = newNode('A')

	n.children['W'].children['A'].children['I'] = newNode('I')
	n.children['W'].children['A'].children['I'].children['T'] = newNode('T')

	n.children['W'].children['A'].children['K'] = newNode('K')
	n.children['W'].children['A'].children['K'].children['E'] = newNode('E')

	// Node for word: AWAIT & AWAKE
	m = newNode('A')
	m.children['W'] = newNode('W')
	m.children['W'].children['A'] = newNode('A')

	m.children['W'].children['A'].children['I'] = newNode('I')
	m.children['W'].children['A'].children['I'].children['T'] = newNode('T')

	m.children['W'].children['A'].children['K'] = newNode('K')
	m.children['W'].children['A'].children['K'].children['E'] = newNode('E')

	// Node for word: AWAIT
	p = newNode('A')
	p.children['W'] = newNode('W')
	p.children['W'].children['A'] = newNode('A')
	p.children['W'].children['A'].children['I'] = newNode('I')
	p.children['W'].children['A'].children['I'].children['T'] = newNode('T')

	if !n.Equals(m) {
		t.Errorf("Node(AWAIT, AWAKE) & Node(AWAIT, AWAKE) are equal")
	}

	if n.Equals(p) || m.Equals(p) {
		t.Errorf("NODE(AWAIT, AWAKE) & NODE(AWAIT) are not equal")
	}

}
