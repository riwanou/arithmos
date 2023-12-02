package lib_test

import (
	"arithmos/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinomialTreeUnion(t *testing.T) {
	keys := genKeys()

	tree := lib.NewBinomialTree(keys[0])
	assert.Equal(t, "(0-10)", tree.String())
	tree.Union(lib.NewBinomialTree(keys[1]))
	assert.Equal(t, "(0-10, (0-20))", tree.String())

	tree2 := lib.NewBinomialTree(keys[2])
	assert.Equal(t, "(0-30)", tree2.String())
	tree2.Union(lib.NewBinomialTree(keys[3]))
	assert.Equal(t, "(0-30, (0-40))", tree2.String())

	cp_tree := *tree
	cp_tree2 := *tree2

	tree.Union(&cp_tree2)
	assert.Equal(t, "(0-10, (0-20), (0-30, (0-40)))", tree.String())

	vizBytes(tree.Viz(), "binom_tree")

	tree2.Union(&cp_tree)
	assert.Equal(t, "(0-30, (0-40), (0-10, (0-20)))", tree2.String())

	tree.Union(tree2)
	assert.Equal(t,
		"(0-10, (0-20), (0-30, (0-40)), (0-30, (0-40), (0-10, (0-20))))",
		tree.String())
}
