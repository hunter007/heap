package main

import (
	"fmt"
	"os"

	"github.com/hunter007/heap"
)

func main() {
	esMap := map[string]int{
		"zheng": 94,
		"sun":   98,
		"zhou":  96,
		"qian":  99,
		"li":    97,
		"wu":    95,
		"zhao":  100,
		"wang":  93,
	}

	es := make([]*element, 0, len(esMap))
	for k, v := range esMap {
		e := &element{heap.NewElement(v), k}
		es = append(es, e)
	}

	fmt.Println("\nOld:")
	for _, e := range es {
		fmt.Println(e)
	}

	h, err := heap.NewMinHeap[*element](es)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(-1)
	}
	fmt.Println("\nMin Heap:")
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Println(e)
	}

	es = h.Sort()
	fmt.Println("\nSorted:")
	for _, e := range es {
		fmt.Println(e)
	}

	fmt.Println("\nNow:")
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Println(e)
	}
}

type element struct {
	*heap.Element

	name string
}

func (e *element) ID() string {
	return e.name
}

func (e *element) String() string {
	return fmt.Sprintf("%s(%d)", e.name, e.GetKey())
}
