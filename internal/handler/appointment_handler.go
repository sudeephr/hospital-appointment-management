package handler

import (
	"context"
	"fmt"

	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/service"
)

type AppointmentHandler struct {
	pb.UnimplementedAppointmentServiceServer
	service *service.AppointmentService
}

func NewAppointmentHandler(service *service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		service: service,
	}
}

func (h *AppointmentHandler) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {
	appointment, err := h.service.CreateAppointment(ctx, req)
	if err != nil {
		return nil, err
	}

	protoAppointment := &pb.Appointment{
		Id:              fmt.Sprint(appointment.ID),
		PatientId:       fmt.Sprint(appointment.PatientID),
		DoctorId:        fmt.Sprint(appointment.DoctorID),
		AppointmentDate: appointment.AppointmentDate.Format("2006-01-02"),
		AppointmentTime: appointment.AppointmentTime.Format("15:04"),
		Status:          appointment.Status,
	}

	return &pb.CreateAppointmentResponse{
		Appointment: protoAppointment,
	}, nil
}

func (h *AppointmentHandler) GetAppointments(ctx context.Context, req *pb.GetAppointmentsRequest) (*pb.GetAppointmentsResponse, error) {
	appointments, err := h.service.GetAppointments(ctx)
	if err != nil {
		return nil, err
	}

	var protoAppointments []*pb.Appointment

	for _, appointment := range appointments {
		protoAppointment := &pb.Appointment{
			Id:              fmt.Sprint(appointment.ID),
			PatientId:       fmt.Sprint(appointment.PatientID),
			DoctorId:        fmt.Sprint(appointment.DoctorID),
			AppointmentDate: appointment.AppointmentDate.Format("2006-01-02"),
			AppointmentTime: appointment.AppointmentTime.Format("15:04"),
			Status:          appointment.Status,
		}

		protoAppointments = append(protoAppointments, protoAppointment)
	}

	return &pb.GetAppointmentsResponse{
		Appointments: protoAppointments,
	}, nil
}

func (h *AppointmentHandler) GetAppointmentById(ctx context.Context, req *pb.GetAppointmentByIdRequest) (*pb.GetAppointmentByIdResponse, error) {
	appointment, err := h.service.GetAppointmentById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	protoAppointment := &pb.Appointment{
		Id:              fmt.Sprint(appointment.ID),
		PatientId:       fmt.Sprint(appointment.PatientID),
		DoctorId:        fmt.Sprint(appointment.DoctorID),
		AppointmentDate: appointment.AppointmentDate.Format("2006-01-02"),
		AppointmentTime: appointment.AppointmentTime.Format("15:04"),
		Status:          appointment.Status,
	}

	return &pb.GetAppointmentByIdResponse{
		Appointment: protoAppointment,
	}, nil
}

func (h *AppointmentHandler) UpdateAppointmentById(ctx context.Context, req *pb.UpdateAppointmentByIdRequest) (*pb.UpdateAppointmentByIdResponse, error) {
	appointment, err := h.service.UpdateAppointmentById(ctx, req)
	if err != nil {
		return nil, err
	}

	protoAppointment := &pb.Appointment{
		Id:              fmt.Sprint(appointment.ID),
		PatientId:       fmt.Sprint(appointment.PatientID),
		DoctorId:        fmt.Sprint(appointment.DoctorID),
		AppointmentDate: appointment.AppointmentDate.Format("2006-01-02"),
		AppointmentTime: appointment.AppointmentTime.Format("15:04"),
		Status:          appointment.Status,
	}

	return &pb.UpdateAppointmentByIdResponse{
		Appointment: protoAppointment,
	}, nil
}

func (h *AppointmentHandler) DeleteAppointmentById(ctx context.Context, req *pb.DeleteAppointmentByIdRequest) (*pb.DeleteAppointmentByIdResponse, error) {
	err := h.service.DeleteAppointmentById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAppointmentByIdResponse{
		Message: "Appointment deleted successfully",
	}, nil
}
