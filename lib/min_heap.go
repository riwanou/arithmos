package lib

type MinHeap interface {
	SupprMin() *KeyInt
	Ajout(key KeyInt)
	AjoutsIteratif(keys []KeyInt)
	String() string
}
