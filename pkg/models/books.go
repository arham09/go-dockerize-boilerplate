package models

import (
	"app/pkg/configs"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Books struct {
	ID          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Author      string `form:"author" json:"author"`
	Description string `form:"description" json:"description"`
	Status      int    `form:"status" json:"status"`
	CreatedAt   string `form:"created_at" json:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at"`
}

func GetBooks() (*[]Books, error) {
	var book Books
	var books []Books

	db := configs.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, author, description, status, created_at, updated_at FROM books WHERE status = 1")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.Status, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Fatal(err.Error())

		} else {
			books = append(books, book)
		}
	}

	return &books, err
}

func GetBook(bookId *int) (*Books, error) {
	var book Books

	db := configs.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, author, description, status, created_at, updated_at FROM books WHERE status = 1 AND id = ?", *bookId)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.Status, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Fatal(err.Error())
		}
	}

	return &book, err
}

func InsertBook(name *string, author *string, description *string) (*Books, error) {
	var book Books

	db := configs.Connect()
	defer db.Close()

	status := 1
	createdAt := time.Now()
	updatedAt := time.Now()

	res, err := db.Exec("INSERT INTO books(name, author, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", *name, *author, *description, status, createdAt, updatedAt)

	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()

	book.ID = id
	book.Name = *name
	book.Author = *author
	book.Description = *description
	book.Status = status
	book.CreatedAt = createdAt.String()
	book.UpdatedAt = updatedAt.String()

	return &book, err
}

func EditBook(bookId *int, name *string, author *string, description *string) error {
	db := configs.Connect()
	defer db.Close()

	updatedAt := time.Now()

	_, err := db.Exec("UPDATE books SET name = ?, author = ?, description = ?, updated_at = ? WHERE id = ?", *name, *author, *description, updatedAt, *bookId)

	if err != nil {
		return err
	}

	return nil
}

func RemoveBook(bookId *int) error {
	db := configs.Connect()
	defer db.Close()

	status := 0
	updatedAt := time.Now()

	_, err := db.Exec("UPDATE books SET status = ?, updated_at = ? WHERE id = ?", status, updatedAt, *bookId)

	if err != nil {
		return err
	}

	return nil
}
