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
	// Initialize the  repository
	patientRepo := repository.NewPatientRepository(pool)
	departmentRepo := repository.NewDepartmentRepository(pool)
	doctorRepo := repository.NewDoctorRepository(pool)
	appointmentRepo := repository.NewAppointmentRepository(pool)

	// Initialize the  service
	patientService := service.NewPatientService(patientRepo)
	departmentService := service.NewDepartmentService(departmentRepo)
	doctorService := service.NewDoctorService(doctorRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo)

	// Initialize the  handler
	patientHandler := handler.NewPatientHandler(patientService)
	departmentHandler := handler.NewDepartmentHandler(departmentService)
	doctorHandler := handler.NewDoctorHandler(doctorService)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)

	// Create a new gRPC server and register the  handler
	grpcServer := grpc.NewServer()
	pb.RegisterPatientServiceServer(grpcServer, patientHandler)
	pb.RegisterDepartmentServiceServer(grpcServer, departmentHandler)
	pb.RegisterDoctorServiceServer(grpcServer, doctorHandler)
	pb.RegisterAppointmentServiceServer(grpcServer, appointmentHandler)

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
