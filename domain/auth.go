package domain

import (
	"context"
	"todo-list/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
	Registrasi(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
