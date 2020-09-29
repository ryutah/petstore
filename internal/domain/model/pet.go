package model

type PetStatus string

const (
	PetStatusValiable PetStatus = "vailable"
	PetStatusPending  PetStatus = "pending"
	PetStatusSold     PetStatus = "sold"
)

type PetID int64

type Pet struct {
	ID        PetID
	Category  *Category
	Name      string
	PhotoURLs []string
	Tags      []*Tag
	Status    PetStatus
}
