package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}

	fmt.Printf("%v\n",e)
	e.id = 0
	e.name ="Name"

	fmt.Printf("%v\n",e)
}