package heap

import (
	"errors"
	"math"
)

type priorityQueue[E Elementer] interface {
	Extract() E
	IncreaseKey(i, key int) error
	Insert(e E)
}

type MaxPriorityQueue[E Elementer] interface {
	MaxHeaper[E]
	priorityQueue[E]
}

type MinPriorityQueue[E Elementer] interface {
	MinHeaper[E]
	priorityQueue[E]
}

func (d *heap[E]) Extract() E {
	v := d.values[1]
	d.values = d.values[1:]
	d.Heapify()
	return v
}

func (d *heap[E]) IncreaseKey(i, key int) error {
	if i > d.Size() {
		return errors.New("no this element")
	}

	if d.max {
		return d.maxIncreaseKey(i, key)
	} else {
		return d.minIncreaseKey(i, key)
	}
}

func (d *heap[E]) maxIncreaseKey(i, key int) error {
	if key <= d.values[i].GetKey() {
		return errors.New("key is smaller")
	}

	d.values[i].SetKey(key)
	for p := parent(i); i > 1; p = parent(i) {
		if d.values[p].GetKey() >= d.values[i].GetKey() {
			return nil
		}

		d.values[0] = d.values[i]
		d.values[i] = d.values[p]
		d.values[p] = d.values[0]
		i = p
	}
	return nil
}

func (d *heap[E]) minIncreaseKey(i, key int) error {
	if key >= d.values[i].GetKey() {
		return errors.New("key is bigger")
	}

	d.values[i].SetKey(key)
	for p := parent(i); i > 1; p = parent(i) {
		if d.values[p].GetKey() <= d.values[i].GetKey() {
			return nil
		}

		d.values[0] = d.values[i]
		d.values[i] = d.values[p]
		d.values[p] = d.values[0]
		i = p
	}
	return nil
}

func (d *heap[E]) Insert(e E) {
	if d.max {
		d.maxInsert(e)
	} else {
		d.minInsert(e)
	}
}

func (d *heap[E]) maxInsert(e E) {
	key := e.GetKey()
	e.innerSetKey(math.MinInt)
	d.values = append(d.values, e)
	d.maxIncreaseKey(d.Size(), key)
}

func (d *heap[E]) minInsert(e E) {
	key := e.GetKey()
	e.innerSetKey(math.MaxInt)
	d.values = append(d.values, e)
	d.minIncreaseKey(d.Size(), key)
}
