package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GET /campaigns/{id}
func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := chi.URLParam(r, "id")
	campaigns, err := h.CampaignService.GetByID(id)

	return campaigns, 200, err
}
