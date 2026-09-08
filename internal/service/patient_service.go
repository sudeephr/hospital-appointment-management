package service

import (
	"context"
	"errors"
	"fmt"
	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/model"
	"hospital-appointment-management/internal/repository"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

func (s *PatientService) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*model.Patient, error) {
	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, err
	}

	patient := &model.Patient{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: dateOfBirth,
		Gender:      req.Gender,
		Phone:       req.Phone,
		Email:       req.Email,
		Address:     req.Address,
	}

	err = s.repo.CreatePatient(ctx, patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (s *PatientService) GetPatientData() ([]model.Patient, error) {
	return s.repo.PatientData()
}

func (s *PatientService) GetPatientById(ctx context.Context, id string) (*model.Patient, error) {
	patientID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	patient, err := s.repo.GetPatientByID(ctx, patientID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("patient not found with id %d", patientID)
		}

		return nil, err
	}
	return patient, nil
}
func (s *PatientService) UpdatePatientById(ctx context.Context, req *pb.UpdatePatientByIdRequest) (*model.Patient, error) {
	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, err
	}
	patient := &model.Patient{
		ID:          int(req.Id),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: dateOfBirth,
		Gender:      req.Gender,
		Phone:       req.Phone,
		Email:       req.Email,
		Address:     req.Address,
	}

	err = s.repo.UpdatePatient(ctx, patient)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("patient not found with id %d", patient.ID)
		}
		return nil, err
	}

	return patient, nil
}

func (s *PatientService) DeletePatientById(ctx context.Context, id string) error {
	patientID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = s.repo.DeletePatientByID(ctx, patientID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("patient not found with id %d", patientID)
		}
		return err
	}
	return nil
}
