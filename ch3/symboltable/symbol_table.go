package symboltable

type (
	Key   = string
	Value = int
)

type SymbolTable interface {
	Put(key Key, value Value)
	Get(key Key) (value Value, exists bool)
	Delete(key Key)
	Contains(key Key) bool
	IsEmpty() bool
	Size() int
	Keys() []Key
}

type OrderedSymbolTable interface {
	SymbolTable

	Min() (key Key, exists bool)
	Max() (key Key, exists bool)

	Floor(key Key) (Key, bool)
	Ceiling(key Key) (Key, bool)

	Rank(key Key) int
	Select(k int) (key Key, exists bool)

	DeleteMin()
	DeleteMax()

	SizeBetween(lo, hi Key) int
	KeysBetween(lo, hi Key) []Key
}
