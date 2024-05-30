package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureIn Счет-фактура полученный
// Ключевое слово: facturein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-poluchennyj
type FactureIn struct {
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`      // ID учетной записи
	Agent          *Counterparty `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool         `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     *Attributes   `json:"attributes,omitempty"`     // Коллекция метаданных доп. полей объекта
	Code           *string       `json:"code,omitempty"`           // Код выданного Счета-фактуры полученного
	Contract       *Contract     `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp    `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp    `json:"deleted,omitempty"`        // Момент последнего удаления Счета-фактуры полученного
	Description    *string       `json:"description,omitempty"`    // Комментарий выданного Счета-фактуры полученного
	ExternalCode   *string       `json:"externalCode,omitempty"`   // Внешний код выданного Счета-фактуры полученного
	Files          *Files        `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group        `json:"group,omitempty"`          // Отдел сотрудника
	ID             *uuid.UUID    `json:"id,omitempty"`             // ID сущности
	Meta           *Meta         `json:"meta,omitempty"`           // Метаданные
	Moment         *Timestamp    `json:"moment,omitempty"`         // Дата документа
	Name           *string       `json:"name,omitempty"`           // Наименование
	Organization   *Organization `json:"organization,omitempty"`   // Метаданные юрлица
	Owner          *Employee     `json:"owner,omitempty"`          // Владелец (Сотрудник)
	Printed        *bool         `json:"printed,omitempty"`        // Напечатан ли документ
	Published      *bool         `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *Rate         `json:"rate,omitempty"`           // Валюта
	Shared         *bool         `json:"shared,omitempty"`         // Общий доступ
	State          *State        `json:"state,omitempty"`          // Метаданные статуса Счета-фактуры полученного
	Sum            *Decimal      `json:"sum,omitempty"`            // Сумма
	SyncID         *uuid.UUID    `json:"syncId,omitempty"`         // ID синхронизации. После заполнения недоступен для изменения
	Updated        *Timestamp    `json:"updated,omitempty"`        // Момент последнего обновления
	Supplies       *Supplies     `json:"supplies,omitempty"`       // Массив ссылок на связанные приемки в формате Метаданных
	Payments       *Payments     `json:"payments,omitempty"`       // Связанные исходящие платежи и расходные ордеры
	IncomingNumber *string       `json:"incomingNumber,omitempty"` // Входящий номер
	IncomingDate   *Timestamp    `json:"incomingDate,omitempty"`   // Входящая дата
}

func (f FactureIn) String() string {
	return Stringify(f)
}

func (f FactureIn) MetaType() MetaType {
	return MetaTypeFactureIn
}

// FactureInService
// Сервис для работы со счетами-фактурами полученными.
type FactureInService interface {
	GetList(ctx context.Context, params *Params) (*List[FactureIn], *resty.Response, error)
	Create(ctx context.Context, factureIn *FactureIn, params *Params) (*FactureIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, factureInList []*FactureIn, params *Params) (*[]FactureIn, *resty.Response, error)
	DeleteMany(ctx context.Context, factureInList []*FactureIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*FactureIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, factureIn *FactureIn, params *Params) (*FactureIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	//endpointTemplate[FactureIn]
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*FactureIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewFactureInService(client *Client) FactureInService {
	e := NewEndpoint(client, "entity/facturein")
	return newMainService[FactureIn, any, MetadataAttributeSharedStates, any](e)
}
