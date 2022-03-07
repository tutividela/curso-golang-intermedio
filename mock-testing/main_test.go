package main

import "testing"

func TestMock(t *testing.T) {
	table := []struct{
		id int 
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id: 1,
			dni:"1",
			mockFunc: func ()  {
				GetEmployeeById = func(id int) (Employee,error){
					return Employee{
						id: id,
						
					},nil
				}
				GetPersonByDni=  func (dni string) (Person,error) {
					return Person{
						dni: dni,
					},nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					dni: "1",
				},
				Employee: Employee{
					id: 1,
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDni

	for _,v := range table{
		v.mockFunc()
		ft,err := GetFullTimeEmployeeById(v.id ,v.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}
		if ft.id != v.id {
			t.Errorf("Error , got %v expected %v",v.id , ft.id)
		}
		if ft.dni != v.dni {
			t.Errorf("Error , got %v expected %v",v.dni , ft.dni)
		}
	}
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDni =originalGetPersonByDni
}