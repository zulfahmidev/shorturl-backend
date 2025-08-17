package validator

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/zulfahmidev/shorturl/pkg/database"
	"github.com/zulfahmidev/shorturl/pkg/logger"
	"go.uber.org/zap"
)

var validate = validator.New()

func InitValidator() {
	validate.RegisterValidation("exists", ExistsValidator)
	validate.RegisterValidation("unique", UniqueValidator)
}

func ValidateStruct(s interface{}) *map[string]string {
	errors := make(map[string]string)

	if err := validate.Struct(s); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			logger.Default().Info("eyyo2", zap.Any("test", fieldErr))
			errors[fieldErr.Field()] = CustomMessage(fieldErr.Tag(), fieldErr.Field())
		}
	}
	return &errors
}

func CustomMessage(tag string, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "exists":
		return fmt.Sprintf("%s does'nt exist", field)
	case "unique":
		return fmt.Sprintf("%s already exist", field)
	case "email":
		return fmt.Sprintf("%s is must be an email", field)
	case "image":
		return fmt.Sprintf("%s is must be an image", field)
	case "url":
		return fmt.Sprintf("%s is must be an url", field)
	default:
		return tag
	}
}

func ExistsValidator(fl validator.FieldLevel) bool {
	param := fl.Param()
	args := strings.Split(param, ".")

	if len(args) != 2 {
		return false
	}

	tableName := args[0]
	fieldName := args[1]
	value := fl.Field().Interface()

	var count int64
	db := database.GetDB()
	err := db.Table(tableName).Where(fieldName+"= $1", value).Count(&count).Error
	if err != nil {
		return false
	}

	return count > 0
}

func UniqueValidator(fl validator.FieldLevel) bool {
	param := fl.Param()
	args := strings.Split(param, ".")

	if len(args) != 2 {
		return false
	}

	tableName := args[0]
	fieldName := args[1]
	value := fl.Field().Interface()

	var count int64
	db := database.GetDB()
	err := db.Table(tableName).Where(fieldName+"= $1", value).Count(&count).Error
	if err != nil {
		return false
	}

	return count == 0
}

func IsValidURL(str string) bool {
	u, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	if u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}