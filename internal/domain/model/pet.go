package model

type PetStatus string

const (
	PetStatusValiable PetStatus = "vailable"
	PetStatusPending  PetStatus = "pending"
	PetStatusSold     PetStatus = "sold"
)

type PetID int64

type Pet struct {
	ID        PetID     `validate:"required"`
	Category  *Category `validate:"required"`
	Name      string    `validate:"required"`
	PhotoURLs []string
	Tags      []*Tag
	Status    PetStatus `validate:"required"`
}

func NewPet(id PetID, category *Category, name string, photoURLs []string, tags []*Tag, status PetStatus) (*Pet, error) {
	pet := Pet{
		ID:        id,
		Category:  category,
		Name:      name,
		PhotoURLs: photoURLs,
		Tags:      tags,
		Status:    status,
	}
	if err := validate(pet); err != nil {
		return nil, err
	}
	return &pet, nil
}

func (p *Pet) Rename(name string) error {
	newPet := *p
	newPet.Name = name
	if err := validate(newPet); err != nil {
		return err
	}
	*p = newPet
	return nil
}

func (p *Pet) ChangePhotos(photoURLs []string) error {
	newPet := *p
	newPet.PhotoURLs = photoURLs
	if err := validate(newPet); err != nil {
		return err
	}
	*p = newPet
	return nil
}
