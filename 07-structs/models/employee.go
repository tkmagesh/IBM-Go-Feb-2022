package models

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    Organization
}

func NewEmployee(id int, name string, salary float32, orgId int, orgName string, city string) Employee {
	return Employee{
		Id:     id,
		Name:   name,
		Salary: salary,
		Org: Organization{
			Id:   orgId,
			Name: orgName,
			City: city,
		},
	}
}
