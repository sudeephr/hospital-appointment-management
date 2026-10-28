package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/model"
	"hospital-appointment-management/internal/repository"

	"github.com/jackc/pgx/v5"
)

type AppointmentService struct {
	repo *repository.AppointmentRepository
}

func NewAppointmentService(repo *repository.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		repo: repo,
	}
}

func (s *AppointmentService) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*model.Appointment, error) {
	appointmentDate, err := time.Parse("2006-01-02", req.AppointmentDate)
	if err != nil {
		return nil, err
	}

	appointmentTime, err := time.Parse("15:04", req.AppointmentTime)
	if err != nil {
		return nil, err
	}

	appointment := &model.Appointment{
		PatientID:       int(req.PatientId),
		DoctorID:        int(req.DoctorId),
		AppointmentDate: appointmentDate,
		AppointmentTime: appointmentTime,
		Status:          req.Status,
	}

	err = s.repo.CreateAppointment(ctx, appointment)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (s *AppointmentService) GetAppointments(ctx context.Context) ([]model.Appointment, error) {
	return s.repo.GetAppointments(ctx)
}

func (s *AppointmentService) GetAppointmentById(ctx context.Context, id string) (*model.Appointment, error) {
	appointmentID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	appointment, err := s.repo.GetAppointmentByID(ctx, appointmentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("appointment not found with id %d", appointmentID)
		}

		return nil, err
	}

	return appointment, nil
}

func (s *AppointmentService) UpdateAppointmentById(ctx context.Context, req *pb.UpdateAppointmentByIdRequest) (*model.Appointment, error) {
	appointmentDate, err := time.Parse("2006-01-02", req.AppointmentDate)
	if err != nil {
		return nil, err
	}

	appointmentTime, err := time.Parse("15:04", req.AppointmentTime)
	if err != nil {
		return nil, err
	}

	appointment := &model.Appointment{
		ID:              int(req.Id),
		PatientID:       int(req.PatientId),
		DoctorID:        int(req.DoctorId),
		AppointmentDate: appointmentDate,
		AppointmentTime: appointmentTime,
		Status:          req.Status,
	}

	err = s.repo.UpdateAppointment(ctx, appointment)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("appointment not found with id %d", appointment.ID)
		}

		return nil, err
	}

	return appointment, nil
}

func (s *AppointmentService) DeleteAppointmentById(ctx context.Context, id string) error {
	appointmentID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.repo.DeleteAppointmentByID(ctx, appointmentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("appointment not found with id %d", appointmentID)
		}

		return err
	}

	return nil
}
