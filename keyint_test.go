package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInferiority(t *testing.T) {
	var key1 = NewKeyInt(0, 100)
	var key2 = NewKeyInt(0, 200)
	assert.Equal(t, Inf(key1, key2), true)
}

func TestEqual(t *testing.T) {
	var key1 = NewKeyInt(0, 100)
	var key2 = NewKeyInt(0, 200)
	assert.Equal(t, Eq(key1, key2), false)
}

func TestFiles(t *testing.T) {
	f, err := os.Open("./data/cles_alea/jeu_1_nb_cles_1000.txt")
	assert.NoError(t, err)
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		v, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		assert.NoError(t, err)

		v = strings.TrimSuffix(v, "\n")
		fmt.Println(v)
	}
}
