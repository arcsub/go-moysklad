package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Token Токен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-autentifikaciq-poluchenie-nowogo-tokena
type Token struct {
	AccessToken string `json:"access_token"`
}

func (token Token) String() string {
	return Stringify(token)
}

// SecurityTokenService
// Сервис для получения нового токена
type SecurityTokenService interface {
	GetNewToken(ctx context.Context) (*Token, *resty.Response, error)
}

type securityTokenService struct {
	Endpoint
}

func NewSecurityTokenService(client *Client) SecurityTokenService {
	e := NewEndpoint(client, "security/token")
	return &securityTokenService{e}
}

// GetNewToken Запрос на получение нового токена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-autentifikaciq-poluchenie-nowogo-tokena
func (s *securityTokenService) GetNewToken(ctx context.Context) (*Token, *resty.Response, error) {
	return NewRequestBuilder[Token](s.client, s.uri).Post(ctx, nil)
}
