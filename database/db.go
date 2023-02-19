package database

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/postgres"
	_ "github.com/lib/pq"
)

func init() {
	db, _ := beego.AppConfig.String("db_database")
	user, _ := beego.AppConfig.String("db_username")
	pass, _ := beego.AppConfig.String("db_password")
	host, _ := beego.AppConfig.String("db_host")
	port, _ := beego.AppConfig.String("db_port")

	dns := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, db)

	fmt.Println(dns)

	err := orm.RegisterDataBase("default", "postgres", dns)
	if err != nil {
		logs.Critical("cannot open connect to db:", err)
	}

	mode, _ := beego.AppConfig.String("runmode")
	if mode == "dev" {
		orm.Debug = true
	}
}
