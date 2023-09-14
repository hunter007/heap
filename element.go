package heap

import (
	"errors"
	"math"
)

// Elementer element for heap
type Elementer interface {
	GetKey() int
	SetKey(int) error
	ID() string
	innerSetKey(key int)
}

type Element struct {
	key int
}

func (e *Element) GetKey() int {
	return e.key
}

func (e *Element) SetKey(key int) error {
	if key == math.MaxInt64 || key == math.MinInt64 {
		return errors.New("key is big or small")
	}
	e.key = key
	return nil
}

func (e *Element) innerSetKey(key int) {
	e.key = key
}

// func (e *element) ID() string {
// 	return ""
// }

func NewElement(key int) *Element {
	return &Element{
		key: key,
	}
}
