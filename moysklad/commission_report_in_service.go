package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CommissionReportInService
// Сервис для работы с полученными отчётами комиссионера.
type CommissionReportInService interface {
	GetList(ctx context.Context, params *Params) (*List[CommissionReportIn], *resty.Response, error)
	Create(ctx context.Context, commissionReportIn *CommissionReportIn, params *Params) (*CommissionReportIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, commissionReportInList []*CommissionReportIn, params *Params) (*[]CommissionReportIn, *resty.Response, error)
	DeleteMany(ctx context.Context, commissionReportInList []*CommissionReportIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CommissionReportIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, commissionReportIn *CommissionReportIn, params *Params) (*CommissionReportIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportInPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*CommissionReportInPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *CommissionReportInPosition, params *Params) (*CommissionReportInPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *CommissionReportInPosition) (*CommissionReportInPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*CommissionReportInPosition) (*[]CommissionReportInPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CommissionReportIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	// Template(ctx context.Context) (*CommissionReportIn, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetReturnPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportInReturnPosition], *resty.Response, error)
	GetReturnPositionById(ctx context.Context, id, positionID *uuid.UUID, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error)
	CreateReturnPosition(ctx context.Context, id *uuid.UUID, position *CommissionReportInReturnPosition) (*CommissionReportInReturnPosition, *resty.Response, error)
	UpdateReturnPosition(ctx context.Context, id, positionID *uuid.UUID, position *CommissionReportInReturnPosition, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error)
	DeleteReturnPosition(ctx context.Context, id, positionID *uuid.UUID) (bool, *resty.Response, error)
}

type commissionReportInService struct {
	Endpoint
	endpointGetList[CommissionReportIn]
	endpointCreate[CommissionReportIn]
	endpointCreateUpdateMany[CommissionReportIn]
	endpointDeleteMany[CommissionReportIn]
	endpointDelete
	endpointGetById[CommissionReportIn]
	endpointUpdate[CommissionReportIn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[CommissionReportInPosition]
	endpointAttributes
	endpointSyncID[CommissionReportIn]
	endpointNamedFilter
	endpointPublication
	endpointRemove
}

func NewCommissionReportInService(client *Client) CommissionReportInService {
	e := NewEndpoint(client, "entity/commissionreportin")
	return &commissionReportInService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[CommissionReportIn]{e},
		endpointCreate:           endpointCreate[CommissionReportIn]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[CommissionReportIn]{e},
		endpointDeleteMany:       endpointDeleteMany[CommissionReportIn]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[CommissionReportIn]{e},
		endpointUpdate:           endpointUpdate[CommissionReportIn]{e},
		endpointMetadata:         endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:        endpointPositions[CommissionReportInPosition]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[CommissionReportIn]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		//endpointTemplate:         endpointTemplate[CommissionReportIn]{e},
		endpointPublication: endpointPublication{e},
		endpointRemove:      endpointRemove{e},
	}
}

// GetReturnPositions Получить позиции возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchit-pozicii-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) GetReturnPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportInReturnPosition], *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[MetaArray[CommissionReportInReturnPosition]](s.client, path).SetParams(params).Get(ctx)
}

// GetReturnPositionById Получить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchit-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) GetReturnPositionById(ctx context.Context, id, positionID *uuid.UUID, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionID)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).SetParams(params).Get(ctx)
}

// CreateReturnPosition Создать позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-sozdat-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) CreateReturnPosition(ctx context.Context, id *uuid.UUID, position *CommissionReportInReturnPosition) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).Post(ctx, position)
}

// UpdateReturnPosition Изменить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-izmenit-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) UpdateReturnPosition(ctx context.Context, id, positionID *uuid.UUID, position *CommissionReportInReturnPosition, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionID)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).SetParams(params).Put(ctx, position)
}

// DeleteReturnPosition Удалить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-udalit-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) DeleteReturnPosition(ctx context.Context, id, positionID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/positions/%s", id, positionID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
