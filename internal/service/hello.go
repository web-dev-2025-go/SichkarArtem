package service

import "strconv"

type HelloService struct{}

func (h *HelloService) Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello " + name
}

func (h *HelloService) TypeOf(number string) int {
	typeOf, err := strconv.Atoi(number)
	if err != nil {
		typeOf = 0
	}
	typeOf = typeOf * 3
	return typeOf
}
