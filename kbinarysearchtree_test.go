package klib

import (
	// "fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	var tree *KBinarySearchTree = NewBinarySearchTree()
	tree.Insert(1)

	if tree.FindMax() == nil || tree.FindMax().element != 1 {
		t.Errorf("Tree Max: %d", tree.FindMax().element)
	}

	tree.Insert(2)
	if tree.FindMax().element != 2 {
		t.Errorf("Tree Max: %d", tree.FindMax().element) // ref element
	}

	tree.Insert(3)
	if tree.FindMax().element != 3 {
		t.Errorf("Tree Max: %d", tree.FindMax())
	}

	// if tree.FindMax().element != 1 {
	// 	t.Errorf("Tree Min: %d", tree.FindMax())
	// }
}
