package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/go-todoList/controllers"
	"github.com/minacio00/go-todoList/database"
)

func main() {
	app := fiber.New()
	database.Connectdb()
	app.Post("/newUser", controllers.NewUser)
	app.Listen(":8080")

}
