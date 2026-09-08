package repository

import (
	"context"
	"hospital-appointment-management/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PatientRepository struct {
	db *pgxpool.Pool
}

func NewPatientRepository(db *pgxpool.Pool) *PatientRepository {
	return &PatientRepository{
		db: db,
	}
}

func (r *PatientRepository) CreatePatient(ctx context.Context, patient *model.Patient) error {
	query := `INSERT INTO patients(first_name, last_name, date_of_birth, gender, phone, email, address)VALUES ($1, $2, $3, $4, $5, $6, $7)RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, patient.FirstName, patient.LastName, patient.DateOfBirth,
		patient.Gender, patient.Phone, patient.Email, patient.Address).Scan(&patient.ID, &patient.CreatedAt, &patient.UpdatedAt)
	return err
}

func (r *PatientRepository) PatientData() ([]model.Patient, error) {
	query := `SELECT * FROM patients`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var patients []model.Patient
	for rows.Next() {
		var patient model.Patient
		err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.DateOfBirth, &patient.Gender, &patient.Phone, &patient.Email, &patient.Address, &patient.CreatedAt, &patient.UpdatedAt)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *PatientRepository) GetPatientByID(ctx context.Context, id int) (*model.Patient, error) {
	query := `
		SELECT id, first_name, last_name, date_of_birth, gender, phone, email, address, created_at, updated_at FROM patients WHERE id = $1`
	var patient model.Patient
	err := r.db.QueryRow(ctx, query, id).Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.DateOfBirth, &patient.Gender, &patient.Phone, &patient.Email, &patient.Address, &patient.CreatedAt, &patient.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) UpdatePatient(ctx context.Context, patient *model.Patient) error {
	query := `UPDATE patients SET first_name = $1, last_name = $2, date_of_birth = $3, gender = $4, phone = $5, email = $6, address = $7, updated_at = NOW() WHERE id = $8`
	result, err := r.db.Exec(ctx, query, patient.FirstName, patient.LastName, patient.DateOfBirth, patient.Gender, patient.Phone, patient.Email, patient.Address, patient.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *PatientRepository) DeletePatientByID(ctx context.Context, id int) error {
	query := `DELETE FROM patients WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
