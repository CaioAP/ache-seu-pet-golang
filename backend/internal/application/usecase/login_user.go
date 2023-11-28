package usecase

import (
	"log"

	"github.com/caioap/ache-seu-pet/backend/internal/application/repository"
	"github.com/caioap/ache-seu-pet/backend/internal/domain/entity"
)

type LoginUserInput struct {
	Email string
	Password string
}

type LoginUserUsecase struct {
	UserRepository repository.UserRepository
}

func NewLoginUserUsecase(userRepository repository.UserRepository) *LoginUserUsecase {
	return &LoginUserUsecase{UserRepository: userRepository}
}

func (u *LoginUserUsecase) Execute(input LoginUserInput) string {
	user := u.UserRepository.FindByEmail(input.Email)
	if &user == nil {
		log.Fatalln("User does not exist!")
	}
	isPasswordMatch := entity.Compare(user.Password, input.Password); if !isPasswordMatch {
		log.Fatalln("Password is incorrect!")
	}
	token, err := entity.Generate(user.ID); if err != nil {
		log.Fatalln("Error creating accessToken")
	}
	return token
}