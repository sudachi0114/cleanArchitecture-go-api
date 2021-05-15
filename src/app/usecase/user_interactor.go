package usecase

import "github.com/sudachi0114/cleanArchitecture-go-api/src/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) ListUsers() (users domain.Users, err error) {
	users, err = interactor.UserRepository.List()
	return
}

func (interactor *UserInteractor) GetUser(id int) (users domain.User, err error) {
	users, err = interactor.UserRepository.FindById(id)
	return
}
