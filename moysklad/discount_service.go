package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
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
func (s *DiscountService) UpdateRoundOffDiscount(ctx context.Context, id *uuid.UUID, entity *RoundOffDiscount) (*RoundOffDiscount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[RoundOffDiscount](s.client, path).Put(ctx, entity)
}

// GetAccumulationDiscounts Получить все накопительные скидки.
func (s *DiscountService) GetAccumulationDiscounts(ctx context.Context, params *Params) (*List[AccumulationDiscount], *resty.Response, error) {
	path := "entity/accumulationdiscount"
	return NewRequestBuilder[List[AccumulationDiscount]](s.client, path).SetParams(params).Get(ctx)
}

// CreateAccumulationDiscount Создать накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-nakopitel-nuu-skidku
func (s *DiscountService) CreateAccumulationDiscount(ctx context.Context, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	path := "entity/accumulationdiscount"
	return NewRequestBuilder[AccumulationDiscount](s.client, path).Post(ctx, entity)
}

// GetByIdAccumulationDiscount Получить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-nakopitel-nuu-skidku
func (s *DiscountService) GetByIdAccumulationDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[AccumulationDiscount](s.client, path).SetParams(params).Get(ctx)
}

// UpdateAccumulationDiscount Изменить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-nakopitel-nuu-skidku
func (s *DiscountService) UpdateAccumulationDiscount(ctx context.Context, id *uuid.UUID, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[AccumulationDiscount](s.client, path).Put(ctx, entity)
}

// DeleteAccumulationDiscount Удалить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-nakopitel-nuu-skidku
func (s *DiscountService) DeleteAccumulationDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetPersonalDiscounts Получить все персональные скидки.
func (s *DiscountService) GetPersonalDiscounts(ctx context.Context, params *Params) (*List[PersonalDiscount], *resty.Response, error) {
	path := "entity/personaldiscount"
	return NewRequestBuilder[List[PersonalDiscount]](s.client, path).SetParams(params).Get(ctx)
}

// CreatePersonalDiscount Создать персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-personal-nuu-skidku
func (s *DiscountService) CreatePersonalDiscount(ctx context.Context, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	path := "entity/personaldiscount"
	return NewRequestBuilder[PersonalDiscount](s.client, path).Post(ctx, entity)
}

// GetByIdPersonalDiscount Получить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-personal-nuu-skidku
func (s *DiscountService) GetByIdPersonalDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[PersonalDiscount](s.client, path).SetParams(params).Get(ctx)
}

// UpdatePersonalDiscount Изменить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-personal-nuu-skidku
func (s *DiscountService) UpdatePersonalDiscount(ctx context.Context, id *uuid.UUID, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[PersonalDiscount](s.client, path).Put(ctx, entity)
}

// DeletePersonalDiscount Удалить персональную скидку.
func (s *DiscountService) DeletePersonalDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetSpecialPriceDiscounts Получить все специальные цены.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-personal-nuu-skidku
func (s *DiscountService) GetSpecialPriceDiscounts(ctx context.Context, params *Params) (*List[SpecialPriceDiscount], *resty.Response, error) {
	path := "entity/specialpricediscount"
	return NewRequestBuilder[List[SpecialPriceDiscount]](s.client, path).SetParams(params).Get(ctx)
}

// CreateSpecialPriceDiscount Создать специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-special-nuu-cenu
func (s *DiscountService) CreateSpecialPriceDiscount(ctx context.Context, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	path := "entity/specialpricediscount"
	return NewRequestBuilder[SpecialPriceDiscount](s.client, path).Post(ctx, entity)
}

// GetByIdSpecialPriceDiscount Получить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-special-nuu-cenu
func (s *DiscountService) GetByIdSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[SpecialPriceDiscount](s.client, path).SetParams(params).Get(ctx)
}

// UpdateSpecialPriceDiscount Изменить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-special-nuu-cenu
func (s *DiscountService) UpdateSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[SpecialPriceDiscount](s.client, path).Put(ctx, entity)
}

// DeleteSpecialPriceDiscount Удалить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-special-nuu-cenu
func (s *DiscountService) DeleteSpecialPriceDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
