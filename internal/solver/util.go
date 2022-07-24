package solver

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
