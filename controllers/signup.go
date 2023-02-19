package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// SignupController operations for Signup
type SignupController struct {
	beego.Controller
}

func (c *SignupController) Post() {

}

func (c *SignupController) View() {
	c.TplName = "signup.tpl"
	c.Data["Title"] = "Create new account"
}
