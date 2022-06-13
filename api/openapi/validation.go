package openapi

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CustomValidator struct{}

func (cv *CustomValidator) Validate(i interface{}) error {
	if c, ok := i.(validation.Validatable); ok {
		return c.Validate()
	}
	return nil
}

func (u UserPostRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Name,
			validation.RuneLength(1, 10).Error("名前は 1 - 10 文字です"),
		),
		validation.Field(
			&u.Password,
			validation.RuneLength(5, 10).Error("パスワードは 5 - 10 文字です"),
		),
	)
}
