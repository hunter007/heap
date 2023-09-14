package heap

import (
	"errors"
	"math"
)

// Elementer element for heap
type Elementer interface {
	GetKey() int
	SetKey(int) error
	innerSetKey(key int)
}

type element struct {
	key int
}

func (e *element) GetKey() int {
	return e.key
}

func (e *element) SetKey(key int) error {
	if key == math.MaxInt64 || key == math.MinInt64 {
		return errors.New("key is big or small")
	}
	e.key = key
	return nil
}

func (e *element) innerSetKey(key int) {
	e.key = key
}

func Element(key int) Elementer {
	return &element{
		key: key,
	}
}
