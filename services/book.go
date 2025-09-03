// service/book_service.go
package service

import (
	"github.com/google/uuid"
	"github.com/ilivestrong/audiobook_store/models"
	"github.com/ilivestrong/audiobook_store/repositories"
)

type BookService interface {
	List(page, size int) ([]models.Book, error)
	GetByID(id uuid.UUID) (*models.Book, error)
	UploadAudio(bookID uuid.UUID, filename string) error
}

type bookServiceImpl struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookServiceImpl{repo: repo}
}

func (s *bookServiceImpl) List(page, size int) ([]models.Book, error) {
	offset := (page - 1) * size
	return s.repo.List(offset, size)
}

func (s *bookServiceImpl) GetByID(id uuid.UUID) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookServiceImpl) UploadAudio(bookID uuid.UUID, filename string) error {
	// TODO: implement S3 upload
	return nil
}
