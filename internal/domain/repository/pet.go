package repository

import (
	"context"

	"github.com/ryutah/petstore/internal/domain/model"
)

type Pet interface {
	NextID(context.Context) (model.PetID, error)
	Store(context.Context, model.Pet) error
}
