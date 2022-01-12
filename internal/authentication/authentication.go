package authentication

import (
	"github.com/ekinbulut-yemeksepeti/auth-api/internal/token"
)

type Service struct {
	// you can depend on a database instance
	token *token.Service
}

type Authentication struct {
	Username string
	Password string
}

type AuthenticationService interface {
	CreateJWTToken(a *Authentication) (map[string]string, error)
}

func NewService(token *token.Service) *Service {
	return &Service{
		token: token,
	}
}

func (s *Service) CreateJWTToken(a *Authentication) (map[string]string, error) {


	// TODO : check if user exists and valid password
	// TODO : get claims from database of the user

	token, err := s.token.CreateJWTToken(a.Username)
	if err != nil {
		return nil, err
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	return tokens, nil
}
