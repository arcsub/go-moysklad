package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Token Токен.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-autentifikaciq-poluchenie-nowogo-tokena
type Token struct {
	AccessToken string `json:"access_token"`
}

// String реализует интерфейс [fmt.Stringer].
func (token Token) String() string {
	return Stringify(token)
}

// SecurityTokenService описывает метод сервиса для получения нового токена.
type SecurityTokenService interface {
	// GetNewToken выполняет запрос на получение нового токена.
	// Принимает контекст.
	// Возвращает новый токен.
	GetNewToken(ctx context.Context) (*Token, *resty.Response, error)
}

const (
	EndpointToken = EndpointSecurity + string(MetaTypeToken)
)

type securityTokenService struct {
	Endpoint
}

func (service *securityTokenService) GetNewToken(ctx context.Context) (*Token, *resty.Response, error) {
	return NewRequestBuilder[Token](service.client, service.uri).Post(ctx, nil)
}

// NewSecurityTokenService принимает [Client] и возвращает сервис для получения нового токена.
func NewSecurityTokenService(client *Client) SecurityTokenService {
	return &securityTokenService{NewEndpoint(client, EndpointToken)}
}
