package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/model"
	"hospital-appointment-management/internal/repository"

	"github.com/jackc/pgx/v5"
)

type DoctorService struct {
	repo *repository.DoctorRepository
}

func NewDoctorService(repo *repository.DoctorRepository) *DoctorService {
	return &DoctorService{
		repo: repo,
	}
}

func (s *DoctorService) CreateDoctor(ctx context.Context, req *pb.CreateDoctorRequest) (*model.Doctor, error) {
	doctor := &model.Doctor{
		Name:           req.Name,
		Specialization: req.Specialization,
		AvailableDays:  req.AvailableDays,
		DepartmentID:   int(req.DepartmentId),
	}

	err := s.repo.CreateDoctor(ctx, doctor)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

func (s *DoctorService) GetDoctors(ctx context.Context) ([]model.Doctor, error) {
	return s.repo.GetDoctors(ctx)
}

func (s *DoctorService) GetDoctorById(ctx context.Context, id string) (*model.Doctor, error) {
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	doctor, err := s.repo.GetDoctorByID(ctx, doctorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("doctor not found with id %d", doctorID)
		}

		return nil, err
	}

	return doctor, nil
}

func (s *DoctorService) UpdateDoctorById(ctx context.Context, req *pb.UpdateDoctorByIdRequest) (*model.Doctor, error) {
	doctor := &model.Doctor{
		ID:             int(req.Id),
		Name:           req.Name,
		Specialization: req.Specialization,
		AvailableDays:  req.AvailableDays,
		DepartmentID:   int(req.DepartmentId),
	}

	err := s.repo.UpdateDoctor(ctx, doctor)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("doctor not found with id %d", doctor.ID)
		}

		return nil, err
	}

	return doctor, nil
}

func (s *DoctorService) DeleteDoctorById(ctx context.Context, id string) error {
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.repo.DeleteDoctorByID(ctx, doctorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("doctor not found with id %d", doctorID)
		}

		return err
	}

	return nil
}
