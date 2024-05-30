package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingProcessService
// Сервис для работы с тех процессами.
type ProcessingProcessService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingProcess], *resty.Response, error)
	Create(ctx context.Context, processingProcess *ProcessingProcess, params *Params) (*ProcessingProcess, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingProcessList []*ProcessingProcess, params *Params) (*[]ProcessingProcess, *resty.Response, error)
	DeleteMany(ctx context.Context, processingProcessList []*ProcessingProcess) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingProcess, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingProcess *ProcessingProcess, params *Params) (*ProcessingProcess, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingProcessPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*ProcessingProcessPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *ProcessingProcessPosition, params *Params) (*ProcessingProcessPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *ProcessingProcessPosition) (*ProcessingProcessPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*ProcessingProcessPosition) (*[]ProcessingProcessPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingProcessService(client *Client) ProcessingProcessService {
	e := NewEndpoint(client, "entity/processingprocess")
	return newMainService[ProcessingProcess, ProcessingProcessPosition, any, any](e)
}
