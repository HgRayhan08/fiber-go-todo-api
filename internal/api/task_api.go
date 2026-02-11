package api

import (
	"context"
	"time"
	"todo-list/domain"
	"todo-list/dto"
	"todo-list/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type taskApi struct {
	todoService domain.TaskService
}

func NewTaskApi(app *fiber.App, todoService domain.TaskService, jwtMidd fiber.Handler) {
	task := taskApi{todoService: todoService}

	app.Get("/todo", jwtMidd, task.Index)
	app.Post("/todo", jwtMidd, task.Create)
	app.Delete("/todo", jwtMidd, task.Delete)
}

func (t *taskApi) Index(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := t.todoService.Index(c, ctx)
	if err != nil {
		return err
	}
	status := ctx.Response().StatusCode()

	return ctx.Status(200).JSON(dto.ResponseSucsessData(status, "Success Get All Todo", res))
}

func (t *taskApi) Create(ctx fiber.Ctx) error {

	// panic("unimplemented")
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateTaskRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return err
	}
	fails := utils.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, "Validation failed, please check your input data"))
	}
	err := t.todoService.Create(c, ctx, req)
	if err != nil {
		return err
	}
	return ctx.Status(201).JSON(dto.ResponseSucsess(fiber.StatusCreated, "Success Create Task"))
}

func (t *taskApi) Delete(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.DeleteTaskRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return err
	}
	fails := utils.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, "Validation failed, please check your input data"))
	}

	err := t.todoService.Delete(c, ctx, req.Id)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(dto.ResponseSucsess(fiber.StatusOK, "Success Delete Task"))
}
