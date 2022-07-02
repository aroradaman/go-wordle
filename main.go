package main

import (
	"encoding/json"
	"os"
)

func main() {
	var words []string

	data, _ := os.ReadFile("words.json")
	_ = json.Unmarshal(data, &words)

	tree := createTree(words)
	play(tree, words)

}
