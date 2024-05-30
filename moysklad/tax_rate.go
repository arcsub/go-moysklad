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

func (t TaxRate) String() string {
	return Stringify(t)
}

func (t TaxRate) MetaType() MetaType {
	return MetaTypeTaxRate
}

// TaxRateService
// Сервис для работы со ставками НДС.
type TaxRateService interface {
	GetList(ctx context.Context, params *Params) (*List[TaxRate], *resty.Response, error)
	Create(ctx context.Context, taxRate *TaxRate, params *Params) (*TaxRate, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, taxRateList []*TaxRate, params *Params) (*[]TaxRate, *resty.Response, error)
	DeleteMany(ctx context.Context, taxRateList []*TaxRate) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*TaxRate, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, taxRate *TaxRate, params *Params) (*TaxRate, *resty.Response, error)
}

func NewTaxRateService(client *Client) TaxRateService {
	e := NewEndpoint(client, "entity/taxrate")
	return newMainService[TaxRate, any, any, any](e)
}
