package controllers

import (
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

type BaseUIController struct {
	beego.Controller
}

//line by line
func (c *BaseUIController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ct, controllerName, actionName, app)
	c.Data["app_name"] = "Todo List - Beego"
	c.Layout = "layout.html"
	c.TplExt = "html"
}
