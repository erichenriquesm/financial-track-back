package utils

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON(ctx *gin.Context, obj interface{}) map[string]string {
	if err := ctx.ShouldBind(obj); err != nil {
		out := make(map[string]string)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				out[fe.Field()] = validationMessage(fe)
			}
			return out
		}

		var tpe *time.ParseError
		if errors.As(err, &tpe) {
			addTimeFormatErrors(obj, out)
			return out
		}

		var ute *json.UnmarshalTypeError
		if errors.As(err, &ute) {
			field := ute.Field
			if field == "" {
				field = "body"
			}
			out[field] = "Invalid type"
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
	case "datetime":
		return "Invalid datetime format. Expected: 2006-01-02 15:04"
	}
	return "Invalid value"
}

func addTimeFormatErrors(obj interface{}, out map[string]string) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if field.Type == reflect.TypeOf(time.Time{}) {
			format := field.Tag.Get("time_format")
			if format != "" {
				out[field.Name] = "Invalid datetime format. Expected: " + format
			}
		}
	}
}
