package factory

import "fmt"

// Person - used as Person entity
type Person struct {
	Name string
	Age  int
}

// GetName - used for get Person name
func (p Person) GetName() string {
	return fmt.Sprintf("I'am %s and i'am %d", p.Name, p.Age)
}

// NewPerson - used for create instance of Person entity
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}
