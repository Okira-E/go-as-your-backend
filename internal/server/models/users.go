package models

type Users struct {
	ID           int    `db:"id" json:"id" validate:"required"`
	Username     string `db:"username" json:"username" validate:"required"`
	Email        string `db:"email" json:"email" validate:"required"`
	PasswordHash string `db:"password_hash" json:"password_hash" validate:"required"`
	FirstName    string `db:"first_name" json:"first_name"`
	LastName     string `db:"last_name" json:"last_name"`
	IsActive     bool   `db:"is_active" json:"is_active"`
	CreatedAt    string `db:"created_at" json:"created_at"`
	UpdatedAt    string `db:"updated_at" json:"updated_at"`
}

type UsersDto struct {
	ID        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username" validate:"required"`
	Email     string `db:"email" json:"email" validate:"required"`
	FirstName string `db:"first_name" json:"first_name" validate:"required"`
	LastName  string `db:"last_name" json:"last_name" validate:"required"`
	IsActive  bool   `db:"is_active" json:"is_active"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
