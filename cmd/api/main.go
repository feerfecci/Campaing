package main

import (
	"campaing/internal/domain/campaign"
	"campaing/internal/endpoints"
	"campaing/internal/infrastruct/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaign", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaign/{id}", endpoints.HandlerError(handler.CampaignGetById))

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
