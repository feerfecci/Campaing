package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// "campaing/internal/domain/campaign"

// "github.com/go-playground/validator/v10"

type product struct {
	ID   int
	Name string
}

func main() {

	//ROUTES
	r := chi.NewRouter()

	//USAR MIDDLEWARE PRÓPRIO
	// r.Use(myMiddleware2)
	// r.Use(myMiddleware)

	//OU USAR OS DO CHI
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		println("endpoint")
		// product := r.URL.Query().Get("product")
		// id := r.URL.Query().Get("id")
		// if product != "" {
		// 	w.Write([]byte(product + id))

		// } else {

		// 	w.Write([]byte("paramteste"))
		// }
	})

	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "aplication/json")
		// obj := map[string]string{"message": "sucess"}
		// b, _ := json.Marshal(obj) // trasnforma em bite
		// w.Write(b)
		obj := map[string]string{"message": "sucess"}
		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
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
