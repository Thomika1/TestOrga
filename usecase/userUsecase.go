package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/Thomika1/TestOrga/model"
	"github.com/Thomika1/TestOrga/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func GenerateJWT(email string, user_id int) (string, error) {
	// Pega a secret da variável de ambiente
	secret := os.Getenv("jwtSecret")
	if secret == "" {
		return "", jwt.ErrInvalidKey
	}

	// Define os dados do token
	claims := jwt.MapClaims{
		"email":   email,
		"user_id": user_id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	// Cria e assina o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
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
		return nil, err
	}

	return user, nil

}

func (pu *UserUsecase) UserLogin(user_email string, plain_password string) (string, error) {
	password_hash, user_id, err := pu.repository.UserLogin(user_email, plain_password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(plain_password))
	if err != nil {
		return "", fmt.Errorf("senha incorreta")
	}

	token, err := GenerateJWT(user_email, user_id)
	if err != nil {
		return "", err
	}

	return token, nil
}
