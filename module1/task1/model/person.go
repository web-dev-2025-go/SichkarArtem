package model

import "fmt"

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

func NewPerson(name string, age uint8, educationalInstitution string) IPerson {
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
	return fmt.Sprintf(
		"Name: %s\nAge: %d\nEducation: %s",
		p.Name, p.Age, p.EducationalInstitution,
	)
}

func (p Person) Describe() {
	fmt.Println(p)
}

func (p *Person) ChangeName(name string) string {
	p.Name = name
	return p.Name
}
