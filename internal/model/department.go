package model

type Department struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
