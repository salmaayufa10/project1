package handler

import (
	"library/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *BookHandler) CreateUser(c echo.Context) error {
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

func (h *BookHandler) Login(c echo.Context) error {
	var user model.Lib_user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	data, err := h.UserService.GetUserByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, data)
}
