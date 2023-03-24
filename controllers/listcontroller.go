package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/go-todoList/database"
	"github.com/minacio00/go-todoList/models"
)

func NewList(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var user models.User
	var list models.List
	// list belongs to a user, so is expected to have an user id to create a new list
	body := struct {
		UserId int    `json:"userId"`
		Tittle string `json:"title"`
	}{}
	c.BodyParser(&body)
	if body.UserId == 0 || body.Tittle == "" {
		return c.SendStatus(400)
	}

	list.Title = body.Tittle
	user.ID = uint(body.UserId)

	database.Db.Find(&user)
	if user.ID == 0 {
		println(body.UserId)
		return c.Status(404).JSON(struct {
			Message string
		}{
			Message: "User not found",
		})

	}
	list.UserID = user.ID
	database.Db.Create(&list)
	return c.Status(200).JSON(struct{ Success string }{Success: "List created"})
}
