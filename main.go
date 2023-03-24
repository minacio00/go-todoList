package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/minacio00/go-todoList/controllers"
	"github.com/minacio00/go-todoList/database"
)

func StartServer() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	database.Connectdb()
	app.Post("/users", controllers.NewUser)
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUser)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)

	app.Post("/lists", controllers.NewList)
	app.Listen(":8080")
	return app
}
func main() {
	StartServer()
}
