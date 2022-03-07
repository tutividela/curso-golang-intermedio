package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func newEmployee(id int , nombre string) *Employee{
	return &Employee{
		id: id,
		name: nombre,

	}
}

func main() {
	e := Employee{}
	e.id = 0
	e.name ="Name"

	e1:= Employee{
		id: 1,
		name: "Name",
	}
	
	e3 := new(Employee)

	fmt.Printf("%v , %T\n",*e3,e3)

	fmt.Printf("%v\n",e)
	

//	fmt.Printf("%v\n",e)

	e4 := newEmployee(1,"Tuti")
	fmt.Printf("%v\n",*e4)
}