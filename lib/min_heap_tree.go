package lib

import (
	"fmt"
	"math/bits"

	"golang.org/x/exp/slices"
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
	path []byte
}

func NewMinHeapTree() *MinHeapTree {
	return &MinHeapTree{
		root: &MinHeapNode{},
		size: 0,
		path: make([]byte, 0, 25),
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

// Compute the path to the last node based on the size
func (heap *MinHeapTree) pathTo(size uint32) []byte {
	len := max(bits.Len32(size)-1, 0)
	heap.path = heap.path[:0]

	for i := 0; i < len; i++ {
		heap.path = append(heap.path, byte(size&1))
		size >>= 1
	}
	slices.Reverse(heap.path)

	return heap.path
}

// Return the last node of the binary tree (full)
func (heap *MinHeapTree) nodeFromPath(path []byte) *MinHeapNode {
	currNode := heap.root
	for _, bitValue := range path {
		if bitValue == 0 {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}
	return currNode
}

// Add node to the tree, keep it full
func (heap *MinHeapTree) addAtEnd(key *KeyInt) *MinHeapNode {
	heap.size += 1

	if heap.root.isNil() {
		heap.root.data = key
		return heap.root
	}

	beforeLastNode := heap.root
	bitPath := heap.pathTo(heap.size)
	if len(bitPath) > 1 {
		beforeLastNode = heap.nodeFromPath(bitPath[:len(bitPath)-1])
	}

	var insertedNode *MinHeapNode
	if beforeLastNode.left.isNil() {
		beforeLastNode.left = &MinHeapNode{data: key, parent: beforeLastNode}
		insertedNode = beforeLastNode.left
	} else {
		beforeLastNode.right = &MinHeapNode{data: key, parent: beforeLastNode}
		insertedNode = beforeLastNode.right
	}

	if insertedNode.isNil() {
		panic("Failed to insert node in min heap tree")
	}

	return insertedNode
}

func (heap *MinHeapTree) Ajout(key *KeyInt) {
	insertedNode := heap.addAtEnd(key)
	heap.bubbleUpNode(insertedNode)
}

func (heap *MinHeapTree) AjoutIteratif(keys []*KeyInt) {
	for _, key := range keys {
		heap.Ajout(key)
	}
}

func (heap *MinHeapTree) heapify(node *MinHeapNode) {
	if node.left != nil && node.left.left != nil {
		heap.heapify(node.left)
	}
	if node.right != nil && node.right.left != nil {
		heap.heapify(node.right)
	}
	heap.sinkNode(node)
}

func (heap *MinHeapTree) Construction(keys []*KeyInt) {
	for _, key := range keys {
		heap.addAtEnd(key)
	}
	if heap.size > 0 {
		heap.heapify(heap.root)
	}
}

// Swap the given node with one of its smaller children recursivly
// For example, if we insert a big key at the top, it will lower it to the bottom
func (heap *MinHeapTree) sinkNode(node *MinHeapNode) {
	var minNode *MinHeapNode

	if !node.left.isNil() && node.left.data.Inf(node.data) {
		minNode = node.left
	}
	if !node.right.isNil() && node.right.data.Inf(node.data) {
		if minNode == nil || (minNode != nil && node.right.data.Inf(minNode.data)) {
			minNode = node.right
		}
	}

	if minNode != nil {
		minNode.data, node.data = node.data, minNode.data
		heap.sinkNode(minNode)
	}
}

func (heap *MinHeapTree) SupprMin() *KeyInt {
	last := heap.nodeFromPath(heap.pathTo(heap.size))
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
