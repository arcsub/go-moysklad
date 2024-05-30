package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// AssortmentService
// Сервис для работы с ассортиментом.
type AssortmentService interface {
	Get(ctx context.Context, params *Params) (*AssortmentResult, *resty.Response, error)
	GetAsync(ctx context.Context) (AsyncResultService[AssortmentResult], *resty.Response, error)
	DeleteMany(ctx context.Context, entities []*AssortmentPosition) (*DeleteManyResponse, *resty.Response, error)
	GetSettings(ctx context.Context) (*AssortmentSettings, *resty.Response, error)
	UpdateSettings(ctx context.Context, settings *AssortmentSettings) (*AssortmentSettings, *resty.Response, error)
	GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)
	GetEmbeddedTemplateByID(ctx context.Context, id *uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)
	GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)
	GetCustomTemplateByID(ctx context.Context, id *uuid.UUID) (*CustomTemplate, *resty.Response, error)
}

type assortmentService struct {
	endpointGetOne[AssortmentResult]
	endpointGetOneAsync[AssortmentResult]
	endpointDeleteMany[AssortmentPosition]
	endpointSettings[AssortmentSettings]
	endpointPrintTemplates
}

func NewAssortmentService(client *Client) AssortmentService {
	e := NewEndpoint(client, "entity/assortment")
	return &assortmentService{
		endpointGetOne:         endpointGetOne[AssortmentResult]{e},
		endpointGetOneAsync:    endpointGetOneAsync[AssortmentResult]{e},
		endpointDeleteMany:     endpointDeleteMany[AssortmentPosition]{e},
		endpointSettings:       endpointSettings[AssortmentSettings]{e},
		endpointPrintTemplates: endpointPrintTemplates{e},
	}
}
