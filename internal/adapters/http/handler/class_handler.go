package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	ClassUsecase usecase.ClassUseCase
}

func (handler ClassHandler) GetAllClasses() echo.HandlerFunc {
	return func(e echo.Context) error {
		var classes []entity.Class

		classes, err := handler.ClassUsecase.GetAllClasses()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all class",
			"data":   classes,
		})
	}
}

func (handler ClassHandler) GetClass() echo.HandlerFunc {
	return func(e echo.Context) error {
		var class entity.Class
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		class, err = handler.ClassUsecase.GetClass(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get class by id",
			"data":   class,
		})
	}
}

func (handler ClassHandler) CreateClass() echo.HandlerFunc {
	return func(e echo.Context) error {
		var class entity.Class
		if err := e.Bind(&class); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(class); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}

		return e.JSON(
			http.StatusCreated, map[string]interface{}{
			"message": "success create new class",
			"data":   class,
		})
	}
}

func (handler ClassHandler) DeleteClass() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.ClassUsecase.DeleteClass(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Class`",
		})
	}
}
