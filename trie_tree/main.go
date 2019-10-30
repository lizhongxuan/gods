package main

import "fmt"

type TrieNode struct {
	data         int32
	children     [26]*TrieNode
	isEndingChar bool
}

func main() {
	root := NewTrieTree('/')
	insertNode(root, "abc")
	fmt.Println(findNode(root, "abc"))
}

func NewTrieTree(data int32) *TrieNode {
	return &TrieNode{
		data:         data,
		isEndingChar: false,
	}
}

func insertNode(root *TrieNode, str string) {
	for _, s := range str {
		fmt.Println(s)
		index := s - 'a'
		if root.children[index] == nil {
			newNode := NewTrieTree(s)
			root.children[index] = newNode
		}
		root = root.children[index]
	}
	root.isEndingChar = true
}

func findNode(root *TrieNode, str string) bool {
	for _, s := range str {
		index := s - 'a'
		if root.children[index] == nil {
			return false
		}
		root = root.children[index]
	}
	if root.isEndingChar == false {
		return false
	}
	return true
}
