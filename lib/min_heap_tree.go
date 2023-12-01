package lib

import "fmt"

type MinHeapNode struct {
	data  *KeyInt
	left  *MinHeapNode
	right *MinHeapNode
}

type MinHeapTree struct {
	root *MinHeapNode
}

func NewMinHeapTree() *MinHeapTree {
	return &MinHeapTree{
		root: &MinHeapNode{},
	}
}

// Stop the level order search if nodeOp return false
func (heap *MinHeapTree) levelOrder(nodeOp func(*MinHeapNode) bool) {
	queue := make([]*MinHeapNode, 0)
	queue = append(queue, heap.root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if nodeOp(node) == false {
			return
		}

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func (heap *MinHeapTree) String() string {
	text := "["
	last := ""

	heap.levelOrder(func(node *MinHeapNode) bool {
		if node.data != nil {
			if last != "" {
				text += last + ", "
			}
			last = node.data.String()
		}
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

func (heap *MinHeapTree) Ajout(key KeyInt) {
	if heap.root.data == nil {
		heap.root.data = &key
		return
	}

	heap.levelOrder(func(node *MinHeapNode) bool {
		if node.left == nil {
			node.left = &MinHeapNode{data: &key}
			return false
		}
		if node.right == nil {
			node.right = &MinHeapNode{data: &key}
			return false
		}
		return true
	})
}

// TODO: use an implementation with better complexity
func (heap *MinHeapTree) AjoutsIteratif(keys []KeyInt) {
	for _, key := range keys {
		heap.Ajout(key)
	}
}

func (*MinHeapTree) SupprMin() *KeyInt {
	panic("unimplemented")
}
