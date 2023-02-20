package forms

import (
	"scylla/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationForm struct {
	Email    string `form:"email" valid:"Required; Email; MaxSize(50)"`
	Password string `form:"password" valid:"Required; MinSize(6); MaxSize(50)"`
}

func (r *RegistrationForm) Valid(v *validation.Validation) {
	user := models.User{Email: r.Email}

	if user.Existed() {
		v.SetError("Email", "Account is already taken")
	}
}

func (r *RegistrationForm) Save() (*models.User, error) {

	validation := validation.Validation{}

	valid, err := validation.Valid(r)

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if !valid {
		for _, err := range validation.Errors {
			return nil, err
		}
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	user := &models.User{
		Email:    r.Email,
		Password: string(hashed),
	}

	_, err = models.AddUser(user)
	if err != nil {
		return nil, err
	}

	return user, err
}
