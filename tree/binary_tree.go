package tree

import "fmt"

type Node struct {
	Parent *Node
	Data int
	Left *Node
	Right *Node
}

func New(data int) *Node {
	return &Node{Data: data}
}

func (root *Node) Add(data int) {
	if data != root.Data {
		if data < root.Data {
			if root.Left == nil {
				root.Left = &Node{Data: data}
			} else {
				root.Left.Add(data)
			}
		} else if data > root.Data {
			if root.Right == nil {
				root.Right = &Node{Data: data}
			} else {
				root.Right.Add(data)
			}
		}
	}
}

func (root *Node) Traverse() {
	if root != nil {
		root.Left.Traverse()
		fmt.Print(" ", root.Data)
		root.Right.Traverse()
	}
}