package main

import (
	"context"
	"fmt"
	"log"

	pb "hospital-appointment-management/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPatientServiceClient(conn)
	// createResponse, err := client.CreatePatient(context.Background(), &pb.CreatePatientRequest{
	// 	FirstName:   "John",
	// 	LastName:    "Doe",
	// 	DateOfBirth: "1995-05-10",
	// 	Gender:      "Male",
	// 	Phone:       "9876543210",
	// 	Email:       "john@gmail.com",
	// 	Address:     "Bengaluru",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Patient created successfully with ID:", createResponse.Patient.Id)
	// fmt.Println("Name: ", createResponse.Patient.FirstName, createResponse.Patient.LastName)
	fmt.Println("-----------------------------")

	// gRPC client sends a request to the server to retrieve all patients and prints their details.
	response, err := client.GetPatients(context.Background(), &pb.GetPatientsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for _, patient := range response.Patients {
		fmt.Println("ID:", patient.Id)
		fmt.Println("Name:", patient.FirstName, patient.LastName)
		fmt.Println("Date of Birth:", patient.DateOfBirth)
		fmt.Println("Gender:", patient.Gender)
		fmt.Println("Phone:", patient.Phone)
		fmt.Println("Email:", patient.Email)
		fmt.Println("Address:", patient.Address)
		fmt.Println("-----------------------------")
	}
	getResponse, err := client.GetPatientById(context.Background(), &pb.GetPatientByIdRequest{Id: "6"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved patient with ID:", getResponse.Patient.Id)
	fmt.Println("Name: ", getResponse.Patient.FirstName, getResponse.Patient.LastName)

	updateResponse, err := client.UpdatePatientById(context.Background(), &pb.UpdatePatientByIdRequest{
		Id:          6,
		FirstName:   "Updated",
		LastName:    "Patient",
		DateOfBirth: "1995-05-10",
		Gender:      "Male",
		Phone:       "9876543210",
		Email:       "updated@gmail.com",
		Address:     "Bengaluru",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Patient updated successfully")
	fmt.Println("ID:", updateResponse.Patient.Id)
	fmt.Println("Name:", updateResponse.Patient.FirstName, updateResponse.Patient.LastName)
	fmt.Println("Date of Birth:", updateResponse.Patient.DateOfBirth)
	fmt.Println("Gender:", updateResponse.Patient.Gender)
	fmt.Println("Phone:", updateResponse.Patient.Phone)
	fmt.Println("Email:", updateResponse.Patient.Email)
	fmt.Println("Address:", updateResponse.Patient.Address)

	// responseDelete, err := client.DeletePatientById(context.Background(), &pb.DeletePatientByIdRequest{Id: "4"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(responseDelete.Message)
}
