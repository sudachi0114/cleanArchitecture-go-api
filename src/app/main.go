package main

import (
	"fmt"

	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/domain"
)

func main() {
	fmt.Println("Hello, cleanArchitecture!")

	user := domain.User{ID: 1, FirstName: "suzuki", LastName: "taro"}
	fmt.Println(user)
}
