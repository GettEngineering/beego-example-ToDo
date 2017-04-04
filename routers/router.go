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


func init() {
	apiv1 := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&apiv1.TaskController{}, //annotation routing
			),
		),
	)
	beego.AddNamespace(apiv1)
	beego.Router("/", &controllers.Tasks{}, "GET:Index") //regular routing root
	beego.AutoRouter(&controllers.Tasks{})                //auto routing

	beego.BConfig.WebConfig.DirectoryIndex = true //static files
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	initFilters()
}

func initFilters() {
	beego.InsertFilter("/*", beego.BeforeRouter, func(c *context.Context) {
		c.Input.SetData("time", time.Now())
	}, false)

	beego.InsertFilter("/*", beego.AfterExec, func(c *context.Context) {
		if startTime, ok := c.Input.GetData("time").(time.Time); ok {
			log.Println(c.Request.RequestURI, c.Request.Method, " execTime:", time.Now().Sub(startTime).Nanoseconds())
		}
	}, false)
}
