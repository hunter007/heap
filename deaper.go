package deap

type deaper[E Elementer] interface {
	Length() int
	Size() int
	Get(i int) E
	Heapify()
	Sort()
}

type MaxDeaper[E Elementer] interface {
	deaper[E]
	Max() E
}

type MinDeaper[E Elementer] interface {
	deaper[E]
	Min() E
}
