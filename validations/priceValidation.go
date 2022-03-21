package validations

import (
	"demir/models"

	"github.com/go-playground/validator/v10"
)

func PriceValidation(price models.Price) (map[string]string, error) {
	v := validator.New()
	data := map[string]string{"message": ""}

	err := v.Struct(price)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			data["message"] += e.Field() + ","
		}
		data["message"] += " alanları hatalı"
		return data, err
	}
	data["message"] += "Success"
	return data, err
}
