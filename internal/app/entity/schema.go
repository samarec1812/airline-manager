package entity

type Schema struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Providers []int64
}
