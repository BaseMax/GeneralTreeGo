# General Tree Go

Implentation of a general tree in **Go Programming Language**. A general tree is a tree data structure in which each node can have an arbitrary number of children. We are storing children in an array.

## Structure

**Types:**

```go
type Node struct 
	Value    interface{}
	Children []*Node
}

type GeneralTree struct 
	Root *Node
}
```

**Functions:**

```go
// Create a new general tree
func NewGeneralTree() *GeneralTree 
	return &GeneralTree{}
}

// Create a new node
func NewNode() *Node 
	return &Node{}
}
```

**Methods:**

```go
// Add a new node to the tree
func (t *GeneralTree) AddNode(value interface{})

// Add a new node to the tree
func (t *GeneralTree) AddNodeToParent(value int, parent *Node)

// Get the root node of the tree
func (t *GeneralTree) GetRoot() *Node 

// Get the children of a node
func (t *GeneralTree) GetChildren(node *Node) []*Node 

// Get the value of a node
func (t *GeneralTree) GetValue(node *Node) interface{} 

// Set the value of a node
func (t *GeneralTree) SetValue(node *Node, value interface{})

// Get the number of nodes in the tree
func (t *GeneralTree) GetSize() int 
func (t *GeneralTree) getSize(node *Node) int 

// Get the height of the tree
func (t *GeneralTree) getHeight() int 
func (t *Node) GetHeight(node *Node) int 

// Print the tree
func (t *GeneralTree) PrintTree()
func (t *GeneralTree) printTree(node *Node, level int)
```

## Usage

```go
package main

import "fmt"

func main()
	// Create a new tree
	tree := NewGeneralTree()

	// Add a root node
	tree.Root = &Node{Value: 1}

	// Add a node to the root
	tree.AddNode(2)

	// Add a node to the root
	tree.AddNode(3)

	// Add a node to the root
	tree.AddNode(4)

	// Add a node to the root
	tree.AddNode(5)

	// Add a node to the root
	tree.AddNode(6)

	// Adding a child to the first node
	tree.AddNodeToParent(7, tree.Root.Children[0])

	// Adding a child to the first node
	tree.AddNodeToParent(8, tree.Root.Children[0])

	// Adding a child to the first node
	tree.AddNodeToParent(9, tree.Root.Children[0])

	// Print the tree
	tree.PrintTree()

	// Get the size of the tree
	fmt.Println("Size of the tree: ", tree.GetSize())

	// Get the height of the tree
	fmt.Println("Height of the tree: ", tree.getHeight())

	// Get the root node
	fmt.Println("Root node: ", tree.GetRoot().Value)

	// Get the children of the root node
	child := tree.GetChildren(tree.Root)
	fmt.Println("Children of the root node: ", tree.GetChildren(tree.Root))
	for _, c := range child 
		fmt.Println("\tChildren of the root node: ", c.Value)
	}

	// Get the value of the root node
	fmt.Println("Value of the root node: ", tree.GetValue(tree.Root))

	// Set the value of the root node
	tree.SetValue(tree.Root, 10)

	// Get the value of the root node
	fmt.Println("Value of the root node: ", tree.GetValue(tree.Root))
}
```

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2022, Max Base
