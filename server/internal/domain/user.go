package domain

import (
	"context"
	"server/internal/ent"
)

type User struct {
	ID int `json:"id" example:"1"`

	Username string  `json:"username"`
	Tasks    []*Task `json:"-"`
	Tags     []*Tag  `json:"tags"`
} //@name User

func UserFromEnt(usr *ent.User) *User {
	if usr == nil {
		return nil
	}

	return &User{
		ID:       usr.ID,
		Username: usr.Username,
		Tasks:    TasksFromEnt(usr.Edges.Tasks),
		Tags:     TagsFromEnt(usr.Edges.Tags),
	}
}

func UsersFromEnt(slice []*ent.User) []*User {
	users := make([]*User, len(slice))
	for i, e := range slice {
		users[i] = UserFromEnt(e)
	}
	return users
}

type UserService interface {
	JWTToken(user *User) (string, error)
	DataFromJWT(tokenStr string) (int, error)
	Signup(ctx context.Context, username string) (string, error)
	Login(ctx context.Context, username string) (string, error)

	ByID(ctx context.Context, dto GetUserDTO) (*User, error)
}

// maybe it's better to replace such little DTOs with just arguments

type AccessToken struct {
	AccessToken string `json:"accessToken"`
}

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
}

type SignupDTO struct {
	Username string `json:"username" binding:"required"`
}

type GetUserDTO struct {
	ID int
}
