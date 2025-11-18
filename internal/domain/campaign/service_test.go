package campaign

import (
	"campaing/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Teste y",
		Content: "Body",
		Emails:  []string{"teste1@em.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)

}
func Test_Create_SaveCampaign(t *testing.T) {
	// assert := assert.New(t)
	newCampaign := contract.NewCampaign{
		Name:    "Teste y",
		Content: "Body",
		Emails:  []string{"teste1@em.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)

	// assert.NotNil(id)
	// assert.Nil(err)

	repositoryMock.AssertExpectations(t)

}
