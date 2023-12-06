package lib

type MinHeapArray struct {
	array []KeyInt
	size  int
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (2 * i) + 1
}

func right(i int) int {
	return (2 * i) + 2
}

func NewMinHeapArray() *MinHeapArray {
	panic(("unimplemented"))
}

func (heap *MinHeapArray) SupprMin() KeyInt {
	panic(("unimplemented"))
}

func (heap *MinHeapArray) Ajout(key KeyInt) {
	panic(("unimplemented"))
}

func (heap *MinHeapArray) AjoutsIteratifs(keys []KeyInt) {
	panic(("unimplemented"))
}

func (heap *MinHeapArray) Construction(keys []*KeyInt) {
	panic(("unimplemented"))
}

func (heap *MinHeapArray) String() string {
	panic(("unimplemented"))
}

func (heap *MinHeapArray) Viz() []byte {
	panic(("unimplemented"))
}
