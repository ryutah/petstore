package repository

import (
	"context"

	"github.com/ryutah/petstore/internal/domain/model"
)

type Tag interface {
	GetMulti(context.Context, []model.TagID) ([]*model.Tag, error)
}
