package endpoints

import "campaing/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.ServiceImp
}
