package solver

const (
	Fixed      = iota // GREEN	// Fixed when the position of the character is fixed
	Included          // YELLOW // Included when character exists but not in this position
	Excluded          // GREY	// Excluded when character is not included at this position
	Unexplored        // NA		// Unexplored initialization, of the characters for exploration
)

type Char struct {
	val        rune
	annotation int
}

func (c Char) GetAnnotation() int {
	return c.annotation
}

func (c Char) GetVal() string {
	return string(c.val)
}

type Feedback []*Char

func NewFeedback(word string) Feedback {
	data := make([]*Char, len(word))
	for i, r := range word {
		data[i] = &Char{
			val:        r,
			annotation: Unexplored,
		}
	}
	return data
}

func (f Feedback) String() string {
	word := make([]rune, len(f))
	for i, char := range f {
		word[i] = char.val
	}
	return string(word)
}

func (f Feedback) Equals(g Feedback) bool {
	if len(f) != len(g) {
		return false
	}

	for i := 0; i < len(f); i++ {
		if f[i].val != g[i].val {
			return false
		}

		if f[i].annotation != g[i].annotation {
			return false
		}
	}

	return true
}
