package main

import (
	"github.com/gofiber/fiber/v2"
)

type hello struct {
	Message string
}
type Task struct {
	Title       string
	Description string
	IsCompleted bool
}
type Todos struct {
	Title string
	Tasks *[]Task
}

func main() {
	var tasks []Task
	todos := Todos{Title: "my list", Tasks: &tasks}
	app := fiber.New()
	Connectdb()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.JSON(hello{Message: "hello world"})
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		// tasks := []Task{
		// 	{Title: "my Day", Description: "", IsCompleted: false},
		// }
		tasks = append(tasks, Task{Title: "your task", Description: "this your task not mine", IsCompleted: true})
		// todos := Todos{Title: "my list", Tasks: &tasks}
		return c.JSON(todos)
	})
	app.Post("/newTodo", func(c *fiber.Ctx) error {
		tasks = append(tasks, Task{Title: "New todo", Description: "this your task not mine", IsCompleted: true})
		return c.JSON(todos)
	})

	app.Listen(":8080")

}
