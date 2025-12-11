package endpoints

import (
	"campaing/internal/contract"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)
	email := r.Context().Value("email").(string)
	request.CreatedBy = email
	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err
}

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaigns, err := h.CampaignService.Get()

	return campaigns, 200, err
}

// GET /campaigns/{id}
func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := chi.URLParam(r, "id")
	campaign, err := h.CampaignService.GetByID(id)
	if err == nil && campaign == nil {
		return nil, http.StatusNotFound, err
	}

	return campaign, 200, err
}

func (h *Handler) CampaignDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := chi.URLParam(r, "id")
	err := h.CampaignService.DeleteByID(id)

	return nil, 200, err
}
