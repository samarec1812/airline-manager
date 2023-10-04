package entity

type Account struct {
	ID       int64 `db:"id"`
	SchemaID int64 `db:"schema_id"`
}
