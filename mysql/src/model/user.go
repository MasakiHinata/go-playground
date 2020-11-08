package model

type User struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:32"`
	Age  int
}
