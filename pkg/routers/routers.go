package routers

import (
	"github.com/gorilla/mux"

	"app/pkg/handlers"
	"app/pkg/middlewares"
)

//ampas
func New() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/books", middlewares.Authorized(handlers.GetBooks)).Methods("GET")
	api.HandleFunc("/books/{bookId}", middlewares.Authorized(handlers.GetBook)).Methods("GET")
	api.HandleFunc("/books/add", middlewares.Authorized(handlers.CreateBook)).Methods("POST")
	api.HandleFunc("/books/edit/{bookId}", middlewares.Authorized(handlers.UpdateBook)).Methods("PUT")
	api.HandleFunc("/books/delete/{bookId}", middlewares.Authorized(handlers.DeleteBook)).Methods("DELETE")

	api.HandleFunc("/users/register", handlers.RegisterUser).Methods("POST")
	api.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	return router
}
