package main

import (
	"fmt"

	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/domain"
	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/infrastructure"
)

func main() {
	fmt.Println("Hello, cleanArchitecture!")

	user := domain.User{ID: 1, FirstName: "suzuki", LastName: "taro"}
	fmt.Println(user)

	conn := infrastructure.NewSqlHandler()
	fmt.Println(conn)
}
