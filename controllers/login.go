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
	session := c.StartSession()
	defer session.SessionRelease(c.Ctx.Request.Context(), c.Ctx.ResponseWriter)

	c.ParseForm(&form)

	user, err := form.Login()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	session.Set(c.Ctx.Request.Context(), "UserID", user.Id)
	flash.Success("Wellcome to the town")
	flash.Store(&c.Controller)
	c.Redirect("/signup", 302) // TODO: change to home page
}

func (c *LoginController) View() {
	session := c.StartSession()
	defer session.SessionRelease(c.Ctx.Request.Context(), c.Ctx.ResponseWriter)
	userID := session.Get(c.Ctx.Request.Context(), "UserID")
	if userID != nil {
		c.Redirect("/signup", 302) // TODO: change to home page
	}

	flash := beego.ReadFromRequest(&c.Controller)

	fmt.Println(flash)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	c.TplName = "login.tpl"
}
