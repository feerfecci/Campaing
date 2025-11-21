package campaign

import (
	"campaing/internal/contract"
	internalerrors "campaing/internal/internalErrors"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	Get() ([]contract.CampaignReponse, error)
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

	err = s.Repository.Create(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s ServiceImp) Get() ([]contract.CampaignReponse, error) {
	campaigns, err := s.Repository.Get()
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	// Converter para o contrato de resposta
	var response []contract.CampaignReponse
	for _, c := range campaigns {
		response = append(response, contract.CampaignReponse{
			ID:            c.ID,
			Name:          c.Name,
			Status:        c.Status,
			AmountOfEmail: len(c.Contacts),
		})
	}

	return response, nil
}

func (s *ServiceImp) GetByID(idCampaign string) (*contract.CampaignReponse, error) {
	campaign, err := s.Repository.GetByID(idCampaign)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.CampaignReponse{
		ID:            campaign.ID,
		Name:          campaign.Name,
		Content:       campaign.Content,
		Status:        campaign.Status,
		AmountOfEmail: len(campaign.Contacts),
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
	err = s.Repository.Update(campaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil

}
func (s *ServiceImp) DeleteByID(idCampaign string) error {
	campaign, err := s.Repository.GetByID(idCampaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil

}
