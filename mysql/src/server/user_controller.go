package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mysql/src/model"
	"mysql/src/repository"
	"net/http"
)

type UserController struct {
	rep repository.UserRepository
}

type AddUserRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type DeleteUserRequest struct {
	Id *int `json:"id"`
}

type GetUserRequest struct {
	Id *int `json:"id"`
}

type RenameUserRequest struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

func CreateUserController(rep repository.UserRepository) *UserController {
	return &UserController{
		rep: rep,
	}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users, _ := uc.rep.GetUsers()
	uj, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req GetUserRequest
	decoder.Decode(&req)

	u, _ := uc.rep.GetUser(*req.Id)

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req AddUserRequest
	decoder.Decode(&req)

	u := &model.User{Name: *req.Name, Age: *req.Age}
	uc.rep.AddUser(u)
	log.Printf("User(%v) was added\n", *u)

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req DeleteUserRequest
	json.NewDecoder(r.Body).Decode(&req)

	uc.rep.DeleteUser(*req.Id)
	log.Printf("User(id: %d) was deleted\n", *req.Id)

	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) RenameUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req RenameUserRequest
	decoder.Decode(&req)

	uc.rep.RenameUser(*req.Id, *req.Name)
	log.Printf("User(id: %d) was renamed to %s\n", *req.Id, *req.Name)

	w.WriteHeader(http.StatusOK)
}
