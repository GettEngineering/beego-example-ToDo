package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"] = append(beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"],
		beego.ControllerComments{
			Method: "NewTask",
			Router: `/task/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"] = append(beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"],
		beego.ControllerComments{
			Method: "GetTaskById",
			Router: `/task/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"] = append(beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/task/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"] = append(beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/task/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"] = append(beego.GlobalControllerRouter["github.com/gettaxi/beego-example-ToDo/controllers/apiv1:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/task/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
