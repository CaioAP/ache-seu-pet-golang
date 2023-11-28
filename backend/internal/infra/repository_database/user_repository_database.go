package repository_database

import (
	"database/sql"

	"github.com/caioap/ache-seu-pet/backend/internal/domain/entity"
)

type UserRepositoryDatabase struct {
	DB *sql.DB
}

func NewUserRepositoryDatabase(db *sql.DB) *UserRepositoryDatabase {
	return &UserRepositoryDatabase{DB: db}
}

func (r *UserRepositoryDatabase) Create(user *entity.User) error {
	var userId int
	err := r.DB.QueryRow(
		"INSERT INTO users (name, email, phone, birthday, created_at, password) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id",
		user.Name, user.Email, user.Phone, user.Birthday, user.CreatedAt, user.Password,
	).Scan(&userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryDatabase) FindByEmail(email string) *entity.User {
	var user entity.User
	row := r.DB.QueryRow("SELECT * FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Birthday, &user.CreatedAt, &user.Password)
	if err != nil {
		return nil
	}
	return &user
}