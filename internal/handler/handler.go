package handler

import (
	"library/internal/model"
	"library/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookService *service.BookService
}

type UserHandler struct {
	UserService *service.UserService
}

func NewBookHandler(BookService *service.BookService) *BookHandler {
	return &BookHandler{
		BookService: BookService,
	}
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: UserService,
	}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	err := h.BookService.CreateBook(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, nil)
}

func (h *BookHandler) ListBooks(c echo.Context) error {
	datas, err := h.BookService.ListBooks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": datas})
}

func (h *BookHandler) GetBookByID(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	datas, err := h.BookService.GetBookByID(int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": datas})
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	var book model.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	err = h.BookService.UpdateBook(&book, int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = h.BookService.DeleteBook(int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user model.Lib_user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	err := h.UserService.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, nil)
}
