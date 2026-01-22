package validation

import (
	"fmt"
	"strings"

	"github.com/binhbeng/goex/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) map[string]string {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationError {
			root := strings.Split(e.Namespace(), ".")[0]
			rawPath := strings.TrimPrefix(e.Namespace(), root+".")
			parts := strings.Split(rawPath, ".")[1:]
			fieldPath := strings.Join(parts, ".")

			switch e.Tag() {
			case "gt":
				errors[fieldPath] = fmt.Sprintf("%s must be greater than %s", fieldPath, e.Param())
			case "lt":
				errors[fieldPath] = fmt.Sprintf("%s must be less than %s", fieldPath, e.Param())
			case "gte":
				errors[fieldPath] = fmt.Sprintf("%s must be greater than or equal to %s", fieldPath, e.Param())
			case "lte":
				errors[fieldPath] = fmt.Sprintf("%s must be less than or equal to %s", fieldPath, e.Param())
			case "min":
				errors[fieldPath] = fmt.Sprintf("%s must be longer than %s characters", fieldPath, e.Param())
			case "max":
				errors[fieldPath] = fmt.Sprintf("%s must be shorter than %s characters", fieldPath, e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[fieldPath] = fmt.Sprintf("%s must be one of the values: %s", fieldPath, allowedValues)
			case "required":
				errors[fieldPath] = fmt.Sprintf("%s is required", fieldPath)
			}
		}

		return errors
	}

	return nil
}

func ValidateQueryParams(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		errs := HandleValidationErrors(err)
		utils.HttpBadRequest(c, "Invalid params", errs)
		return err
	}

	return nil
}

func ValidateBodyParams(c *gin.Context, obj any) error {
	if err := c.ShouldBind(obj); err != nil {
		errs := HandleValidationErrors(err)
		utils.HttpBadRequest(c, "Invalid body request", errs)
	}

	return nil
}
