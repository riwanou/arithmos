package lib

type MinHeap interface {
	SupprMin() *KeyInt
	Ajout(key *KeyInt)
	AjoutIteratif(keys []*KeyInt)
	Construction(keys []*KeyInt)
	String() string
	Viz() []byte
}
