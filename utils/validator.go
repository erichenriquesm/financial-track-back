package utils

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON(ctx *gin.Context, obj interface{}) map[string]string {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		out := make(map[string]string)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				out[fe.Field()] = validationMessage(fe)
			}
			return out
		}

		val := reflect.ValueOf(obj).Elem()
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("binding")
			if tag == "required" || containsRequired(tag) {
				out[field.Name] = "This field is required"
			}
		}

		return out
	}

	return nil
}

func containsRequired(tag string) bool {
	for _, t := range splitTag(tag) {
		if t == "required" {
			return true
		}
	}
	return false
}

// Separar múltiplas tags
func splitTag(tag string) []string {
	var parts []string
	current := ""
	for _, t := range tag {
		if t == ',' {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(t)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "gt":
		return "Must be greater than 0"
	case "email":
		return "Must be a valid email address"
	}
	return "Invalid value"
}
