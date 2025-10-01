package usecase

import (
	"errors"
	"financial-track/model"
	"financial-track/repository"
	"financial-track/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) RegisterUser(input model.RegisterUserInput) (*model.User, error) {
	existing, err := u.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error to hash password")
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := u.repo.Create(&user); err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}

func (u *UserUseCase) LoginUser(input model.LoginUserInput) (string, error) {
	user, err := u.repo.FindByEmail(input.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
