package main

import (
	"context"
	"fmt"
	"hospital-appointment-management/internal/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hospital Appointment Management System")
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	fmt.Println("Database connected successfully!")
}
