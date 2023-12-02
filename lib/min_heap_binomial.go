package lib

import (
	"bytes"

	"github.com/bradleyjkemp/memviz"
)

/**
* Binomial tree
 */

type BinomialTree struct {
	order    uint32
	data     *KeyInt
	children []*BinomialTree
}

func NewBinomialTree(data *KeyInt) *BinomialTree {
	return &BinomialTree{
		order:    0,
		data:     data,
		children: make([]*BinomialTree, 0),
	}
}

func (tree *BinomialTree) Union(other *BinomialTree) {
	if tree.order != other.order {
		panic("Failed to merge binomial trees, they must have the same order")
	}

	if tree.order == 0 {
		tree.order = 1
	} else {
		tree.order += other.order
	}

	tree.children = append(tree.children, other)
}

/**
* Binomial Queue
 */

type MinHeapBinomial struct{}

func NewMinHeapBinomial() *MinHeapBinomial {
	return &MinHeapBinomial{}
}

/**
* Vizualisation
 */

func (tree *BinomialTree) String() string {
	text := "("
	text += tree.data.String()
	if len(tree.children) > 0 {
		text += ", "
	}
	for i, child := range tree.children {
		text += child.String()
		if i < len(tree.children)-1 {
			text += ", "
		}
	}
	text += ")"
	return text
}

func (tree *BinomialTree) Viz() []byte {
	buf := &bytes.Buffer{}
	memviz.Map(buf, tree)
	return buf.Bytes()
}
