package lib_test

import (
	"arithmos/lib"
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"runtime/debug"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const keysDirName = "../data/cles_alea/"

func vizBytes(data []byte, filename string) {
	DirPath := "../test-output/"
	_ = os.Mkdir(DirPath, 0755)
	path := DirPath + filename
	cmd := exec.Command("dot", "-Tpng", "-Gdpi=300", "-o", path+".png")
	cmd.Stdin = bytes.NewReader(data)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func vizHeap(heap lib.MinHeap, filename string) {
	vizBytes(heap.Viz(), filename)
}

func genKeys() []*lib.KeyInt {
	return []*lib.KeyInt{
		lib.NewKeyInt(0, 10),
		lib.NewKeyInt(0, 20),
		lib.NewKeyInt(0, 30),
		lib.NewKeyInt(0, 40),
		lib.NewKeyInt(0, 50),
	}
}

func runTestHeaps(test func(lib.MinHeap), withBinomial bool) {
	heaps := []lib.MinHeap{lib.NewMinHeapTree()}
	if withBinomial {
		heaps = append(heaps, lib.NewMinHeapBinomial())
	}
	for _, heap := range heaps {
		test(heap)
	}
}

func getKeysFromFile(path string) []*lib.KeyInt {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	keys := make([]*lib.KeyInt, 0, 80000)
	s := bufio.NewScanner(f)
	for s.Scan() {
		keyInt, err := lib.NewKeyIntFromString(s.Text())
		if err != nil {
			panic(err)
		}
		keys = append(keys, keyInt)
	}

	return keys
}

/**
 * Tests
 */

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
	}, false)
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
	}, false)
}

func TestAjoutIteratif(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		assert.Equal(t, "[]", heap.String())
		heap.AjoutIteratif(keys)
		assert.Equal(t, "[0-10, 0-20, 0-30, 0-40, 0-50]", heap.String())
	}, false)
}

func TestSupprMin(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		heap.Ajout(keys[0])
		assert.Equal(t, keys[0], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
	}, false)
}

func TestSupprMinEmpty(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		assert.Nil(t, heap.SupprMin())
	}, true)
}

func TestSupprMultiple(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		heap.Ajout(keys[4])
		heap.Ajout(keys[3])
		heap.Ajout(keys[2])
		heap.Ajout(keys[1])
		heap.Ajout(keys[0])
		assert.Equal(t, keys[0], heap.SupprMin())
		assert.Equal(t, keys[1], heap.SupprMin())
		assert.Equal(t, keys[2], heap.SupprMin())
		assert.Equal(t, keys[3], heap.SupprMin())
		assert.Equal(t, keys[4], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
		heap.Ajout(keys[0])
		heap.Ajout(keys[1])
		assert.Equal(t, keys[0], heap.SupprMin())
		assert.Equal(t, keys[1], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
	}, false)
}

func TestSupprFile(t *testing.T) {
	keys := getKeysFromFile(keysDirName + "jeu_1_nb_cles_1000.txt")
	heapTree := lib.NewMinHeapTree()
	heapTree.AjoutIteratif(keys)

	heapBinomial := lib.NewMinHeapBinomial()
	heapBinomial.Construction(keys)

	for i := 0; i < len(keys); i++ {
		assert.Equal(t, heapTree.SupprMin(), heapBinomial.SupprMin())
	}
}

/**
 * Benchmarks
 */

func benchmarkHeaps(
	b *testing.B,
	bench func(heap lib.MinHeap, keys []*lib.KeyInt),
	withBinomial bool,
) {
	debug.SetGCPercent(800)
	dirEntries, err := os.ReadDir(keysDirName)
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		keys := getKeysFromFile(keysDirName + entry.Name())
		b.Run("heapTree/"+entry.Name(), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				bench(lib.NewMinHeapTree(), keys)
			}
		})
		if withBinomial {
			b.Run("heapBinomial/"+entry.Name(), func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					bench(lib.NewMinHeapBinomial(), keys)
				}
			})
		}
	}
}

func BenchmarkAjoutIteratif(b *testing.B) {
	benchmarkHeaps(b, func(heap lib.MinHeap, keys []*lib.KeyInt) {
		heap.AjoutIteratif(keys)
	}, false)
}

func BenchmarkConstruction(b *testing.B) {
	benchmarkHeaps(b, func(heap lib.MinHeap, keys []*lib.KeyInt) {
		heap.Construction(keys)
	}, true)
}

/**
 * Union benchmarks
 */

func BenchmarkUnion(b *testing.B) {
	dataSizes := []int{1000, 5000, 20000, 50000, 80000, 120000, 200000}

	for _, dataSize := range dataSizes {
		keysGroups := make([][]*lib.KeyInt, 0, 5)
		for keysData := 1; keysData <= 5; keysData++ {
			path := keysDirName + "jeu_" + strconv.Itoa(keysData) +
				"_nb_cles_" + strconv.Itoa(dataSize) + ".txt"
			keysGroups = append(keysGroups, getKeysFromFile(path))
		}

		name := "cles_" + strconv.Itoa(dataSize)
		heaps := make([]*lib.MinHeapBinomial, len(keysGroups))

		b.Run("heapBinomial/"+name, func(b *testing.B) {
			for i, keys := range keysGroups {
				heaps[i] = lib.NewMinHeapBinomial()
				heaps[i].Construction(keys)
			}
			for n := 0; n < b.N; n++ {
				binoHeap := lib.NewMinHeapBinomial()
				for _, heap := range heaps {
					binoHeap.Union(heap)
				}
			}
		})
	}
}
