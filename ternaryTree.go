package goTernaryTree

import (
	"errors"
	"strings"
)

// Item represents a single object in the Node.
type Item interface{}

// Node represents a single object in the tree.
type Node struct {
	c                uint8
	left, mid, right *Node
	value            Item
}

//Ternary Tree
type TernaryTree struct {
	size int
	root *Node
}

//Creates new instance of TernaryTree
func New() *TernaryTree {
	return &TernaryTree{size: 0}
}

//Adds new key to the tree.
func (t *TernaryTree) add(key string, value Item) {
	if len(strings.TrimSpace(key)) > 1 {

		if !t.contains(key) {
			t.size++
			t.root = insert(t.root, key, value, 0)
		}
	}

}

//Searches if tree contains key.
func (t *TernaryTree) contains(key string) bool {
	node := search(t.root, key, 0)
	if node == nil {
		return false
	}

	return true
}

//Gets saved key from the tree.
func (t *TernaryTree) get(key string) (Item, error) {
	if len(strings.TrimSpace(key)) < 1 {
		return nil, nil
	}
	node := search(t.root, key, 0)
	if node == nil {
		return nil, errors.New("Key not found.")
	}
	return node.value, nil
}

//Searches tree for keys with common prefix.
func (t *TernaryTree) prefixMatch(prefix string) []Item {
	if t.root == nil {
		return nil
	}
	bucket := []Item{}
	node := search(t.root, prefix, 0)
	if node == nil {
		return nil
	}
	if node.value != nil {
		bucket = append(bucket, prefix)
	}
	collect(node.mid, prefix, &bucket)
	return bucket
}

//Searches tree with a wild card to match  keys.
func (t *TernaryTree) wildcardMatch(pattern string) []Item {
	if t.root == nil {
		return nil
	}
	bucket := []Item{}
	collect2(t.root, "", 0, pattern, &bucket)
	return bucket
}

func search(node *Node, key string, charIndex int) *Node {
	if node == nil {
		return nil
	}
	c := key[charIndex]
	if c < node.c {
		return search(node.left, key, charIndex)
	} else if c > node.c {
		return search(node.right, key, charIndex)
	} else if charIndex < len(key)-1 {
		return search(node.mid, key, charIndex+1)
	} else {
		return node
	}
}

func insert(node *Node, key string, value Item, charIndex int) *Node {
	ci := key[charIndex]

	if node == nil {
		node = &Node{c: ci}
	}

	if ci < node.c {
		node.left = insert(node.left, key, value, charIndex)
	} else if ci > node.c {
		node.right = insert(node.right, key, value, charIndex)
	} else if charIndex < len(key)-1 {
		node.mid = insert(node.mid, key, value, charIndex+1)
	} else {
		node.value = value
	}
	return node
}

func collect(node *Node, prefix string, bucket *[]Item) {

	if node == nil {
		return
	}
	collect(node.left, prefix, bucket)
	if node.value != nil {
		*bucket = append(*bucket, prefix+string(node.c))
	}
	collect(node.mid, prefix+string(node.c), bucket)
	collect(node.right, prefix, bucket)
}

func collect2(node *Node, prefix string, charIndex int, pattern string, bucket *[]Item) {
	if node == nil {
		return
	}
	ci := pattern[charIndex]

	if ci == '.' || ci < node.c {
		collect2(node.left, prefix, charIndex, pattern, bucket)
	}

	if ci == '.' || ci == node.c {
		if charIndex == len(pattern)-1 && node.value != nil {
			*bucket = append(*bucket, prefix+string(node.c))
		}
		if charIndex < len(pattern)-1 {
			collect2(node.mid, prefix+string(node.c), charIndex+1, pattern, bucket)
		}
	}

	if ci == '.' || ci > node.c {
		collect2(node.right, prefix, charIndex, pattern, bucket)
	}

}
