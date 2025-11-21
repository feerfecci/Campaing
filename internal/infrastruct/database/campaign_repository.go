package database

import (
	"campaing/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	// campaigns []campaign.Campaign
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	// c.campaigns = append(c.campaigns, *campaign)
	tx := c.Db.Create(campaign)
	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Find(&campaigns)
	return campaigns, tx.Error
}

func (c *CampaignRepository) GetByID(ID string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := c.Db.First(&campaign, ID)
	return &campaign, tx.Error
}
