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
	return []model.User{}, nil

}
