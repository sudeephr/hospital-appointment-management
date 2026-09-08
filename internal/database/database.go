package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	connString := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Error creating database pool:", err)
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		fmt.Println("Error connecting to the database:", err)
		pool.Close()
		return nil, err
	}
	return pool, nil
}
