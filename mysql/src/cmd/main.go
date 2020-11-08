package main

import (
	"fmt"
	"log"
	"mysql/src/model"
	"mysql/src/repository"
)

func main() {
	rep, err := repository.CreateDatabase()
	if err != nil {
		log.Fatalln("Can not create database:", err)
	}
	defer rep.CloseDatabase()

	user := model.User{Name: "Alice", Age: 20}
	err = rep.AddUser(&user)
	if err != nil {
		log.Fatalln("Can not add user:", err)
	}

	users, err := rep.GetUsers()
	if err != nil {
		log.Fatalf("There is no user.")
	}
	for _, user := range users {
		fmt.Println(*user)
	}
}
