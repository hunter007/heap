package main

import (
	"fmt"

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
	// es = append(es, &element{})
	for k, v := range esMap {
		e := &element{heap.Element(v), k}
		es = append(es, e)
	}

	fmt.Println("\nOld:")
	for _, e := range es {
		fmt.Println(e)
	}

	h := heap.NewMaxPriorityQueue[*element](es)
	fmt.Println("\nMax Queue:")
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Println(e)
	}

	fmt.Printf("\nInsert yang(200): Size=%d\n-----------------------\n", h.Size())
	h.Insert(&element{heap.Element(200), "yang"})
	fmt.Printf("After insert yang(200): Size=%d\n-----------------------\n", h.Size())

	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Printf("%d: %s\n", i, e)
	}
	fmt.Printf("=====Max: %s\n=====Extract: %s\n", h.Max(), h.Extract())

	fmt.Println("\nInsert qin(65)\n-----------------------")
	h.Insert(&element{heap.Element(65), "qin"})
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Printf("%d: %s\n", i, e)
	}
	// fmt.Printf("=====Max: %s\n=====Extract: %s\n", h.Max(), h.Extract())

	_ = h.IncreaseKey(5, 300)
	fmt.Println("\nIncreaseKey Max Queue:")
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Printf("%d: %s\n", i, e)
	}
	fmt.Printf("\nIncreaseKey 80:\n%s\n%s\n", h.Max(), h.Extract())
}

type element struct {
	heap.Elementer

	name string
}

func (e *element) String() string {
	return fmt.Sprintf("%s(%d)", e.name, e.GetKey())
}
