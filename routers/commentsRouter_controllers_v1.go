package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"] = append(beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"],
		beego.ControllerComments{
			Method: "NewTask",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"] = append(beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"],
		beego.ControllerComments{
			Method: "GetTaskById",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"] = append(beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"] = append(beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"] = append(beego.GlobalControllerRouter["github.com/BorisBorshvesky/meetup/controllers/v1:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
