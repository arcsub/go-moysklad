package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingStageService
// Сервис для работы с этапами производства.
type ProcessingStageService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingStage], *resty.Response, error)
	Create(ctx context.Context, processingStage *ProcessingStage, params *Params) (*ProcessingStage, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingStageList []*ProcessingStage, params *Params) (*[]ProcessingStage, *resty.Response, error)
	DeleteMany(ctx context.Context, processingStageList []*ProcessingStage) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingStage, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingStage *ProcessingStage, params *Params) (*ProcessingStage, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingStageService(client *Client) ProcessingStageService {
	e := NewEndpoint(client, "entity/processingstage")
	return newMainService[ProcessingStage, any, any, any](e)
}
