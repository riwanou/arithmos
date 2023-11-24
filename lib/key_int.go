package lib

import (
	"fmt"
	"strconv"
)

// KeyInt is a 128b unsigned int composed of a low and high 64b unsigned int
type KeyInt struct {
	u1 uint64
	u2 uint64
}

// Create a key from a low and high 64b unsigned int
func NewKeyInt(u1 uint64, u2 uint64) *KeyInt {
	return &KeyInt{u1, u2}
}

// Create a key from a given hexadecimal string, the format shall be
// similar to: 0xdf6943ba6d51464f6b02157933bdd9ad
func NewKeyIntFromString(str string) (*KeyInt, error) {
	u1, err := strconv.ParseUint(str[2:18], 16, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing u1: %v", err)
	}

	u2, err := strconv.ParseUint(str[18:], 16, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing u2: %v", err)
	}

	return &KeyInt{u1, u2}, nil
}

// Return whether the current key is inferior to the given other key
func (key *KeyInt) Inf(other *KeyInt) bool {
	return key.u1 < other.u1 || (key.u1 == other.u1 && key.u2 < other.u2)
}

// Return whether the current key is equal to the given other key
func (key *KeyInt) Eq(other *KeyInt) bool {
	return key.u1 == other.u1 && key.u2 == other.u2
}
