package deap

func (d *deap[E]) Max() E {
	return d.values[1]
}

func (d *deap[E]) maxHeapify(i, size int) {
	left, right, max := left(i), right(i), 0

	if left <= size && d.values[left].GetKey() > d.values[i].GetKey() {
		max = left
	} else {
		max = i
	}

	if right <= size && d.values[right].GetKey() > d.values[max].GetKey() {
		max = right
	}

	if max != i {
		e := d.values[max]
		d.values[max] = d.values[i]
		d.values[i] = e
		d.maxHeapify(max, size)
	}
}

func (d *deap[E]) maxSort() {
	for i := d.Size(); i > 1; i-- {
		d.values[0] = d.values[1]
		d.values[1] = d.values[i]
		d.values[i] = d.values[0]

		d.maxHeapify(1, i-1)
	}
}
