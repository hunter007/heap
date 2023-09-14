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

	es := make([]*element2, 0, len(esMap))
	for k, v := range esMap {
		e := &element2{heap.NewElement(v), k}
		es = append(es, e)
	}

	fmt.Println("\nOld:")
	for _, e := range es {
		fmt.Println(e)
	}

	h, err := heap.NewMaxPriorityQueue[*element2](es)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(-1)
	}
	fmt.Println("\nMax Queue:")
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Println(e)
	}

	fmt.Printf("\nInsert yang(200): Size=%d\n-----------------------\n", h.Size())
	h.Insert(&element2{heap.NewElement(200), "yang"})
	fmt.Printf("After insert yang(200): Size=%d\n-----------------------\n", h.Size())

	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Printf("%d: %s\n", i, e)
	}
	fmt.Printf("=====Max: %s\n=====Extract: %s\n", h.Max(), h.Extract())

	fmt.Println("\nInsert qin(65)\n-----------------------")
	h.Insert(&element2{heap.NewElement(65), "qin"})
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Printf("%d: %s\n", i, e)
	}

	_ = h.IncreaseKey("zhao", 300)
	fmt.Println("\nIncreaseKey Max Queue:")
	for i := 1; i <= h.Size(); i++ {
		e := h.Get(i)
		fmt.Printf("%d: %s\n", i, e)
	}
	fmt.Printf("\nIncreaseKey 80:\n%s\n%s\n", h.Max(), h.Extract())
}

type element2 struct {
	*heap.Element

	name string
}

func (e *element2) String() string {
	return fmt.Sprintf("%s(%d)", e.name, e.GetKey())
}

func (e *element2) ID() string {
	return e.name
}
