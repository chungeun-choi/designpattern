package iterator


type Iterator interface {
	GetNext()
	HashNext() bool
}