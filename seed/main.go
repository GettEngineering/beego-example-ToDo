package main

import "github.com/gettaxi/meetup/models"

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

