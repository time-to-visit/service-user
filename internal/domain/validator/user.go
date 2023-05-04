package validator

import (
	"net/http"

	"service-user/internal/domain/entity"
	objectValues "service-user/internal/domain/object_values"
	validatorPer "service-user/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func ValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		user := new(entity.User)

		_ = c.Bind(&user)
		if err := v.Struct(user); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("user-data", user)
		return next(c)
	}
}

func ValidateVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		verify := new(objectValues.Verify)

		_ = c.Bind(&verify)
		if err := v.Struct(verify); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("verify", verify)
		return next(c)
	}
}
