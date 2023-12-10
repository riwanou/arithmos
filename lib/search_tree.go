package lib

type SearchTreeNode struct {
	data  *KeyInt
	left  *SearchTreeNode
	right *SearchTreeNode
}

func (node *SearchTreeNode) isNil() bool {
	return (node == nil || node.data == nil)
}

type SearchTree struct {
	root *SearchTreeNode
}

func NewSearchTree() *SearchTree {
	return &SearchTree{
		root: &SearchTreeNode{},
	}
}

func (tree *SearchTree) insertNode(node *SearchTreeNode, key *KeyInt) {
	if key.Inf(node.data) {
		if node.left.isNil() {
			node.left = &SearchTreeNode{data: key}
		} else {
			tree.insertNode(node.left, key)
		}
	} else {
		if node.right.isNil() {
			node.right = &SearchTreeNode{data: key}
		} else {
			tree.insertNode(node.right, key)
		}
	}
}

func (tree *SearchTree) Insert(key *KeyInt) {
	if tree.root.data == nil {
		tree.root.data = key
		return
	}

	tree.insertNode(tree.root, key)
}

func (tree *SearchTree) getNode(node *SearchTreeNode, key *KeyInt) *KeyInt {
	if key.Eq(node.data) {
		return node.data
	}
	if key.Inf(node.data) && !node.left.isNil() {
		return tree.getNode(node.left, key)
	}
	if !node.right.isNil() {
		return tree.getNode(node.right, key)
	}
	return nil
}

func (tree *SearchTree) Get(key *KeyInt) *KeyInt {
	if tree.root.data == nil {
		return nil
	}
	return tree.getNode(tree.root, key)
}
