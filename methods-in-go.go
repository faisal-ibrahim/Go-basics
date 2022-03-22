package main

import "fmt"

type Employee struct {
	Name            string
	Age             int
	EmployeeProject EmployeeProject
}

type EmployeeProject struct {
	Name         string
	Priority     string
	Technologies []string
}

func (e Employee) Print() {
	fmt.Println("Employee Information")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.EmployeeProject.Name)
}

func (e *Employee) UpdateAge() {
	e.Age += 5
}

func (e *Employee) GetProjectPriority() string {
	return e.EmployeeProject.Priority
}

func main() {
	fmt.Println("Methods in go")
	employee := Employee{
		Name: "Elliot",
		Age:  30,
		EmployeeProject: EmployeeProject{
			Name:         "Beginner's Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go", "PHP"},
		},
	}

	employee.UpdateAge()
	employee.Print()

	fmt.Println(employee.GetProjectPriority())
}
