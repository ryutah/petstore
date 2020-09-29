package repository

import (
	"context"

	"github.com/ryutah/petstore/internal/domain/model"
)

type Category interface {
	Get(context.Context, model.CategoryID) (*model.Category, error)
}
