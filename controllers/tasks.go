package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gettaxi/meetup/models"
)

type Tasks struct {
	BaseUIController
}

// Index - Main page
func (c *Tasks) Index() {
	tasks, err := models.GetAllTasks(0, 0)
	if err != nil {
		c.Data["error"] = err.Error()
	}
	c.Data["tasks"] = tasks

	c.Render()
}

func (c *Tasks) Edit() {
	ID, err := c.GetInt64("id")
	if err != nil {
		c.Redirect(beego.URLFor("Tasks.Index"), http.StatusFound)
		return
	}
	task, _ := models.GetTaskById(ID)
	c.Data["task"] = task
	c.Render()
}

func (c *Tasks) New() {
	c.Render()
}

func (c *Tasks) Update() {
	task := models.Task{}
	c.ParseForm(&task)
	models.UpdateTask(task.Id, &task)
	c.Redirect(beego.URLFor("Tasks.Index"), http.StatusFound)
}

func (c *Tasks) Create() {
	task := models.Task{}
	c.ParseForm(&task)

	models.AddTask(&task)
	c.Redirect(beego.URLFor("Tasks.Index"), http.StatusFound)
}

func (c *Tasks) Delete() {
	task := models.Task{}
	c.ParseForm(&task)

	models.DeleteTask(task.Id)
	c.Redirect(beego.URLFor("Tasks.Index"), http.StatusFound)
}
