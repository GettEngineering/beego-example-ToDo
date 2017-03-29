package controllers

import (
	"html/template"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

type BaseUIController struct {
	beego.Controller
}

func (c *BaseUIController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ct, controllerName, actionName, app)
	c.EnableXSRF = true
	c.XSRFExpire = 7200
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["app_name"] = "Todo List - Beego"
	c.Layout = "layout.html"
	c.TplExt = "html"
}
