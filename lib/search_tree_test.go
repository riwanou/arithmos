package lib_test

import (
	"arithmos/lib"
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestGet(t *testing.T) {
	keys := genKeys()
	tree := lib.NewSearchTree()

	assert.Nil(t, tree.Get(keys[0]))

	tree.Insert(keys[0])
	assert.Equal(t, keys[0], tree.Get(keys[0]))

	tree.Insert(keys[4])
	tree.Insert(keys[3])
	tree.Insert(keys[2])
	tree.Insert(keys[1])
	assert.Equal(t, keys[1], tree.Get(keys[1]))
	assert.Equal(t, keys[2], tree.Get(keys[2]))
	assert.Equal(t, keys[3], tree.Get(keys[3]))
	assert.Equal(t, keys[4], tree.Get(keys[4]))
}

/**
 * Shakespeare
 */

const shakespeareDir = "../data/Shakespeare/"

func getShakespeareWords(wordCallback func(word string, filename string)) {
	dirEntries, err := os.ReadDir(shakespeareDir)
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		f, err := os.Open(shakespeareDir + entry.Name())
		if err != nil {
			panic(err)
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		for s.Scan() {
			wordCallback(s.Text(), entry.Name())
		}
	}
}

func getShakespeareUniqueWordsFiles() [][]string {
	wordSet := lib.NewSearchTree()
	words := make([][]string, 0)
	files := make([]string, 0)
	fileIndex := -1

	getShakespeareWords(func(word string, filename string) {
		if !slices.Contains(files, filename) {
			files = append(files, filename)
			words = append(words, make([]string, 0))
			fileIndex++
		}

		hash := lib.MD5([]byte(word))
		key := lib.NewKeyIntFromBytes(hash)

		// not already here
		if wordSet.Get(key) == nil {
			wordSet.Insert(key)
			words[fileIndex] = append(words[fileIndex], word)
		}
	})

	return words
}

func getShakespeareUniqueWords() []string {
	words := make([]string, 0)
	wordsFiles := getShakespeareUniqueWordsFiles()

	for _, wordsFile := range wordsFiles {
		words = append(words, wordsFile...)
	}

	return words
}

/**
 * Tests
 */

func TestShakespeareUniqueWords(t *testing.T) {
	words := getShakespeareUniqueWords()
	assert.Equal(t, 23086, len(words))
}

func TestShakespeareUniqueCollisionWords(t *testing.T) {
	words := make([]string, 0)
	wordSet := lib.NewSearchTree()
	collisionWords := make([]string, 0)

	getShakespeareWords(func(word string, _ string) {
		hash := lib.MD5([]byte(word))
		key := lib.NewKeyIntFromBytes(hash)

		// not already here
		if wordSet.Get(key) == nil {
			wordSet.Insert(key)
			words = append(words, word)
		} else {
			// collision on the key, check if the words already exist
			if !slices.Contains(words, word) {
				collisionWords = append(collisionWords, word)
			}
		}
	})

	assert.Equal(t, 0, len(collisionWords))
}

/**
 * Benchmarks
 */

func getShakespeareKeysFiles() [][]*lib.KeyInt {
	wordsFiles := getShakespeareUniqueWordsFiles()
	keysFiles := make([][]*lib.KeyInt, 0, len(wordsFiles))

	for _, words := range wordsFiles {
		keys := make([]*lib.KeyInt, 0, len(words))
		for _, word := range words {
			hash := lib.MD5([]byte(word))
			keys = append(keys, lib.NewKeyIntFromBytes(hash))
		}
		keysFiles = append(keysFiles, keys)
	}

	return keysFiles
}

func getShakespeareKeys() []*lib.KeyInt {
	filesKeys := getShakespeareKeysFiles()
	keys := make([]*lib.KeyInt, 0)

	for _, fileKeys := range filesKeys {
		keys = append(keys, fileKeys...)
	}

	return keys
}

func benchmarkHeapsWords(b *testing.B, size int, bench func(heap lib.MinHeap)) {
	name := "cles_" + strconv.Itoa(size)
	b.Run("heapTree/"+name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bench(lib.NewMinHeapTree())
		}
	})
	b.Run("heapArray/"+name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bench(lib.NewMinHeapArray())
		}
	})
	b.Run("heapBinomial/"+name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bench(lib.NewMinHeapBinomial())
		}
	})
}

func BenchmarkAjoutWords(b *testing.B) {
	keys := getShakespeareKeys()
	benchmarkHeapsWords(b, len(keys), func(heap lib.MinHeap) {
		for _, key := range keys {
			heap.Ajout(key)
		}
	})
}

func BenchmarkConstructionWords(b *testing.B) {
	keys := getShakespeareKeys()
	benchmarkHeapsWords(b, len(keys), func(heap lib.MinHeap) {
		heap.Construction(keys)
	})
}

func BenchmarkSupprMinWords(b *testing.B) {
	keys := getShakespeareKeys()
	benchmarkHeapsWords(b, len(keys), func(heap lib.MinHeap) {
		heap.Construction(keys)
		for i := 0; i < len(keys); i++ {
			heap.SupprMin()
		}
	})
}

func BenchmarkUnionWords(b *testing.B) {
	keysFiles := getShakespeareKeysFiles()

	size := 0
	for _, keys := range keysFiles {
		size += len(keys)
	}
	name := "cles_" + strconv.Itoa(size)

	b.Run("heapBinomial/"+name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			binoHeap := lib.NewMinHeapBinomial()
			for _, keysFile := range keysFiles {
				heap := lib.NewMinHeapBinomial()
				heap.Construction(keysFile)
				binoHeap.Union(heap)
			}
		}
	})

	b.Run("heapTree/"+name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			treeHeap := lib.NewMinHeapTree()
			for _, keysFile := range keysFiles {
				heap := lib.NewMinHeapTree()
				heap.Construction(keysFile)
				treeHeap = lib.HeapTreeUnion(treeHeap, heap)
			}
		}
	})

	b.Run("heapArray/"+name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			arrayHeap := lib.NewMinHeapArray()
			for _, keysFile := range keysFiles {
				heap := lib.NewMinHeapArray()
				heap.Construction(keysFile)
				arrayHeap = lib.HeapArrayUnion(arrayHeap, heap)
			}
		}
	})
}
