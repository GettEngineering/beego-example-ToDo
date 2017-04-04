package apiv1

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/gettaxi/beego-example-ToDo/models"
)

//  TaskController operations for Task
type TaskController struct {
	beego.Controller  //embedding beego.controller!!!
}

// Create new task
// @Title Create new
// @Description create new task
// @Param	body		body 	models.Task	true	"body for Task content"
// @Success 200 {object} models.Task
// @Failure 400 bad request
// @router /task/ [post]
func (c *TaskController) NewTask() {
	var v models.Task
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if id, err := models.AddTask(&v); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"id": id}

	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}
	c.ServeJSON()
}

// GetOne
// @Title Get task by id
// @Description get Task by id
// @Param	id		path 	string	true	"Id of the task"
// @Success 200 {object} models.Task
// @Failure 400 :id is empty
// @router /task/:id [get]
func (c *TaskController) GetTaskById() {
	// c. functions
	// c.Ctx.
	id, err := c.GetInt64(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	task, err := models.GetTaskById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = task
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get all Tasks
// @Param	limit	query	number	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	number	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Task
// @Failure 404
// @router /task/ [get]
func (c *TaskController) GetAll() {
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")
	tasks, err := models.GetAllTasks(limit, offset)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"tasks": tasks}

	}
	c.ServeJSON()
}

// Put update the task
// @Title update the task
// @Description update the Task
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Task	true	"body for Task content"
// @Success 200 {object} models.Task
// @Failure 403 :id is not int
// @router /task/:id [put]
func (c *TaskController) Put() {
	id, err := c.GetInt64("id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	v := models.Task{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateTask(id, &v); err == nil {
		c.Data["json"] = map[string]interface{}{"status": "OK"}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}
	c.ServeJSON()
}

// Delete delete the Task
// @Title delete the Task
// @Description delete the Task
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 400 id is empty
// @router /task/:id [delete]
func (c *TaskController) Delete() {
	id, err := c.GetInt64("id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	if err := models.DeleteTask(id); err == nil {
		c.Data["json"] = map[string]interface{}{"status": "OK"}
	} else {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}
	c.ServeJSON()
}
