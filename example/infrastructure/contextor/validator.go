package contextor

import (
	"fmt"
	"github.com/aksara-tech/aksarabase/utils/stringtor"
	"github.com/go-playground/validator/v10"
)

type ErrorValidation struct {
	Fields []string `json:"fields"`
	Errors []string `json:"errors"`
}

func (v ErrorValidation) Error() string {
	return v.Errors[0]
}

func (c Contextor) Validate(dest interface{}) error {
	errr := ErrorValidation{
		Fields: []string{},
		Errors: []string{},
	}

	validate := validator.New()
	err := validate.Struct(dest)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return fmt.Errorf("invalid validation")
		}
		for _, validationError := range validationErrors {
			errr.Fields = append(errr.Fields, validationError.Field())
			switch validationError.Tag() {
			case "required":
				errr.Errors = append(errr.Errors, fmt.Sprintf("%v field is required", stringtor.ToSnakeCase(validationError.Field())))
				break
			case "email":
				errr.Errors = append(errr.Errors, fmt.Sprintf("%v field is not valid email", stringtor.ToSnakeCase(validationError.Field())))
				break
			case "gte":
				errr.Errors = append(errr.Errors, fmt.Sprintf("%s value must be greater than %s",
					stringtor.ToSnakeCase(validationError.Field()), validationError.Param()))
				break
			case "lte":
				errr.Errors = append(errr.Errors, fmt.Sprintf("%s value must be less than %s",
					stringtor.ToSnakeCase(validationError.Field()), validationError.Param()))
				break
			}

			break

		}
		return errr
	}

	return nil
}
