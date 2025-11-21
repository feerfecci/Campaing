package campaign

import (
	"campaing/internal/contract"
	internalerrors "campaing/internal/internalErrors"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetByID(idCampaign int) (*contract.CampaignReponse, error)
	CancelByID(idCampaign string) error
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetByID(idCampaign string) (*contract.CampaignReponse, error) {
	campaign, err := s.Repository.GetByID(idCampaign)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.CampaignReponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil

}
func (s *ServiceImp) CancelByID(idCampaign string) error {
	campaign, err := s.Repository.GetByID(idCampaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Cancel()
	err = s.Repository.Save(campaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil

}
