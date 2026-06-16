package main

import (
	_ "GESTOCK_API_BEEGO_INVENTARIO/routers"

	beego "github.com/beego/beego/v2/server/web"
	beeLogger "github.com/beego/bee/v2/logger"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func main() {
	sqlConn,err := beego.AppConfig.String("sqlconn")
	if err != nil {
		beeLogger.Log.Fatal(err.Error())
	}
	orm.RegisterDataBase("default", "postgres", sqlConn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

