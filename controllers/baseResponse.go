package controllers

import (
	"github.com/labstack/echo/v4"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewResponseSuccess[T any](c echo.Context, statusCode int, message string, data T) error {
	return c.JSON(statusCode, Response[T]{
		Message: message,
		Data:    data,
	})
}
