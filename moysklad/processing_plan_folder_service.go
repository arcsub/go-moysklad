package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingPlanFolderService
// Сервис для работы с группами техкарт.
type ProcessingPlanFolderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingPlanFolder], *resty.Response, error)
	Create(ctx context.Context, processingPlanFolder *ProcessingPlanFolder, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingPlanFolder *ProcessingPlanFolder, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingPlanFolderService(client *Client) ProcessingPlanFolderService {
	e := NewEndpoint(client, "entity/processingplanfolder")
	return newMainService[ProcessingPlanFolder, any, MetadataAttributeSharedStates, any](e)
}
