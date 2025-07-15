package usecase

import (
	"github.com/Thomika1/TestOrga/model"
	"github.com/Thomika1/TestOrga/repository"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{repository: repo}
}

func (pu *UserUsecase) GetUsers() ([]model.User, error) {
	return pu.repository.GetUsers()

}

func (pu *UserUsecase) CreateUser(user model.User) (model.User, error) {
	userId, err := pu.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userId

	return user, nil
}
