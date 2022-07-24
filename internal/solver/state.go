package solver

import (
	"fmt"
	. "github.com/daman1807/go-wordle/pkg/container"
	. "github.com/daman1807/go-wordle/pkg/set"
	"reflect"
)

type PositionalState struct {
	fixed    bool
	excludes *Set
	val      rune
}

type State struct {
	includes         *Container
	positionalStates []*PositionalState
}

func newPositionalState() *PositionalState {
	return &PositionalState{
		fixed:    false,
		excludes: NewSet(),
	}
}

func newState() *State {
	positionalStates := make([]*PositionalState, LENGTH)
	for i := 0; i < LENGTH; i++ {
		positionalStates[i] = newPositionalState()
	}
	return &State{
		includes:         NewContainer(),
		positionalStates: positionalStates,
	}
}

func (s *State) Print() {
	fmt.Println("----------------------------------")
	fmt.Println("State")
	fmt.Printf("\tIncludes: %s\n", s.includes.String())
	fmt.Println("\tPositional States:")

	for i := 0; i < len(s.positionalStates); i++ {
		fmt.Printf("\t\tPosition: %d\n", i)
		fmt.Printf("\t\tFixed: %t\n", s.positionalStates[i].fixed)
		fmt.Printf("\t\tVal: %s\n", string(s.positionalStates[i].val))
		if s.positionalStates[i].excludes != nil {
			fmt.Printf("\t\tExcludes: %s\n", s.positionalStates[i].excludes.String())
		} else {
			fmt.Printf("\t\tExcludes: \n")
		}

		fmt.Println()

	}
	fmt.Println("----------------------------------")
}

func (s *State) Update(attempt string, feedback Feedback) {
	excludes := make([]rune, 0)
	attemptContainer := NewContainer()
	for _, r := range attempt {
		attemptContainer.Add(r)
	}

	for i, r := range attempt {

		switch feedback[i].annotation {
		case Fixed:
			s.positionalStates[i].fixed = true
			s.positionalStates[i].val = r
		case Included:
			s.includes.Add(r)
			s.includes.UpdateCount(r, IntMin(s.includes.GetCount(r), attemptContainer.GetCount(r)))
			s.positionalStates[i].excludes.Add(r)
		case Excluded:
			if s.includes.GetCount(r) == 0 {
				excludes = append(excludes, r)
			}
		}
	}

	for i := 0; i < len(attempt); i++ {
		if s.positionalStates[i].fixed {
			s.positionalStates[i].excludes = nil
		} else {
			s.positionalStates[i].excludes.Add(excludes...)
		}
	}
}

func (s *State) Equals(t *State) bool {
	return reflect.DeepEqual(s, t)
}
