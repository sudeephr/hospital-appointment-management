package model

import "time"

type Appointment struct {
	ID              int       `db:"id"`
	PatientID       int       `db:"patient_id"`
	DoctorID        int       `db:"doctor_id"`
	AppointmentDate time.Time `db:"appointment_date"`
	AppointmentTime time.Time `db:"appointment_time"`
	Status          string    `db:"status"`
}
