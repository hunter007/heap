package heap

func NewMaxHeap[E Elementer](arrA []E) MaxHeaper[E] {
	return newHeap(arrA, true)
}

func NewMinHeap[E Elementer](arrA []E) MinHeaper[E] {
	return newHeap(arrA, false)
}

func NewMaxPriorityQueue[E Elementer](arrA []E) MaxPriorityQueue[E] {
	return newHeap(arrA, true)
}

func NewMinPriorityQueue[E Elementer](arrA []E) MinPriorityQueue[E] {
	return newHeap(arrA, false)
}

func newHeap[E Elementer](arrA []E, max bool) *heap[E] {
	a := make([]E, 0, len(arrA)+1)
	a = append(a, arrA[0])
	a = append(a, arrA...)

	d := &heap[E]{
		values: a,
		max:    max,
	}

	d.Heapify()
	return d
}

type heap[E Elementer] struct {
	values []E
	max    bool
}

func (d *heap[E]) Length() int {
	return cap(d.values) - 1
}

func (d *heap[E]) Size() int {
	return len(d.values) - 1
}

func (d *heap[E]) Get(i int) E {
	return d.values[i]
}

func (d *heap[E]) Heapify() {
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

func (d *heap[E]) Sort() []E {
	if d.max {
		return d.maxSort()
	}

	return d.minSort()
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
