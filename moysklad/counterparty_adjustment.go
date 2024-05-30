package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CounterPartyAdjustment Корректировка баланса контрагента.
// Ключевое слово: counterpartyadjustment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-korrektirowka-balansa-kontragenta
type CounterPartyAdjustment struct {
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`    // ID учетной записи
	ID           *uuid.UUID    `json:"id,omitempty"`           // ID сущности
	Name         *string       `json:"name,omitempty"`         // Наименование
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
	Applicable   *bool         `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes   `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	Created      *Timestamp    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp    `json:"deleted,omitempty"`      // Момент последнего удаления Корректировки баланса контрагента
	Description  *string       `json:"description,omitempty"`  // Комментарий Корректировки баланса контрагента
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код Корректировки баланса контрагента
	Files        *Files        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group        `json:"group,omitempty"`        // Отдел сотрудника
	Moment       *Timestamp    `json:"moment,omitempty"`       // Дата документа
	Organization *Organization `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Printed      *bool         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool         `json:"shared,omitempty"`       // Общий доступ
	Sum          *Decimal      `json:"sum,omitempty"`          // Сумма Корректировки баланса контрагента в копейках
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления Корректировки баланса контрагента
}

func (c CounterPartyAdjustment) String() string {
	return Stringify(c)
}

func (c CounterPartyAdjustment) MetaType() MetaType {
	return MetaTypeCounterPartyAdjustment
}

// CounterPartyAdjustmentService
// Сервис для работы с корректировками баланса контрагента.
type CounterPartyAdjustmentService interface {
	GetList(ctx context.Context, params *Params) (*List[CounterPartyAdjustment], *resty.Response, error)
	Create(ctx context.Context, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, counterPartyAdjustmentList []*CounterPartyAdjustment, params *Params) (*[]CounterPartyAdjustment, *resty.Response, error)
	DeleteMany(ctx context.Context, counterPartyAdjustmentList []*CounterPartyAdjustment) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCounterPartyAdjustmentService(client *Client) CounterPartyAdjustmentService {
	e := NewEndpoint(client, "entity/counterpartyadjustment")
	return newMainService[CounterPartyAdjustment, any, MetadataAttributeSharedStates, any](e)
}
