package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Task     string `json:"task"`
	Finished bool   `json:"finished"`
	Body     string `json:"body"`
}

func main() {

	app := fiber.New()

	todos := []Todo{}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/healthchecl", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/api/todoslist", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Post("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Finished = true
				break
			}
		}

		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
