package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// DiscountService
// Сервис для работы со скидками.
type DiscountService struct {
	Endpoint
	endpointGetList[Discount]
}

func NewDiscountService(client *Client) *DiscountService {
	e := NewEndpoint(client, "entity/discount")
	return &DiscountService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Discount]{e},
	}
}

// UpdateRoundOffDiscount Изменить округление копеек.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-okruglenie-kopeek
func (s *DiscountService) UpdateRoundOffDiscount(ctx context.Context, id *uuid.UUID, entity *RoundOffDiscount) (*RoundOffDiscount, *Response, error) {
	path := id.String()
	return NewRequestBuilder[RoundOffDiscount](s.Endpoint, ctx).WithPath(path).WithBody(entity).Put()
}

// GetAccumulationDiscounts Получить все накопительные скидки.
func (s *DiscountService) GetAccumulationDiscounts(ctx context.Context, params *Params) (*List[AccumulationDiscount], *Response, error) {
	e := NewEndpoint(s.client, "entity/accumulationdiscount")
	return NewRequestBuilder[List[AccumulationDiscount]](e, ctx).WithParams(params).Get()
}

// CreateAccumulationDiscount Создать накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-nakopitel-nuu-skidku
func (s *DiscountService) CreateAccumulationDiscount(ctx context.Context, entity *AccumulationDiscount) (*AccumulationDiscount, *Response, error) {
	e := NewEndpoint(s.client, "entity/accumulationdiscount")
	return NewRequestBuilder[AccumulationDiscount](e, ctx).WithBody(entity).Post()
}

// GetByIdAccumulationDiscount Получить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-nakopitel-nuu-skidku
func (s *DiscountService) GetByIdAccumulationDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*AccumulationDiscount, *Response, error) {
	uri := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[AccumulationDiscount](e, ctx).WithParams(params).Get()
}

// UpdateAccumulationDiscount Изменить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-nakopitel-nuu-skidku
func (s *DiscountService) UpdateAccumulationDiscount(ctx context.Context, id *uuid.UUID, entity *AccumulationDiscount) (*AccumulationDiscount, *Response, error) {
	uri := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[AccumulationDiscount](e, ctx).WithBody(entity).Put()
}

// DeleteAccumulationDiscount Удалить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-nakopitel-nuu-skidku
func (s *DiscountService) DeleteAccumulationDiscount(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	uri := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[any](e, ctx).Delete()
}

// GetPersonalDiscounts Получить все персональные скидки.
func (s *DiscountService) GetPersonalDiscounts(ctx context.Context, params *Params) (*List[PersonalDiscount], *Response, error) {
	e := NewEndpoint(s.client, "entity/personaldiscount")
	return NewRequestBuilder[List[PersonalDiscount]](e, ctx).WithParams(params).Get()
}

// CreatePersonalDiscount Создать персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-personal-nuu-skidku
func (s *DiscountService) CreatePersonalDiscount(ctx context.Context, entity *PersonalDiscount) (*PersonalDiscount, *Response, error) {
	e := NewEndpoint(s.client, "entity/personaldiscount")
	return NewRequestBuilder[PersonalDiscount](e, ctx).WithBody(entity).Post()
}

// GetByIdPersonalDiscount Получить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-personal-nuu-skidku
func (s *DiscountService) GetByIdPersonalDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*PersonalDiscount, *Response, error) {
	uri := fmt.Sprintf("entity/personaldiscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[PersonalDiscount](e, ctx).WithParams(params).Get()
}

// UpdatePersonalDiscount Изменить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-personal-nuu-skidku
func (s *DiscountService) UpdatePersonalDiscount(ctx context.Context, id *uuid.UUID, entity *PersonalDiscount) (*PersonalDiscount, *Response, error) {
	uri := fmt.Sprintf("entity/personaldiscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[PersonalDiscount](e, ctx).WithBody(entity).Put()
}

// DeletePersonalDiscount Удалить персональную скидку.
func (s *DiscountService) DeletePersonalDiscount(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	uri := fmt.Sprintf("entity/personaldiscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[any](e, ctx).Delete()
}

// GetSpecialPriceDiscounts Получить все специальные цены.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-personal-nuu-skidku
func (s *DiscountService) GetSpecialPriceDiscounts(ctx context.Context, params *Params) (*List[SpecialPriceDiscount], *Response, error) {
	e := NewEndpoint(s.client, "entity/specialpricediscount")
	return NewRequestBuilder[List[SpecialPriceDiscount]](e, ctx).WithParams(params).Get()
}

// CreateSpecialPriceDiscount Создать специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-special-nuu-cenu
func (s *DiscountService) CreateSpecialPriceDiscount(ctx context.Context, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *Response, error) {
	e := NewEndpoint(s.client, "entity/specialpricediscount")
	return NewRequestBuilder[SpecialPriceDiscount](e, ctx).WithBody(entity).Post()
}

// GetByIdSpecialPriceDiscount Получить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-special-nuu-cenu
func (s *DiscountService) GetByIdSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*SpecialPriceDiscount, *Response, error) {
	uri := fmt.Sprintf("entity/specialpricediscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[SpecialPriceDiscount](e, ctx).WithParams(params).Get()
}

// UpdateSpecialPriceDiscount Изменить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-special-nuu-cenu
func (s *DiscountService) UpdateSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *Response, error) {
	uri := fmt.Sprintf("entity/specialpricediscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[SpecialPriceDiscount](e, ctx).WithBody(entity).Put()
}

// DeleteSpecialPriceDiscount Удалить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-special-nuu-cenu
func (s *DiscountService) DeleteSpecialPriceDiscount(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	uri := fmt.Sprintf("entity/specialpricediscount/%s", id)
	e := NewEndpoint(s.client, uri)
	return NewRequestBuilder[any](e, ctx).Delete()
}
