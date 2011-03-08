package main

import (
	. "maps"
	"fmt"
)


//---
type Person struct {
	Name    string
	Surname string
	Age     int
}

func (p *Person) Key() interface{} {
	return p.Name + "#" + p.Surname
}

//---


func main() {

	p1 := &Person{"John", "Rambo", 60}
	p2 := &Person{"John", "Doe", 30}
	p3 := &Person{"Jane", "Doe", 30}

	//Create map:
	M := make(Map)

	//Insert into map:
	M.Insert(p1, "Manager")
	M.Insert(p2, "Employee")
	M.Insert(p3, "Secretary")

	//Action to perform during iteration:
	printValues := func(v Pair) {
		fmt.Printf("Person, Job = %v, %s\n", v.Key, v.Value)
	}

	//Iterate over Map:
	M.Do(printValues)

	println("---")

	//Query map:
	p := &Person{"John", "Rambo", 60}
	Value, Ok := M.Get(p)
	if Ok {
		fmt.Printf("Job = %s\n", Value)
	}

	println("---")

	//Delete from map
	M.Delete(p1)

	//Check the result of the deletion,
	//second iteration over map:
	M.Do(printValues)

	fmt.Printf("Size of map M = %d\n", M.Len())

}
