package usecase

import (
	"context"

	"github.com/ryutah/petstore/internal/domain/model"
	"github.com/ryutah/petstore/internal/domain/repository"
	"github.com/ryutah/petstore/internal/errors"
)

// Defenitions for Pet#Add
type (
	PetAddRequest struct {
		Name     string
		Category struct {
			ID int64
		}
		PhotoURLs []string
		Tags      []struct {
			ID int64
		}
		Status string
	}

	PetAddInputPort interface {
		Add(context.Context, PetAddRequest, PetAddOutputPort)
	}

	PetAddOutputPort interface {
		Suceceded(context.Context)
		InvalidInput(context.Context)
		ServerError(context.Context)
	}
)

func (p PetAddRequest) tagIDs() []model.TagID {
	ids := make([]model.TagID, len(p.Tags))
	for i, t := range p.Tags {
		ids[i] = model.TagID(t.ID)
	}
	return ids
}

type Pet struct {
	repository struct {
		category repository.Category
		tag      repository.Tag
		pet      repository.Pet
	}
}

var _ PetAddInputPort = (*Pet)(nil)

func (p *Pet) Add(ctx context.Context, req PetAddRequest, out PetAddOutputPort) {
	handleError := func(err error) {
		if errors.Is(err, errors.ErrNoSuchEntity) || errors.Is(err, errors.ErrInvalidInput) {
			out.InvalidInput(ctx)
		} else {
			out.ServerError(ctx)
		}
	}

	category, err := p.repository.category.Get(ctx, model.CategoryID(req.Category.ID))
	if err != nil {
		handleError(err)
		return
	}
	tags, err := p.repository.tag.GetMulti(ctx, req.tagIDs())
	if err != nil {
		handleError(err)
		return
	}
	petID, err := p.repository.pet.NextID(ctx)
	if err != nil {
		handleError(err)
		return
	}
	newPet, err := model.NewPet(petID, category, req.Name, req.PhotoURLs, tags, model.PetStatus(req.Status))
	if err != nil {
		handleError(err)
		return
	}

	if err := p.repository.pet.Store(ctx, *newPet); err != nil {
		handleError(err)
		return
	}

	out.Suceceded(ctx)
}
