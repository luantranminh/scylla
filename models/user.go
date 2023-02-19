package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Base

	Email    string `orm:"unique"`
	Password string
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{
		Base: Base{Id: id},
	}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (u *User) Existed() bool {
	if user, _ := GetUserByEmail(u.Email); user == nil {
		return false
	}

	return true
}

func GetUserByEmail(email string) (v *User, err error) {
	o := orm.NewOrm()
	var user User

	if err = o.QueryTable(user.TableName()).Filter("email", email).One(&user); err == nil {
		return &user, nil
	}

	return nil, err
}
