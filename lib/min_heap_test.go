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
	"golang.org/x/exp/slices"
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
	heaps := []lib.MinHeap{lib.NewMinHeapTree(), lib.NewMinHeapArray()}
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

func genAscendingKeys(nbKeys uint64) []*lib.KeyInt {
	keys := make([]*lib.KeyInt, 0, nbKeys)
	var i uint64 = 0
	for ; i < nbKeys; i++ {
		keys = append(keys, lib.NewKeyInt(0, i))
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
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		slices.Reverse(keys)
		heap.AjoutIteratif(keys)
		assert.Equal(t, "[0-10, 0-20, 0-40, 0-50, 0-30]", heap.String())
	}, false)
}

func TestConstruction(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		heap.Construction(keys[:0])
		assert.Equal(t, "[]", heap.String())
	}, false)
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		assert.Equal(t, "[]", heap.String())
		heap.Construction(keys)
		assert.Equal(t, "[0-10, 0-20, 0-30, 0-40, 0-50]", heap.String())
	}, false)
}

func TestUnion(t *testing.T) {
	keys := genKeys()

	heap1 := lib.NewMinHeapTree()
	heap1.Construction(keys[:2])
	heap2 := lib.NewMinHeapTree()
	heap2.Construction(keys[2:])

	heap := lib.HeapTreeUnion(heap1, heap2)
	assert.Equal(t, "[0-10, 0-20, 0-30, 0-40, 0-50]", heap.String())
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
	}, true)
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		heap.Construction(keys)
		assert.Equal(t, keys[0], heap.SupprMin())
		assert.Equal(t, keys[1], heap.SupprMin())
		assert.Equal(t, keys[2], heap.SupprMin())
		assert.Equal(t, keys[3], heap.SupprMin())
		assert.Equal(t, keys[4], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
		heap.Construction(keys[:2])
		assert.Equal(t, keys[0], heap.SupprMin())
		assert.Equal(t, keys[1], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
	}, true)
}

func TestSupprFile(t *testing.T) {
	keys := getKeysFromFile(keysDirName + "jeu_1_nb_cles_1000.txt")

	heapArray := lib.NewMinHeapArray()
	heapArray.AjoutIteratif(keys)

	heapArrayCons := lib.NewMinHeapArray()
	heapArrayCons.Construction(keys)

	heapTree := lib.NewMinHeapTree()
	heapTree.AjoutIteratif(keys)

	heapTreeCons1 := lib.NewMinHeapTree()
	heapTreeCons2 := lib.NewMinHeapTree()
	heapTreeCons1.Construction(keys[500:])
	heapTreeCons2.Construction(keys[:500])
	heapTreeCons := lib.HeapTreeUnion(heapTreeCons1, heapTreeCons2)

	heapBinomial := lib.NewMinHeapBinomial()
	heapBinomial.Construction(keys)

	for i := 0; i < len(keys); i++ {
		binoMin := heapBinomial.SupprMin()
		assert.Equal(t, heapArray.SupprMin(), binoMin)
		assert.Equal(t, heapArrayCons.SupprMin(), binoMin)
		assert.Equal(t, heapTree.SupprMin(), binoMin)
		assert.Equal(t, heapTreeCons.SupprMin(), binoMin)
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
	run := func(name string, keys []*lib.KeyInt) {
		b.Run("heapTree/"+name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				bench(lib.NewMinHeapTree(), keys)
			}
		})
		if withBinomial {
			b.Run("heapBinomial/"+name, func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					bench(lib.NewMinHeapBinomial(), keys)
				}
			})
		}
	}

	debug.SetGCPercent(800)
	dirEntries, err := os.ReadDir(keysDirName)
	if err != nil {
		panic(err)
	}
	for _, entry := range dirEntries {
		keys := getKeysFromFile(keysDirName + entry.Name())
		run(entry.Name(), keys)
	}

	withBinomial = false
	run_extra := func(nbKeys uint64) {
		keys := genAscendingKeys(nbKeys)
		run("extra_jeu_nb_cles_"+strconv.FormatUint(nbKeys, 10), keys)
	}

	run_extra(250000)
	run_extra(300000)
	run_extra(500000)
	run_extra(1000000)
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
	run := func(name string, keysGroups [][]*lib.KeyInt) {
		heaps := make([]*lib.MinHeapBinomial, len(keysGroups))
		b.Run("heapBinomial/"+name, func(b *testing.B) {
			for i, keys := range keysGroups {
				heaps[i] = lib.NewMinHeapBinomial()
				heaps[i].Construction(keys)
			}
			for n := 0; n < b.N; n++ {
				binoHeap := lib.NewMinHeapBinomial()
				binoHeap.Union(heaps[0])
				binoHeap.Union(heaps[1])
			}
		})
	}

	dataSizes := []int{1000, 5000, 20000, 50000, 80000, 120000, 200000}

	for _, dataSize := range dataSizes {
		keysGroups := make([][]*lib.KeyInt, 0, 2)
		for keysData := 1; keysData <= 2; keysData++ {
			path := keysDirName + "jeu_" + strconv.Itoa(keysData) +
				"_nb_cles_" + strconv.Itoa(dataSize) + ".txt"
			keysGroups = append(keysGroups, getKeysFromFile(path))
		}
		name := "cles_" + strconv.Itoa(dataSize)
		run(name, keysGroups)
	}

	run_extra := func(nbKeys uint64) {
		keysGroups := make([][]*lib.KeyInt, 0, 2)
		for j := 0; j < 2; j++ {
			keys := genAscendingKeys(nbKeys)
			keysGroups = append(keysGroups, keys)
		}
		run("cles_"+strconv.FormatUint(nbKeys, 10), keysGroups)
	}

	run_extra(250000)
	run_extra(300000)
}
