package api

import (
	"context"
	"time"
	"todo-list/domain"
	"todo-list/dto"
	"todo-list/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type categoryApi struct {
	categoryService domain.CategoryService
}

func NewCategoryApi(app *fiber.App, categoryService domain.CategoryService, jwtMiddleware fiber.Handler) {
	api := categoryApi{categoryService: categoryService}
	app.Get("/category", jwtMiddleware, api.IndexByUser) // show all category user
	app.Post("/category", jwtMiddleware, api.Create)     // create category
	app.Delete("/category", jwtMiddleware, api.Delete)   // delete category
}

func (ca *categoryApi) IndexByUser(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	res, err := ca.categoryService.IndexUser(c, ctx)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(dto.ResponseSucsessData(200, "Success Get All Category", res))
}

func (ca *categoryApi) Create(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCategoryRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return err
	}
	fails := utils.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, "Validation failed, please check your input data"))
	}
	err := ca.categoryService.Create(c, ctx, req)

	if err != nil {
		return err
	}
	return ctx.Status(201).JSON(dto.ResponseSucsess(fiber.StatusCreated, "Success Create Category"))

}

func (ca *categoryApi) Delete(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.IdCategoryRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return err
	}
	fails := utils.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, "Validation failed, please check your input data"))
	}
	err := ca.categoryService.Delete(c, ctx, req.Id)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(dto.ResponseSucsess(fiber.StatusOK, "Success Delete Category"))
}
