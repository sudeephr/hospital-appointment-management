package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"hospital-appointment-management/internal/model"
)

type DoctorRepository struct {
	db *pgxpool.Pool
}

func NewDoctorRepository(db *pgxpool.Pool) *DoctorRepository {
	return &DoctorRepository{
		db: db,
	}
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doctor *model.Doctor) error {
	query := `INSERT INTO doctors ( name, specialization, available_days, department_id) VALUES ($1, $2, $3, $4) RETURNING id `
	return r.db.QueryRow(ctx, query, doctor.Name, doctor.Specialization, doctor.AvailableDays, doctor.DepartmentID).Scan(&doctor.ID)
}

func (r *DoctorRepository) GetDoctors(ctx context.Context) ([]model.Doctor, error) {
	query := `SELECT id, name, specialization, available_days, department_id FROM doctors ORDER BY id `
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var doctors []model.Doctor
	for rows.Next() {
		var doctor model.Doctor
		err := rows.Scan(
			&doctor.ID,
			&doctor.Name,
			&doctor.Specialization,
			&doctor.AvailableDays,
			&doctor.DepartmentID,
		)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return doctors, nil
}

func (r *DoctorRepository) GetDoctorByID(ctx context.Context, id int) (*model.Doctor, error) {
	query := `SELECT id,name, specialization, available_days, department_id FROM doctors WHERE id = $1`
	var doctor model.Doctor
	err := r.db.QueryRow(ctx, query, id).Scan(&doctor.ID, &doctor.Name, &doctor.Specialization, &doctor.AvailableDays, &doctor.DepartmentID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("doctor not found with id %d", id)
		}
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepository) UpdateDoctor(ctx context.Context, doctor *model.Doctor) error {
	query := `UPDATE doctors SET name = $1,specialization = $2,available_days = $3,department_id = $4WHERE id = $5`

	result, err := r.db.Exec(ctx, query, doctor.Name, doctor.Specialization, doctor.AvailableDays, doctor.DepartmentID, doctor.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *DoctorRepository) DeleteDoctorByID(ctx context.Context, id int) error {
	query := `DELETE FROM doctorsWHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
