package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"

	"time"
)

// TaxRate Ставка НДС.
//
// Код сущности: taxrate
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-stawka-nds
type TaxRate struct {
	AccountID *string    `json:"accountId,omitempty"` // ID учётной записи
	Archived  *bool      `json:"archived,omitempty"`  // Флаг принадлежности ставки к архивным ставкам
	Comment   *string    `json:"comment,omitempty"`   // Комментарий к налоговой ставке
	Group     *Group     `json:"group,omitempty"`     // Отдел сотрудника
	ID        *string    `json:"id,omitempty"`        // ID налоговой ставки
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные налоговой ставки
	Rate      *int       `json:"rate,omitempty"`      // Значение налоговой ставки
	Owner     *Employee  `json:"owner,omitempty"`     // Метаданные владельца (Сотрудника)
	Shared    *bool      `json:"shared,omitempty"`    // Общий доступ
	Updated   *Timestamp `json:"updated,omitempty"`   // Момент последнего обновления налоговой ставки
}

// GetAccountID возвращает ID учётной записи.
func (taxRate TaxRate) GetAccountID() string {
	return Deref(taxRate.AccountID)
}

// GetArchived возвращает флаг нахождения в архиве.
func (taxRate TaxRate) GetArchived() bool {
	return Deref(taxRate.Archived)
}

// GetComment возвращает Комментарий к налоговой ставке.
func (taxRate TaxRate) GetComment() string {
	return Deref(taxRate.Comment)
}

// GetGroup возвращает Отдел сотрудника.
func (taxRate TaxRate) GetGroup() Group {
	return Deref(taxRate.Group)
}

// GetID возвращает ID налоговой ставки.
func (taxRate TaxRate) GetID() string {
	return Deref(taxRate.ID)
}

// GetMeta возвращает Метаданные налоговой ставки.
func (taxRate TaxRate) GetMeta() Meta {
	return Deref(taxRate.Meta)
}

// GetRate возвращает Значение налоговой ставки.
func (taxRate TaxRate) GetRate() int {
	return Deref(taxRate.Rate)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (taxRate TaxRate) GetOwner() Employee {
	return Deref(taxRate.Owner)
}

// GetShared возвращает флаг Общего доступа.
func (taxRate TaxRate) GetShared() bool {
	return Deref(taxRate.Shared)
}

// GetUpdated возвращает Момент последнего обновления налоговой ставки.
func (taxRate TaxRate) GetUpdated() time.Time {
	return Deref(taxRate.Updated).Time()
}

// SetArchived устанавливает флаг нахождения в архиве.
func (taxRate *TaxRate) SetArchived(archived bool) *TaxRate {
	taxRate.Archived = &archived
	return taxRate
}

// SetComment устанавливает Комментарий к налоговой ставке.
func (taxRate *TaxRate) SetComment(comment string) *TaxRate {
	taxRate.Comment = &comment
	return taxRate
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (taxRate *TaxRate) SetGroup(group *Group) *TaxRate {
	if group != nil {
		taxRate.Group = group.Clean()
	}
	return taxRate
}

// SetMeta устанавливает Метаданные налоговой ставки.
func (taxRate *TaxRate) SetMeta(meta *Meta) *TaxRate {
	taxRate.Meta = meta
	return taxRate
}

// SetRate устанавливает Значение налоговой ставки.
func (taxRate *TaxRate) SetRate(rate int) *TaxRate {
	taxRate.Rate = &rate
	return taxRate
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (taxRate *TaxRate) SetOwner(owner *Employee) *TaxRate {
	if owner != nil {
		taxRate.Owner = owner.Clean()
	}
	return taxRate
}

// SetShared устанавливает флаг общего доступа.
func (taxRate *TaxRate) SetShared(shared bool) *TaxRate {
	taxRate.Shared = &shared
	return taxRate
}

// String реализует интерфейс [fmt.Stringer].
func (taxRate TaxRate) String() string {
	return Stringify(taxRate)
}

// MetaType возвращает код сущности.
func (TaxRate) MetaType() MetaType {
	return MetaTypeTaxRate
}

// Update shortcut
func (taxRate *TaxRate) Update(ctx context.Context, client *Client, params ...func(*Params)) (*TaxRate, *resty.Response, error) {
	return NewTaxRateService(client).Update(ctx, taxRate.GetID(), taxRate, params...)
}

// Create shortcut
func (taxRate *TaxRate) Create(ctx context.Context, client *Client, params ...func(*Params)) (*TaxRate, *resty.Response, error) {
	return NewTaxRateService(client).Create(ctx, taxRate, params...)
}

// Delete shortcut
func (taxRate *TaxRate) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewTaxRateService(client).Delete(ctx, taxRate)
}

// TaxRateService описывает методы сервиса для работы со ставками НДС.
type TaxRateService interface {
	// GetList выполняет запрос на получение списка налоговых ставок.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...func(*Params)) (*List[TaxRate], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех налоговых ставок в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...func(*Params)) (*Slice[TaxRate], *resty.Response, error)

	// Create выполняет запрос на создание налоговой ставки.
	// Обязательные поля для заполнения:
	//	- rate (Значение налоговой ставки)
	// Принимает контекст, налоговую ставку и опционально объект параметров запроса Params.
	// Возвращает созданную налоговую ставку.
	Create(ctx context.Context, taxRate *TaxRate, params ...func(*Params)) (*TaxRate, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение налоговых ставок.
	// Изменяемые налоговые ставки должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список налоговых ставок и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых налоговых ставок.
	CreateUpdateMany(ctx context.Context, taxRateList Slice[TaxRate], params ...func(*Params)) (*Slice[TaxRate], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление налоговых ставок.
	// Принимает контекст и множество налоговых ставок.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*TaxRate) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление налоговой ставки по ID.
	// Принимает контекст и ID налоговой ставки.
	// Возвращает «true» в случае успешного удаления налоговой ставки.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление налоговой ставки.
	// Принимает контекст и налоговую ставку.
	// Возвращает «true» в случае успешного удаления налоговой ставки.
	Delete(ctx context.Context, entity *TaxRate) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной налоговой ставки по ID.
	// Принимает контекст, ID налоговой ставки и опционально объект параметров запроса Params.
	// Возвращает налоговую ставку.
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*TaxRate, *resty.Response, error)

	// Update выполняет запрос на изменение налоговой ставки.
	// Принимает контекст, налоговую ставку и опционально объект параметров запроса Params.
	// Возвращает изменённую налоговую ставку.
	Update(ctx context.Context, id string, taxRate *TaxRate, params ...func(*Params)) (*TaxRate, *resty.Response, error)
}

const (
	EndpointTaxRate = EndpointEntity + string(MetaTypeTaxRate)
)

// NewTaxRateService принимает [Client] и возвращает сервис для работы со ставками НДС.
func NewTaxRateService(client *Client) TaxRateService {
	return newMainService[TaxRate, any, any, any](client, EndpointTaxRate)
}
