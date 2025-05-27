package main

import (
	"fmt"
	"go-http/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}
