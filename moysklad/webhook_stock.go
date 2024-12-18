package moysklad

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

// WebhookStock Вебхук на изменение остатков.
//
// Код сущности: webhookstock
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuk-na-izmenenie-ostatkow
type WebhookStock struct {
	AccountID         *string       `json:"accountId,omitempty"`         // ID учётной записи
	AuthorApplication *Meta         `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего вебхук на изменение остатков
	Enabled           *bool         `json:"enabled,omitempty"`           // Флажок состояния вебхука на изменение остатков (включен / отключен)
	StockType         *string       `json:"stockType,omitempty"`         // Тип остатков, изменение которых вызывает вебхук на изменение остатков
	ReportUrl         *string       `json:"reportUrl,omitempty"`         // URL на получения данных по изменившейся номенклатуре за указанный период
	ID                *string       `json:"id,omitempty"`                // ID вебхука на изменение остатков
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные вебхука на изменение остатков
	URL               *string       `json:"url,omitempty"`               // URL, по которому будет происходить обработка вебхука. Допустимая длина до 255 символов
	ReportType        WebhookReport `json:"reportType,omitempty"`        // Тип отчета остатков, к которым привязан вебхук на изменение остатков
}

// GetAccountID возвращает ID учётной записи.
func (webhookStock WebhookStock) GetAccountID() string {
	return Deref(webhookStock.AccountID)
}

// GetAuthorApplication возвращает Метаданные Приложения, создавшего вебхук на изменение остатков.
func (webhookStock WebhookStock) GetAuthorApplication() Meta {
	return Deref(webhookStock.AuthorApplication)
}

// GetEnabled возвращает флаг состояния вебхука на изменение остатков (включен / отключен).
func (webhookStock WebhookStock) GetEnabled() bool {
	return Deref(webhookStock.Enabled)
}

// GetStockType возвращает Тип остатков, изменение которых вызывает вебхук на изменение остатков.
func (webhookStock WebhookStock) GetStockType() string {
	return Deref(webhookStock.StockType)
}

// GetReportUrl возвращает URL на получения данных по изменившейся номенклатуре за указанный период.
func (webhookStock WebhookStock) GetReportUrl() string {
	return Deref(webhookStock.ReportUrl)
}

// GetID возвращает ID вебхука на изменение остатков.
func (webhookStock WebhookStock) GetID() string {
	return Deref(webhookStock.ID)
}

// GetMeta возвращает Метаданные вебхука на изменение остатков.
func (webhookStock WebhookStock) GetMeta() Meta {
	return Deref(webhookStock.Meta)
}

// GetURL возвращает URL, по которому будет происходить обработка вебхука.
func (webhookStock WebhookStock) GetURL() string {
	return Deref(webhookStock.URL)
}

// GetReportType возвращает Тип отчета остатков, к которым привязан вебхук на изменение остатков.
func (webhookStock WebhookStock) GetReportType() WebhookReport {
	return webhookStock.ReportType
}

// SetEnabled устанавливает флаг состояния вебхука на изменение остатков (включен / отключен).
func (webhookStock *WebhookStock) SetEnabled(enabled bool) *WebhookStock {
	webhookStock.Enabled = &enabled
	return webhookStock
}

// SetReportUrl устанавливает URL на получения данных по изменившейся номенклатуре за указанный период.
func (webhookStock *WebhookStock) SetReportUrl(reportUrl string) *WebhookStock {
	webhookStock.ReportUrl = &reportUrl
	return webhookStock
}

// SetMeta устанавливает Метаданные вебхука на изменение остатков.
func (webhookStock *WebhookStock) SetMeta(meta *Meta) *WebhookStock {
	webhookStock.Meta = meta
	return webhookStock
}

// SetURL устанавливает URL, по которому будет происходить обработка вебхука.
func (webhookStock *WebhookStock) SetURL(url string) *WebhookStock {
	webhookStock.URL = &url
	return webhookStock
}

// SetReportType устанавливает Тип отчета остатков, к которым привязан вебхук на изменение остатков.
func (webhookStock *WebhookStock) SetReportType(reportType WebhookReport) *WebhookStock {
	webhookStock.ReportType = reportType
	return webhookStock
}

// SetReportTypeAll устанавливает Тип отчета остатков, к которым привязан вебхук на изменение остатков в значение [WebhookReportAll].
func (webhookStock *WebhookStock) SetReportTypeAll() *WebhookStock {
	webhookStock.ReportType = WebhookReportAll
	return webhookStock
}

// SetReportTypeByStore устанавливает Тип отчета остатков, к которым привязан вебхук на изменение остатков в значение [WebhookReportByStore].
func (webhookStock *WebhookStock) SetReportTypeByStore() *WebhookStock {
	webhookStock.ReportType = WebhookReportByStore
	return webhookStock
}

// String реализует интерфейс [fmt.Stringer].
func (webhookStock WebhookStock) String() string {
	return Stringify(webhookStock)
}

// MetaType возвращает код сущности.
func (WebhookStock) MetaType() MetaType {
	return MetaTypeWebhookStock
}

// Update shortcut
func (webhookStock *WebhookStock) Update(ctx context.Context, client *Client, params ...func(*Params)) (*WebhookStock, *resty.Response, error) {
	return NewWebhookStockService(client).Update(ctx, webhookStock.GetID(), webhookStock, params...)
}

// Create shortcut
func (webhookStock *WebhookStock) Create(ctx context.Context, client *Client, params ...func(*Params)) (*WebhookStock, *resty.Response, error) {
	return NewWebhookStockService(client).Create(ctx, webhookStock, params...)
}

// Delete shortcut
func (webhookStock *WebhookStock) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewWebhookStockService(client).Delete(ctx, webhookStock)
}

// MarshalJSON реализует интерфейс [json.Marshaler].
func (webhookStock WebhookStock) MarshalJSON() ([]byte, error) {
	webhookStock.StockType = String("stock")
	return json.Marshal(webhookStock)
}

// WebhookReport Тип отчёта остатков, к которым привязан вебхук на изменение остатков.
//
// Возможные варианты:
//   - WebhookReportAll     – все
//   - WebhookReportByStore – по складам
type WebhookReport string

const (
	WebhookReportAll     WebhookReport = "all"     // все
	WebhookReportByStore WebhookReport = "bystore" // по складам
)

// WebhookStockService описывает методы сервиса для работы с вебхуками на изменение остатков.
type WebhookStockService interface {
	// GetList выполняет запрос на получение списка вебхуков на изменение остатков.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...func(*Params)) (*List[WebhookStock], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех вебхуков на изменение остатков в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...func(*Params)) (*Slice[WebhookStock], *resty.Response, error)

	// Create выполняет запрос на создание вебхука на изменение остатков.
	// Обязательные поля для заполнения:
	//	- reportType (Тип отчета остатков, к которым привязан вебхук на изменение остатков)
	//	- url (URL, по которому будет происходить обработка вебхука)
	// Принимает контекст, вебхук на изменение остатков и опционально объект параметров запроса Params.
	// Возвращает созданный вебхук на изменение остатков.
	Create(ctx context.Context, webhookStock *WebhookStock, params ...func(*Params)) (*WebhookStock, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение вебхуков на изменение остатков.
	// Изменяемые вебхуки на изменение остатков должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список вебхуков на изменение остатков и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых вебхуков на изменение остатков.
	CreateUpdateMany(ctx context.Context, webhookStockList Slice[WebhookStock], params ...func(*Params)) (*Slice[WebhookStock], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление вебхуков на изменение остатков.
	// Принимает контекст и множество вебхуков на изменение остатков.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*WebhookStock) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление вебхука на изменение остатков по ID.
	// Принимает контекст и ID вебхука на изменение остатков.
	// Возвращает «true» в случае успешного удаления вебхука на изменение остатков.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление вебхука на изменение остатков.
	// Принимает контекст и вебхук на изменение остатков.
	// Возвращает «true» в случае успешного удаления вебхука на изменение остатков.
	Delete(ctx context.Context, entity *WebhookStock) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного вебхука на изменение остатков по ID.
	// Принимает контекст, ID вебхука на изменение остатков и опционально объект параметров запроса Params.
	// Возвращает найденный вебхук на изменение остатков.
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*WebhookStock, *resty.Response, error)

	// Update выполняет запрос на изменение вебхука на изменение остатков.
	// Принимает контекст, вебхук на изменение остатков и опционально объект параметров запроса Params.
	// Возвращает изменённый вебхук на изменение остатков.
	Update(ctx context.Context, id string, webhookStock *WebhookStock, params ...func(*Params)) (*WebhookStock, *resty.Response, error)
}

const (
	EndpointWebhookStock = EndpointEntity + string(MetaTypeWebhookStock)
)

// NewWebhookStockService принимает [Client] и возвращает сервис для работы с вебхуками на изменение остатков.
func NewWebhookStockService(client *Client) WebhookStockService {
	return newMainService[WebhookStock, any, any, any](client, EndpointWebhookStock)
}
