package main

import (
	"log"

	"github.com/emrekas/go-react-todo/backend/database"
	"github.com/emrekas/go-react-todo/backend/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")
	todo.Register(api, database.DB)

	log.Fatal(app.Listen(":5000"))
}
