package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ilivestrong/audiobook_store/models"
)

type BookRepository interface {
	List(offset, limit int) ([]models.Book, error)
	GetByID(id uuid.UUID) (*models.Book, error)
}

type bookRepository struct {
	books []models.Book
}

func NewBookRepository() BookRepository {
	books := []models.Book{
		{ID: uuid.New(), Title: "Book One", Description: "First book"},
		{ID: uuid.New(), Title: "Book Two", Description: "Second book"},
		{ID: uuid.New(), Title: "Book Three", Description: "Third book"},
		{ID: uuid.New(), Title: "Book Four", Description: "Fourth book"},
	}
	fmt.Println("Books: ", books)
	return &bookRepository{
		books: books,
	}
}

func (r *bookRepository) List(offset, limit int) ([]models.Book, error) {
	if offset >= len(r.books) {
		return []models.Book{}, nil
	}

	end := offset + limit
	if end > len(r.books) {
		end = len(r.books)
	}

	return r.books[offset:end], nil
}

func (r *bookRepository) GetByID(id uuid.UUID) (*models.Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}
