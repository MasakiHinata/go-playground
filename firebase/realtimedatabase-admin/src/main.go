package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"fmt"
	"log"
)

type User struct {
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Name        string `json:"name,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
}

func main() {
	ctx := context.Background()

	conf := &firebase.Config{
		DatabaseURL: "https://javascriptfirebase-98e79-default-rtdb.firebaseio.com/",
	}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	ref := client.NewRef("/server")
	usersRef := ref.Child("users")
	_ = save(ctx, usersRef)
	_ = update(ctx, usersRef)
	_ = read(ctx, usersRef.Child("alice"))
	_ = readAll(ctx, usersRef)
}

func save(ctx context.Context, ref *db.Ref) error {
	// 値を設定
	john := &User{
		DateOfBirth: "August 9, 1906",
		Name:        "John",
	}
	if err := ref.Child("john").Set(ctx, john); err != nil {
		return fmt.Errorf("error setting value: %v", err)
	}

	// 複数の値を設定
	// マップのキーが自動的にオブジェクトのキーになる
	err := ref.Set(ctx, map[string]*User{
		"alice": {
			DateOfBirth: "June 23, 1912",
			Name:        "Alice",
		},
		"bob": {
			DateOfBirth: "December 9, 1906",
			Name:        "Bob",
		},
	})
	if err != nil {
		return fmt.Errorf("error setting value: %v", err)
	}

	jack := &User{
		DateOfBirth: "August 9, 1906",
		Name:        "Jack",
	}
	if _, err := ref.Push(ctx, jack); err != nil {
		return fmt.Errorf("error setting value: %v", err)
	}

	return nil
}

func update(ctx context.Context, ref *db.Ref) error {
	err := ref.Child("john").Update(ctx, map[string]interface{}{
		"nickname": "Johnie",
	})
	return err
}

func read(ctx context.Context, ref *db.Ref) error {
	var user User
	if err := ref.Get(ctx, &user); err != nil {
		return fmt.Errorf("error reading value: %v", err)
	}
	fmt.Println(user)
	return nil
}

func readAll(ctx context.Context, ref *db.Ref) error {
	var data map[string]User

	if err := ref.Get(ctx, &data); err != nil {
		return fmt.Errorf("error reading values: %v", err)
	}

	for k, d := range data {
		fmt.Println(k, d)
	}

	return nil
}
