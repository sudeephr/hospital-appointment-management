package model

type Doctor struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	Specialization string `db:"specialization"`
	AvailableDays  string `db:"available_days"`
	DepartmentID   int    `db:"department_id"`
}
