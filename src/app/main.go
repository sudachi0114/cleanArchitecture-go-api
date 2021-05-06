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

	// user_repo := database.UserRepository{}
	// id, err := user_repo.Store(taro) <- contoller 層を経由して呼び出していないため、実装に結びついていない。なので、実際に走る処理は nil になる。
	// if err != nil {
	// 	panic(err.Error)
	// }

	uc := controllers.NewUserController(infrastructure.NewSqlHandler()) // <- こっちは controller で infrastructure.SqlHandler と database.SqlHandler が結びつくので nil にならない
	uc.Create(taro)
}
