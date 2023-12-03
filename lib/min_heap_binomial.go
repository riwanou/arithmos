package lib

import (
	"bytes"
	"cmp"
	"math"

	"github.com/bradleyjkemp/memviz"
	"golang.org/x/exp/slices"
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

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func (tree *BinomialTree) addSubtree(other *BinomialTree) {
	tree.order = max(other.order, tree.order) + 1
	tree.children = append(tree.children, other)
}

func BinomialTreeUnion(lhs *BinomialTree, rhs *BinomialTree) *BinomialTree {
	if lhs.data.Inf(rhs.data) {
		lhs.addSubtree(rhs)
		return lhs
	} else {
		rhs.addSubtree(lhs)
		return rhs
	}
}

/**
* Binomial Queue
 */

type MinHeapBinomial struct {
	trees []*BinomialTree
	size  int
}

func NewMinHeapBinomial() *MinHeapBinomial {
	return &MinHeapBinomial{
		trees: make([]*BinomialTree, 0),
		size:  0,
	}
}

func (heap *MinHeapBinomial) Union(other *MinHeapBinomial) {
	trees := append(heap.trees, other.trees...)
	slices.SortFunc(trees, func(a, b *BinomialTree) int {
		return cmp.Compare(a.order, b.order)
	})

	heap.size += other.size
	maxOrder := int(math.Log2(float64(heap.size))) + 1
	merged := make([]*BinomialTree, maxOrder)

	for _, tree := range trees {
		order := tree.order
		for merged[order] != nil {
			tree = BinomialTreeUnion(tree, merged[order])
			merged[order] = nil
			order += 1
		}
		merged[order] = tree
	}

	heap.trees = make([]*BinomialTree, 0, len(trees))
	for _, tree := range merged {
		if tree != nil {
			heap.trees = append(heap.trees, tree)
		}
	}
}

func (heap *MinHeapBinomial) Ajout(key *KeyInt) {
	addHeap := NewMinHeapBinomial()
	addHeap.trees = append(addHeap.trees, NewBinomialTree(key))
	addHeap.size += 1
	heap.Union(addHeap)
}

func (heap *MinHeapBinomial) SupprMin() *KeyInt {
	if len(heap.trees) == 0 {
		return nil
	}

	// remove min tree from binomial heap list
	minTree := heap.trees[0]
	minTreeIndex := 0
	for i, tree := range heap.trees[1:] {
		if tree.data.Inf(minTree.data) {
			minTree = tree
			minTreeIndex = i
		}
	}
	heap.trees = append(heap.trees[:minTreeIndex],
		heap.trees[minTreeIndex+1:]...)

	// merge the children of the min tree into the heap list
	minKey := minTree.data
	childrenHeap := NewMinHeapBinomial()
	childrenHeap.trees = minTree.children
	childrenHeap.size = len(minTree.children)
	heap.Union(childrenHeap)

	return minKey
}

// not needed for binomial min heap
func (heap *MinHeapBinomial) AjoutIteratif(keys []*KeyInt) {
	panic("unimplemented")
}

func (heap *MinHeapBinomial) Construction(keys []*KeyInt) {
	for _, key := range keys {
		heap.Ajout(key)
	}
}

/**
 * Heap Vizualisation
 */

func (heap *MinHeapBinomial) String() string {
	text := "["
	for i, tree := range heap.trees {
		text += tree.String()
		if i < len(heap.trees)-1 {
			text += ", "
		}
	}
	text += "]"
	return text
}

func (heap *MinHeapBinomial) Viz() []byte {
	buf := &bytes.Buffer{}
	memviz.Map(buf, heap)
	return buf.Bytes()
}

/**
 * Tree Vizualisation
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
