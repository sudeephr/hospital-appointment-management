package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"hospital-appointment-management/internal/model"
)

type AppointmentRepository struct {
	db *pgxpool.Pool
}

func NewAppointmentRepository(db *pgxpool.Pool) *AppointmentRepository {
	return &AppointmentRepository{
		db: db,
	}
}

func (r *AppointmentRepository) CreateAppointment(ctx context.Context, appointment *model.Appointment) error {
	query := `INSERT INTO appointments (patient_id, doctor_id, appointment_date, appointment_time, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	return r.db.QueryRow(ctx, query, appointment.PatientID, appointment.DoctorID, appointment.AppointmentDate, appointment.AppointmentTime, appointment.Status).Scan(&appointment.ID)
}

func (r *AppointmentRepository) GetAppointments(ctx context.Context) ([]model.Appointment, error) {
	query := `SELECT id, patient_id, doctor_id, appointment_date, appointment_time, status FROM appointments ORDER BY id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []model.Appointment

	for rows.Next() {
		var appointment model.Appointment

		err := rows.Scan(&appointment.ID, &appointment.PatientID, &appointment.DoctorID, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.Status)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (r *AppointmentRepository) GetAppointmentByID(ctx context.Context, id int) (*model.Appointment, error) {
	query := `SELECT id, patient_id, doctor_id, appointment_date, appointment_time, status FROM appointments WHERE id = $1`

	var appointment model.Appointment

	err := r.db.QueryRow(ctx, query, id).Scan(&appointment.ID, &appointment.PatientID, &appointment.DoctorID, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.Status)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("appointment not found with id %d", id)
		}

		return nil, err
	}

	return &appointment, nil
}

func (r *AppointmentRepository) UpdateAppointment(ctx context.Context, appointment *model.Appointment) error {
	query := `UPDATE appointments SET patient_id = $1, doctor_id = $2, appointment_date = $3, appointment_time = $4, status = $5 WHERE id = $6`

	result, err := r.db.Exec(ctx, query, appointment.PatientID, appointment.DoctorID, appointment.AppointmentDate, appointment.AppointmentTime, appointment.Status, appointment.ID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *AppointmentRepository) DeleteAppointmentByID(ctx context.Context, id int) error {
	query := `DELETE FROM appointments WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
