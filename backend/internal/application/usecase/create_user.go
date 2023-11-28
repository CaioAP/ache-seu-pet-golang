package usecase

import (
	"time"

	"github.com/caioap/ache-seu-pet/backend/internal/application/repository"
	"github.com/caioap/ache-seu-pet/backend/internal/domain/entity"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Phone    string
	Birthday string
	Password string
}

type CreateUserUsecase struct {
	UserRepository repository.UserRepository
}

func NewCreateUserUsecase(userRepository repository.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{UserRepository: userRepository}
}

func (u *CreateUserUsecase) Execute(input CreateUserInput) (*entity.User, error) {
	user := entity.NewUser(nil, input.Name, input.Email, input.Phone, input.Birthday, time.Now())
	password, err := entity.Hash(input.Password); if err != nil {
		return nil, err
	}
	user.Password = password
	userId, err := u.UserRepository.Create(user); if err != nil {
		return nil, err
	}
	user.ID = userId
	return user, nil
}
