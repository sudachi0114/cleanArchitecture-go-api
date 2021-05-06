package controllers

import (
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

func (controller *UserController) Create(u domain.User) { // 本当はダメ..
	err := controller.Interactor.Add(u)
	if err != nil {
		panic(err)
	}
}
