package value

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p *Password) ConvertHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (p *Password) Verify(r string) error {
	err := bcrypt.CompareHashAndPassword([]byte(string(*p)), []byte(r))
	if err != nil {
		return err
	}
	return nil
}
