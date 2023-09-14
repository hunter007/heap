package heap

type Integer interface {
	comparable
	int
	byte
	int16
	int32
	int64
}

type deaper[E Elementer] interface {
	Length() int
	Size() int
	Get(i int) E
	Heapify()
	Sort() []E
}

type MaxHeaper[E Elementer] interface {
	deaper[E]
	Max() E
}

type MinHeaper[E Elementer] interface {
	deaper[E]
	Min() E
}
