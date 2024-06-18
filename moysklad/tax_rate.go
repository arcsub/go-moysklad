package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// TaxRate Ставка НДС.
// Ключевое слово: taxrate
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-stawka-nds
type TaxRate struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	Archived  *bool      `json:"archived,omitempty"`  // Флаг принадлежности ставки к архивным ставкам
	Comment   *string    `json:"comment,omitempty"`   // Комментарий к налоговой ставке
	Group     *Group     `json:"group,omitempty"`     // Отдел сотрудника
	ID        *uuid.UUID `json:"id,omitempty"`        // ID налоговой ставки
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные налоговой ставки
	Rate      *int       `json:"rate,omitempty"`      // Значение налоговой ставки
	Owner     *Employee  `json:"owner,omitempty"`     // Владелец (Сотрудник)
	Shared    *bool      `json:"shared,omitempty"`    // Общий доступ
	Updated   *Timestamp `json:"updated,omitempty"`   // Момент последнего обновления сущности
}

func (taxRate TaxRate) GetAccountID() uuid.UUID {
	return Deref(taxRate.AccountID)
}

func (taxRate TaxRate) GetArchived() bool {
	return Deref(taxRate.Archived)
}

func (taxRate TaxRate) GetComment() string {
	return Deref(taxRate.Comment)
}

func (taxRate TaxRate) GetGroup() Group {
	return Deref(taxRate.Group)
}

func (taxRate TaxRate) GetID() uuid.UUID {
	return Deref(taxRate.ID)
}

func (taxRate TaxRate) GetMeta() Meta {
	return Deref(taxRate.Meta)
}

func (taxRate TaxRate) GetRate() int {
	return Deref(taxRate.Rate)
}

func (taxRate TaxRate) GetOwner() Employee {
	return Deref(taxRate.Owner)
}

func (taxRate TaxRate) GetShared() bool {
	return Deref(taxRate.Shared)
}

func (taxRate TaxRate) GetUpdated() Timestamp {
	return Deref(taxRate.Updated)
}

func (taxRate *TaxRate) SetArchived(archived bool) *TaxRate {
	taxRate.Archived = &archived
	return taxRate
}

func (taxRate *TaxRate) SetComment(comment string) *TaxRate {
	taxRate.Comment = &comment
	return taxRate
}

func (taxRate *TaxRate) SetGroup(group *Group) *TaxRate {
	taxRate.Group = group.Clean()
	return taxRate
}

func (taxRate *TaxRate) SetMeta(meta *Meta) *TaxRate {
	taxRate.Meta = meta
	return taxRate
}

func (taxRate *TaxRate) SetRate(rate int) *TaxRate {
	taxRate.Rate = &rate
	return taxRate
}

func (taxRate *TaxRate) SetOwner(owner *Employee) *TaxRate {
	taxRate.Owner = owner.Clean()
	return taxRate
}

func (taxRate *TaxRate) SetShared(shared bool) *TaxRate {
	taxRate.Shared = &shared
	return taxRate
}

func (taxRate TaxRate) String() string {
	return Stringify(taxRate)
}

func (taxRate TaxRate) MetaType() MetaType {
	return MetaTypeTaxRate
}

// Update shortcut
func (taxRate TaxRate) Update(ctx context.Context, client *Client, params ...*Params) (*TaxRate, *resty.Response, error) {
	return client.Entity().TaxRate().Update(ctx, taxRate.GetID(), &taxRate, params...)
}

// Create shortcut
func (taxRate TaxRate) Create(ctx context.Context, client *Client, params ...*Params) (*TaxRate, *resty.Response, error) {
	return client.Entity().TaxRate().Create(ctx, &taxRate, params...)
}

// Delete shortcut
func (taxRate TaxRate) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().TaxRate().Delete(ctx, taxRate.GetID())
}

// TaxRateService
// Сервис для работы со ставками НДС.
type TaxRateService interface {
	GetList(ctx context.Context, params ...*Params) (*List[TaxRate], *resty.Response, error)
	Create(ctx context.Context, taxRate *TaxRate, params ...*Params) (*TaxRate, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, taxRateList Slice[TaxRate], params ...*Params) (*Slice[TaxRate], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...TaxRate) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*TaxRate, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, taxRate *TaxRate, params ...*Params) (*TaxRate, *resty.Response, error)
}

func NewTaxRateService(client *Client) TaxRateService {
	e := NewEndpoint(client, "entity/taxrate")
	return newMainService[TaxRate, any, any, any](e)
}
