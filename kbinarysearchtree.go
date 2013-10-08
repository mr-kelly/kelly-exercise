package klib

import (
// "fmt"
)

type KBinaryNode struct {
	element int // _T
	left    *KBinaryNode
	right   *KBinaryNode
}

type KBinarySearchTree struct {
	root *KBinaryNode // top node
}

func NewBinarySearchTree() *KBinarySearchTree {
	newTree := new(KBinarySearchTree)
	return newTree
}

func (this *KBinarySearchTree) Insert(element int) {
	this.insert(element, &this.root)
}

func (this *KBinarySearchTree) Remove(element int) {
	this.remove(element, this.root)
}

func (this *KBinarySearchTree) FindMin() *KBinaryNode {
	return this.findMin(this.root)
}

func (this *KBinarySearchTree) FindMax() *KBinaryNode {
	return this.findMax(this.root)
}

func (this *KBinarySearchTree) remove(element int, startNode *KBinaryNode) {
	if startNode == nil {
		return // 必须开始结点
	} else if element < startNode.element {
		this.remove(element, startNode.left)
	} else if element > startNode.element {
		this.remove(element, startNode.right)
	} else if startNode.left != nil && startNode.right != nil {
		// 找不到需要移除的，但它有left，right结点，需要移动它们
		// 删除右子树最小的结点
		startNode.element = this.findMin(startNode.right).element
		this.remove(startNode.element, startNode.right)
	} else {
		// 需要移除的
	}
}

//		传入指针的指针，用于改变上层
// 传入指针为空，因此要传入指针的指针，获取外层地址
// 另一方法是： 采用返回值，代替指针指针传递(类似Java)
func (this *KBinarySearchTree) insert(element int, startNode **KBinaryNode) {
	//refStartNode := *startNode

	if *startNode == nil {
		*startNode = new(KBinaryNode)
		(*startNode).element = element

		if this.root == nil {
			this.root = *startNode // 初始化顶点root
		}

		// 设置上级绑定
	} else if element < (*startNode).element {
		this.insert(element, &((*startNode).left)) // 比它小，继续插右
	} else if element > (*startNode).element {
		this.insert(element, &((*startNode).right))
	} else {
		// 重复
	}
}

/////////////////////////////////// 下方未完善

func (this *KBinarySearchTree) findMin(startNode *KBinaryNode) *KBinaryNode {
	if startNode != nil {
		for startNode.left != nil {
			startNode = startNode.left
		} // else startNode.left == nil,  means no left, the min
	}

	return startNode
}

func (this *KBinarySearchTree) findMax(startNode *KBinaryNode) *KBinaryNode {
	if startNode != nil {
		for startNode.right != nil {
			startNode = startNode.right
		} // else startNode.right == nil,  means no right, the max
	}

	return startNode
}
