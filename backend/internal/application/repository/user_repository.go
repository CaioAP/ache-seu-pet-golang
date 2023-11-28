package repository

import "github.com/caioap/ache-seu-pet/backend/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) (int, error)
	FindByEmail(email string) entity.User
}
