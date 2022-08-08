package main

import (
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "paninti-region-service/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func main() {
	// Db Connection
	sqlConn, _ := beego.AppConfig.String("sqlconn")
	sqlMaxIdle, _ := beego.AppConfig.Int("sqlmaxidle")
	sqlMaxConn, _ := beego.AppConfig.Int("sqlmaxconn")
	err := orm.RegisterDataBase("default", "postgres", sqlConn)
	if err != nil {
		panic(err)
	}
	orm.SetMaxIdleConns("default", sqlMaxIdle)
	orm.SetMaxOpenConns("default", sqlMaxConn)

	// Swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Cors
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}
