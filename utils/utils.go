package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/natalizhy/pupils_grade_sheet/models"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"regexp"
)


var userError = map[string]map[string]string{
	"username": {
		"required": "Обязательное поле",
		"cyr":      "Можна использовать только буквы",
		"min":      "Слишком мало символов",
	},
	"password": {
		"required": "Обязательное поле",
		"cyr":      "Можна использовать только буквы",
		"min":      "Слишком мало символов",
	},
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GenerateId() string  {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

var validate = validator.New()

func ValidateCyr(fl validator.FieldLevel) bool {
	result := regexp.MustCompile("^[a-zA-ZА-Яа-я]+$")
	return result.MatchString(fl.Field().String())
}

func ValidateUser(user models.User) (errors map[string]map[string]string, err error) {
	err = validate.RegisterValidation("cyr", ValidateCyr)
	if err != nil {
		fmt.Println(err)
	}

	err = validate.Struct(user)

	errors = make(map[string]map[string]string)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			if _, ok := errors[err.Field()]; !ok {
				errors[err.Field()] = make(map[string]string)
			}

			errors[err.Field()][err.Tag()] = userError[err.Field()][err.Tag()]
		}
	}

	return
}
