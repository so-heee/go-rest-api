package controllers

import (
	"errors"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
	"golang.org/x/crypto/bcrypt"
)

type Error2 struct {
	err  error
	code string
}

var (
	ErrNilID         = errors.New("nil id")
	ErrInvalidID     = errors.New("invalid uuid")
	ErrNotFound      = errors.New("not found")
	ErrForbidden     = errors.New("forbidden")
	ErrAlreadyExists = errors.New("already exists")
	ErrInvalidArg    = errors.New("argument error")
	ErrBind          = errors.New("bind error")
	ErrValidate      = errors.New("validate error")
)

func NewError(status string, err error, sub ...oapi.SubError) *oapi.ApiError {
	return &oapi.ApiError{
		oapi.Error{
			Code:    status,
			Message: err.Error(),
			Errors:  &sub,
		},
	}
}

func convertError(err error, sub ...oapi.SubError) error {
	switch {
	case errors.Is(err, ErrNilID):
		fallthrough
	case errors.Is(err, ErrInvalidID):
		fallthrough
	case errors.Is(err, ErrInvalidArg):
		fallthrough
	case errors.Is(err, ErrBind):
		fallthrough
	case errors.Is(err, ErrValidate):
		// return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("bad request: %w", err).Error())
		return echo.NewHTTPError(http.StatusBadRequest, NewError(http.StatusText(http.StatusBadRequest), err, sub...))

	case errors.Is(err, ErrAlreadyExists):
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("conflicts: %w", err).Error())

	case errors.Is(err, ErrForbidden):
		return echo.NewHTTPError(http.StatusForbidden, fmt.Errorf("forbideen: %w", err).Error())

	case errors.Is(err, ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, NewError(http.StatusText(http.StatusNotFound), err))

	case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
		return echo.NewHTTPError(http.StatusNotFound, NewError(http.StatusText(http.StatusForbidden), err))

	default:
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("%w", err).Error())
	}
}

func convertValidationError(errs validation.Errors) []oapi.SubError {
	e := []oapi.SubError{}
	for k, err := range errs {
		s := oapi.SubError{
			Parameter: k,
			Message:   err.Error(),
		}
		e = append(e, s)
	}
	return e
}
