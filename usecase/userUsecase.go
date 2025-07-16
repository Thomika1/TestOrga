package usecase

import (
	"fmt"

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

func (pu *UserUsecase) GetUserById(user_email string) (*model.User, error) {
	user, err := pu.repository.GetUserByEmail(user_email)
	if err != nil {
		fmt.Println("###dentrousecase")
		return nil, err
	}

	return user, nil

}
