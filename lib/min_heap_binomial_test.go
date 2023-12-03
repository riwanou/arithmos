package lib_test

import (
	"arithmos/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinomialTreeUnion(t *testing.T) {
	keys := genKeys()

	tree := lib.NewBinomialTree(keys[3])
	assert.Equal(t, "(0-40)", tree.String())
	tree = lib.BinoTreeUnion(tree, lib.NewBinomialTree(keys[2]))
	assert.Equal(t, "(0-30, (0-40))", tree.String())

	second_tree := lib.NewBinomialTree(keys[1])
	assert.Equal(t, "(0-20)", second_tree.String())
	second_tree = lib.BinoTreeUnion(second_tree, lib.NewBinomialTree(keys[0]))
	assert.Equal(t, "(0-10, (0-20))", second_tree.String())

	tree = lib.BinoTreeUnion(tree, second_tree)
	assert.Equal(t, "(0-10, (0-20), (0-30, (0-40)))", tree.String())
	vizBytes(tree.Viz(), "binom_tree")
}

func TestBinomialAjout(t *testing.T) {
	keys := append(genKeys(), []*lib.KeyInt{
		lib.NewKeyInt(0, 60),
		lib.NewKeyInt(0, 70),
		lib.NewKeyInt(0, 80),
	}...)

	heap := lib.NewMinHeapBinomial()
	assert.Equal(t, "[]", heap.String())
	heap.Ajout(keys[7])
	assert.Equal(t, "[(0-80)]", heap.String())
	heap.Ajout(keys[6])
	assert.Equal(t, "[(0-70, (0-80))]", heap.String())
	heap.Ajout(keys[5])
	assert.Equal(t, "[(0-60), (0-70, (0-80))]", heap.String())
	heap.Ajout(keys[4])
	assert.Equal(t, "[(0-50, (0-60), (0-70, (0-80)))]", heap.String())
	heap.Ajout(keys[3])
	assert.Equal(t, "[(0-40), (0-50, (0-60), (0-70, (0-80)))]", heap.String())
	heap.Ajout(keys[2])
	assert.Equal(t,
		"[(0-30, (0-40)), (0-50, (0-60), (0-70, (0-80)))]",
		heap.String())
	heap.Ajout(keys[1])
	assert.Equal(t,
		"[(0-20), (0-30, (0-40)), (0-50, (0-60), (0-70, (0-80)))]",
		heap.String())
	heap.Ajout(keys[0])
	assert.Equal(t,
		"[(0-10, (0-20), (0-30, (0-40)), (0-50, (0-60), (0-70, (0-80))))]",
		heap.String())
}

func TestBinomialMerge(t *testing.T) {
	keys := append(genKeys())

	heaps1 := lib.NewMinHeapBinomial()
	heaps1.Ajout(keys[0])
	heaps1.Ajout(keys[1])

	heaps2 := lib.NewMinHeapBinomial()
	heaps2.Ajout(keys[2])

	heaps3 := lib.NewMinHeapBinomial()
	heaps3.Ajout(keys[3])

	heaps4 := lib.NewMinHeapBinomial()
	heaps4.Ajout(keys[4])

	heaps3.Merge(heaps4)
	heaps3_1 := lib.NewMinHeapBinomial()
	heaps3_1.Ajout(lib.NewKeyInt(0, 100))
	heaps3.Merge(heaps3_1)
	assert.Equal(t, "[(0-100), (0-40, (0-50))]", heaps3.String())
	vizBytes(heaps3.Viz(), "binomial_heap_merge")

	heaps1.Merge(heaps2)
	assert.Equal(t, "[(0-30), (0-10, (0-20))]", heaps1.String())

	heaps1.Merge(heaps3)
	assert.Equal(t, "[(0-40, (0-50)), (0-10, (0-20), (0-30, (0-100)))]",
		heaps1.String())

}
