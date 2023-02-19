package controllers

import (
	"fmt"
	"scylla/forms"

	beego "github.com/beego/beego/v2/server/web"
)

// SignupController operations for Signup
type SignupController struct {
	beego.Controller
}

func (c *SignupController) Signup() {
	form := forms.RegistrationForm{}
	flash := beego.NewFlash()

	c.ParseForm(&form)

	_, err := form.Save()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/signup", 302)
		return
	}

	flash.Success("Wellcome to the town")
	flash.Store(&c.Controller)
	c.Redirect("/signup", 302)
}

func (c *SignupController) View() {
	flash := beego.ReadFromRequest(&c.Controller)

	fmt.Println(flash)

	c.TplName = "signup.tpl"
	c.Data["Title"] = "Create new account"
}
