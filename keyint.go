package main

type KeyInt struct {
	u1 uint64
	u2 uint64
}

func NewKeyInt(u1 uint64, u2 uint64) *KeyInt {
	return &KeyInt{u1, u2}
}

func Inf(key *KeyInt, other *KeyInt) bool {
	if key.u1 == other.u1 {
		return key.u2 < other.u2
	}
	return key.u1 < other.u1
}

func Eq(key *KeyInt, other *KeyInt) bool {
	return key.u1 == other.u1 && key.u2 == other.u2
}
