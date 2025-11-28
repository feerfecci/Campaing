package main

import (
	"campaing/internal/domain/campaign"
	"campaing/internal/endpoints"
	"campaing/internal/infrastruct/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaign", func(r chi.Router) {
		r.Use(endpoints.Auth)

		r.Post("/new_campaign", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/", endpoints.HandlerError(handler.CampaignGet))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		// r.Patch("/campaign/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))
	})

	http.ListenAndServe(":3000", r)
}
