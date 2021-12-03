package user

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/ent"
	userent "server/internal/ent/user"
	"server/internal/platform"
)

type Service struct {
	client    *ent.UserClient
	jwtConfig config.Jwt
}

func NewService(client *ent.Client, config config.Config) *Service {
	return &Service{
		client:    client.User,
		jwtConfig: config.Jwt,
	}
}

func (s Service) JWTToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": user.ID})
	return token.SignedString([]byte(s.jwtConfig.Secret))
}

func (s Service) DataFromJWT(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signing method")
		}
		return []byte(s.jwtConfig.Secret), nil
	})
	if err != nil {
		return 0, platform.WrapInternal(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, platform.WrapInternal(err)
	}

	return int(claims["id"].(float64)), nil
}

func (s Service) Signup(ctx context.Context, username, password string) (string, error) {
	panic("implement me")
}

func (s Service) Login(ctx context.Context, userID int) (string, error) {
	panic("implement me")
}

func (s Service) ByID(ctx context.Context, dto domain.GetUserDTO) (*domain.User, error) {
	usr, err := s.client.Query().
		Where(userent.ID(dto.ID)).
		Only(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	return domain.UserFromEnt(usr), nil
}
