package controllers

import (
	"html/template"
	"net/http"
	"scylla/forms"

	beego "github.com/beego/beego/v2/server/web"
)

// HomeController operations for Home
type HomeController struct {
	beego.Controller
}

func (c *HomeController) View() {
	// flash := beego.ReadFromRequest(&c.Controller)

	// fmt.Println(flash)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	c.TplName = "home.tpl"
}

func (c *HomeController) UploadFile() {
	file, _, err := c.GetFile("csvfile")

	if err != nil {
		c.Ctx.WriteString("Error getting file")
		return
	}
	defer file.Close()

	csv := forms.Parse(file)
	csv.Save()
	// Do something with the CSV data (e.g. display it in a template)
	c.Data["records"] = csv.Keywords

	c.Ctx.Redirect(http.StatusFound, "/home")
}
