package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Application Серверное приложение.
// Ключевое слово: application
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-serwernye-prilozheniq
type Application struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	ID        *uuid.UUID `json:"id,omitempty"`        // ID сущности
	Name      *string    `json:"name,omitempty"`      // Наименование
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные
	AppUID    *uuid.UUID `json:"appUid,omitempty"`    // UID приложения
}

func (application Application) GetAccountID() uuid.UUID {
	return Deref(application.AccountID)
}

func (application Application) GetID() uuid.UUID {
	return Deref(application.ID)
}

func (application Application) GetName() string {
	return Deref(application.Name)
}

func (application Application) GetMeta() Meta {
	return Deref(application.Meta)
}

func (application Application) GetAppUID() uuid.UUID {
	return Deref(application.AppUID)
}

func (application Application) String() string {
	return Stringify(application)
}

func (application Application) MetaType() MetaType {
	return MetaTypeApplication
}

// ApplicationService
// Сервис для работы с приложениями.
type ApplicationService interface {
	// GetList выполняет запрос на получение списка сущностей установленных приложений.
	GetList(ctx context.Context, params *Params) (*List[Application], *resty.Response, error)

	// GetByID выполняет запрос на получение сущности установленного приложения.
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Application, *resty.Response, error)
}

func NewApplicationService(client *Client) ApplicationService {
	e := NewEndpoint(client, "entity/application")
	return newMainService[Application, any, any, any](e)
}
