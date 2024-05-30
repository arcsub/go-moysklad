package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailDrawerCashOut Выплата денег.
// Ключевое слово: retaildrawercashout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vyplata-deneg
type RetailDrawerCashOut struct {
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`    // ID учетной записи
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
	Applicable   *bool         `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes   `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей
	Code         *string       `json:"code,omitempty"`         // Код Выплаты денег
	Created      *Timestamp    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp    `json:"deleted,omitempty"`      // Момент последнего удаления Выплаты денег
	Description  *string       `json:"description,omitempty"`  // Комментарий Выплаты денег
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код Выплаты денег
	Files        *Files        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group        `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID    `json:"id,omitempty"`           // ID сущности
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp    `json:"moment,omitempty"`       // Дата документа
	Name         *string       `json:"name,omitempty"`         // Наименование
	Organization *Organization `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Printed      *bool         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool         `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *Rate         `json:"rate,omitempty"`         // Валюта
	Shared       *bool         `json:"shared,omitempty"`       // Общий доступ
	State        *State        `json:"state,omitempty"`        // Метаданные статуса Выплаты денег
	Sum          *Decimal      `json:"sum,omitempty"`          // Сумма Выплаты денег установленной валюте
	SyncID       *uuid.UUID    `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления
	RetailShift  *RetailShift  `json:"retailShift,omitempty"`  // Ссылка на розничную смену, в рамках которой было выполнено Внесение денег в формате Метаданных
}

func (r RetailDrawerCashOut) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailDrawerCashOut) GetMeta() *Meta {
	return r.Meta
}

func (r RetailDrawerCashOut) MetaType() MetaType {
	return MetaTypeRetailDrawerCashOut
}

// RetailDrawerCashOutService
// Сервис для работы с выплатами денег.
type RetailDrawerCashOutService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailDrawerCashOut], *resty.Response, error)
	Create(ctx context.Context, retailDrawerCashOut *RetailDrawerCashOut, params *Params) (*RetailDrawerCashOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDrawerCashOutList []*RetailDrawerCashOut, params *Params) (*[]RetailDrawerCashOut, *resty.Response, error)
	DeleteMany(ctx context.Context, retailDrawerCashOutList []*RetailDrawerCashOut) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailDrawerCashOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, retailDrawerCashOut *RetailDrawerCashOut, params *Params) (*RetailDrawerCashOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	//endpointTemplate[RetailDrawerCashOut]
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*RetailDrawerCashOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailDrawerCashOutService(client *Client) RetailDrawerCashOutService {
	e := NewEndpoint(client, "entity/retaildrawercashout")
	return newMainService[RetailDrawerCashOut, any, MetadataAttributeSharedStates, any](e)
}
