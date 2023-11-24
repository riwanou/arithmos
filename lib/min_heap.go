package lib

type MinHeap interface {
	SupprMin() KeyInt
	Ajout(key KeyInt)
	AjoutsIteratifs(keys []KeyInt)
}

type MinHeapNode struct {
	data  KeyInt
	left  *MinHeapNode
	right *MinHeapNode
}

type MinHeapTree struct {
	root *MinHeapNode
}

func NewMinHeapTree() *MinHeapTree {
	return &MinHeapTree{
		root: nil,
	}
}

func (*MinHeapTree) Ajout(key KeyInt) {
	panic("unimplemented")
}

func (*MinHeapTree) AjoutsIteratifs(keys []KeyInt) {
	panic("unimplemented")
}

func (*MinHeapTree) SupprMin() KeyInt {
	panic("unimplemented")
}
