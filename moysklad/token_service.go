package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// SecurityTokenService
// Сервис для получения нового токена
type SecurityTokenService struct {
	Endpoint
}

func NewSecurityTokenService(client *Client) *SecurityTokenService {
	e := NewEndpoint(client, "security/token")
	return &SecurityTokenService{e}
}

// GetNewToken Запрос на получение нового токена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-autentifikaciq-poluchenie-nowogo-tokena
func (s *SecurityTokenService) GetNewToken(ctx context.Context) (*Token, *resty.Response, error) {
	return NewRequestBuilder[Token](s.client, s.uri).Post(ctx, nil)
}
