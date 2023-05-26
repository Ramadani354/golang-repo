package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryUsecase usecase.CategoryUseCase
}

func (handler CategoryHandler) GetAllCategories() echo.HandlerFunc {
	return func(e echo.Context) error {
		var categoryes []entity.Category

		categoryes, err := handler.CategoryUsecase.GetAllCategories()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all category",
			"data":   categoryes,
		})
	}
}

func (handler CategoryHandler) GetCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		var category entity.Category
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		category, err = handler.CategoryUsecase.GetCategory(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get category by id",
			"data":   category,
		})
	}
}

func (handler CategoryHandler) CreateCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		var category entity.Category
		if err := e.Bind(&category); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(category); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}

		return e.JSON(
			http.StatusCreated, map[string]interface{}{
			"message": "success create new category",
			"data":   category,
		})
	}
}

func (handler CategoryHandler) DeleteCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.CategoryUsecase.DeleteCategory(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Category`",
		})
	}
}
