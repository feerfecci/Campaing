package campaign

import "campaing/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {
	return nil
}
