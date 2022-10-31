package validator

import (
	"fmt"
	"github.com/aksara-tech/aksarabase/utils/stringtor"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func SetupValidation(r *echo.Echo) {
	r.Validator = &CustomValidator{
		Validator: validator.New(),
	}

	r.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			var errorMessages interface{}

			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					errorMessages = fmt.Sprintf("%s is required",
						stringtor.ToSnakeCase(err.Field()))
					break
				case "email":
					errorMessages = fmt.Sprintf("%s is not valid email",
						stringtor.ToSnakeCase(err.Field()))
					break
				case "gte":
					errorMessages = fmt.Sprintf("%s value must be greater than %s",
						stringtor.ToSnakeCase(err.Field()), err.Param())
					break
				case "lte":
					errorMessages = fmt.Sprintf("%s value must be lower than %s",
						stringtor.ToSnakeCase(err.Field()), err.Param())
					break
				}

				break
			}
			report.Message = errorMessages
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

}
