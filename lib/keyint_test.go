package lib_test

import (
	"arithmos/lib"
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInferiority(t *testing.T) {
	var key1 *lib.KeyInt
	var key2 *lib.KeyInt

	key1 = lib.NewKeyInt(0, 100)
	key2 = lib.NewKeyInt(0, 200)
	assert.Equal(t, key1.Inf(key2), true)

	key1 = lib.NewKeyInt(10, 400)
	key2 = lib.NewKeyInt(11, 200)
	assert.Equal(t, key1.Inf(key2), true)
}

func TestEqual(t *testing.T) {
	var key1 *lib.KeyInt
	var key2 *lib.KeyInt

	key1 = lib.NewKeyInt(10, 100)
	key2 = lib.NewKeyInt(10, 100)
	assert.Equal(t, key1.Eq(key2), true)

	key1 = lib.NewKeyInt(30, 100)
	key2 = lib.NewKeyInt(20, 300)
	assert.Equal(t, key1.Eq(key2), false)
}

func TestDataset1Keys1000(t *testing.T) {
	f, err := os.Open("../data/cles_alea/jeu_1_nb_cles_1000.txt")
	assert.NoError(t, err)
	defer f.Close()

	s := bufio.NewScanner(f)
	var last string

	for s.Scan() {
		curr := s.Text()
		if len(last) > 0 {
			// compare key int
			currKey, err := lib.NewKeyIntFromString(curr)
			assert.NoError(t, err)
			lastKey, err := lib.NewKeyIntFromString(last)
			assert.NoError(t, err)
			keyEqComp := currKey.Eq(lastKey)
			keyInfComp := currKey.Inf(lastKey)

			// compare hexadecimal strings
			valEqComp := curr == last
			valInfComp := curr < last

			// should yield same results
			assert.Equal(t, keyEqComp, valEqComp)
			assert.Equal(t, keyInfComp, valInfComp)
		}
		last = curr
	}
}
