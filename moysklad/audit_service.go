package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// AuditService
// Сервис для работы с аудитом.
type AuditService interface {
	GetContexts(ctx context.Context, params *Params) (*List[Audit], *resty.Response, error)
	GetEvents(ctx context.Context, id *uuid.UUID) (*List[AuditEvent], *resty.Response, error)
	GetFilters(ctx context.Context) (*AuditFilters, *resty.Response, error)
}
type auditService struct {
	Endpoint
}

func NewAuditService(client *Client) AuditService {
	e := NewEndpoint(client, "audit")
	return &auditService{e}
}

// GetContexts Получить Контексты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-kontexty
func (s *auditService) GetContexts(ctx context.Context, params *Params) (*List[Audit], *resty.Response, error) {
	return NewRequestBuilder[List[Audit]](s.client, s.uri).SetParams(params).Get(ctx)
}

// GetEvents Получить События по Контексту.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-sobytiq-po-kontextu
func (s *auditService) GetEvents(ctx context.Context, id *uuid.UUID) (*List[AuditEvent], *resty.Response, error) {
	path := fmt.Sprintf("audit/%s/events", id)
	return NewRequestBuilder[List[AuditEvent]](s.client, path).Get(ctx)
}

// GetFilters Получить Фильтры.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-fil-try
func (s *auditService) GetFilters(ctx context.Context) (*AuditFilters, *resty.Response, error) {
	path := "audit/metadata/filters"
	return NewRequestBuilder[AuditFilters](s.client, path).Get(ctx)
}
