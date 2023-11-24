package lib

type HeapMin interface {
	SupprMin() KeyInt
	Ajout(key KeyInt)
	AjoutsIteratifs(keys []KeyInt)
}

type HeapMinNode struct {
	data  KeyInt
	left  *HeapMinNode
	right *HeapMinNode
}

type HeapMinTree struct {
	root *HeapMinNode
}

func NewHeapMinTree() *HeapMinTree {
	return &HeapMinTree{
		root: nil,
	}
}

func (*HeapMinTree) Ajout(key KeyInt) {
	panic("unimplemented")
}

func (*HeapMinTree) AjoutsIteratifs(keys []KeyInt) {
	panic("unimplemented")
}

func (*HeapMinTree) SupprMin() KeyInt {
	panic("unimplemented")
}
