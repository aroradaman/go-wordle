package solver

import (
	. "github.com/daman1807/go-wordle/pkg/container"
	"math/rand"
	"time"
)

var LENGTH = 5

type Solver struct {
	tree   *Node
	answer string
	state  *State
}

func NewSolver(words []string, answer string) *Solver {
	rand.Seed(time.Now().UnixNano())
	solver := Solver{
		tree:   newNode('.'),
		answer: answer,
		state:  newState(),
	}
	solver.build(words)
	return &solver
}

func (s Solver) Print() {
	s.tree.print()
}

func (s *Solver) build(words []string) {
	exists := false
	var parent, child *Node

	for _, word := range words {
		parent = s.tree

		for _, r := range word {
			child, exists = parent.children[r]
			if !exists {
				child = newNode(r)
				parent.children[r] = child
			}
			parent = child
		}
	}
}

// Validate
/*
Generates a feedback, if the attempt has repeating character, then it will only be annotated
Fixed & Included only the number of times it appears in the answer, rest occurrences will be
annotated as Excluded
*/
func (s *Solver) Validate(attempt string) (Feedback, bool) {
	feedback := NewFeedback(attempt)

	answerContainer := NewContainer()
	for _, r := range s.answer {
		answerContainer.Add(r)
	}
	for i := 0; i < len(attempt); i++ {
		a := rune(attempt[i])
		s := rune(s.answer[i])

		if a == s {
			feedback[i].annotation = Fixed
			answerContainer.Pop(a)
		} else {
			feedback[i].annotation = Excluded
		}
	}

	for i := 0; i < len(attempt); i++ {
		a := rune(attempt[i])
		if feedback[i].annotation != Fixed && answerContainer.Contains(a) {
			answerContainer.Pop(a)
			feedback[i].annotation = Included
		}
	}

	s.Update(attempt, feedback)
	return feedback, attempt == s.answer
}

// update prunes branches based on feedback
func (s *Solver) update(node *Node, depth int) {
	for r, child := range node.children {
		if s.state.positionalStates[depth].fixed && s.state.positionalStates[depth].val != r {
			delete(node.children, r)
		} else if s.state.positionalStates[depth].excludes != nil && s.state.positionalStates[depth].excludes.Contains(r) {
			delete(node.children, r)
		} else {
			s.update(child, depth+1)
		}
	}
}

func (s *Solver) Update(attempt string, feedback Feedback) {
	s.state.Update(attempt, feedback)
	s.update(s.tree, 0)
}

func (s Solver) getHints(node *Node, path []rune, hints *[]string) {
	if len(path) == LENGTH {
		container := s.state.includes.Copy()
		for _, r := range path {

			container.Pop(r)

		}
		if container.IsEmpty() {
			*hints = append(*hints, string(path))
		}
	} else {
		for r, child := range node.children {
			nextPath := make([]rune, len(path))
			copy(nextPath, path)

			nextPath = append(nextPath, r)
			s.getHints(child, nextPath, hints)
		}
	}
}

func (s Solver) GetHints() []string {
	hints := make([]string, 0)
	s.getHints(s.tree, make([]rune, 0), &hints)

	for i := 0; i < 10; i++ {
		rand.Shuffle(len(hints), func(i, j int) { hints[i], hints[j] = hints[j], hints[i] })
	}

	return hints
}

func (s Solver) IsValidAttempt(attempt string) bool {
	parent := s.tree
	for _, r := range attempt {
		child, exists := parent.children[r]
		if !exists {
			return false
		}
		parent = child
	}
	return true
}
