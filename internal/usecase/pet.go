package usecase

import (
	"context"

	"github.com/ryutah/petstore/internal/domain/model"
	"github.com/ryutah/petstore/internal/domain/repository"
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
		Add(context.Context, PetAddRequest, PetAddOutputPort) (succeeded bool)
	}

	PetAddOutputPort interface {
		errorOutputPort
		Succeeded(context.Context)
	}
)

func (p PetAddRequest) tagIDs() []model.TagID {
	ids := make([]model.TagID, len(p.Tags))
	for i, t := range p.Tags {
		ids[i] = model.TagID(t.ID)
	}
	return ids
}

type (
	PetFindRequest struct {
		ID int64
	}

	PetFindResult struct {
		ID       int64
		Name     string
		Category struct {
			ID   int64
			Name string
		}
		PhotoURLs []string
		Tags      []struct {
			ID   int64
			Name string
		}
		Status string
	}

	PetFindInputPort interface {
		Find(context.Context, PetFindRequest, PetFindOutputPort) (succeeded bool)
	}

	PetFindOutputPort interface {
		errorOutputPort
		Succeeded(context.Context, PetFindResult)
	}
)

func newPetFindResult(pet model.Pet) PetFindResult {
	ret := PetFindResult{
		ID:        int64(pet.ID),
		Name:      pet.Name,
		PhotoURLs: pet.PhotoURLs,
		Status:    string(pet.Status),
	}
	if pet.Category != nil {
		ret.Category = struct {
			ID   int64
			Name string
		}{
			ID:   int64(pet.Category.ID),
			Name: pet.Category.Name,
		}
	}

	tags := make([]struct {
		ID   int64
		Name string
	}, len(pet.Tags))
	for i, tag := range pet.Tags {
		tags[i] = struct {
			ID   int64
			Name string
		}{
			ID:   int64(tag.ID),
			Name: tag.Name,
		}
	}
	return ret
}

type Pet struct {
	repository struct {
		category repository.Category
		tag      repository.Tag
		pet      repository.Pet
	}
}

var (
	_ PetAddInputPort  = (*Pet)(nil)
	_ PetFindInputPort = (*Pet)(nil)
)

func (p *Pet) Add(ctx context.Context, req PetAddRequest, out PetAddOutputPort) (succeeded bool) {
	handleNoSuchEntity := errorHandlerFunc(func(ctx context.Context, err error, out errorOutputPort) {
		out.InvalidInput(ctx)
	})

	category, err := p.repository.category.Get(ctx, model.CategoryID(req.Category.ID))
	if err != nil {
		return handleError(ctx, err, out, withNoSuchEntityFunc(handleNoSuchEntity))
	}
	tags, err := p.repository.tag.GetMulti(ctx, req.tagIDs())
	if err != nil {
		return handleError(ctx, err, out, withNoSuchEntityFunc(handleNoSuchEntity))
	}

	petID, err := p.repository.pet.NextID(ctx)
	if err != nil {
		return handleError(ctx, err, out)
	}
	newPet, err := model.NewPet(petID, category, req.Name, req.PhotoURLs, tags, model.PetStatus(req.Status))
	if err != nil {
		return handleError(ctx, err, out)
	}

	if err := p.repository.pet.Store(ctx, *newPet); err != nil {
		return handleError(ctx, err, out)
	}

	out.Succeeded(ctx)
	return true
}

func (p *Pet) Find(ctx context.Context, req PetFindRequest, out PetFindOutputPort) (succeeded bool) {
	pet, err := p.repository.pet.Get(ctx, model.PetID(req.ID))
	if err != nil {
		return handleError(ctx, err, out)
	}

	out.Succeeded(ctx, newPetFindResult(*pet))
	return true
}
