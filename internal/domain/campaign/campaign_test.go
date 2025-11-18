package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body Hi"
	contacts = []string{"email1@e.com", "email2@e.com"}
	fake     = faker.New()
)

func Test_NewCampaign(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	campaign, _ := NewCampaign(name, content, contacts)

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
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)

}

func Test_NewCampaign_CreatedOnMustBeBow(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	//act (primeiras declarações de objetos)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)

}

// func Test_NewCampaign_MustValidateName(t *testing.T) {
// 	//arrange (escopo de variáveis)
// 	assert := assert.New(t)

// 	//act (primeiras declarações de objetos)
// 	_, err := NewCampaign("", content, contacts)

// 	assert.Equal("name is required", err.Error())

// }
func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	_, err := NewCampaign("aa", content, contacts)

	assert.Equal("name is required min 5", err.Error())

}
func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required max 24", err.Error())

}

// func Test_NewCampaign_MustValidateContent(t *testing.T) {
// 	//arrange (escopo de variáveis)
// 	assert := assert.New(t)

// 	//act (primeiras declarações de objetos)
// 	_, err := NewCampaign(name, "", contacts)

// 	assert.Equal("content is required", err.Error())

// }
func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required min 5", err.Error())

}
func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is required max 1024", err.Error())

}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())

}
func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	//arrange (escopo de variáveis)
	assert := assert.New(t)

	//act (primeiras declarações de objetos)
	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required min 1", err.Error())

}
