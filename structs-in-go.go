package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func main() {
	fmt.Println("Structs basics")

	engineer := Engineer{
		Name: "Faisal Ibrahim",
		Age:  28,
		Project: Project{
			Name:         "New project using Go",
			Priority:     "High",
			Technologies: []string{"Go", "HTML"},
		},
	}

	fmt.Printf("%+v\n", engineer)

	fmt.Println(engineer.Name)
	fmt.Println(engineer.Project.Name)
}
