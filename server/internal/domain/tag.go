package domain

import (
	"context"
	"server/internal/ent"
)

type Tag struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Mult  float64 `json:"mult"`
}

func TagFromEnt(tag *ent.Tag) *Tag {
	if tag == nil {
		return nil
	}

	return &Tag{
		ID:    tag.ID,
		Title: tag.Title,
		Mult:  tag.Mult,
	}
}

func TagsFromEnt(slice []*ent.Tag) []*Tag {
	tags := make([]*Tag, len(slice))
	for i, e := range slice {
		tags[i] = TagFromEnt(e)
	}
	return tags
}

type TagService interface {
	All(ctx context.Context, dto AllTagsDTO) ([]*Tag, error)
	ByID(ctx context.Context, dto TagsByIDDTO) (*Tag, error)
	Create(ctx context.Context, dto CreateTagDTO) (*Tag, error)
	Update(ctx context.Context, dto TagPutDTO) (*Tag, error)
	Delete(ctx context.Context, dto DeleteTagDTO) error
}

type AllTagsDTO struct {
	UserID int `json:"-"`
}

type TagsByIDDTO struct {
	UserID int `json:"-"`
	TagID  int
}

type CreateTagDTO struct {
	UserID int     `json:"-"`
	Title  string  `json:"title"`
	Mult   float64 `json:"mult"`
}

type TagPutDTO struct {
	UserID int     `json:"-"`
	TagID  int     `json:"-"`
	Title  string  `json:"title"`
	Mult   float64 `json:"mult"`
}

type DeleteTagDTO struct {
	UserID int `json:"-"`
	TagID  int
}
