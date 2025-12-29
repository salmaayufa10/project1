package service

import (
	"errors"
	"fmt"
	"library/internal/model"
	"library/internal/repository"
)

type BookService struct {
	BookRepository *repository.BookRepository
}

func NewBookService(BookRepository *repository.BookRepository) *BookService {
	return &BookService{
		BookRepository: BookRepository,
	}
}

func (s *BookService) CreateBook(book *model.Book) error {
	if book.Title == "" || book.Author == "" || book.Publisher == "" {
		return errors.New("Title, Author, Publisher are required")
	}

	if book.TotalCopies < 1 || book.AvailableCopies < 1 {
		return errors.New("Total Copies and Available Copies must greater than 0")
	}

	if book.TotalCopies > 100 || book.AvailableCopies > 100 {
		return errors.New("Total Copies and Available Copies must less than 100")
	}

	err := s.BookRepository.CreateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetBookByID(id int64) (*model.Book, error) {
	if id < 1 {
		return nil, errors.New("Invalid ID. ID cannot less than 1")
	}

	data, err := s.BookRepository.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *BookService) ListBooks() ([]model.Book, error) {
	fmt.Println("masuk")
	datas, err := s.BookRepository.ListBooks()
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (s *BookService) UpdateBook(book *model.Book, id int64) error {
	if id < 1 {
		return errors.New("Invalid ID. ID cannot less than 1")
	}
	if book.Title == "" || book.Author == "" || book.Publisher == "" {
		return errors.New("Title, Author, Publisher are required")
	}

	if book.TotalCopies < 1 || book.AvailableCopies < 1 {
		return errors.New("Total Copies and Available Copies must greater than 0")
	}

	if book.TotalCopies > 100 || book.AvailableCopies > 100 {
		return errors.New("Total Copies and Available Copies must less than 100")
	}

	err := s.BookRepository.UpdateBook(book, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) DeleteBook(id int64) error {
	if id < 1 {
		return errors.New("invalid id")
	}

	err := s.BookRepository.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}
