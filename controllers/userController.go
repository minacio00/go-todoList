package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/go-todoList/database"
	"github.com/minacio00/go-todoList/models"
)

func NewUser(c *fiber.Ctx) error {
	var db = database.Db
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	db.Create(&user)
	return c.Status(200).JSON(user)
}
