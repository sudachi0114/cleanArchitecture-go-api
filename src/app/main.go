package main

import (
	"fmt"

	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/domain"
	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/infrastructure"
	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/interfaces/controllers"
)

func main() {
	fmt.Println("Hello, cleanArchitecture!")

	taro := domain.User{ID: 1, FirstName: "suzuki", LastName: "taro"}
	fmt.Println(taro)

	conn := infrastructure.NewSqlHandler()
	fmt.Println(conn)

	uc := controllers.NewUserController(infrastructure.NewSqlHandler())
	uc.Create(taro)
}
