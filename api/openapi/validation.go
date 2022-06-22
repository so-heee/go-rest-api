package openapi

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CustomValidator struct{}

const (
	ErrMsgRequired   = "%sは必須入力です"
	ErrMsgRuneLengrh = "%sは %d - %d 文字です"
	ErrMsgEmail      = "メールアドレスを入力して下さい"
)

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
			validation.Required.Error("名前は必須入力です"),
			validation.RuneLength(1, 20).Error("名前は 1 - 20 文字です"),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("メールアドレスは必須入力です"),
			validation.RuneLength(8, 20).Error("パスワードは 8 - 20 文字です"),
			is.Email.Error(ErrMsgEmail),
		),
	)
}

func (u UserPatchRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Name,
			validation.Required.Error("名前は必須入力です"),
			validation.RuneLength(1, 20).Error("名前は 1 - 20 文字です"),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("メールアドレスは必須入力です"),
			validation.RuneLength(8, 20).Error("パスワードは 8 - 20 文字です"),
			is.Email.Error(ErrMsgEmail),
		),
	)
}
