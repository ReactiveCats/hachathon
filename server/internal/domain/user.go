package domain

import (
	"context"
	"server/internal/ent"
)

type User struct {
	ID int `json:"id" example:"1"`

	Username string `json:"username"`
} //@name User

func UserFromEnt(usr *ent.User) *User {
	if usr == nil {
		return nil
	}

	return &User{
		ID:       usr.ID,
		Username: usr.Username,
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

type GetUserDTO struct {
	ID int
}
