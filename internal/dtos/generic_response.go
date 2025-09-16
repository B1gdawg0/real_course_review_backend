package dtos

import (
	"errors"
	"net/http"

	errs "github.com/B1gdawg0/real_course_review_backend/internal/errors"
	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Status string `json:"status"`
	Data   *T     `json:"data,omitempty"`
	Error  *Error `json:"error,omitempty"`
	Meta   any    `json:"meta,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func HTTPStatusFromError(err error) int {
	if err == nil {
		return http.StatusOK
	}
	

	var appErr *errs.AppError
	if errors.As(err, &appErr) {
		return appErr.HTTPStatus
	}

	return http.StatusInternalServerError
}

func WrapResult[T any](data T, err error) Response[T] {
	if err != nil {
		var appErr *errs.AppError
		if !errors.As(err, &appErr) {
			appErr = errs.New("INTERNAL_ERROR", "Something went wrong", http.StatusInternalServerError, err)
		}

		return Response[T]{
			Status: "error",
			Error: &Error{
				Code:    appErr.Code,
				Message: appErr.Message,
			},
		}
	}

	return Response[T]{
		Status: "success",
		Data:   &data,
	}
}

func WrapResultWithMeta[T any](data T, err error, meta any) Response[T] {
	response := WrapResult(data, err)
	if err == nil && meta != nil {
		response.Meta = meta
	}
	return response
}

func Respond[T any](c *fiber.Ctx, data T, err error) error {
	statusCode := HTTPStatusFromError(err)
	response := WrapResult(data, err)
	return c.Status(statusCode).JSON(response)
}

func RespondWithMeta[T any](c *fiber.Ctx, data T, err error, meta any) error {
	statusCode := HTTPStatusFromError(err)
	response := WrapResultWithMeta(data, err, meta)
	return c.Status(statusCode).JSON(response)
}