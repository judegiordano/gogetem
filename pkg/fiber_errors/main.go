package fiber_errors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/judegiordano/gogetem/pkg/logger"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func parseCodeError(code int) string {
	switch code {
	case 400:
		return "BAD_REQUEST"
	case 401:
		return "UNAUTHORIZED"
	case 403:
		return "FORBIDDEN"
	case 404:
		return "NOT_FOUND"
	default:
		return "INTERNAL_SERVER_ERROR"
	}
}

func BadRequest(c *fiber.Ctx, err error) error {
	logger.Error("[BAD REQUEST]", err)
	return c.Status(400).JSON(ErrorResponse{Error: parseCodeError(400), Message: err.Error()})
}

func Unauthorized(c *fiber.Ctx, err error) error {
	logger.Error("[UNAUTHORIZED]", err)
	return c.Status(401).JSON(ErrorResponse{Error: parseCodeError(401), Message: err.Error()})
}

func Forbidden(c *fiber.Ctx, err error) error {
	logger.Error("[FORBIDDEN]", err)
	return c.Status(403).JSON(ErrorResponse{Error: parseCodeError(403), Message: err.Error()})
}

func NotFound(c *fiber.Ctx, err error) error {
	logger.Error("[NOT FOUND]", err)
	return c.Status(404).JSON(ErrorResponse{Error: parseCodeError(404), Message: err.Error()})
}

func InternalServerError(c *fiber.Ctx, err error) error {
	logger.Error("[INTERNAL SERVER ERROR]", err)
	return c.Status(500).JSON(ErrorResponse{Error: parseCodeError(500), Message: err.Error()})
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e string
	message := err.Error()
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}
	logger.Error(e, code, message)
	return ctx.Status(code).JSON(ErrorResponse{
		Error:   parseCodeError(code),
		Message: message,
	})
}
