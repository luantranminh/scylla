package routers

import (
	"github.com/beego/beego/v2/server/web/context"
)

func Auth(ctx *context.Context) {
	userID := ctx.Input.Session("UserID")
	if userID == nil {
		ctx.Redirect(302, "/login")
	}
}
