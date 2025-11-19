package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// "campaing/internal/domain/campaign"

// "github.com/go-playground/validator/v10"

func main() {

	//ROUTES
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Query().Get("product")
		id := r.URL.Query().Get("id")
		if product != "" {
			w.Write([]byte(product + id))

		} else {

			w.Write([]byte("paramteste"))
		}
	})
	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})
	http.ListenAndServe(":3000", r)

	//AULA VALIDAÇÃO
	// campaign := campaign.Campaign{}
	// validate := validator.New()
	// err := validate.Struct(campaign)

	// if err == nil {
	// 	println("nenhum erro")
	// } else {
	// 	validateError := err.(validator.ValidationErrors)
	// 	for _, v := range validateError {
	// 		switch v.Tag() {
	// 		case "required":
	// 			println(v.StructField() + "is required:" + v.Param())

	// 		case "min":
	// 			println(v.StructField() + "is required min " + v.Param())

	// 		case "max":
	// 			println(v.StructField() + "is required min " + v.Param())
	// 		case "email":
	// 			println(v.StructField() + "is invalid:" + v.Tag())
	// 		}
	// 	}
	// }
}
