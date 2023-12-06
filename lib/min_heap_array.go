package lib

type MinHeapArray struct {
	array       []KeyInt
	currentSize int
}

/*
Checks if heap is empty.
*/
func (heap *MinHeapArray) isEmpty() bool {
	return heap.currentSize == 0
}

/*
Checks if an index exists.
*/
func (heap *MinHeapArray) isExists(i int) bool {
	return heap.currentSize >= i
}

/*
Checks if an index has children.
*/
func (heap *MinHeapArray) hasChildren(i int, considerOneChild bool) bool {
	if considerOneChild {
		return heap.hasLeftChild(i) || heap.hasRightChild(i)
	}
	return heap.hasLeftChild(i) && heap.hasRightChild(i)
}

/*
Checks if an index has left child.
*/
func (heap *MinHeapArray) hasLeftChild(i int) bool {
	return heap.isExists(heap.left(i))
}

/*
Checks if an index has right child.
*/
func (heap *MinHeapArray) hasRightChild(i int) bool {
	return heap.isExists(heap.right(i))
}

/*
Returns KeyInt from an index.
*/
func (heap *MinHeapArray) key(i int) KeyInt {
	if heap.isEmpty() {
		panic("Error: Unable to get root key because heap is empty!")
	}

	if heap.isExists(i) {
		panic("Error: Unable to get key because index match any existing key!")
	}

	return heap.array[i]
}

/*
Returns parent from an index.
*/
func (heap *MinHeapArray) parent(i int) int {
	return (i - 1) / 2
}

/*
Returns left child from an index.
*/
func (heap *MinHeapArray) left(i int) int {
	return (2 * i) + 1
}

/*
Returns right child from an index.
*/
func (heap *MinHeapArray) right(i int) int {
	return (2 * i) + 2
}

func NewMinHeapArray() *MinHeapArray {
	heap := &MinHeapArray{}
	heap.array = make([]KeyInt, 0)
	return heap
}

/*
SupprMin removes key with the minimum value.
*/
func (heap *MinHeapArray) SupprMin() KeyInt {
	if heap.isEmpty() {
		panic("Error: Unable to remove minimum because heap is empty!")
	}

	minKey := heap.array[0]
	heap.array[0] = heap.array[heap.currentSize-1]
	heap.currentSize--
	heap.algorithm(0)

	return minKey
}

// algorithm
//
//  1. Copy the last value in the array to the root
//
//  2. Decrease heap's size by 1
//
//  3. Sift down root's value. Sifting is done as following:
//
//     - if current node has no children:
//     sifting is over;
//
//     - if current node has one child:
//
//     if heap property is broken:
//     then swap current node's value and child value
//     sift down the child;
//
//     - if current node has two children:
//     find the smallest of them.
//     if heap property is broken:
//     then swap current node's value and selected child value
//     sift down the child.
func (heap *MinHeapArray) algorithm(i int) {
	// Check if current index has no children
	if !heap.hasChildren(i, false) {
		return
	}

	// Check if current index has one child
	if heap.hasChildren(i, true) {

		// Check if heap has no broken heap property
		childIndex, hasBrokenHeapProperty := heap.hasBrokenHeapProperty(i)

		if !hasBrokenHeapProperty {
			return
		}

		// Swap current node's value and child value
		heap.array[i] = heap.key(childIndex)
		heap.array[childIndex] = heap.key(i)

		heap.algorithm(childIndex)
	}

	// Check if current index has two children
	if heap.hasChildren(i, false) {

		leftKeyIndex := heap.left(i)
		leftKeyValue := heap.key(leftKeyIndex)
		rightKeyIndex := heap.right(i)
		rightKeyValue := heap.key(rightKeyIndex)

		if leftKeyValue.Inf(&rightKeyValue) {
			heap.array[i] = heap.key(leftKeyIndex)
			heap.array[leftKeyIndex] = heap.key(i)
		} else {
			heap.array[i] = heap.key(rightKeyIndex)
			heap.array[rightKeyIndex] = heap.key(i)
		}
	}

}

func (heap *MinHeapArray) hasBrokenHeapProperty(i int) (int, bool) {
	currentKey := heap.key(i)

	if heap.hasLeftChild(i) {
		leftKeyIndex := heap.left(i)
		leftKey := heap.key(leftKeyIndex)

		if leftKey.Inf(&currentKey) || leftKey.Eq(&currentKey) {
			return leftKeyIndex, true
		}
	}

	if heap.hasRightChild(i) {
		rightKeyIndex := heap.right(i)
		rightKey := heap.key(rightKeyIndex)

		if rightKey.Inf(&currentKey) || rightKey.Eq(&currentKey) {
			return rightKeyIndex, true
		}
	}

	return -1, false
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
