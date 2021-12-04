package tag

import (
	"context"
	"server/internal/domain"
	"server/internal/ent"
)

type Service struct {
	client *ent.TagClient
}

func NewService(client *ent.Client) *Service {
	return &Service{client: client.Tag}
}

func (s Service) All(ctx context.Context, dto domain.AllTagsDTO) ([]*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ByID(ctx context.Context, dto domain.TagsByIDDTO) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create(ctx context.Context, dto domain.CreateTagDTO) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(ctx context.Context, dto domain.TagPutDTO) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, dto domain.DeleteTagDTO) error {
	//TODO implement me
	panic("implement me")
}
