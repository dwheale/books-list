package controllers

import (
	"books-list/models"
	bookRepository "books-list/repository/book"
	"books-list/utils"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controller struct {}

var books []models.Book



func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		books, err := bookRepo.GetBooks(db, book, books)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Error - Not Found"
				utils.SendError(w, http.StatusNotFound, error)
			} else {
				error.Message = "Server Error - Could Not Get Books"
				utils.SendError(w, http.StatusInternalServerError, error)
			}
			return
		}

		if r.Method == http.MethodOptions {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Error - Not Found"
				utils.SendError(w, http.StatusNotFound, error)
			} else {
				error.Message = "Server Error - Could Not Get Book"
				utils.SendError(w, http.StatusInternalServerError, error)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error
		err := json.NewDecoder(r.Body).Decode(&book)
		utils.LogFatal(err)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter Missing Fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err = bookRepo.AddBook(db, book)
		if err != nil {
			error.Message = "Server Error - could not add book"
			utils.SendError(w, http.StatusInternalServerError, error)
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		err := json.NewDecoder(r.Body).Decode(&book)
		utils.InternalError(w, err)

		if  book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}

		rowsUpdated, err := bookRepo.UpdateBook(db, book)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Error - Not Found"
				utils.SendError(w, http.StatusNotFound, error)
			} else {
				error.Message = "Server Error - Could Not Update Book"
				utils.SendError(w, http.StatusInternalServerError, error)
			}
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var error models.Error

		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := bookRepo.RemoveBook(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Error - Not Found"
				utils.SendError(w, http.StatusNotFound, error)
			} else {
				error.Message = "Server Error - Could Not Remove Book"
				utils.SendError(w, http.StatusInternalServerError, error)
			}
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}

