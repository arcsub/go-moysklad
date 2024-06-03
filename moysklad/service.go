package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Service Услуга.
// Ключевое слово: service
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-usluga
type Service struct {
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`
	Group               *Group         `json:"group,omitempty"`
	Barcodes            Barcodes       `json:"barcodes,omitempty"`
	Description         *string        `json:"description,omitempty"`
	ExternalCode        *string        `json:"externalCode,omitempty"`
	ID                  *uuid.UUID     `json:"id,omitempty"`
	Meta                *Meta          `json:"meta,omitempty"`
	Name                *string        `json:"name,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	Files               *Files         `json:"files,omitempty"`
	BuyPrice            *BuyPrice      `json:"buyPrice,omitempty"`
	DiscountProhibited  *bool          `json:"discountProhibited,omitempty"`
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"`
	UseParentVat        *bool          `json:"useParentVat,omitempty"`
	Code                *string        `json:"code,omitempty"`
	MinPrice            *MinPrice      `json:"minPrice,omitempty"`
	Owner               *Employee      `json:"owner,omitempty"`
	PathName            *string        `json:"pathName,omitempty"`
	AccountID           *uuid.UUID     `json:"accountId,omitempty"`
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`
	SalePrices          *SalePrices    `json:"salePrices,omitempty"`
	Shared              *bool          `json:"shared,omitempty"`
	SyncID              *uuid.UUID     `json:"syncId,omitempty"`
	Vat                 *int           `json:"vat,omitempty"`
	Uom                 *Uom           `json:"uom,omitempty"`
	Updated             *Timestamp     `json:"updated,omitempty"`
	PaymentItemType     PaymentItem    `json:"paymentItemType,omitempty"`
	TaxSystem           TaxSystem      `json:"taxSystem,omitempty"`
	Attributes          Attributes     `json:"attributes,omitempty"`
}

func (s Service) String() string {
	return Stringify(s)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (s Service) GetMeta() Meta {
	return Deref(s.Meta)
}

func (s Service) MetaType() MetaType {
	return MetaTypeService
}

// ServiceService
// Сервис для работы с услугами.
type ServiceService interface {
	GetList(ctx context.Context, params *Params) (*List[Service], *resty.Response, error)
	Create(ctx context.Context, service *Service, params *Params) (*Service, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, serviceList []*Service, params *Params) (*[]Service, *resty.Response, error)
	DeleteMany(ctx context.Context, serviceList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Service, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, service *Service, params *Params) (*Service, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Service, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewServiceService(client *Client) ServiceService {
	e := NewEndpoint(client, "entity/service")
	return newMainService[Service, any, any, any](e)
}
