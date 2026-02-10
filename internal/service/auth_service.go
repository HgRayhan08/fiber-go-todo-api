package service

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"todo-list/domain"
	"todo-list/dto"
	"todo-list/internal/config"

	"github.com/golang-jwt/jwt/v5"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuthService(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &AuthService{
		conf:           cnf,
		userRepository: userRepository,
	}
}

// Login implements [domain.AuthService].
func (a *AuthService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	cekEmail, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, errors.New("Email tidak terdaftar")
	}
	if cekEmail.Id == "" {
		return dto.AuthResponse{}, errors.New("Failed to login, email tidak terdaftar")
	}
	err = bcrypt.CompareHashAndPassword([]byte(cekEmail.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("Password salah")
	}

	claim := jwt.MapClaims{
		"id":  cekEmail.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Expire) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("Failed to generate token")
	}
	return dto.AuthResponse{
		Code:    200,
		Message: "Login Success",
		User: dto.User{
			ID:    cekEmail.Id,
			Email: cekEmail.Email,
		},
		Token: tokenStr,
	}, nil
}

// Registrasi implements [domain.AuthService].
func (a *AuthService) Registrasi(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	cekEmail, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, errors.New("Email sudah terdaftar")
	}

	if cekEmail.Id != "" {
		return dto.AuthResponse{}, errors.New("Failed to registrasi, email sudah terdaftar")
	}

	passwordBycrip, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	user := domain.User{
		Id:        uuid.New().String(),
		Email:     req.Email,
		Password:  string(passwordBycrip),
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
		UpdatedAt: sql.NullTime{Valid: false},
	}
	err = a.userRepository.Save(ctx, user)

	if err != nil {
		return dto.AuthResponse{}, err
	}
	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Expire) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	return dto.AuthResponse{
		Code:    201,
		Message: "Register is Success",
		User: dto.User{
			ID:    user.Id,
			Email: user.Email,
		},
		Token: tokenStr,
	}, nil

}
