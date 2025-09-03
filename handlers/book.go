// handler/book_handler.go
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	service "github.com/ilivestrong/audiobook_store/services"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	books, _ := h.service.List(page, size)
	c.JSON(http.StatusOK, gin.H{"data": books, "page": page, "size": size})
}

func (h *BookHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	book, _ := h.service.GetByID(id)

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UploadAudio(c *gin.Context) {
	idParam := c.Param("id")
	bookID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	_ = h.service.UploadAudio(bookID, file.Filename)

	c.JSON(http.StatusOK, gin.H{
		"message": "Audio uploaded successfully",
		"book_id": bookID,
		"file":    file.Filename,
	})
}
