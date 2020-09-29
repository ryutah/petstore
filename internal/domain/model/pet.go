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
	PhotoURLs []string  `validate:"required"`
	Tags      []*Tag
	Status    PetStatus `validate:"required"`
}
