package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// WebhookStock Вебхук на изменение остатков.
// Ключевое слово: webhookstock
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuk-na-izmenenie-ostatkow
type WebhookStock struct {
	AccountID         *uuid.UUID    `json:"accountId,omitempty"`
	AuthorApplication *Meta         `json:"authorApplication,omitempty"`
	Enabled           *bool         `json:"enabled,omitempty"`
	StockType         *string       `json:"stockType,omitempty"`
	ReportUrl         *string       `json:"reportUrl,omitempty"`
	ID                *uuid.UUID    `json:"id,omitempty"`
	Meta              *Meta         `json:"meta,omitempty"`
	URL               *string       `json:"url,omitempty"`
	ReportType        WebhookReport `json:"reportType,omitempty"`
}

func (w WebhookStock) String() string {
	return Stringify(w)
}

func (w WebhookStock) MetaType() MetaType {
	return MetaTypeWebhookStock
}

// WebhookReport Тип отчета остатков, к которым привязан вебхук на изменение остатков
type WebhookReport string

const (
	WebhookReportAll     WebhookReport = "all"
	WebhookReportByStore WebhookReport = "bystore"
)

// WebhookStockService
// Сервис для работы с вебхуками на изменение остатков.
type WebhookStockService interface {
	GetList(ctx context.Context, params *Params) (*List[WebhookStock], *resty.Response, error)
	Create(ctx context.Context, webhookStock *WebhookStock, params *Params) (*WebhookStock, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, webhookStockList []*WebhookStock, params *Params) (*[]WebhookStock, *resty.Response, error)
	DeleteMany(ctx context.Context, webhookStockList []*WebhookStock) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*WebhookStock, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, webhookStock *WebhookStock, params *Params) (*WebhookStock, *resty.Response, error)
}

func NewWebhookStockService(client *Client) WebhookStockService {
	e := NewEndpoint(client, "entity/webhookstock")
	return newMainService[WebhookStock, any, any, any](e)
}
