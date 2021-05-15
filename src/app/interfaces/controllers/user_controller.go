package controllers

import (
	"fmt"

	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/domain"
	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/interfaces/database"
	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, u)
}

func (controller *UserController) List(c Context) {
	users, err := controller.Interactor.ListUsers()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, users)
}

func (controller *UserController) Get(c Context) {
	id := c.Param("id")
	fmt.Println(id)
	/* user, err := controller.Interactor.GetUser()
	if err != nil {
		c.JSON(500, err)
		return
	} */
	c.Status(500)
}
