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
