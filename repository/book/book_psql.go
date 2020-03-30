package bookRepository

import (
	"books-list/models"
	"database/sql"
)

type BookRepository struct {}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("select * from books")
	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}
	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
	rows := db.QueryRow("select * from books where id=$1;", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	return book, err
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {
	var bookID int
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) returning id;", book.Title, book.Author, book.Year).Scan(&bookID)
	return bookID, err
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int, error) {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 returning id", &book.Title, &book.Author, &book.Year, &book.ID)
	rowsUpdated, err := result.RowsAffected()
	return int(rowsUpdated), err
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (int, error) {
	result, err := db.Exec("delete from books where id = $1", id)
	rowsDeleted, err := result.RowsAffected()
	return int(rowsDeleted), err
}