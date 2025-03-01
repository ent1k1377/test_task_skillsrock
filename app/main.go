package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Handler struct {
	db *pgx.Conn
}

func main() {
	time.Sleep(10 * time.Second) // Ждем когда бд прогрузиться
	loadEnv()

	handler := Handler{
		db: ConnectDB(),
	}
	defer handler.db.Close(context.Background())

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/tasks", handler.CreateTask)
	app.Get("/tasks", handler.GetTask)
	app.Put("/tasks/:id", handler.UpdateTask)
	app.Delete("/tasks/:id", handler.DeleteTask)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось считать .env файл")
	}
}
