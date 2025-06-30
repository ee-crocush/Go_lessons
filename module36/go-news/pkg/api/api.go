package api

import "github.com/gofiber/fiber/v2"

// Error представляет ошибку с кодом.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrWithCode возвращает структуру Error с заданным кодом и сообщением.
func ErrWithCode(code, msg string) Error {
	return Error{Code: code, Message: msg}
}

// Err оборачивает error в стандартный ответ с кодом "internal-error".
func Err(err error) Error {
	return Error{Code: "internal-error", Message: err.Error()}
}

// Req читает JSON‑тело и возвращает DTO.
func Req[T any](c *fiber.Ctx) (T, error) {
	var req T
	if err := c.BodyParser(&req); err != nil {
		return req, err
	}
	return req, nil
}

// Resp возвращает данные ответа
func Resp[T any](data T) T { return data }
