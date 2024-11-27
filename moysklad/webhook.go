package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

// Webhook Вебхук.
//
// Код сущности: webhook
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki
type Webhook struct {
	AccountID         *string       `json:"accountId,omitempty"`         // ID учётной записи
	Action            WebhookAction `json:"action,omitempty"`            // Действие, которое отслеживается веб-хуком. Задать значение PROCESSED возможно только для асинхронных задач
	AuthorApplication *Application  `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего веб-хук
	DiffType          WebhookDiff   `json:"diffType,omitempty"`          // Режим отображения изменения сущности. Указывается только для действия UPDATE
	Enabled           *bool         `json:"enabled,omitempty"`           // Флажок состояние веб-хука (включен / отключен)
	EntityType        MetaType      `json:"entityType,omitempty"`        // Тип сущности, к которой привязан веб-хук
	ID                *string       `json:"id,omitempty"`                // ID Веб-хука
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные веб-хука
	Method            *string       `json:"method,omitempty"`            // HTTP метод, с которым будет происходить запрос
	URL               *string       `json:"url,omitempty"`               // URL, по которому будет происходить запрос. Допустимая длина до 255 символов
	UpdatedFields     Slice[string] `json:"updatedFields,omitempty"`     // Поля сущности, измененные пользователем
}

// GetAccountID возвращает ID учётной записи.
func (webhook Webhook) GetAccountID() string {
	return Deref(webhook.AccountID)
}

// GetAction возвращает Действие, которое отслеживается веб-хуком.
func (webhook Webhook) GetAction() WebhookAction {
	return webhook.Action
}

// GetAuthorApplication возвращает Метаданные Приложения, создавшего веб-хук.
func (webhook Webhook) GetAuthorApplication() Application {
	return Deref(webhook.AuthorApplication)
}

// GetDiffType возвращает Режим отображения изменения сущности.
//
// Указывается только для действия UPDATE.
func (webhook Webhook) GetDiffType() WebhookDiff {
	return webhook.DiffType
}

// GetEnabled возвращает Флажок состояние веб-хука (включен / отключен).
func (webhook Webhook) GetEnabled() bool {
	return Deref(webhook.Enabled)
}

// GetEntityType возвращает Тип сущности, к которой привязан веб-хук.
func (webhook Webhook) GetEntityType() MetaType {
	return webhook.EntityType
}

// GetID возвращает ID Веб-хука.
func (webhook Webhook) GetID() string {
	return Deref(webhook.ID)
}

// GetMeta возвращает Метаданные веб-хука.
func (webhook Webhook) GetMeta() Meta {
	return Deref(webhook.Meta)
}

// GetMethod возвращает HTTP метод, с которым будет происходить запрос.
func (webhook Webhook) GetMethod() string {
	return Deref(webhook.Method)
}

// GetURL возвращает URL, по которому будет происходить запрос.
func (webhook Webhook) GetURL() string {
	return Deref(webhook.URL)
}

// GetUpdatedFields возвращает Поля сущности, измененные пользователем.
func (webhook Webhook) GetUpdatedFields() Slice[string] {
	return webhook.UpdatedFields
}

// SetAction устанавливает Действие, которое отслеживается веб-хуком.
func (webhook *Webhook) SetAction(action WebhookAction) *Webhook {
	webhook.Action = action
	return webhook
}

// SetActionCreate устанавливает Действие, которое отслеживается веб-хуком в значение [WebhookActionCreate].
func (webhook *Webhook) SetActionCreate() *Webhook {
	webhook.Action = WebhookActionCreate
	return webhook
}

// SetActionUpdate устанавливает Действие, которое отслеживается веб-хуком в значение [WebhookActionUpdate].
func (webhook *Webhook) SetActionUpdate() *Webhook {
	webhook.Action = WebhookActionUpdate
	return webhook
}

// SetActionDelete устанавливает Действие, которое отслеживается веб-хуком в значение [WebhookActionDelete].
func (webhook *Webhook) SetActionDelete() *Webhook {
	webhook.Action = WebhookActionDelete
	return webhook
}

// SetDiffType устанавливает Режим отображения изменения сущности.
func (webhook *Webhook) SetDiffType(diffType WebhookDiff) *Webhook {
	webhook.DiffType = diffType
	return webhook
}

// SetDiffTypeDefault устанавливает Режим отображения изменения сущности в значение [WebhookDiffNone].
func (webhook *Webhook) SetDiffTypeDefault() *Webhook {
	webhook.DiffType = WebhookDiffNone
	return webhook
}

// SetDiffTypeFields устанавливает Режим отображения изменения сущности в значение [WebhookDiffFields].
func (webhook *Webhook) SetDiffTypeFields() *Webhook {
	webhook.DiffType = WebhookDiffFields
	return webhook
}

// SetEnabled устанавливает Флажок состояние веб-хука (включен / отключен).
func (webhook *Webhook) SetEnabled(enabled bool) *Webhook {
	webhook.Enabled = &enabled
	return webhook
}

// SetEntityType устанавливает Тип сущности, к которой привязан веб-хук.
func (webhook *Webhook) SetEntityType(entityType MetaType) *Webhook {
	webhook.EntityType = entityType
	return webhook
}

// SetMeta устанавливает Метаданные веб-хука.
func (webhook *Webhook) SetMeta(meta *Meta) *Webhook {
	webhook.Meta = meta
	return webhook
}

// SetURL устанавливает URL, по которому будет происходить запрос.
//
// Допустимая длина до 255 символов.
func (webhook *Webhook) SetURL(url string) *Webhook {
	webhook.URL = &url
	return webhook
}

// String реализует интерфейс [fmt.Stringer].
func (webhook Webhook) String() string {
	return Stringify(webhook)
}

// MetaType возвращает код сущности.
func (Webhook) MetaType() MetaType {
	return MetaTypeWebhook
}

// Update shortcut
func (webhook *Webhook) Update(ctx context.Context, client *Client, params ...*Params) (*Webhook, *resty.Response, error) {
	return NewWebhookService(client).Update(ctx, webhook.GetID(), webhook, params...)
}

// Create shortcut
func (webhook *Webhook) Create(ctx context.Context, client *Client, params ...*Params) (*Webhook, *resty.Response, error) {
	return NewWebhookService(client).Create(ctx, webhook, params...)
}

// Delete shortcut
func (webhook *Webhook) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewWebhookService(client).Delete(ctx, webhook)
}

// MarshalJSON реализует интерфейс [json.Marshaler].
func (webhook Webhook) MarshalJSON() ([]byte, error) {
	webhook.Method = String("POST")
	return json.Marshal(webhook)
}

// WebhookAction Действие, которое отслеживается веб-хуком.
//
// Задать значение [WebhookActionProcessed] возможно только для асинхронных задач.
//
// Возможные значения:
//   - WebhookActionCreate    – создание
//   - WebhookActionUpdate    – изменение
//   - WebhookActionDelete    – удаление
//   - WebhookActionProcessed – значение только для асинхронных задач
type WebhookAction string

const (
	WebhookActionCreate    WebhookAction = "CREATE"    // создание
	WebhookActionUpdate    WebhookAction = "UPDATE"    // изменение
	WebhookActionDelete    WebhookAction = "DELETE"    // удаление
	WebhookActionProcessed WebhookAction = "PROCESSED" // значение для асинхронных задач
)

// WebhookDiff Режим отображения изменения сущности.
//
// Указывается только для действия UPDATE.
//
// По умолчанию [WebhookDiffNone].
//
// Возможные значения:
//   - WebhookDiffNone   – NONE
//   - WebhookDiffFields – FIELDS
type WebhookDiff string

const (
	WebhookDiffNone   WebhookDiff = "NONE"
	WebhookDiffFields WebhookDiff = "FIELDS"
)

// WebhookService описывает методы сервиса для работы с вебхуками.
type WebhookService interface {
	// GetList выполняет запрос на получение списка вебхуков.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Webhook], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех вебхуков в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Webhook], *resty.Response, error)

	// Create выполняет запрос на создание вебхука.
	// Обязательные поля для заполнения:
	//	- entityType (Тип сущности, к которой привязан вебхук)
	//	- action (Действие, которое отслеживается вебхуком)
	//	- url (URL, по которому будет происходить запрос)
	// Принимает контекст, вебхук и опционально объект параметров запроса Params.
	// Возвращает созданный вебхук.
	Create(ctx context.Context, webhook *Webhook, params ...*Params) (*Webhook, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение вебхуков.
	// Изменяемые вебхуки должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список вебхуков и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых вебхуков.
	CreateUpdateMany(ctx context.Context, webhookList Slice[Webhook], params ...*Params) (*Slice[Webhook], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление вебхуков.
	// Принимает контекст и множество вебхуков.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Webhook) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление вебхука по ID.
	// Принимает контекст и ID вебхука.
	// Возвращает «true» в случае успешного удаления вебхука.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление вебхука.
	// Принимает контекст и вебхук.
	// Возвращает «true» в случае успешного удаления вебхука.
	Delete(ctx context.Context, entity *Webhook) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного вебхука по ID.
	// Принимает контекст, ID вебхука и опционально объект параметров запроса Params.
	// Возвращает найденный вебхук.
	GetByID(ctx context.Context, id string, params ...*Params) (*Webhook, *resty.Response, error)

	// Update выполняет запрос на изменение вебхука.
	// Принимает контекст, вебхук и опционально объект параметров запроса Params.
	// Возвращает изменённый вебхук.
	Update(ctx context.Context, id string, webhook *Webhook, params ...*Params) (*Webhook, *resty.Response, error)
}

const (
	EndpointWebhook = EndpointEntity + string(MetaTypeWebhook)
)

// NewWebhookService принимает [Client] и возвращает сервис для работы с вебхуками.
func NewWebhookService(client *Client) WebhookService {
	return newMainService[Webhook, any, any, any](client, EndpointWebhook)
}
