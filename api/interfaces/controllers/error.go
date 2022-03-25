package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
	"github.com/so-heee/go-rest-api/api/usecase/repository"
	"golang.org/x/crypto/bcrypt"
)

type Error struct {
	Message string
}

func NewError(status string, err error) *oapi.ApiError {
	return &oapi.ApiError{
		oapi.Error{
			Code:    status,
			Message: err.Error(),
		},
	}
}

func convertError(err error) error {
	switch {
	case errors.Is(err, repository.ErrNilID):
		fallthrough
	case errors.Is(err, repository.ErrInvalidID):
		fallthrough
	case errors.Is(err, repository.ErrInvalidArg):
		fallthrough
	case errors.Is(err, repository.ErrBind):
		fallthrough
	case errors.Is(err, repository.ErrValidate):
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("bad request: %w", err).Error())

	case errors.Is(err, repository.ErrAlreadyExists):
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("conflicts: %w", err).Error())

	case errors.Is(err, repository.ErrForbidden):
		return echo.NewHTTPError(http.StatusForbidden, fmt.Errorf("forbideen: %w", err).Error())

	case errors.Is(err, repository.ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, NewError(http.StatusText(http.StatusNotFound), err))

	case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
		return echo.NewHTTPError(http.StatusNotFound, NewError(http.StatusText(http.StatusForbidden), err))

	default:
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("%w", err).Error())
	}
}
