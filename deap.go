package deap

func NewMaxDeap[E Elementer](arrA []E) MaxDeaper[E] {
	return newDeap(arrA, true)
}

func NewMinDeap[E Elementer](arrA []E) MinDeaper[E] {
	return newDeap(arrA, false)
}

func NewMaxPriorityQueue[E Elementer](arrA []E) MaxPriorityQueue[E] {
	return newDeap(arrA, true)
}

func NewMinPriorityQueue[E Elementer](arrA []E) MinPriorityQueue[E] {
	return newDeap(arrA, false)
}

func newDeap[E Elementer](arrA []E, max bool) *deap[E] {
	d := &deap[E]{
		values: arrA,
		max:    false,
	}

	d.Heapify()
	return d
}

type deap[E Elementer] struct {
	values []E
	max    bool
}

func (d *deap[E]) Length() int {
	return cap(d.values) - 1
}

func (d *deap[E]) Size() int {
	return len(d.values) - 1
}

func (d *deap[E]) Get(i int) E {
	return d.values[i]
}

func (d *deap[E]) Heapify() {
	size := d.Size()
	y := yeaf(size)

	if d.max {
		for i := y; i >= 1; i-- {
			d.maxHeapify(i, size)
		}
	} else {
		for i := y; i >= 1; i-- {
			d.minHeapify(i, size)
		}
	}
}

func (d *deap[E]) Sort() {
	if d.max {
		d.maxSort()
	}

	d.minSort()
}

func parent(i int) int {
	return i >> 1
}

func left(i int) int {
	return i << 1
}

func right(i int) int {
	return i<<1 + 1
}

func yeaf(size int) int {
	return size >> 1
}
