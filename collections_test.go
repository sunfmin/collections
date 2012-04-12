package collections

import (
	"fmt"
)

func ExampleStringContainsAny() {
	fmt.Println(StringContainsAny([]string{"1", "2", "3"}, "2"))
	//Output: true
}

func ExampleInt64Select() {
	r := Int64Select([]int64{1, 2, 3, 4}, func(i int64) bool { return i%2 == 0 })
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

func ExampleCollect() {
	source := []*Person{
		{"Juice", true},
		{"Felix", false},
		{"Bin", true},
	}
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

	Collect(source, m, &pigs)

	fmt.Println(len(pigs))
	fmt.Println(pigs[1].Name)
	//Output:
	//2
	//Bin

}

func ExampleFind() {
	people := []*Person{
		{"Juice", true},
		{"Felix", false},
		{"Bin", true},
	}
	var found *Person
	Find(people, func(v interface{}) bool {
		p := v.(*Person)
		if p.Name == "Felix" {
			return true
		}
		return false
	}, &found)

	fmt.Println(found.Name)
	//Output: Felix
}

func ExampleStringInGroupsOf() {
	numbers := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}
	r := StringInGroupsOf(numbers, 2)
	fmt.Println(r[0])
	fmt.Println(r[1])
	fmt.Println(r[2])

	r = StringInGroupsOf(numbers, 3)
	fmt.Println(r[0])
	fmt.Println(r[1])
	//Output:
	//[1 2]
	//[3 4]
	//[5]
	//[1 2 3]
	//[4 5]
}

func ExampleRandomPick() {
	source := []*Person{
		{"Juice", true},
		{"Felix", false},
		{"Bin", true},
	}
	var rs []*Person
	var r *Person
	for i := 0; i < 10; i++ {
		RandomPick(source, &r)
		rs = append(rs, r)
	}
	fmt.Println(len(rs))
	//Output: 10
}

func ExampleIntShuffle() {
	source := []int{1, 2, 3, 4, 5}
	r := IntShuffle(source)
	fmt.Println(len(r))
	//Output: 5
}

func ExampleShuffle() {
	source := []*Person{
		{"Juice", true},
		{"Felix", false},
		{"Bin", true},
	}
	var rs []*Person
	Shuffle(source, &rs)
	fmt.Println(len(rs))
	//Output: 3
}
