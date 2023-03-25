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
		return c.Status(400).JSON(struct{ Message string }{Message: "body expects both userId, and tittle fields"})
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
func FindList(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := c.ParamsInt("id", -1)
	if err != nil {
		return c.SendStatus(400)
	}

	list := models.List{}

	database.Db.Find(&list, "id = ? ", id)
	if list.ID == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(list)

}
func GetLists(c *fiber.Ctx) error {
	c.Accepts("application/json")

	lists := []models.List{}
	database.Db.Find(&lists)
	return c.Status(200).JSON(&lists)
}
func UpdateList(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	body := struct {
		Title string `json:"title"`
	}{}
	c.BodyParser(&body)
	if body.Title == "" {
		return c.SendStatus(400)
	}

	list := models.List{}

	database.Db.Find(&list, "id = ?", id)
	if list.ID == 0 {
		return c.SendStatus(404)
	}

	list.Title = body.Title
	database.Db.Save(&list)
	return c.Status(200).JSON(list)
}

func DeleteList(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	list := models.List{}

	database.Db.Find(&list, "id = ?", id)
	if list.ID == 0 {
		return c.SendStatus(404)
	}

	database.Db.Delete(&list)
	return c.SendStatus(204)
}
