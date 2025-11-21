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
	r.Post("/new_campaign", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaign", endpoints.HandlerError(handler.CampaignGet))
	r.Get("/campaign/{id}", endpoints.HandlerError(handler.CampaignGetById))
	r.Patch("/campaign/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
	r.Delete("/campaign/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))

	http.ListenAndServe(":3000", r)
}
