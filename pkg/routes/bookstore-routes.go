package routes

import (
	"go-bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("Put")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
