package deap

func (d *deap[E]) Min() E {
	return d.values[1]
}

func (d *deap[E]) minHeapify(i, size int) {
	left, right, min := left(i), right(i), 0

	if left <= size && d.values[left].GetKey() < d.values[i].GetKey() {
		min = left
	} else {
		min = i
	}

	if right <= size && d.values[right].GetKey() < d.values[min].GetKey() {
		min = right
	}

	if min != i {
		d.values[0] = d.values[min]
		d.values[min] = d.values[i]
		d.values[i] = d.values[0]
		d.minHeapify(min, size)
	}
}

func (d *deap[E]) minSort() {
	for i := d.Size(); i > 1; i-- {
		d.values[0] = d.values[1]
		d.values[1] = d.values[i]
		d.values[i] = d.values[0]

		d.minHeapify(1, i-1)
	}
}
