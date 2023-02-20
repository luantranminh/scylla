package controllers

import (
	"fmt"
	"html/template"
	"scylla/forms"

	beego "github.com/beego/beego/v2/server/web"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	form := forms.LoginForm{}
	flash := beego.NewFlash()

	c.ParseForm(&form)

	_, err := form.Login()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	flash.Success("Wellcome to the town")
	flash.Store(&c.Controller)
	c.Redirect("/login", 302)
}

func (c *LoginController) View() {
	flash := beego.ReadFromRequest(&c.Controller)

	fmt.Println(flash)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	c.TplName = "login.tpl"
}
