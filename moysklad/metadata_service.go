package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// MetadataService
// Сервис для работы с метаданными.
type MetadataService interface {
	Get(ctx context.Context, params *Params) (*Metadata, *resty.Response, error)
}

func NewMetadataService(client *Client) MetadataService {
	e := NewEndpoint(client, "entity/metadata")
	return newMainService[Metadata, any, any, any](e)
}
