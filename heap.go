package heap

import (
	"errors"
)

func NewMaxHeap[E Elementer](arrA []E) (MaxHeaper[E], error) {
	return newHeap(arrA, true)
}

func NewMinHeap[E Elementer](arrA []E) (MinHeaper[E], error) {
	return newHeap(arrA, false)
}

func NewMaxPriorityQueue[E Elementer](arrA []E) (MaxPriorityQueue[E], error) {
	return newHeap(arrA, true)
}

func NewMinPriorityQueue[E Elementer](arrA []E) (MinPriorityQueue[E], error) {
	return newHeap(arrA, false)
}

var errDup = errors.New("dup element")

func newHeap[E Elementer](arrA []E, max bool) (*heap[E], error) {
	a := make([]E, 1, len(arrA)+1)
	a = append(a, arrA...)

	d := &heap[E]{
		values: a,
		m:      map[string]E{},
		max:    max,
	}
	for i := 1; i < len(a); i++ {
		id := a[i].ID()
		if _, exists := d.m[id]; exists {
			if exists {
				return nil, errDup
			}
		}
		d.m[id] = a[i]
	}

	d.Heapify()
	return d, nil
}

type heap[E Elementer] struct {
	values []E
	m      map[string]E
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
