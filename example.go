package main

import "fmt"

func main() {
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
	for _, c := range child {
		fmt.Println("\tChildren of the root node: ", c.Value)
	}

	// Get the value of the root node
	fmt.Println("Value of the root node: ", tree.GetValue(tree.Root))

	// Set the value of the root node
	tree.SetValue(tree.Root, 10)

	// Get the value of the root node
	fmt.Println("Value of the root node: ", tree.GetValue(tree.Root))
}
