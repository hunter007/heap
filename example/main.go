package example

import (
	"fmt"

	"github.com/hunter007/deap"
)

func main() {
	es := []*element{
		{},
		{"zheng", 94},
		{"sun", 98},
		{"zhou", 96},
		{"qian", 99},
		{"li", 97},
		{"wu", 95},
		{"zhao", 100},
		{"wang", 93},
	}

	d := deap.NewMaxDeap[*element](es)

	e := &element{"yang", 200}
	d.Insert(e)
	// d.Print()
	// for d.Size() > 0 {
	// 	fmt.Println(d.Size())
	// 	m := d.ExtractAndMax()
	// 	m.Print()
	// }
}

type element struct {
	*deap.Element
	name string
}

func (e *element) String() string {
	return fmt.Sprintf("%s(%d)\n", e.name, e.GetKey())
}

func Print(d deap.Deaper) {
	fmt.Println("---------------")
	for i := 1; i <= d.Size(); i++ {
		d.values[i].String()
	}
}
