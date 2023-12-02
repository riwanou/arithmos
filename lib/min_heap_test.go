package lib_test

import (
	"arithmos/lib"
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
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

func genKeys() []*lib.KeyInt {
	return []*lib.KeyInt{
		lib.NewKeyInt(0, 10),
		lib.NewKeyInt(0, 20),
		lib.NewKeyInt(0, 30),
		lib.NewKeyInt(0, 40),
		lib.NewKeyInt(0, 50),
	}
}

func runTestHeaps(test func(lib.MinHeap)) {
	heaps := []lib.MinHeap{lib.NewMinHeapTree()}
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

	var keys []*lib.KeyInt
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

func benchmarkHeaps(
	b *testing.B,
	bench func(heap lib.MinHeap, keys []*lib.KeyInt),
) {
	debug.SetGCPercent(800)
	dirName := "../data/cles_alea/"
	dirEntries, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	ignore := []string{
		"jeu_1_nb_cles_200000.txt",
		"jeu_2_nb_cles_200000.txt",
		"jeu_3_nb_cles_200000.txt",
		"jeu_4_nb_cles_200000.txt",
		"jeu_5_nb_cles_200000.txt",
		"jeu_1_nb_cles_120000.txt",
		"jeu_2_nb_cles_120000.txt",
		"jeu_3_nb_cles_120000.txt",
		"jeu_4_nb_cles_120000.txt",
		"jeu_5_nb_cles_120000.txt",
	}

	for _, entry := range dirEntries {
		if !slices.Contains(ignore, entry.Name()) {
			b.Run("heapTree/"+entry.Name(), func(b *testing.B) {
				keys := getKeysFromFile(dirName + entry.Name())
				for n := 0; n < b.N; n++ {
					bench(lib.NewMinHeapTree(), keys)
				}
			})
		}
	}
}

func BenchmarkAjout(b *testing.B) {
	benchmarkHeaps(b, func(heap lib.MinHeap, keys []*lib.KeyInt) {
		heap.AjoutIteratif(keys)
	})
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
		heap.AjoutIteratif(keys)
		assert.Equal(t, "[0-10, 0-20, 0-30, 0-40, 0-50]", heap.String())
	})
}

func TestSupprMin(t *testing.T) {
	runTestHeaps(func(heap lib.MinHeap) {
		keys := genKeys()
		heap.Ajout(keys[0])
		assert.Equal(t, keys[0], heap.SupprMin())
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
		assert.Equal(t, keys[0], heap.SupprMin())
		vizHeap(heap, "heap-2")
		assert.Equal(t, keys[1], heap.SupprMin())
		vizHeap(heap, "heap-3")
		assert.Equal(t, keys[2], heap.SupprMin())
		vizHeap(heap, "heap-4")
		assert.Equal(t, keys[3], heap.SupprMin())
		vizHeap(heap, "heap-5")
		assert.Equal(t, keys[4], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
		heap.Ajout(keys[0])
		heap.Ajout(keys[1])
		assert.Equal(t, keys[0], heap.SupprMin())
		assert.Equal(t, keys[1], heap.SupprMin())
		assert.Nil(t, heap.SupprMin())
	})
}
