package handler

import (
	"context"
	"fmt"

	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/service"
)

type PatientHandler struct {
	pb.UnimplementedPatientServiceServer
	service *service.PatientService
}

func NewPatientHandler(service *service.PatientService) *PatientHandler {
	return &PatientHandler{
		service: service,
	}
}

func (h *PatientHandler) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.CreatePatientResponse, error) {
	patient, err := h.service.CreatePatient(ctx, req)
	if err != nil {
		return nil, err
	}
	protoPatient := &pb.Patient{
		Id:          fmt.Sprint(patient.ID),
		FirstName:   patient.FirstName,
		LastName:    patient.LastName,
		DateOfBirth: patient.DateOfBirth.Format("2006-01-02"),
		Gender:      patient.Gender,
		Phone:       patient.Phone,
		Email:       patient.Email,
		Address:     patient.Address,
		CreatedAt:   patient.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   patient.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &pb.CreatePatientResponse{Patient: protoPatient}, nil
}

func (h *PatientHandler) GetPatients(ctx context.Context, req *pb.GetPatientsRequest) (*pb.GetPatientsResponse, error) {
	patients, err := h.service.GetPatientData()
	if err != nil {
		return nil, err
	}
	var protoPatients []*pb.Patient
	for _, patient := range patients {
		protoPatient := &pb.Patient{
			Id:          fmt.Sprint(patient.ID),
			FirstName:   patient.FirstName,
			LastName:    patient.LastName,
			DateOfBirth: patient.DateOfBirth.Format("2006-01-02"),
			Gender:      patient.Gender,
			Phone:       patient.Phone,
			Email:       patient.Email,
			Address:     patient.Address,
			CreatedAt:   patient.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   patient.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		protoPatients = append(protoPatients, protoPatient)
	}
	return &pb.GetPatientsResponse{Patients: protoPatients}, nil
}

func (h *PatientHandler) GetPatientById(ctx context.Context, req *pb.GetPatientByIdRequest) (*pb.GetPatientByIdResponse, error) {
	patient, err := h.service.GetPatientById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	protoPatient := &pb.Patient{
		Id:          fmt.Sprint(patient.ID),
		FirstName:   patient.FirstName,
		LastName:    patient.LastName,
		DateOfBirth: patient.DateOfBirth.Format("2006-01-02"),
		Gender:      patient.Gender,
		Phone:       patient.Phone,
		Email:       patient.Email,
		Address:     patient.Address,
		CreatedAt:   patient.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   patient.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &pb.GetPatientByIdResponse{Patient: protoPatient}, nil
}

func (h *PatientHandler) UpdatePatientById(ctx context.Context, req *pb.UpdatePatientByIdRequest) (*pb.UpdatePatientByIdResponse, error) {
	patient, err := h.service.UpdatePatientById(ctx, req)
	if err != nil {
		return nil, err
	}
	protoPatient := &pb.Patient{
		Id:          fmt.Sprint(patient.ID),
		FirstName:   patient.FirstName,
		LastName:    patient.LastName,
		DateOfBirth: patient.DateOfBirth.Format("2006-01-02"),
		Gender:      patient.Gender,
		Phone:       patient.Phone,
		Email:       patient.Email,
		Address:     patient.Address,
	}
	return &pb.UpdatePatientByIdResponse{Patient: protoPatient}, nil
}

func (h *PatientHandler) DeletePatientById(ctx context.Context, req *pb.DeletePatientByIdRequest) (*pb.DeletePatientByIdResponse, error) {
	err := h.service.DeletePatientById(ctx, req.Id)
	if err != nil {
		return nil, err
	} else {
		return &pb.DeletePatientByIdResponse{Message: "Patient deleted successfully"}, nil
	}
}
