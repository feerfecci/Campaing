package campaign

type Repository interface {
	Create(campaign *Campaign) error
	Delete(campaign *Campaign) error
	Update(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetByID(ID string) (*Campaign, error)
}
