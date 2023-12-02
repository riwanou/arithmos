package lib

import (
	"fmt"
)

type MinHeapNode struct {
	data   *KeyInt
	left   *MinHeapNode
	right  *MinHeapNode
	parent *MinHeapNode
}

func (node *MinHeapNode) isNil() bool {
	return node == nil || node.data == nil
}

type MinHeapTree struct {
	root *MinHeapNode
	size uint32
}

func NewMinHeapTree() *MinHeapTree {
	return &MinHeapTree{
		root: &MinHeapNode{},
		size: 0,
	}
}

// Stop the level order search if nodeOp return false
func (heap *MinHeapTree) levelOrder(nodeOp func(*MinHeapNode) bool) {
	queue := make([]*MinHeapNode, 0, heap.size)
	if !heap.root.isNil() {
		queue = append(queue, heap.root)
	}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if nodeOp(node) == false {
			return
		}

		if !node.left.isNil() {
			queue = append(queue, node.left)
		}
		if !node.right.isNil() {
			queue = append(queue, node.right)
		}
	}
}

// Swap the given node with its parent recursivly
// For example, if we insert a low key at the bottom, it will raise it to the top
func (*MinHeapTree) bubbleUpNode(node *MinHeapNode) {
	currNode := node
	for !currNode.parent.isNil() {
		parentNode := currNode.parent
		if currNode.data.Inf(parentNode.data) {
			currNode.data, parentNode.data = parentNode.data, currNode.data
		}
		currNode = parentNode
	}
}

func (heap *MinHeapTree) Ajout(key *KeyInt) {
	if heap.root.isNil() {
		heap.root.data = key
		heap.size += 1
		return
	}

	var insertedNode *MinHeapNode
	heap.levelOrder(func(node *MinHeapNode) bool {
		if node.left.isNil() {
			node.left = &MinHeapNode{data: key, parent: node}
			insertedNode = node.left
			return false
		}
		if node.right.isNil() {
			node.right = &MinHeapNode{data: key, parent: node}
			insertedNode = node.right
			return false
		}
		return true
	})

	if insertedNode.isNil() {
		panic("Failed to insert node in min heap tree")
	}

	heap.bubbleUpNode(insertedNode)
	heap.size += 1
}

func (heap *MinHeapTree) AjoutIteratif(keys []*KeyInt) {
	for _, key := range keys {
		heap.Ajout(key)
	}
}

// Swap the given node with one of its smaller children recursivly
// For example, if we insert a big key at the top, it will lower it to the bottom
func (heap *MinHeapTree) sinkNode(node *MinHeapNode) {
	if !node.left.isNil() && node.left.data.Inf(node.data) {
		node.left.data, node.data = node.data, node.left.data
		heap.sinkNode(node.left)
	} else if !node.right.isNil() && node.right.data.Inf(node.data) {
		node.right.data, node.data = node.data, node.right.data
		heap.sinkNode(node.right)
	}
}

func (heap *MinHeapTree) SupprMin() *KeyInt {
	var last *MinHeapNode
	heap.levelOrder(func(node *MinHeapNode) bool {
		last = node
		return true
	})
	if last.isNil() {
		return nil
	}

	// extract root data, then swap it with the last heap node
	data := heap.root.data
	heap.root.data = last.data
	last.data = nil

	heap.sinkNode(heap.root)
	heap.size -= 1

	return data
}

/**
* Vizualisation
 */

func (heap *MinHeapTree) String() string {
	text := "["
	last := ""

	heap.levelOrder(func(node *MinHeapNode) bool {
		if last != "" {
			text += last + ", "
		}
		last = node.data.String()
		return true
	})

	text += last + "]"
	return text
}

func (heap *MinHeapTree) Viz() []byte {
	text := "digraph structs {\n"

	heap.levelOrder(func(node *MinHeapNode) bool {
		if node.data != nil {
			curr := node.data.String()
			text += fmt.Sprintf("\"%s\"; \n", curr)
			if node.left != nil && node.left.data != nil {
				text += fmt.Sprintf("\"%s\" -> \"%s\";\n",
					curr, node.left.data.String())
			}
			if node.right != nil && node.right.data != nil {
				text += fmt.Sprintf("\"%s\" -> \"%s\";\n",
					curr, node.right.data.String())
			}
		}
		return true
	})

	text += "}"
	return []byte(text)
}
