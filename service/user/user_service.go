package user_service

import (
	"errors"
	"strings"

	"github.com/andrewidianto12/Manufacture-Rental/util"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
	Fullname string `json:"full_name" validate:"omitempty"`
	Phone    string `json:"phone" validate:"omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(input UserRegisterRequest) (*User, error) {
	if strings.TrimSpace(input.Username) == "" || strings.TrimSpace(input.Email) == "" || strings.TrimSpace(input.Password) == "" {
		return nil, errors.New("semua field harus diisi")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal memproses password")
	}

	user := &User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		Fullname: input.Fullname,
		Phone:    input.Phone,
		Role:     "user",
	}

	err = s.repo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) LoginUser(input LoginRequest) (string, error) {
	if strings.TrimSpace(input.Email) == "" || strings.TrimSpace(input.Password) == "" {
		return "", errors.New("semua field harus diisi")
	}

	user, err := s.repo.LoginUser(input.Email, input.Password)
	if err != nil {
		return "", err
	}

	token, err := util.GenerateToken(uint(user.ID), user.Fullname)
	if err != nil {
		return "", errors.New("gagal membuat token")
	}

	return token, nil
}

func (s *userService) DeleteUser(ID uint) error {
	return s.repo.DeleteUser(ID)
}
