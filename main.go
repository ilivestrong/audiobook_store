package main

import (
	handler "github.com/ilivestrong/audiobook_store/handlers"
	middleware "github.com/ilivestrong/audiobook_store/middlewares"
	repository "github.com/ilivestrong/audiobook_store/repositories"
	router "github.com/ilivestrong/audiobook_store/router"
	service "github.com/ilivestrong/audiobook_store/services"
)

func main() {

	// Repositories
	bookRepo := repository.NewBookRepository()

	// Services
	bookService := service.NewBookService(bookRepo)

	// Handlers
	bookHandler := handler.NewBookHandler(bookService)

	// Routers
	r := router.SetupRouter(bookHandler, middleware.AdminOnly())

	// Run server
	r.Run(":8080")
}
