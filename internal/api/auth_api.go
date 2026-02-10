package api

import (
	"context"
	"time"
	"todo-list/domain"
	"todo-list/dto"
	"todo-list/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuthApi(app *fiber.App, authService domain.AuthService) {
	auth := authApi{authService: authService}
	app.Post("/login", auth.Login)
	app.Post("/registrasi", auth.Registrasi)
}

func (a *authApi) Login(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.AuthRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return err
	}
	user, err := a.authService.Login(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, err.Error()))
	}
	return ctx.Status(200).JSON(user)
}

func (a *authApi) Registrasi(ctx fiber.Ctx) error {

	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.AuthRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return err
	}

	fails := utils.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, "Validation failed, please check your input data"))
	}

	if len(req.Password) <= 8 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, "Password harus lebih dari 8 karakter"))
	}

	user, err := a.authService.Registrasi(c, req)
	if err != nil {
		return err
	}
	return ctx.Status(201).JSON(user)
}
