package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	PublishedAt   time.Time `json:"published_at"`
	ISBN          string    `json:"isbn"`
	CoverImageURL string    `json:"cover_image_url"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     uuid.UUID `json:"created_by"`

	Authors []Author   `json:"authors,omitempty"`
	Audio   *BookAudio `json:"audio,omitempty"`
}

type BookAudio struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"book_id"`
	AudioURL  string    `json:"audio_url"`
	Duration  int       `json:"duration"`
	Format    string    `json:"format"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookTranscription struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"book_id"`
	ESDocID   string    `json:"es_doc_id"`
	Status    string    `json:"status"` // pending, completed, failed
	CreatedAt time.Time `json:"created_at"`
}
