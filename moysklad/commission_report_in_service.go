package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CommissionReportInService
// Сервис для работы с полученными отчётами комиссионера.
type CommissionReportInService struct {
	Endpoint
	endpointGetList[CommissionReportIn]
	endpointCreate[CommissionReportIn]
	endpointCreateUpdateDeleteMany[CommissionReportIn]
	endpointDelete
	endpointGetById[CommissionReportIn]
	endpointUpdate[CommissionReportIn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[CommissionReportInPosition]
	endpointAttributes
	endpointSyncID[CommissionReportIn]
	endpointNamedFilter
	endpointTemplate[CommissionReportIn]
	endpointPublication
	endpointRemove
}

func NewCommissionReportInService(client *Client) *CommissionReportInService {
	e := NewEndpoint(client, "entity/commissionreportin")
	return &CommissionReportInService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[CommissionReportIn]{e},
		endpointCreate:                 endpointCreate[CommissionReportIn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[CommissionReportIn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[CommissionReportIn]{e},
		endpointUpdate:                 endpointUpdate[CommissionReportIn]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[CommissionReportInPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointSyncID:                 endpointSyncID[CommissionReportIn]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointTemplate:               endpointTemplate[CommissionReportIn]{e},
		endpointPublication:            endpointPublication{e},
		endpointRemove:                 endpointRemove{e},
	}
}

// GetReturnPositions Получить позиции возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchit-pozicii-wozwrata-na-sklad-komissionera
func (s *CommissionReportInService) GetReturnPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportInReturnPosition], *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[MetaArray[CommissionReportInReturnPosition]](s.client, path).SetParams(params).Get(ctx)
}

// GetReturnPositionById Получить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchit-poziciu-wozwrata-na-sklad-komissionera
func (s *CommissionReportInService) GetReturnPositionById(ctx context.Context, id, positionId uuid.UUID, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionId)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).SetParams(params).Get(ctx)
}

// CreateReturnPosition Создать позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-sozdat-poziciu-wozwrata-na-sklad-komissionera
func (s *CommissionReportInService) CreateReturnPosition(ctx context.Context, id *uuid.UUID, position *CommissionReportInReturnPosition) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).Post(ctx, position)
}

// UpdateReturnPosition Изменить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-izmenit-poziciu-wozwrata-na-sklad-komissionera
func (s *CommissionReportInService) UpdateReturnPosition(ctx context.Context, id, positionId uuid.UUID, position *CommissionReportInReturnPosition, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionId)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).SetParams(params).Put(ctx, position)
}

// DeleteReturnPosition Удалить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-udalit-poziciu-wozwrata-na-sklad-komissionera
func (s *CommissionReportInService) DeleteReturnPosition(ctx context.Context, id, positionId uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/positions/%s", id, positionId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
