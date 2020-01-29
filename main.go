package main

import (
	_ "golemonilo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"

	lemonilo "github.com/astaxie/beego"
)

func main() {
	if lemonilo.BConfig.RunMode == "dev" {
		lemonilo.BConfig.WebConfig.DirectoryIndex = true
		lemonilo.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.Debug = true

	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	lemonilo.Run()
}
