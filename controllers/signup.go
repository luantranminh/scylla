package controllers

import (
	"fmt"
	"html/template"
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
	c.Redirect("/home", 302)
}

func (c *SignupController) View() {
	session := c.StartSession()
	defer session.SessionRelease(c.Ctx.Request.Context(), c.Ctx.ResponseWriter)
	userID := session.Get(c.Ctx.Request.Context(), "UserID")
	if userID != nil {
		c.Redirect("/home", 302)
	}

	flash := beego.ReadFromRequest(&c.Controller)

	fmt.Println(flash)

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	c.TplName = "signup.tpl"
	c.Data["Title"] = "Create new account"
}
