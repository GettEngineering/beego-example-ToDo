package main

import (
	"github.com/BorisBorshvesky/meetup/models"
	_ "github.com/BorisBorshvesky/meetup/routers"
	"github.com/astaxie/beego"
)

func main() {
	//seed()
	beego.Run()
}

func seed() {
	models.AddTask(&models.Task{
		Title: "Read Harry Potter",
	})

	models.AddTask(&models.Task{
		Title: "Write Blog post about beego",
	})

	models.AddTask(&models.Task{
		Title:       "Arrenge things for Amsterdam",
		Description: "Get List from Barak",
	})
}
