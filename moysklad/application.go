package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Application Серверное приложение.
//
// Код сущности: application
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-serwernye-prilozheniq
type Application struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учётной записи
	ID        *uuid.UUID `json:"id,omitempty"`        // ID Серверного приложения
	Name      *string    `json:"name,omitempty"`      // Наименование Серверного приложения
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные Серверного приложения
	AppUID    *uuid.UUID `json:"appUid,omitempty"`    // UID Серверного приложения
}

// GetAccountID возвращает ID учётной записи.
func (application Application) GetAccountID() uuid.UUID {
	return Deref(application.AccountID)
}

// GetID возвращает ID Серверного приложения.
func (application Application) GetID() uuid.UUID {
	return Deref(application.ID)
}

// GetName возвращает Наименование Серверного приложения.
func (application Application) GetName() string {
	return Deref(application.Name)
}

// GetMeta возвращает Метаданные Серверного приложения.
func (application Application) GetMeta() Meta {
	return Deref(application.Meta)
}

// GetAppUID возвращает UID Серверного приложения.
func (application Application) GetAppUID() uuid.UUID {
	return Deref(application.AppUID)
}

// String реализует интерфейс [fmt.Stringer].
func (application Application) String() string {
	return Stringify(application)
}

// MetaType возвращает код сущности.
func (Application) MetaType() MetaType {
	return MetaTypeApplication
}

// ApplicationService методы сервиса для работы с серверными приложениями.
type ApplicationService interface {
	// GetList выполняет запрос на получение списка сущностей установленных приложений.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Application], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех установленных приложений в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (Slice[Application], *resty.Response, error)

	// GetByID выполняет запрос на получение сущности установленного приложения.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект Application.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Application, *resty.Response, error)
}

const (
	EndpointApplication = EndpointEntity + string(MetaTypeApplication)
)

// NewApplicationService принимает [Client] и возвращает сервис для работы с серверными приложениями.
func NewApplicationService(client *Client) ApplicationService {
	return newMainService[Application, any, any, any](client, EndpointApplication)
}
