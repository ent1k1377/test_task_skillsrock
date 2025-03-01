package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, statusCode int, msg string, err error) error {
	fmt.Printf("%s: %v\n", msg, err)
	return c.Status(statusCode).JSON(fiber.Map{"error": msg})
}

func parseID(c *fiber.Ctx) (int, error) {
	id := c.Params("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("id должен быть числом")
	}

	return taskID, nil
}

func (h *Handler) CreateTask(c *fiber.Ctx) error {
	task := new(Task)

	if err := c.BodyParser(task); err != nil {
		return handleError(c, fiber.StatusBadRequest, "Не удалось распарсить данные", err)
	}

	if task.Title == "" {
		return handleError(c, fiber.StatusBadRequest, "Поле title должно быть не пустым", nil)
	}

	query := "INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id, status, created_at, updated_at"
	err := h.db.QueryRow(context.Background(), query, task.Title, task.Description).Scan(&task.ID, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Не удалось создать новую задачу", err)
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *Handler) GetTask(c *fiber.Ctx) error {
	rows, err := h.db.Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Не удалось получить задачи", err)
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return handleError(c, fiber.StatusInternalServerError, "Не удалось обработать данные", err)
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func (h *Handler) UpdateTask(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return handleError(c, fiber.StatusBadRequest, "Не удалось распарсить данные", err)
	}

	if task.Title == "" {
		return handleError(c, fiber.StatusBadRequest, "Поле title должно быть не пустым", nil)
	}

	query := "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4 RETURNING id, title, description, status, created_at, updated_at"
	err = h.db.QueryRow(context.Background(), query, task.Title, task.Description, task.Status, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Не удалось обновить данные", err)
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func (h *Handler) DeleteTask(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	query := "DELETE FROM tasks WHERE id = $1"
	tag, err := h.db.Exec(context.Background(), query, id)

	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Не удалось удалить запись", err)
	}

	if tag.RowsAffected() == 0 {
		return handleError(c, fiber.StatusNotFound, "Записи с таким id не найдено", nil)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
