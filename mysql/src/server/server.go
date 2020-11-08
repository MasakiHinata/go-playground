package main

import (
	"log"
	"mysql/src/repository"
	"net/http"
)

func main() {
	rep, err := repository.CreateDatabase()
	if err != nil {
		log.Fatalln("Can not create database:", err)
	}
	defer rep.CloseDatabase()

	uc := CreateUserController(rep)

	http.HandleFunc("/", uc.GetUsers)
	http.HandleFunc("/user", uc.GetUser)
	http.HandleFunc("/add", uc.AddUser)
	http.HandleFunc("/rename", uc.RenameUser)
	http.HandleFunc("/delete", uc.DeleteUser)
	http.ListenAndServe(":8080", nil)
}
