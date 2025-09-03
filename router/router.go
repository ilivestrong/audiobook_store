package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ilivestrong/audiobook_store/handlers"
	middleware "github.com/ilivestrong/audiobook_store/middlewares"
)

func SetupRouter(bookHandler *handler.BookHandler, mw ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	// Public endpoints
	api.GET("/books", bookHandler.List)
	api.GET("/books/:id", bookHandler.GetByID)

	// Protected endpoint (admin only)
	api.POST("/books/:id/upload", middleware.AdminOnly(), bookHandler.UploadAudio)

	return r
}
