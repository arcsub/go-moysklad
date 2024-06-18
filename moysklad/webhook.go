package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Webhook Вебхук.
// Ключевое слово: webhook
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki
type Webhook struct {
	AccountID         *uuid.UUID    `json:"accountId,omitempty"`         // ID учетной записи
	Action            WebhookAction `json:"action,omitempty"`            // Действие, которое отслеживается веб-хуком. Возможные значения: [CREATE, UPDATE, DELETE, PROCESSED]. Задать значение PROCESSED возможно только для асинхронных задач
	AuthorApplication *Application  `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего веб-хук
	DiffType          WebhookDiff   `json:"diffType,omitempty"`          // Режим отображения изменения сущности. Указывается только для действия UPDATE. Возможные значения: [NONE, FIELDS] (по умолчанию NONE)
	Enabled           *bool         `json:"enabled,omitempty"`           // Флажок состояние веб-хука (включен / отключен)
	EntityType        MetaType      `json:"entityType,omitempty"`        // Тип сущности, к которой привязан веб-хук
	ID                *uuid.UUID    `json:"id,omitempty"`                // ID Веб-хука
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные
	Method            WebhookMethod `json:"method,omitempty"`            // HTTP метод, с которым будет происходить запрос. Возможные значения: POST
	URL               *string       `json:"url,omitempty"`               // URL, по которому будет происходить запрос. Допустимая длина до 255 символов
	UpdatedFields     Slice[string] `json:"updatedFields,omitempty"`     // Поля сущности, измененные пользователем
}

func (webhook Webhook) GetAccountID() uuid.UUID {
	return Deref(webhook.AccountID)
}

func (webhook Webhook) GetAction() WebhookAction {
	return webhook.Action
}

func (webhook Webhook) GetAuthorApplication() Application {
	return Deref(webhook.AuthorApplication)
}

func (webhook Webhook) GetDiffType() WebhookDiff {
	return webhook.DiffType
}

func (webhook Webhook) GetEnabled() bool {
	return Deref(webhook.Enabled)
}

func (webhook Webhook) GetEntityType() MetaType {
	return webhook.EntityType
}

func (webhook Webhook) GetID() uuid.UUID {
	return Deref(webhook.ID)
}

func (webhook Webhook) GetMeta() Meta {
	return Deref(webhook.Meta)
}

func (webhook Webhook) GetMethod() WebhookMethod {
	return webhook.Method
}

func (webhook Webhook) GetURL() string {
	return Deref(webhook.URL)
}

func (webhook Webhook) GetUpdatedFields() Slice[string] {
	return webhook.UpdatedFields
}

func (webhook *Webhook) SetAction(action WebhookAction) *Webhook {
	webhook.Action = action
	return webhook
}

func (webhook *Webhook) SetDiffType(diffType WebhookDiff) *Webhook {
	webhook.DiffType = diffType
	return webhook
}

func (webhook *Webhook) SetEnabled(enabled bool) *Webhook {
	webhook.Enabled = &enabled
	return webhook
}

func (webhook *Webhook) SetEntityType(entityType MetaType) *Webhook {
	webhook.EntityType = entityType
	return webhook
}

func (webhook *Webhook) SetMeta(meta *Meta) *Webhook {
	webhook.Meta = meta
	return webhook
}

func (webhook *Webhook) SetMethod() *Webhook {
	webhook.Method = Post
	return webhook
}

func (webhook *Webhook) SetURL(url string) *Webhook {
	webhook.URL = &url
	return webhook
}

func (webhook Webhook) String() string {
	return Stringify(webhook)
}

func (webhook Webhook) MetaType() MetaType {
	return MetaTypeWebhook
}

type WebhookAction string

const (
	WebhookActionCreate    WebhookAction = "CREATE"
	WebhookActionUpdate    WebhookAction = "UPDATE"
	WebhookActionDelete    WebhookAction = "DELETE"
	WebhookActionProcessed WebhookAction = "PROCESSED"
)

type WebhookDiff string

const (
	WebhookDiffNone   WebhookDiff = "NONE"
	WebhookDiffFields WebhookDiff = "FIELDS"
)

type WebhookMethod string

const (
	Post WebhookMethod = "POST"
)

// WebhookService
// Сервис для работы с вебхуками.
type WebhookService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Webhook], *resty.Response, error)
	Create(ctx context.Context, webhook *Webhook, params ...*Params) (*Webhook, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, webhookList Slice[Webhook], params ...*Params) (*Slice[Webhook], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*Webhook) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Webhook, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, webhook *Webhook, params ...*Params) (*Webhook, *resty.Response, error)
}

func NewWebhookService(client *Client) WebhookService {
	e := NewEndpoint(client, "entity/webhook")
	return newMainService[Webhook, any, any, any](e)
}
