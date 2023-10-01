package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// AuditService
// Сервис для работы с аудитом.
type AuditService struct {
	Endpoint
}

func NewAuditService(client *Client) *AuditService {
	e := NewEndpoint(client, "audit")
	return &AuditService{e}
}

// GetContexts Получить Контексты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-kontexty
func (s *AuditService) GetContexts(ctx context.Context, params *Params) (*List[Audit], *Response, error) {
	return NewRequestBuilder[List[Audit]](s.Endpoint, ctx).WithParams(params).Get()
}

// GetEvents Получить События по Контексту.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-sobytiq-po-kontextu
func (s *AuditService) GetEvents(ctx context.Context, id *uuid.UUID) (*List[AuditEvent], *Response, error) {
	path := fmt.Sprintf("%s/events", id)
	return NewRequestBuilder[List[AuditEvent]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetFilters Получить Фильтры.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-fil-try
func (s *AuditService) GetFilters(ctx context.Context) (*AuditFilters, *Response, error) {
	path := "metadata/filters"
	return NewRequestBuilder[AuditFilters](s.Endpoint, ctx).WithPath(path).Get()
}
