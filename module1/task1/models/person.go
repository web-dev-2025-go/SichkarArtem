package models

import (
	"fmt"
	"log"
)

type Person struct {
	Name                   string
	Age                    uint8
	EducationalInstitution string
}
type IPerson interface {
	GetName() string
	Describe()
	ChangeName(name string) string
}

func NewPerson(name string, age uint8, educationalInstitution string) *Person {
	return &Person{
		Name:                   name,
		Age:                    age,
		EducationalInstitution: educationalInstitution,
	}
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) String() string {
	fmt.Print(3)
	return fmt.Sprintf(
		"Name: %s\nAge: %d\nEducation: %s",
		p.Name, p.Age, p.EducationalInstitution,
	)
}

func (p Person) Describe() {
	log.Println(p)
}

func (p *Person) ChangeName(name string) string {
	p.Name = name
	return p.Name
}
