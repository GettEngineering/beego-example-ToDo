package controllers

import (
	"net/http"

	"github.com/BorisBorshvesky/meetup/models"
	"github.com/astaxie/beego"
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
	c.Ctx.Request.ParseForm()
	id, _ := c.GetInt64("id")
	done, _ := c.GetBool("done")

	task := models.Task{
		Id:          id,
		Description: c.GetString("description"),
		Title:       c.GetString("title"),
		Done:        done,
	}

	models.UpdateTask(task.Id, &task)
	c.Redirect(beego.URLFor("Tasks.Index"), http.StatusFound)
}

func (c *Tasks) Create() {
	c.Ctx.Request.ParseForm()

	task := models.Task{
		Title:       c.GetString("title"),
		Description: c.GetString("description"),
	}
	models.AddTask(&task)
	c.Redirect(beego.URLFor("Tasks.Index"), http.StatusFound)
}

func (c *Tasks) Delete() {
	c.Ctx.Request.ParseForm()
	id, _ := c.GetInt64("id")
	models.DeleteTask(id)
	c.Redirect(beego.URLFor("Tasks.Index"), http.StatusFound)
}
