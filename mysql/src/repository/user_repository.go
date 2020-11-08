package repository

import "mysql/src/model"

type UserRepository interface {
	GetUsers() ([]*model.User, error)

	GetUser(id int) (*model.User, error)

	AddUser(user *model.User) error

	RenameUser(id int, name string) error

	DeleteUser(id int) error
}
