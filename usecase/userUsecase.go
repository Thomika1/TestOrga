package usecase

import "github.com/Thomika1/TestOrga/model"

type UserUsecase struct {
	//repository
}

func NewUserUsecase() UserUsecase {
	return UserUsecase{}
}

func (pu *UserUsecase) GetUser() ([]model.User, error) {
	return []model.User{}, nil

}
