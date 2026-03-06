package handler

import (
	"fmt"
	"log"
	"net/http"

	user_service "github.com/andrewidianto12/Manufacture-Rental/service/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var input user_service.UserRegisterRequest

	if err := c.Bind(&input); err != nil {
		log.Println("Bind error:", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
	}

	if err := c.Validate(&input); err != nil {
		log.Println("Validation error:", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Validasi gagal",
			"error":   err.Error(),
		})
	}

	result, err := h.userService.RegisterUser(input)
	if err != nil {
		log.Println("Service error:", err)
		if err.Error() == "username already exists" {
			return c.JSON(http.StatusConflict, map[string]any{
				"message": "Username sudah digunakan",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Terjadi kesalahan pada server",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Berhasil register user",
		"data":    result,
	})
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	var input user_service.LoginRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Input tidak valid", "error": err.Error()})
	}

	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validasi gagal", "error": err.Error()})
	}

	token, err := h.userService.LoginUser(input)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Berhasil login user",
		"token":   token,
	})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	var ID uint
	_, err := fmt.Sscanf(idParam, "%d", &ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	err = h.userService.DeleteUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal menghapus user", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Berhasil menghapus user",
	})
}
