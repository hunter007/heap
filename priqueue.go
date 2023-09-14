package deap

import (
	"errors"
	"math"
)

type MaxPriorityQueue[E Elementer] interface {
	MaxDeaper[E]

	Extract() E
	IncreaseKey(i, key int) error
	Insert(e E) error
}

type MinPriorityQueue[E Elementer] interface {
	MinDeaper[E]

	Extract() E
	IncreaseKey(i, key int) error
	Insert(e E) error
}

func (d *deap[E]) Extract() E {
	max := d.values[1]
	d.values = d.values[1:]
	d.Heapify()
	return max
}

func (d *deap[E]) IncreaseKey(i, key int) error {
	if d.max {
		return d.maxIncreaseKey(i, key)
	}
	return d.minIncreaseKey(i, key)
}

func (d *deap[E]) maxIncreaseKey(i, key int) error {
	if key <= d.values[i].GetKey() {
		return errors.New("key is smaller")
	}

	d.values[i].SetKey(key)
	for p := parent(i); i > 1; p = parent(i) {
		if d.values[p].GetKey() < d.values[i].GetKey() {
			d.values[0] = d.values[i]
			d.values[i] = d.values[p]
			d.values[p] = d.values[0]
		}
		i = p
	}
	return nil
}

func (d *deap[E]) minIncreaseKey(i, key int) error {
	if key >= d.values[i].GetKey() {
		return errors.New("key is bigger")
	}

	d.values[i].SetKey(key)
	for p := parent(i); i > 1; p = parent(i) {
		if d.values[p].GetKey() > d.values[i].GetKey() {
			d.values[0] = d.values[i]
			d.values[i] = d.values[p]
			d.values[p] = d.values[0]
		}
		i = p
	}
	return nil
}

func (d *deap[E]) Insert(e E) error {
	if d.max {
		return d.maxInsert(e)
	}

	return d.minInsert(e)
}

func (d *deap[E]) maxInsert(e E) error {
	key := e.GetKey()
	if key == math.MinInt {
		return errors.New("key too small")
	}
	e.SetKey(math.MinInt)
	d.values = append(d.values, e)
	return d.maxIncreaseKey(d.Size(), key)
}

func (d *deap[E]) minInsert(e E) error {
	key := e.GetKey()
	if key == math.MaxInt {
		return errors.New("key too big")
	}

	e.SetKey(math.MaxInt)
	d.values = append(d.values, e)
	return d.minIncreaseKey(d.Size(), key)
}
