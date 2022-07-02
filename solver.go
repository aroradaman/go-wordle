package main

import "math/rand"

type Node struct {
	val      rune
	children map[rune]Node
}

func notIn(index int, indices []int) bool {
	for i := 0; i < len(indices); i++ {
		if indices[i] == index {
			return false
		}
	}
	return true
}

func pickOne(validAttempts []string, attemptHistory []string) string {
	var attempt string
	found := false
	for !found {
		found = true
		attempt = validAttempts[rand.Intn(len(validAttempts))]
		for i := 0; i < len(attemptHistory); i++ {
			if attemptHistory[i] == attempt {
				found = false
				break
			}
		}
	}
	return attempt
}

func createNode(val rune) Node {
	return Node{
		val:      val,
		children: make(map[rune]Node, 0),
	}
}

func createTree(words []string) Node {
	var parent *Node
	tree := createNode('.')

	for _, word := range words {
		parent = &tree
		for _, r := range word {
			child, ok := parent.children[r]
			if !ok {
				parent.children[r] = createNode(r)
				child = parent.children[r]
			}
			parent = &child
		}
	}
	return tree
}

func pruneAndGuess(node *Node, fixed map[rune]int, includes map[rune][]int, excludes []rune, attempt []rune, depth int, length int, validAttempts *[]string) {
	if len(node.children) == 0 && depth == length {
		containsAllYellow := true
		// direct exclusion
		for _, r := range attempt {
			for _, er := range excludes {
				if r == er {
					return

				}
			}
		}

		for cr, ignore := range includes {
			containsYellow := false
			for i, r := range attempt {
				if r == cr && notIn(i, ignore) {
					containsYellow = true
					break
				}
			}
			if !containsYellow {
				containsAllYellow = false
				break
			}
		}
		if containsAllYellow {
			*validAttempts = append(*validAttempts, string(attempt))
		}

	} else {
		for _, child := range node.children {
			prune := false

			for r, d := range fixed {
				if r != child.val && d == depth {
					prune = true
					break
				}
			}
			if prune {
				delete(node.children, child.val)
			} else {
				branchAttempt := make([]rune, len(attempt))
				copy(branchAttempt, attempt)
				branchAttempt = append(branchAttempt, child.val)
				pruneAndGuess(&child, fixed, includes, excludes, branchAttempt, depth+1, length, validAttempts)
			}
		}
	}
}
