package repository

import (
	"database/sql"
	"errors"
	"library/internal/model"
	"time"
)

type BookRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(book model.Book) error {
	const query = `
	INSERT INTO books (title, author, publisher, year, isbn, total_copies, available_copies, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := r.db.Exec(
		query,
		book.Title,
		book.Author,
		book.Publisher,
		book.Year,
		book.Isbn,
		book.TotalCopies,
		book.AvailableCopies,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) GetBookByID(id int64) (*model.Book, error) {
	const query = `
	SELECT id, title, author, publisher, year, isbn, total_copies, available_copies
	FROM books
	WHERE id = $1;
	`
	var b model.Book
	err := r.db.QueryRow(query, id).Scan(
		&b.Id,
		&b.Title,
		&b.Publisher,
		&b.Year,
		&b.Isbn,
		&b.TotalCopies,
		&b.AvailableCopies,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Data Not Found")
		}
		return nil, err
	}
	return &b, nil
}

func (r *BookRepository) ListBooks(province string) ([]model.Book, error) {
	rows, err := r.db.Query(`
	SELECT id, title, author, publisher, year, isbn, total_copies, available_copies
	FROM books
	ORDER BY id; 
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		err := rows.Scan(
			&b.Id,
			&b.Title,
			&b.Author,
			&b.Publisher,
			&b.Year,
			&b.Isbn,
			&b.TotalCopies,
			&b.AvailableCopies,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil

}

func (r *BookRepository) UpdateBook(book *model.Book, id int64) error {
	const query = `
	UPDATE books 
	SET title =$1,
		author =$2,
		publisher = $3, 
		year = $4,
		isbn = $5, 
		total_copies = $6,
		available_copies = $7, 
		updated_at = $8
		WHERE id = $9;
		`

	_, err := r.db.Exec(
		query,
		book.Id,
		book.Author,
		book.Publisher,
		book.Year,
		book.Isbn,
		book.TotalCopies,
		book.AvailableCopies,
		time.Now(),
		id,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("Data Not Found")
		}
		return err
	}
	return nil

}

func (r *BookRepository) DeleteBook(id int64) error {
	const query = `
	DELETE FROM books
	WHERE id = $1;
	
	`
	_, err := r.db.Exec(query, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("Data Not Found")
		}
		return err

	}
	return nil
}
