package handler

import (
	"context"
	"fmt"

	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/service"
)

type DoctorHandler struct {
	pb.UnimplementedDoctorServiceServer
	service *service.DoctorService
}

func NewDoctorHandler(service *service.DoctorService) *DoctorHandler {
	return &DoctorHandler{
		service: service,
	}
}

func (h *DoctorHandler) CreateDoctor(ctx context.Context, req *pb.CreateDoctorRequest) (*pb.CreateDoctorResponse, error) {
	doctor, err := h.service.CreateDoctor(ctx, req)
	if err != nil {
		return nil, err
	}

	protoDoctor := &pb.Doctor{
		Id:             fmt.Sprint(doctor.ID),
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		AvailableDays:  doctor.AvailableDays,
		DepartmentId:   fmt.Sprint(doctor.DepartmentID),
	}

	return &pb.CreateDoctorResponse{
		Doctor: protoDoctor,
	}, nil
}

func (h *DoctorHandler) GetDoctors(ctx context.Context, req *pb.GetDoctorsRequest) (*pb.GetDoctorsResponse, error) {
	doctors, err := h.service.GetDoctors(ctx)
	if err != nil {
		return nil, err
	}

	var protoDoctors []*pb.Doctor

	for _, doctor := range doctors {
		protoDoctor := &pb.Doctor{
			Id:             fmt.Sprint(doctor.ID),
			Name:           doctor.Name,
			Specialization: doctor.Specialization,
			AvailableDays:  doctor.AvailableDays,
			DepartmentId:   fmt.Sprint(doctor.DepartmentID),
		}

		protoDoctors = append(protoDoctors, protoDoctor)
	}

	return &pb.GetDoctorsResponse{
		Doctors: protoDoctors,
	}, nil
}

func (h *DoctorHandler) GetDoctorById(ctx context.Context, req *pb.GetDoctorByIdRequest) (*pb.GetDoctorByIdResponse, error) {
	doctor, err := h.service.GetDoctorById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	protoDoctor := &pb.Doctor{
		Id:             fmt.Sprint(doctor.ID),
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		AvailableDays:  doctor.AvailableDays,
		DepartmentId:   fmt.Sprint(doctor.DepartmentID),
	}

	return &pb.GetDoctorByIdResponse{
		Doctor: protoDoctor,
	}, nil
}

func (h *DoctorHandler) UpdateDoctorById(ctx context.Context, req *pb.UpdateDoctorByIdRequest) (*pb.UpdateDoctorByIdResponse, error) {
	doctor, err := h.service.UpdateDoctorById(ctx, req)
	if err != nil {
		return nil, err
	}

	protoDoctor := &pb.Doctor{
		Id:             fmt.Sprint(doctor.ID),
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		AvailableDays:  doctor.AvailableDays,
		DepartmentId:   fmt.Sprint(doctor.DepartmentID),
	}

	return &pb.UpdateDoctorByIdResponse{
		Doctor: protoDoctor,
	}, nil
}

func (h *DoctorHandler) DeleteDoctorById(ctx context.Context, req *pb.DeleteDoctorByIdRequest) (*pb.DeleteDoctorByIdResponse, error) {
	err := h.service.DeleteDoctorById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteDoctorByIdResponse{
		Message: "Doctor deleted successfully",
	}, nil
}
