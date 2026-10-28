package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"hospital-appointment-management/internal/model"
)

type DepartmentRepository struct {
	db *pgxpool.Pool
}

func NewDepartmentRepository(db *pgxpool.Pool) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (r *DepartmentRepository) CreateDepartment(ctx context.Context, department *model.Department) error {
	query := `INSERT INTO departments (name) VALUES ($1) RETURNING id`
	return r.db.QueryRow(ctx, query, department.Name).Scan(&department.ID)
}

func (r *DepartmentRepository) GetDepartments(ctx context.Context) ([]model.Department, error) {
	query := `SELECT id, name FROM departments ORDER BY id`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var departments []model.Department
	for rows.Next() {
		var department model.Department
		err := rows.Scan(&department.ID, &department.Name)
		if err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *DepartmentRepository) GetDepartmentByID(ctx context.Context, id int) (*model.Department, error) {
	query := `SELECT id, name FROM departments WHERE id = $1`
	var department model.Department
	err := r.db.QueryRow(ctx, query, id).Scan(&department.ID, &department.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("department not found with id %d", id)
		}
		return nil, err
	}
	return &department, nil
}

func (r *DepartmentRepository) UpdateDepartment(ctx context.Context, department *model.Department) error {
	query := `UPDATE departments SET name = $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, department.Name, department.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *DepartmentRepository) DeleteDepartmentByID(ctx context.Context, id int) error {
	query := `DELETE FROM departments WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
