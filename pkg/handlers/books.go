package handlers

import (
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"app/pkg/models"
	"app/pkg/responses"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseBooks

	books, err := models.GetBooks()
	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Data = *books

	responses.Response(w, http.StatusOK, response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseBook

	vars := mux.Vars(r)
	bookId, err := strconv.Atoi(vars["bookId"])
	if err != nil {
		log.Print(err)
	}

	book, err := models.GetBook(&bookId)
	if err != nil {
		log.Print(err)
		return
	} else if book.ID == 0 {
		responses.Error(w, 400, "Book Not Found")
		return
	}

	response.Status = http.StatusOK
	response.Data = *book

	responses.Response(w, http.StatusOK, response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseBook

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return
	}

	name := r.Form.Get("name")
	author := r.Form.Get("author")
	description := r.Form.Get("description")

	book, err := models.InsertBook(&name, &author, &description)

	if err != nil {
		log.Print("ada error")
		log.Print(err)
		return
	}

	response.Status = http.StatusOK
	response.Data = *book

	responses.Response(w, http.StatusOK, response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	vars := mux.Vars(r)
	bookId, err := strconv.Atoi(vars["bookId"])
	if err != nil {
		log.Print(err)
	}

	err = r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	name := r.Form.Get("name")
	author := r.Form.Get("author")
	description := r.Form.Get("description")

	err = models.EditBook(&bookId, &name, &author, &description)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Successfully Updated"

	responses.Response(w, http.StatusOK, response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	vars := mux.Vars(r)
	bookId, err := strconv.Atoi(vars["bookId"])
	if err != nil {
		log.Print(err)
	}

	err = models.RemoveBook(&bookId)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Successfully Deleted"

	responses.Response(w, http.StatusOK, response)
}
