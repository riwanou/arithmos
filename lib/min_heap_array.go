package lib

type MinHeapArray struct {
	array       []KeyInt
	currentSize int
}

func (heap *MinHeapArray) root() KeyInt {
	// Check if array is not empty
	if heap.currentSize == 0 {
		panic("Error: Unable to get root key because heap is empty!")
	}
	return heap.array[0]
}

func (heap *MinHeapArray) parent(i int) KeyInt {
	// Check if index i doesn't exists
	if i > heap.currentSize {
		panic("Error: Unable to get parent key because index match any existing key!")
	}
	computeIndex := (i - 1) / 2
	return heap.array[computeIndex]
}

func (heap *MinHeapArray) left(i int) KeyInt {
	// Check if index i doesn't exists
	if i > heap.currentSize {
		panic("Error: Unable to get left key because index match any existing key!")
	}
	computeIndex := (2 * i) + 1
	return heap.array[computeIndex]
}

func (heap *MinHeapArray) right(i int) KeyInt {
	// Check if index i doesn't exists
	if i > heap.currentSize {
		panic("Error: Unable to get right key because index match any existing key!")
	}
	computeIndex := (2 * i) + 2
	return heap.array[computeIndex]
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
