package repository

import (
	"context"

	"github.com/ryutah/petstore/internal/domain/model"
)

type Pet interface {
	NextID(context.Context) (model.PetID, error)
	Get(context.Context, model.PetID) (*model.Pet, error)
	Store(context.Context, model.Pet) error
	Replace(context.Context, model.Pet) error
}
