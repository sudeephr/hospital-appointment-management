package handler

import (
	"context"
	"fmt"

	pb "hospital-appointment-management/api/proto"
	"hospital-appointment-management/internal/service"
)

type DepartmentHandler struct {
	pb.UnimplementedDepartmentServiceServer
	service *service.DepartmentService
}

func NewDepartmentHandler(service *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		service: service,
	}
}

func (h *DepartmentHandler) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentResponse, error) {
	department, err := h.service.CreateDepartment(ctx, req)
	if err != nil {
		return nil, err
	}

	protoDepartment := &pb.Department{
		Id:   fmt.Sprint(department.ID),
		Name: department.Name,
	}

	return &pb.CreateDepartmentResponse{
		Department: protoDepartment,
	}, nil
}

func (h *DepartmentHandler) GetDepartments(ctx context.Context, req *pb.GetDepartmentsRequest) (*pb.GetDepartmentsResponse, error) {
	departments, err := h.service.GetDepartments(ctx)
	if err != nil {
		return nil, err
	}

	var protoDepartments []*pb.Department

	for _, department := range departments {
		protoDepartment := &pb.Department{
			Id:   fmt.Sprint(department.ID),
			Name: department.Name,
		}

		protoDepartments = append(protoDepartments, protoDepartment)
	}

	return &pb.GetDepartmentsResponse{
		Departments: protoDepartments,
	}, nil
}

func (h *DepartmentHandler) GetDepartmentById(ctx context.Context, req *pb.GetDepartmentByIdRequest) (*pb.GetDepartmentByIdResponse, error) {
	department, err := h.service.GetDepartmentById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	protoDepartment := &pb.Department{
		Id:   fmt.Sprint(department.ID),
		Name: department.Name,
	}

	return &pb.GetDepartmentByIdResponse{
		Department: protoDepartment,
	}, nil
}

func (h *DepartmentHandler) UpdateDepartmentById(ctx context.Context, req *pb.UpdateDepartmentByIdRequest) (*pb.UpdateDepartmentByIdResponse, error) {
	department, err := h.service.UpdateDepartmentById(ctx, req)
	if err != nil {
		return nil, err
	}

	protoDepartment := &pb.Department{
		Id:   fmt.Sprint(department.ID),
		Name: department.Name,
	}

	return &pb.UpdateDepartmentByIdResponse{
		Department: protoDepartment,
	}, nil
}

func (h *DepartmentHandler) DeleteDepartmentById(ctx context.Context, req *pb.DeleteDepartmentByIdRequest) (*pb.DeleteDepartmentByIdResponse, error) {
	err := h.service.DeleteDepartmentById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteDepartmentByIdResponse{
		Message: "Department deleted successfully",
	}, nil
}
