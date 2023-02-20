package routers

import (
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		if strings.HasPrefix(ctx.Input.URL(), "/login") {
			return
		}

		_, ok := ctx.Input.Session("UserID").(int)
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/home", beego.BeforeRouter, FilterUser)
}
