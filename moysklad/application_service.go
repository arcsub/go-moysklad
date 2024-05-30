package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ApplicationService
// Сервис для работы с приложениями.
type ApplicationService interface {
	// GetList выполняет запрос на получение списка сущностей установленных приложений.
	GetList(ctx context.Context, params *Params) (*List[Application], *resty.Response, error)
	// GetByID выполняет запрос на получение сущности установленного приложения.
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Application, *resty.Response, error)
}

func NewApplicationService(client *Client) ApplicationService {
	e := NewEndpoint(client, "entity/application")
	return newMainService[Application, any, any, any](e)
}
