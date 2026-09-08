package model

import "time"

type Patient struct {
	ID          int       `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	Gender      string    `db:"gender"`
	Phone       string    `db:"phone"`
	Email       string    `db:"email"`
	Address     string    `db:"address"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
