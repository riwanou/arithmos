package lib_test

import (
	"arithmos/lib"
	"bufio"
	"fmt"
	"os"
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

func getShakespeareWords(wordCallback func(word string)) {
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
			wordCallback(s.Text())
		}
	}
}

func TestShakespeareUniqueWords(t *testing.T) {
	wordSet := lib.NewSearchTree()
	words := make([]string, 0)
	totalWords := 0

	getShakespeareWords(func(word string) {
		hash := lib.MD5([]byte(word))
		key := lib.NewKeyIntFromBytes(hash)
		totalWords++

		// not already here
		if wordSet.Get(key) == nil {
			wordSet.Insert(key)
			words = append(words, word)
		}
	})

	assert.Equal(t, 23086, len(words))
	assert.Equal(t, 905534, totalWords)
}

func TestShakespeareUniqueCollisionWords(t *testing.T) {
	words := make([]string, 0)
	wordSet := lib.NewSearchTree()
	collisionWords := make([]string, 0)

	getShakespeareWords(func(word string) {
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

	fmt.Println(len(collisionWords))
}
