package main

import "fmt"

type Person struct {
	name string
	edad int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	//campo anonimos
	Person
	Employee
	endDate string
}

type PrintInfo interface {
	GetMessage() string
}

type TemporalEmployee struct {
	Person
	Employee
	taxRate int
}

func (ftEmployee FullTimeEmployee) GetMessage() string {
	return "Full Time Employee\n"
}

func (tEmployee TemporalEmployee) GetMessage() string {
	return "Temporal Time Employee\n"
}

func GetMessage(p PrintInfo) {
	fmt.Printf(p.GetMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Nmae"
	ftEmployee.edad = 24
	ftEmployee.id = 5

	tEmployee := TemporalEmployee{}
	GetMessage(tEmployee)
	GetMessage(ftEmployee)
}