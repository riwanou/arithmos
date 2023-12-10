package lib

type MinHeapArray struct {
	array []KeyInt
}

/*
Checks if heap is empty.
*/
func (heap *MinHeapArray) isEmpty() bool {
	return len(heap.array) == 0
}

/*
Checks if an index exists.
*/
func (heap *MinHeapArray) isExists(i int) bool {
	return i < len(heap.array)
}

/*
Checks if an index has children.
*/
func (heap *MinHeapArray) hasChildren(i int) bool {
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

	if !heap.isExists(i) {
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
func (heap *MinHeapArray) SupprMin() *KeyInt {
	if heap.isEmpty() {
		panic("Error: Unable to remove minimum because heap is empty!")
	}

	minKey := heap.array[0]
	heap.array[0] = heap.array[len(heap.array)-1]
	heap.algorithm(0)

	return &minKey
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
func (heap *MinHeapArray) algorithm(keyIndex int) {
	var childKeyIndex int

	childKeyIndex = heap.siftDown(keyIndex)

	// Check if everything is done
	if childKeyIndex == -1 {
		return
	}

	heap.algorithm(childKeyIndex)
}

func (heap *MinHeapArray) siftDown(keyIndex int) int {
	var key KeyInt
	var leftOrRightKeyIndex int
	var leftOrRightKey KeyInt

	key = heap.key(keyIndex)

	if heap.hasLeftChild(keyIndex) {
		leftOrRightKeyIndex = heap.left(keyIndex)
	}

	if heap.hasRightChild(keyIndex) {
		leftOrRightKeyIndex = heap.right(keyIndex)
	}

	leftOrRightKey = heap.key(leftOrRightKeyIndex)

	if leftOrRightKey.Inf(&key) {
		heap.array[keyIndex], heap.array[leftOrRightKeyIndex] = leftOrRightKey, key
		return leftOrRightKeyIndex
	}

	return -1
}

/*
Ajout

Add a new element to the end of an array;

 1. Sift up the new element, while heap property is broken.
 2. Sifting is done as following: compare node's value with parent's value.
    If they are in wrong order, swap them.
*/
func (heap *MinHeapArray) Ajout(key *KeyInt) {
	heap.array = append(heap.array, *key)
	heap.siftUp(len(heap.array) - 1)
}

func (heap *MinHeapArray) siftUp(keyIndex int) {
	if keyIndex == 0 {
		return
	}

	key := heap.key(keyIndex)
	parentKeyIndex := heap.parent(keyIndex)
	parentKey := heap.key(parentKeyIndex)

	// Check if property is broken
	if key.Inf(&parentKey) {
		// Swap key and key's parent
		heap.array[keyIndex], heap.array[parentKeyIndex] = heap.key(parentKeyIndex), heap.key(keyIndex)

		heap.siftUp(parentKeyIndex)
	}
}

func (heap *MinHeapArray) AjoutIteratif(keys []*KeyInt) {
	for _, key := range keys {
		heap.Ajout(key)
	}
}

func (heap *MinHeapArray) Construction(keys []*KeyInt) {
	// Add every key to array
	for _, key := range keys {
		heap.array = append(heap.array, *key)
	}

	// Sift down every tree
	for i := 1; i < len(heap.array)/2; i++ {
		heap.siftDown(i)
	}
}

func (heap *MinHeapArray) String() string {
	text := "["
	last := ""

	for _, key := range heap.array {
		if last != "" {
			text += last + ", "
		}
		last = key.String()
	}

	text += last + "]"
	return text
}

func (heap *MinHeapArray) Viz() []byte {
	panic(("unimplemented"))
}
