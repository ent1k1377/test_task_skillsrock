package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@postgres:5432/test_task_skillsrock_db?sslmode=disable") // не разбириха с портами и @postgres
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не получилось подключиться к бд, %v", err)
		os.Exit(1)
	}

	return conn
}
