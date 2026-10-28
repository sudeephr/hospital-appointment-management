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

type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repo: repo,
	}
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*model.Department, error) {
	department := &model.Department{
		Name: req.Name,
	}
	err := s.repo.CreateDepartment(ctx, department)
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (s *DepartmentService) GetDepartments(ctx context.Context) ([]model.Department, error) {
	return s.repo.GetDepartments(ctx)
}

func (s *DepartmentService) GetDepartmentById(ctx context.Context, id string) (*model.Department, error) {
	departmentID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	department, err := s.repo.GetDepartmentByID(ctx, departmentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("department not found with id %d", departmentID)
		}

		return nil, err
	}

	return department, nil
}

func (s *DepartmentService) UpdateDepartmentById(ctx context.Context, req *pb.UpdateDepartmentByIdRequest) (*model.Department, error) {
	department := &model.Department{
		ID:   int(req.Id),
		Name: req.Name,
	}

	err := s.repo.UpdateDepartment(ctx, department)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("department not found with id %d", department.ID)
		}

		return nil, err
	}

	return department, nil
}

func (s *DepartmentService) DeleteDepartmentById(ctx context.Context, id string) error {
	departmentID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.repo.DeleteDepartmentByID(ctx, departmentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("department not found with id %d", departmentID)
		}

		return err
	}

	return nil
}
