package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/go-todoList/database"
	"github.com/minacio00/go-todoList/models"
)

func NewTask(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var task models.Task
	body := struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ListID      uint   `json:"listId"`
	}{}
	c.BodyParser(&body)

	if body.Title == "" || body.ListID == 0 {
		return c.Status(400).JSON(struct{ Message string }{Message: "body expects both title and listId fields"})
	}

	task.Title = body.Title
	task.Description = body.Description
	task.ListID = body.ListID

	database.Db.Create(&task)
	return c.Status(200).JSON(struct{ Success string }{Success: "Task created"})
}
func FindTask(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := c.ParamsInt("id", -1)
	if err != nil {
		return c.SendStatus(400)
	}

	task := models.Task{}

	database.Db.Find(&task, "id = ?", id)
	if task.ID == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(task)
}
func GetTasks(c *fiber.Ctx) error {
	c.Accepts("application/json")

	listID, err := c.ParamsInt("listId", -1)
	if err != nil {
		return c.SendStatus(400)
	}

	tasks := []models.Task{}
	database.Db.Find(&tasks, "list_id = ?", listID)
	return c.Status(200).JSON(&tasks)
}
func UpdateTask(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	body := struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{}
	c.BodyParser(&body)
	if body.Title == "" {
		return c.SendStatus(400)
	}

	task := models.Task{}

	database.Db.Find(&task, "id = ?", id)
	if task.ID == 0 {
		return c.SendStatus(404)
	}

	task.Title = body.Title
	task.Description = body.Description
	database.Db.Save(&task)
	return c.Status(200).JSON(task)
}
func DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	task := models.Task{}

	database.Db.Find(&task, "id = ?", id)
	if task.ID == 0 {
		return c.SendStatus(404)
	}

	database.Db.Delete(&task)
	return c.SendStatus(204)
}
