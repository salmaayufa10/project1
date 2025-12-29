package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *BookHandler) {
	e.POST("/books", h.CreateBook)
	e.GET("/books", h.ListBooks)
	e.GET("/books/:id", h.GetBookByID)
	e.PUT("/books/:id", h.UpdateBook)
	e.DELETE("/books/:id", h.DeleteBook)

}
