package main

import (
	_ "scylla/database"
	_ "scylla/models"
	_ "scylla/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical("cannot sync models:", err)
	}

	beego.Run()
}
