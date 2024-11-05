package main

import (
	"dsa/tree"
	"math/rand"
)


func main() {
	btree := tree.New(50)

	// insert 100 numbers randomly into the tree and let the tree sort itself
	for i := 0; i < 100; i++ {
		randomNumber := rand.Intn(100)
		btree.Add(randomNumber)
	}
	
	// now traverse through the tree
	btree.Traverse()
}