package models

import (
	"log"
	"time"

	"app/pkg/configs"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	ID        int    `form:"id" json:"id"`
	Name      string `form:"name" json:"name"`
	Email     string `form:"email" json:"email"`
	Password  string `form:"password" json:"password"`
	Status    int    `form:"status" json:"status"`
	CreatedAt string `form:"created_at" json:"created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
}

func GetUser(email *string) (*Users, error) {
	var user Users

	db := configs.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email, password, status, created_at, updated_at FROM users WHERE status = 1 AND email = ?", *email)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Fatal(err.Error())
		}
	}

	return &user, err
}

func CheckUser(email *string) (*int, error) {
	var status int

	db := configs.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT EXISTS(SELECT * FROM users WHERE status = 1 AND email = ?) as status", *email)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&status); err != nil {
			log.Fatal(err.Error())
		}
	}

	return &status, err
}

func CreateUser(name *string, email *string, password *string) error {
	db := configs.Connect()
	defer db.Close()

	status := 1
	createdAt := time.Now()
	updatedAt := time.Now()

	_, err := db.Exec("INSERT INTO users(name, email, password, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", *name, *email, *password, status, createdAt, updatedAt)

	if err != nil {
		return err
	}

	return nil
}
