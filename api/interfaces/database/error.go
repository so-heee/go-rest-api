package database

import (
	"errors"

	"github.com/so-heee/go-rest-api/api/usecase/repository"
	"gorm.io/gorm"
)

func convertError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.ErrNotFound
	}
	return err
}
