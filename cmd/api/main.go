package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// "campaing/internal/domain/campaign"

// "github.com/go-playground/validator/v10"
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	http.ListenAndServe(":3000", r)
}

/*

type product struct {
	ID   int
	Name string
}
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	r.Post("/campaign", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaign
		render.DecodeJSON(r.Body, &request)

		id, err := service.Create(request)

		if err != nil {

			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, 500)
			} else {
				render.Status(r, 400)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, 201) //status de criação
		render.JSON(w, r, map[string]string{"id": id})

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

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("before")
		next.ServeHTTP(w, r)
		println("after")
	})
}
func myMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("request: ", r.Method, " - url: ", r.URL)
		next.ServeHTTP(w, r)
		println("after2")
	})
}
*/
