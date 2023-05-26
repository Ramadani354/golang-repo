package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MajorHandler struct {
	MajorUsecase usecase.MajorUseCase
}

func (handler MajorHandler) GetAllMajors() echo.HandlerFunc {
	return func(e echo.Context) error {
		var majors []entity.Major

		majors, err := handler.MajorUsecase.GetAllMajors()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all major",
			"data":   majors,
		})
	}
}

func (handler MajorHandler) GetMajor() echo.HandlerFunc {
	return func(e echo.Context) error {
		var major entity.Major
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		major, err = handler.MajorUsecase.GetMajor(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get major by id",
			"data":   major,
		})
	}
}

func (handler MajorHandler) CreateMajor() echo.HandlerFunc {
	return func(e echo.Context) error {
		var major entity.Major
		if err := e.Bind(&major); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(major); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}

		return e.JSON(
			http.StatusCreated, map[string]interface{}{
			"message": "success create new major",
			"data":   major,
		})
	}
}

func (handler MajorHandler) DeleteMajor() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.MajorUsecase.DeleteMajor(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Major`",
		})
	}
}
