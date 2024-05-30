package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureOut Счет-фактура выданный.
// Ключевое слово: factureout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-wydannyj
type FactureOut struct {
	AccountID       *uuid.UUID       `json:"accountId,omitempty"`       // ID учетной записи
	Agent           *Counterparty    `json:"agent,omitempty"`           // Метаданные контрагента
	Applicable      *bool            `json:"applicable,omitempty"`      // Отметка о проведении
	Attributes      *Attributes      `json:"attributes,omitempty"`      // Коллекция метаданных доп. полей. Поля объекта
	Code            *string          `json:"code,omitempty"`            // Код выданного Счета-фактуры
	Contract        *Contract        `json:"contract,omitempty"`        // Метаданные договора
	Created         *Timestamp       `json:"created,omitempty"`         // Дата создания
	Deleted         *Timestamp       `json:"deleted,omitempty"`         // Момент последнего удаления Счета-фактуры
	Description     *string          `json:"description,omitempty"`     // Комментарий выданного Счета-фактуры
	ExternalCode    *string          `json:"externalCode,omitempty"`    // Внешний код выданного Счета-фактуры
	Files           *Files           `json:"files,omitempty"`           // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group           *Group           `json:"group,omitempty"`           // Отдел сотрудника
	ID              *uuid.UUID       `json:"id,omitempty"`              // ID сущности
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные
	Moment          *Timestamp       `json:"moment,omitempty"`          // Дата документа
	Name            *string          `json:"name,omitempty"`            // Наименование
	Organization    *Organization    `json:"organization,omitempty"`    // Метаданные юрлица
	Owner           *Employee        `json:"owner,omitempty"`           // Владелец (Сотрудник)
	Printed         *bool            `json:"printed,omitempty"`         // Напечатан ли документ
	Published       *bool            `json:"published,omitempty"`       // Опубликован ли документ
	Rate            *Rate            `json:"rate,omitempty"`            // Валюта
	Shared          *bool            `json:"shared,omitempty"`          // Общий доступ
	State           *State           `json:"state,omitempty"`           // Метаданные статуса Счета-фактуры
	StateContractId *string          `json:"stateContractId,omitempty"` // Идентификатор гос. контракта
	Sum             *Decimal         `json:"sum,omitempty"`             // Сумма
	SyncID          *uuid.UUID       `json:"syncId,omitempty"`          // ID синхронизации. После заполнения недоступен для изменения
	Updated         *Timestamp       `json:"updated,omitempty"`         // Момент последнего обновления
	Demands         *Demands         `json:"demands,omitempty"`         // Массив ссылок на связанные отгрузки в формате Метаданных
	Payments        *Payments        `json:"payments,omitempty"`        // Массив ссылок на связанные входящие платежи в формате Метаданных
	Returns         *PurchaseReturns `json:"returns,omitempty"`         // Массив ссылок на связанные возвраты поставщикам в формате Метаданных
	Consignee       *Counterparty    `json:"consignee,omitempty"`       // Грузополучатель
	PaymentNumber   *string          `json:"paymentNumber,omitempty"`   // Название платежного документа
	PaymentDate     *Timestamp       `json:"paymentDate,omitempty"`     // Дата платежного документа
}

func (f FactureOut) String() string {
	return Stringify(f)
}

func (f FactureOut) MetaType() MetaType {
	return MetaTypeFactureOut
}

// FactureOutService
// Сервис для работы со счетами-фактурами выданными.
type FactureOutService interface {
	GetList(ctx context.Context, params *Params) (*List[FactureOut], *resty.Response, error)
	Create(ctx context.Context, factureOut *FactureOut, params *Params) (*FactureOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, factureOutList []*FactureOut, params *Params) (*[]FactureOut, *resty.Response, error)
	DeleteMany(ctx context.Context, factureOutList []*FactureOut) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*FactureOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, factureOut *FactureOut, params *Params) (*FactureOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*FactureOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewFactureOutService(client *Client) FactureOutService {
	e := NewEndpoint(client, "entity/factureout")
	return newMainService[FactureOut, any, MetadataAttributeSharedStates, any](e)
}
