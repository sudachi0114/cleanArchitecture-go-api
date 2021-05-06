package usecase

import "github.com/sudachi0114/cleanArchitecture-go-api/src/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}
