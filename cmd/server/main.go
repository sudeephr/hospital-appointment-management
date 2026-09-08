package main

import (
	// "context"
	"fmt"
	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/database"
	"hospital-appointment-management/internal/handler"
	"hospital-appointment-management/internal/repository"
	"hospital-appointment-management/internal/service"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hospital Appointment Management System")
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	// Coonect to the database and it will return a connection pool that can be used to interact with the database.
	pool, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	// Initialize the patient repository
	patientRepo := repository.NewPatientRepository(pool)

	// Initialize the patient service
	patientService := service.NewPatientService(patientRepo)

	// Initialize the patient handler
	patientHandler := handler.NewPatientHandler(patientService)

	// Create a new gRPC server and register the patient handler
	grpcServer := grpc.NewServer()
	pb.RegisterPatientServiceServer(grpcServer, patientHandler)

	// opening a network listener on TCP port 8080. This listener will accept incoming network connections on that port.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Failed to serve gRPC server", err)
	}

}
