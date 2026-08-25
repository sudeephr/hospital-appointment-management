package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	connString := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	}
	return conn, err
}
