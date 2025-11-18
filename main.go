package main

import (
	"campaing/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("nenhum erro")
	} else {
		validateError := err.(validator.ValidationErrors)
		for _, v := range validateError {
			switch v.Tag() {
			case "required":
				println(v.StructField() + "is required:" + v.Param())

			case "min":
				println(v.StructField() + "is required min " + v.Param())

			case "max":
				println(v.StructField() + "is required min " + v.Param())
			case "email":
				println(v.StructField() + "is invalid:" + v.Tag())
			}
		}
	}
}
