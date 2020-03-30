package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/models"
	"books-list/utils"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"

)

var books []models.Book
var db *sql.DB

func init() {
	err := gotenv.Load()
	utils.LogFatal(err)
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods(http.MethodGet)
	router.HandleFunc("/books", controller.AddBook(db)).Methods(http.MethodPost)
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods(http.MethodDelete)

	//CORS
	router.Use(mux.CORSMethodMiddleware(router))


	fmt.Println("Server is starting on port", os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))
}

