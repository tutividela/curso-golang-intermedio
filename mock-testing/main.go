package main

import "time"

type Person struct {
	dni string
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

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	e,err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee,err
	}
	ftEmployee.Employee = e
	p,err := GetPersonByDni(dni)
	if err != nil {
		return ftEmployee,err
	}
	ftEmployee.Person=p
	return ftEmployee ,nil
}

 var GetPersonByDni = func(dni string) (Person, error) {
	time.Sleep(5*time.Second)
	//Select * from Person ...
	return Person{},nil
}

 var GetEmployeeById = func(id int) (Employee,error) {
	time.Sleep(5*time.Second)
	return Employee{},nil
}

