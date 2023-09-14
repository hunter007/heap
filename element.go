package deap

import (
	"errors"
	"math"
)

// Elementer element for deap
type Elementer interface {
	GetKey() int
	SetKey(int) error
}

type Element struct {
	key int
}

func (e *Element) GetKey() int {
	return e.key
}

func (e *Element) SetKey(key int) error {
	if key == math.MaxInt || key == math.MinInt {
		return errors.New("key is big or small")
	}
	e.key = key
	return nil
}
