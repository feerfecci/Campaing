package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	campaign := NewCampaign(name, content, contacts)

	//assert (identifiçao de erros)
	// println(campaign.ID)
	// assert.Equal(campaign.ID, "1")

	// if campaign.ID != "1" {
	// 	t.Errorf("expected 1")
	// } else if campaign.Name != name {
	// 	t.Errorf("expected correct Name")
	// } else if campaign.Content != content {
	// 	t.Errorf("expected correct Content")
	// } else if len(campaign.Contacts) != len(contacts) {
	// 	t.Errorf("expected correct Contacts")
	// }

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)

}

func Test_NewCampaign_CreatedOnMustBeBow(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	//act (primeiras declarações de objetos)
	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)

}
