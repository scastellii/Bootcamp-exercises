package structs

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}
type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(
		"ID:", e.ID,
		"Person ID:", e.Person.ID,
		"Position:", e.Position,
		"Employee id:", e.ID,
	)
}
