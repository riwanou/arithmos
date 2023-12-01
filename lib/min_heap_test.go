package lib_test

import (
	"arithmos/lib"
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func vizHeap(heap lib.MinHeap, filename string) {
	DirPath := "../test-output/"
	_ = os.Mkdir(DirPath, 0755)
	path := DirPath + filename
	cmd := exec.Command("dot", "-Tpng", "-Gdpi=300", "-o", path+".png")
	cmd.Stdin = bytes.NewReader(heap.Viz())
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func genKeys() []lib.KeyInt {
	return []lib.KeyInt{
		*lib.NewKeyInt(0, 10),
		*lib.NewKeyInt(0, 20),
		*lib.NewKeyInt(0, 30),
		*lib.NewKeyInt(0, 40),
		*lib.NewKeyInt(0, 50),
	}
}

func runTestHeaps(test func(lib.MinHeap)) {
	heaps := []lib.MinHeap{lib.NewMinHeapTree()}
	for _, heap := range heaps {
		test(heap)
	}
}

func TestAjoutAscending(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		assert.Equal(t, "[]", heap.String())
		heap.Ajout(keys[0])
		assert.Equal(t, "[0-10]", heap.String())
		heap.Ajout(keys[1])
		assert.Equal(t, "[0-10, 0-20]", heap.String())
		heap.Ajout(keys[2])
		assert.Equal(t, "[0-10, 0-20, 0-30]", heap.String())
		heap.Ajout(keys[4])
		assert.Equal(t, "[0-10, 0-20, 0-30, 0-50]", heap.String())
		heap.Ajout(keys[3])
		assert.Equal(t, "[0-10, 0-20, 0-30, 0-50, 0-40]", heap.String())
	})
}

func TestAjoutDescending(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		assert.Equal(t, "[]", heap.String())
		heap.Ajout(keys[4])
		assert.Equal(t, "[0-50]", heap.String())
		heap.Ajout(keys[3])
		assert.Equal(t, "[0-40, 0-50]", heap.String())
		heap.Ajout(keys[2])
		assert.Equal(t, "[0-30, 0-50, 0-40]", heap.String())
		heap.Ajout(keys[0])
		assert.Equal(t, "[0-10, 0-30, 0-40, 0-50]", heap.String())
		heap.Ajout(keys[1])
		assert.Equal(t, "[0-10, 0-20, 0-40, 0-50, 0-30]", heap.String())
	})
}

func TestAjoutIteratif(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		assert.Equal(t, "[]", heap.String())
		heap.AjoutsIteratif(keys)
		assert.Equal(t, "[0-10, 0-20, 0-30, 0-40, 0-50]", heap.String())
	})
}

func TestSupprMin(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		heap.Ajout(keys[0])
		assert.Equal(t, &keys[0], heap.SupprMin())
	})
}

func TestSupprMinEmpty(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		assert.Nil(t, heap.SupprMin())
	})
}

func TestSupprMultiple(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		heap.Ajout(keys[4])
		heap.Ajout(keys[3])
		heap.Ajout(keys[2])
		heap.Ajout(keys[1])
		heap.Ajout(keys[0])
		vizHeap(heap, "heap-1")
		assert.Equal(t, &keys[0], heap.SupprMin())
		vizHeap(heap, "heap-2")
		assert.Equal(t, &keys[1], heap.SupprMin())
		vizHeap(heap, "heap-3")
		assert.Equal(t, &keys[2], heap.SupprMin())
		vizHeap(heap, "heap-4")
		assert.Equal(t, &keys[3], heap.SupprMin())
		vizHeap(heap, "heap-5")
		assert.Equal(t, &keys[4], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
		heap.Ajout(keys[0])
		heap.Ajout(keys[1])
		assert.Equal(t, &keys[0], heap.SupprMin())
		assert.Equal(t, &keys[1], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
	})
}
