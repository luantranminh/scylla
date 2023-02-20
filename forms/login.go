package forms

import (
	"errors"
	"scylla/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUserCredential = errors.New("email or password is incorrect")
)

type LoginForm struct {
	Email    string `form:"email" valid:"Required; Email; MaxSize(50)"`
	Password string `form:"password" valid:"Required; MinSize(6); MaxSize(50)"`
}

func (l *LoginForm) Login() (*models.User, error) {
	// query for duplicated email
	user, err := models.GetUserByEmail(l.Email)
	if err != nil {
		return nil, ErrInvalidUserCredential
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))
	if err != nil {
		return nil, ErrInvalidUserCredential
	}

	return user, nil
}
