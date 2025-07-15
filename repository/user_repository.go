package repository

import (
	"database/sql"
	"fmt"

	"github.com/Thomika1/TestOrga/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT id, email, password_hash, created_at FROM users"
	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.Email,
			&userObj.PasswordHash,
			&userObj.CreatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}
		userList = append(userList, userObj)
	}

	rows.Close()
	return userList, err

}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id int

	query, err := ur.connection.Prepare("INSERT INTO users(email, password_hash) VALUES($1,$2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Email, user.PasswordHash).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
