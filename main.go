package main

import (
	"API_MONGO/app"
	"API_MONGO/configs"
	"API_MONGO/repository"
	"API_MONGO/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()

	configs.ConnectDB()

	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAll)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":8080")
}
