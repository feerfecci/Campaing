package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetByID(ID string) (*Campaign, error)
}
