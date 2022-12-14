/**
 *
 * @file GeneralTree.go
 * @author Max Base (maxbasecode@gmail.com)
 * @brief General Tree Implementation in Go
 * @version 0.1
 * @date 2022-12-14
 * @copyright Copyright (c) 2022
 *
 */

package main

import "fmt"

type Node struct {
	Value    interface{}
	Children []*Node
}

type GeneralTree struct {
	Root *Node
}

// Create a new general tree
func NewGeneralTree() *GeneralTree {
	return &GeneralTree{}
}

// Add a new node to the tree
func (t *GeneralTree) AddNode(value interface{}) {
	newNode := &Node{Value: value}
	t.Root.Children = append(t.Root.Children, newNode)
}

// Add a new node to the tree
func (t *GeneralTree) AddNodeToParent(value int, parent *Node) {
	newNode := &Node{Value: value}
	parent.Children = append(parent.Children, newNode)
}

// Get the root node of the tree
func (t *GeneralTree) GetRoot() *Node {
	return t.Root
}

// Get the children of a node
func (t *GeneralTree) GetChildren(node *Node) []*Node {
	return node.Children
}

// Get the value of a node
func (t *GeneralTree) GetValue(node *Node) interface{} {
	return node.Value
}

// Set the value of a node
func (t *GeneralTree) SetValue(node *Node, value interface{}) {
	node.Value = value
}

// Get the number of nodes in the tree
func (t *GeneralTree) GetSize() int {
	return t.getSize(t.Root)
}

func (t *GeneralTree) getSize(node *Node) int {
	if node == nil {
		return 0
	}
	size := 1
	for _, child := range node.Children {
		size += t.getSize(child)
	}
	return size
}

// Get the height of the tree
func (t *GeneralTree) getHeight() int {
	if t.Root == nil {
		return 0
	}
	return t.Root.GetHeight(t.Root)
}

func (t *Node) GetHeight(node *Node) int {
	if node == nil {
		return 0
	}
	height := 0
	for _, child := range node.Children {
		child_height := t.GetHeight(child)
		if child_height > height {
			height = child_height
		}
	}
	return height + 1
}

// Print the tree
func (t *GeneralTree) PrintTree() {
	t.printTree(t.Root, 0)
}

func (t *GeneralTree) printTree(node *Node, level int) {
	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}
	fmt.Println(node.Value)
	for _, child := range node.Children {
		t.printTree(child, level+1)
	}
}
