package tag

import (
	"context"
	"server/internal/domain"
	"server/internal/ent"
	tagent "server/internal/ent/tag"
	userent "server/internal/ent/user"
	"server/internal/platform"
)

type Service struct {
	client *ent.TagClient
}

func NewService(client *ent.Client) *Service {
	return &Service{client: client.Tag}
}

func (s Service) All(ctx context.Context, dto domain.AllTagsDTO) ([]*domain.Tag, error) {
	tags, err := s.client.Query().
		Where(tagent.HasUserWith(userent.ID(dto.UserID))).
		All(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	return domain.TagsFromEnt(tags), nil
}

func (s Service) ByID(ctx context.Context, dto domain.TagsByIDDTO) (*domain.Tag, error) {
	tag, err := s.client.Query().
		Where(tagent.HasUserWith(userent.ID(dto.UserID)),
			tagent.ID(dto.TagID)).
		Only(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	return domain.TagFromEnt(tag), nil
}

func (s Service) Create(ctx context.Context, dto domain.CreateTagDTO) (*domain.Tag, error) {
	tag, err := s.client.Create().
		SetTitle(dto.Title).
		SetUserID(dto.UserID).
		Save(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	return domain.TagFromEnt(tag), nil
}

func (s Service) Update(ctx context.Context, dto domain.TagPutDTO) error {
	_, err := s.client.Update().
		Where(tagent.ID(dto.TagID)).
		SetTitle(dto.Title).
		SetMult(dto.Mult).
		SetUserID(dto.UserID).
		Save(ctx)
	if err != nil {
		return platform.WrapInternal(err)
	}

	return nil
}

func (s Service) Delete(ctx context.Context, dto domain.DeleteTagDTO) error {
	_, err := s.client.Delete().Where(tagent.ID(dto.TagID)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
