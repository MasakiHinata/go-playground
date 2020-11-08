package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mysql/src/model"
)

const (
	DBMS     = "mysql"
	USER     = "root"
	PASS     = "mysql"
	PROTOCOL = "tcp(localhost:3306)"
	DBNAME   = "user_schema"
)

type MySqlUserRepository struct {
	db *gorm.DB
}

func CreateDatabase() (*MySqlUserRepository, error) {
	connect := fmt.Sprintf("%s:%s@%s/%s", USER, PASS, PROTOCOL, DBNAME)
	db, err := gorm.Open(DBMS, connect)
	if err != nil {
		return nil, err
	}
	rep := &MySqlUserRepository{db: db}
	return rep, nil
}

func (rep *MySqlUserRepository) CloseDatabase() {
	_ = rep.db.Close()
}

func (rep *MySqlUserRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := rep.db.Find(&users).Error
	return users, err
}

func (rep *MySqlUserRepository) GetUser(id int) (*model.User, error) {
	var user model.User
	err := rep.db.Where("id = ?", id).Find(&user).Error
	return &user, err
}

func (rep *MySqlUserRepository) AddUser(user *model.User) error {
	err := rep.db.Save(&user).Error
	return err
}

func (rep *MySqlUserRepository) RenameUser(id int, name string) error {
	err := rep.db.Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
	return err
}

func (rep *MySqlUserRepository) DeleteUser(id int) error {
	user, err := rep.GetUser(id)
	if err != nil {
		return err
	}

	err = rep.db.Delete(*user).Error
	return err
}
