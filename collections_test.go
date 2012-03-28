package collections

import (
	"fmt"
)

func ExampleContains() {
	fmt.Println(StringsContainAny([]string{"1", "2", "3"}, "2"))
	//Output: true
}

func ExampleFilter() {
	r := Int64sFilter([]int64{1, 2, 3, 4}, func(i int64) bool { return i%2 == 0 })
	fmt.Println(r)
	//Output: [2 4]
}

type Person struct {
	Name   string
	Stupid bool
}

type Pig struct {
	Name string
}

func people() []*Person {
	return []*Person{
		{"Juice", true},
		{"Felix", false},
		{"Bin", true},
	}
}

func ExampleMap() {
	source := people()
	var pigs []*Pig

	m := func(p interface{}) (r interface{}, add bool) {
		person := p.(*Person)
		if !person.Stupid {
			return nil, false
		}
		add = true
		r = &Pig{person.Name}
		return
	}

	Map(source, m, &pigs)

	fmt.Println(len(pigs))
	fmt.Println(pigs[1].Name)
	//Output:
	//2
	//Bin

}

func ExampleFind() {
	var found *Person
	Find(people(), func(v interface{}) bool {
		p := v.(*Person)
		if p.Name == "Felix" {
			return true
		}
		return false
	}, &found)

	fmt.Println(found.Name)
	//Output: Felix
}
