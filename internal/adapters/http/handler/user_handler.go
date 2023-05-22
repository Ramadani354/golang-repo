package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserUsecase usecase.UserUseCase
}

func (handler UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(e echo.Context) error {
		var users []entity.User

		users, err := handler.UserUsecase.GetAllUsers()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all users",
			"users":   users,
		})
	}
}

func (handler UserHandler) GetUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		user, err = handler.UserUsecase.GetUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get user",
			"user":    user,
		})
	}
}

func (handler UserHandler) CreateUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}

		// Validasi email unik
		if err := handler.UserUsecase.UniqueEmail(user.Email); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}
		user.Password = string(hashedPassword)
		// Set Role default cutomer
		user.Role = "customer"

		err = handler.UserUsecase.CreateUser(user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}

		return e.JSON(http.StatusCreated, user)
	}
}

func (handler UserHandler) DeleteUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.UserUsecase.DeleteUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete User`",
		})
	}
}
