// @APIVersion 1.0.0
// @Title Tasks Swagger
// @Description  My Tasks Swagger
// @Name Gett
// @Schemes http
package routers

import (
	"time"

	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/gettaxi/beego-example-ToDo/controllers"
	"github.com/gettaxi/beego-example-ToDo/controllers/apiv1"
)

const root  = "/"

func init() {
	apiv1 := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&apiv1.TaskController{},
			),
		),
	)
	beego.AddNamespace(apiv1)
	beego.Router(root, &controllers.Tasks{}, "GET:Index")
	beego.AutoRouter(&controllers.Tasks{})

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	initFilters()
}

func initFilters() {
	beego.InsertFilter(root, beego.BeforeRouter, func(c *context.Context) {
		c.Input.SetData("time", time.Now())
	}, false)

	beego.InsertFilter(root, beego.AfterExec, func(c *context.Context) {
		if startTime, ok := c.Input.GetData("time").(time.Time); ok {
			log.Println("execTime:", time.Now().Sub(startTime).Nanoseconds())
		}
	}, false)

}
