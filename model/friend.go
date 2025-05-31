package model

import "fmt"

type Friend struct {
	Person
	Music []string
}

type IFriend interface {
	IPerson
	GetMusic() string
	AddMusic(music []string) string
}

func NewFriend(name string, age uint, educationalInstitution string, music []string) IFriend {
	return &Friend{
		Person: Person{
			name,
			age,
			educationalInstitution,
		},
		Music: music,
	}
}

func (f Friend) String() string {
	return fmt.Sprintf("Name: %s\nAge: %d\nEducation: %s\nMusic:  %s\n", f.Name, f.Age, f.EducationalInstitution, f.Music)
}

func (f Friend) Describe() {
	fmt.Println(f)
}

func (f *Friend) AddMusic(music []string) string {
	f.Music = append(f.Music, music...)
	fmt.Printf("len: %v,cap:%v", len(f.Music), cap(f.Music))
	if len(f.Music) < cap(f.Music) {
		f.Music = f.Music[:len(f.Music):len(f.Music)]
	}

	return fmt.Sprintf("Updated music list: %v", f.Music)
}

func (f *Friend) GetMusic() string {
	return fmt.Sprintf("Music list: %v", f.Music)
}
