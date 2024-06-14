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
	AccountID         *uuid.UUID       `json:"accountId,omitempty"`
	AuthorApplication *Meta            `json:"authorApplication,omitempty"`
	Enabled           *bool            `json:"enabled,omitempty"`
	StockType         WebhookStockType `json:"stockType,omitempty"`
	ReportUrl         *string          `json:"reportUrl,omitempty"`
	ID                *uuid.UUID       `json:"id,omitempty"`
	Meta              *Meta            `json:"meta,omitempty"`
	URL               *string          `json:"url,omitempty"`
	ReportType        WebhookReport    `json:"reportType,omitempty"`
}

func (webhookStock WebhookStock) GetAccountID() uuid.UUID {
	return Deref(webhookStock.AccountID)
}

func (webhookStock WebhookStock) GetAuthorApplication() Meta {
	return Deref(webhookStock.AuthorApplication)
}

func (webhookStock WebhookStock) GetEnabled() bool {
	return Deref(webhookStock.Enabled)
}

func (webhookStock WebhookStock) GetStockType() WebhookStockType {
	return webhookStock.StockType
}

func (webhookStock WebhookStock) GetReportUrl() string {
	return Deref(webhookStock.ReportUrl)
}

func (webhookStock WebhookStock) GetID() uuid.UUID {
	return Deref(webhookStock.ID)
}

func (webhookStock WebhookStock) GetMeta() Meta {
	return Deref(webhookStock.Meta)
}

func (webhookStock WebhookStock) GetURL() string {
	return Deref(webhookStock.URL)
}

func (webhookStock WebhookStock) GetReportType() WebhookReport {
	return webhookStock.ReportType
}

func (webhookStock *WebhookStock) SetEnabled(enabled bool) *WebhookStock {
	webhookStock.Enabled = &enabled
	return webhookStock
}

func (webhookStock *WebhookStock) SetStockType() *WebhookStock {
	webhookStock.StockType = WebhookTypeStock
	return webhookStock
}

func (webhookStock *WebhookStock) SetReportUrl(reportUrl string) *WebhookStock {
	webhookStock.ReportUrl = &reportUrl
	return webhookStock
}

func (webhookStock *WebhookStock) SetMeta(meta *Meta) *WebhookStock {
	webhookStock.Meta = meta
	return webhookStock
}

func (webhookStock *WebhookStock) SetURL(url string) *WebhookStock {
	webhookStock.URL = &url
	return webhookStock
}

func (webhookStock *WebhookStock) SetReportType(reportType WebhookReport) *WebhookStock {
	webhookStock.ReportType = reportType
	return webhookStock
}

func (webhookStock WebhookStock) String() string {
	return Stringify(webhookStock)
}

func (webhookStock WebhookStock) MetaType() MetaType {
	return MetaTypeWebhookStock
}

// WebhookReport Тип отчета остатков, к которым привязан вебхук на изменение остатков
type WebhookReport string

const (
	WebhookReportAll     WebhookReport = "all"
	WebhookReportByStore WebhookReport = "bystore"
)

type WebhookStockType string

const (
	WebhookTypeStock WebhookStockType = "stock"
)

// WebhookStockService
// Сервис для работы с вебхуками на изменение остатков.
type WebhookStockService interface {
	GetList(ctx context.Context, params ...*Params) (*List[WebhookStock], *resty.Response, error)
	Create(ctx context.Context, webhookStock *WebhookStock, params ...*Params) (*WebhookStock, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, webhookStockList Slice[WebhookStock], params ...*Params) (*Slice[WebhookStock], *resty.Response, error)
	DeleteMany(ctx context.Context, webhookStockList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*WebhookStock, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, webhookStock *WebhookStock, params ...*Params) (*WebhookStock, *resty.Response, error)
}

func NewWebhookStockService(client *Client) WebhookStockService {
	e := NewEndpoint(client, "entity/webhookstock")
	return newMainService[WebhookStock, any, any, any](e)
}
