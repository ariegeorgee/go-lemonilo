package routers

import (
	"golemonilo/controllers"

	lemonilo "github.com/astaxie/beego"
)

func init() {
	ns := lemonilo.NewNamespace("/v1",
		lemonilo.NSNamespace("/user",
			lemonilo.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	lemonilo.AddNamespace(ns)
}
