// @APIVersion 1.0.0
// @Title Tasks Swagger
// @Description  My Tasks Swagger
// @Name Gett
// @Schemes http
package routers

import (
	"github.com/BorisBorshvesky/meetup/controllers"
	"github.com/BorisBorshvesky/meetup/controllers/apiv1"
	"github.com/astaxie/beego"
)

func init() {
	apiv1 := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&apiv1.TaskController{},
			),
		),
	)
	beego.AddNamespace(apiv1)
	beego.Router("/", &controllers.Tasks{}, "GET:Index")
	beego.AutoRouter(&controllers.Tasks{})

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

}
