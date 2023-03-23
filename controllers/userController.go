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
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Db.Find(&users)
	return c.Status(200).JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var user *models.User

	database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(404).JSON("id not found")
	}

	// if err := findUser(id, &user); err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	return c.Status(200).JSON(&user)
}

func UpdateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer value")
	}
	user := models.User{ID: 0}
	database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(404).JSON("User not found")
	}
	c.BodyParser(&user)
	database.Db.Save(&user)
	return c.Status(200).JSON(&user)
}
func DeleteUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer value")
	}
	user := models.User{}
	database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(404).JSON("User not found")
	}
	database.Db.Delete(&user)
	return c.Status(200).JSON(user.ID)
}
