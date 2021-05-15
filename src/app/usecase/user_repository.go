package usecase

import "github.com/sudachi0114/cleanArchitecture-go-api/src/app/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	List() (domain.Users, error)
}
