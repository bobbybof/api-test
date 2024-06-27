package helper

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func validationErrorMsg(field string, param string, tag string, typeName string) string {
	switch tag {
	case "min":
		if typeName == "int16" || typeName == "int64" {
			return fmt.Sprintf("%s(number) value min %s", field, param)
		}
		return fmt.Sprintf("%s character min %s character", field, param)
	case "max":
		if typeName == "int16" || typeName == "int64" {
			return fmt.Sprintf("%s(number) value max is %s", field, param)
		}
		return fmt.Sprintf("%s character max %s character", field, param)
	case "required":
		return fmt.Sprintf("%s is %s", field, tag)
	default:
		return fmt.Sprintf("%s %s", field, tag)
	}
}

func ValidationErrorResponse(err error) gin.H {
	var ves validator.ValidationErrors
	var temp []APIValidationError
	if errors.As(err, &ves) {
		out := make([]APIValidationError, len(ves))
		for i, ve := range ves {
			out[i] = APIValidationError{ve.Field(), validationErrorMsg(ve.Field(), ve.Param(), ve.ActualTag(), ve.Type().Name())}
		}
		temp = out
	}
	return gin.H{
		"message": "Validation error",
		"errors":  temp,
	}
}

func ErrorHttpResponse(err error, msg string) gin.H {
	message := "Something went wrong"

	if len(msg) > 0 {
		message = msg
	}

	return gin.H{
		"error":   err.Error(),
		"message": message,
	}
}
