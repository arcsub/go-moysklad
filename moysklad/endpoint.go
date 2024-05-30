package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

type mainService[E any, P any, M any, S any] struct {
	endpointGetOne[E]
	endpointGetList[E]
	endpointCreate[E]
	endpointCreateUpdateMany[E]
	endpointDeleteMany[E]
	endpointDelete
	endpointGetById[E]
	endpointUpdate[E]
	endpointMetadata[M]
	endpointAttributes
	endpointNamedFilter
	endpointImages
	endpointSyncID[E]
	endpointAudit
	endpointPrintLabel
	endpointPositions[P]
	//endpointTemplate[E]
	endpointPublication
	endpointSettings[S]
	endpointGetOneAsync[E]
	endpointPrintTemplates
	endpointRemove
	endpointPrintDocument
	endpointAccounts
	endpointStates
	endpointFiles
}

func newMainService[E any, P any, M any, S any](e Endpoint) *mainService[E, P, M, S] {
	return &mainService[E, P, M, S]{
		endpointGetOne:           endpointGetOne[E]{e},
		endpointGetList:          endpointGetList[E]{e},
		endpointCreate:           endpointCreate[E]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[E]{e},
		endpointDeleteMany:       endpointDeleteMany[E]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[E]{e},
		endpointUpdate:           endpointUpdate[E]{e},
		endpointMetadata:         endpointMetadata[M]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointImages:           endpointImages{e},
		endpointSyncID:           endpointSyncID[E]{e},
		endpointAudit:            endpointAudit{e},
		endpointPrintLabel:       endpointPrintLabel{e},
		endpointPositions:        endpointPositions[P]{e},
		endpointPublication:      endpointPublication{e},
		endpointSettings:         endpointSettings[S]{e},
		endpointGetOneAsync:      endpointGetOneAsync[E]{e},
		endpointPrintTemplates:   endpointPrintTemplates{e},
		endpointRemove:           endpointRemove{e},
		endpointPrintDocument:    endpointPrintDocument{e},
		endpointAccounts:         endpointAccounts{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
	}
}

type endpointGetList[T any] struct{ Endpoint }

// GetList Запрос на получение списка объектов.
func (s *endpointGetList[T]) GetList(ctx context.Context, params *Params) (*List[T], *resty.Response, error) {
	return NewRequestBuilder[List[T]](s.client, s.uri).SetParams(params).Get(ctx)
}

type endpointDelete struct{ Endpoint }

// Delete Запрос на удаление объекта по id.
func (s *endpointDelete) Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

type endpointGetById[T any] struct{ Endpoint }

// GetByID Запрос на получение объекта по id.
func (s *endpointGetById[T]) GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[T](s.client, path).SetParams(params).Get(ctx)
}

type endpointGetOne[T any] struct{ Endpoint }

// Get Запрос (отдельный) на получение объекта. Например, ассортимент, контекст.
func (s *endpointGetOne[T]) Get(ctx context.Context, params *Params) (*T, *resty.Response, error) {
	return NewRequestBuilder[T](s.client, s.uri).SetParams(params).Get(ctx)
}

type endpointGetOneAsync[T any] struct{ Endpoint }

// GetAsync Запрос на асинхронное выполнение задачи.
// Первым возвращаемым аргументом является сервис для дальнейшей работы с конкретной асинхронной задачей.
func (s *endpointGetOneAsync[T]) GetAsync(ctx context.Context) (AsyncResultService[T], *resty.Response, error) {
	_, resp, err := NewRequestBuilder[PrintFile](s.client, s.uri).
		SetParams(new(Params).withAsync()).
		Get(ctx)

	if err != nil {
		return nil, resp, nil
	}

	async := NewAsyncResultService[T](s.client.R(), resp)
	return async, resp, err
}

type endpointMetadata[T any] struct{ Endpoint }

// GetMetadata Получить метаданные объекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-suschnosti
func (s *endpointMetadata[T]) GetMetadata(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata", s.uri)
	return NewRequestBuilder[T](s.client, path).Get(ctx)
}

type endpointTemplate[T MetaTyper] struct{ Endpoint }

// Template Получить предзаполненный стандартными полями объект.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-shablony-dokumentow
func (s *endpointTemplate[T]) Template(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/new", s.uri)
	return NewRequestBuilder[T](s.client, path).Put(ctx, nil)
}

// TemplateArg типы, которые могут быть использованы в качестве документа-основания
// при запросе на создание шаблона документа
type TemplateArg interface {
	InvoiceOutTemplateArg | SalesReturnTemplateArg | PurchaseReturnTemplateArg |
		PaymentInTemplateArg | ProcessingOrderTemplateArg | PurchaseOrderTemplateArg |
		PaymentOutTemplateArg | EnterTemplateArg | DemandTemplateArg | MoveTemplateArg |
		CashInTemplateArg | CashOutTemplateArg | RetailDemandTemplateArg | LossTemplateArg |
		InvoiceInTemplateArg | ProcessingTemplateArg
}

type endpointTemplateBasedOn[T MetaTyper, A TemplateArg] struct{ Endpoint }

// TemplateBasedOn Получить предзаполненный стандартными полями объект на основании документа(-ов)
func (s *endpointTemplateBasedOn[T, A]) TemplateBasedOn(ctx context.Context, arg *A) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/new", s.uri)
	return NewRequestBuilder[T](s.client, path).Put(ctx, arg)
}

type endpointCreate[T any] struct{ Endpoint }

// Create Запрос на создание объекта.
func (s *endpointCreate[T]) Create(ctx context.Context, entity *T, params *Params) (*T, *resty.Response, error) {
	return NewRequestBuilder[T](s.client, s.uri).SetParams(params).Post(ctx, entity)
}

// DeleteManyResponse объект ответа на запрос удаления нескольких элементов
type DeleteManyResponse []struct {
	Info      string    `json:"info"`
	ApiErrors ApiErrors `json:"errors"`
}

type endpointDeleteMany[T any] struct{ Endpoint }

// DeleteMany Запрос на удаление нескольких объектов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (s *endpointDeleteMany[T]) DeleteMany(ctx context.Context, entities []*T) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/delete", s.uri)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, entities)
}

type endpointCreateUpdateMany[T any] struct{ Endpoint }

// CreateUpdateMany Запрос на создание и обновление нескольких объектов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (s *endpointCreateUpdateMany[T]) CreateUpdateMany(ctx context.Context, entities []*T, params *Params) (*[]T, *resty.Response, error) {
	return NewRequestBuilder[[]T](s.client, s.uri).SetParams(params).Post(ctx, entities)
}

type endpointUpdate[T any] struct{ Endpoint }

// Update Запрос на обновление объекта.
func (s *endpointUpdate[T]) Update(ctx context.Context, id *uuid.UUID, entity *T, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[T](s.client, path).SetParams(params).Put(ctx, entity)
}

type endpointAccounts struct{ Endpoint }

// GetAccounts Получить все счета.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-scheta-kontragenta
func (s *endpointAccounts) GetAccounts(ctx context.Context, id *uuid.UUID) (*List[AgentAccount], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts", s.uri, id)
	return NewRequestBuilder[List[AgentAccount]](s.client, path).Get(ctx)
}

// GetAccountByID Получить отдельный счёт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-schet-kontragenta
func (s *endpointAccounts) GetAccountByID(ctx context.Context, id, accountId *uuid.UUID) (*AgentAccount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts/%s", s.uri, id, accountId)
	return NewRequestBuilder[AgentAccount](s.client, path).Get(ctx)
}

// UpdateAccounts Изменить счета (списком).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico-izmenit-scheta-urlica
func (s *endpointAccounts) UpdateAccounts(ctx context.Context, id *uuid.UUID, accounts []*AgentAccount) (*[]AgentAccount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts", s.uri, id)
	return NewRequestBuilder[[]AgentAccount](s.client, path).Post(ctx, accounts)
}

type endpointAttributes struct{ Endpoint }

// GetAttributes Получить все дополнительные поля для указанного типа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-poluchit-wse-dopolnitel-nye-polq-dlq-ukazannogo-tipa
func (s *endpointAttributes) GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", s.uri)
	return NewRequestBuilder[MetaArray[Attribute]](s.client, path).Get(ctx)
}

// GetAttributeByID Получить дополнительное поле по id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-poluchit-dopolnitel-noe-pole
func (s *endpointAttributes) GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%s", s.uri, id)
	return NewRequestBuilder[Attribute](s.client, path).Get(ctx)
}

// CreateAttribute Создать дополнительное поле.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (s *endpointAttributes) CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", s.uri)
	return NewRequestBuilder[Attribute](s.client, path).Post(ctx, attribute)
}

// CreateAttributes Создать несколько дополнительных полей.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (s *endpointAttributes) CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", s.uri)
	// при передаче массива из 1-го доп поля сервис возвращает 1 доп поле, а не массив доп полей.
	// если количество передаваемых доп полей равняется 1, то дополнительно оборачиваем в срез.
	if len(attributeList) == 1 {
		attribute, resp, err := NewRequestBuilder[Attribute](s.client, path).Post(ctx, attributeList[0])
		return &[]Attribute{*attribute}, resp, err
	}
	return NewRequestBuilder[[]Attribute](s.client, path).Post(ctx, attributeList)
}

// UpdateAttribute Изменить дополнительное поле.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-izmenit-dopolnitel-noe-pole
func (s *endpointAttributes) UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%d", s.uri, id)
	return NewRequestBuilder[Attribute](s.client, path).Put(ctx, attribute)
}

// DeleteAttribute Удалить дополнительное поле.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-udalit-dopolnitel-noe-pole
func (s *endpointAttributes) DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%d", s.uri, id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// DeleteAttributes Удалить несколько дополнительных полей.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-udalit-dopolnitel-nye-polq
func (s *endpointAttributes) DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/delete", s.uri)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, attributeList)
}

type endpointAudit struct{ Endpoint }

// GetAudit Запрос на получение событий по сущности с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-sobytiq-po-suschnosti
func (s *endpointAudit) GetAudit(ctx context.Context, id *uuid.UUID, params *Params) (*List[AuditEvent], *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/audit", s.uri, id)
	return NewRequestBuilder[List[AuditEvent]](s.client, path).SetParams(params).Get(ctx)
}

type endpointFiles struct{ Endpoint }

// GetFiles Получить список Файлов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-poluchit-spisok-fajlow-operacii-nomenklatury-zadachi-ili-kontragenta
func (s *endpointFiles) GetFiles(ctx context.Context, id *uuid.UUID) (*MetaArray[File], *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/files", s.uri, id)
	return NewRequestBuilder[MetaArray[File]](s.client, path).Get(ctx)
}

// CreateFile Добавить Файл.
func (s *endpointFiles) CreateFile(ctx context.Context, id *uuid.UUID, file *File) (*[]File, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/files", s.uri, id)
	return NewRequestBuilder[[]File](s.client, path).Get(ctx)
}

// UpdateFiles Добавить/обновить Файлы.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-dobawit-fajly-k-operacii-nomenklature-ili-kontragentu
func (s *endpointFiles) UpdateFiles(ctx context.Context, id *uuid.UUID, files []*File) (*[]File, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/files", s.uri, id)
	return NewRequestBuilder[[]File](s.client, path).Post(ctx, files)
}

// DeleteFile Удалить Файл.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-udalit-fajl
func (s *endpointFiles) DeleteFile(ctx context.Context, id, fileId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/files/%d", s.uri, id, fileId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// DeleteFiles Удалить несколько Файлов.
func (s *endpointFiles) DeleteFiles(ctx context.Context, id *uuid.UUID, files []*File) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/files/delete", s.uri, id)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, files)
}

type endpointImages struct{ Endpoint }

// GetImages Получить список Изображений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-poluchit-spisok-izobrazhenij-towara-komplekta-i-modifikacii
func (s *endpointImages) GetImages(ctx context.Context, id *uuid.UUID) (*MetaArray[Image], *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/images", s.uri, id)
	return NewRequestBuilder[MetaArray[Image]](s.client, path).Get(ctx)
}

// CreateImage Добавить Изображение.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-dobawit-izobrazhenie-k-towaru-komplektu-ili-modifikacii
func (s *endpointImages) CreateImage(ctx context.Context, id *uuid.UUID, image *Image) (*[]*Image, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/images", s.uri, id)
	return NewRequestBuilder[[]*Image](s.client, path).Post(ctx, image)
}

// UpdateImages Изменение Изображений (списком).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-izmenenie-spiska-izobrazhenij-u-towara-komplekta-ili-modifikacii
func (s *endpointImages) UpdateImages(ctx context.Context, id *uuid.UUID, images []*Image) (*[]Image, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/images", s.uri, id)
	return NewRequestBuilder[[]Image](s.client, path).Post(ctx, images)
}

// DeleteImage Удалить Изображение.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-izobrazhenie
func (s *endpointImages) DeleteImage(ctx context.Context, id, imageId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/images/%d", s.uri, id, imageId)
	return NewRequestBuilder[[]Image](s.client, path).Delete(ctx)
}

// DeleteImages Удалить несколько Изображений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-gruppu-izobrazhenij
func (s *endpointImages) DeleteImages(ctx context.Context, id *uuid.UUID, images []*Image) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%d/images/delete", s.uri, id)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, images)
}

type endpointNamedFilter struct{ Endpoint }

// GetNamedFilters Получить список фильтров.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-spisok-fil-trow
func (s *endpointNamedFilter) GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error) {
	path := fmt.Sprintf("%s/namedfilter", s.uri)
	return NewRequestBuilder[List[NamedFilter]](s.client, path).SetParams(params).Get(ctx)
}

// GetNamedFilterByID Получить отдельный фильтр по id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-fil-tr-po-id
func (s *endpointNamedFilter) GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error) {
	path := fmt.Sprintf("%s/namedfilter/%d", s.uri, id)
	return NewRequestBuilder[NamedFilter](s.client, path).Get(ctx)
}

type endpointPositions[T any] struct{ Endpoint }

// GetPositions Получить все позиции документа.
func (s *endpointPositions[T]) GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[T], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", s.uri, id)
	return NewRequestBuilder[MetaArray[T]](s.client, path).SetParams(params).Get(ctx)
}

// GetPositionByID Получение отдельной позиции.
func (s *endpointPositions[T]) GetPositionByID(ctx context.Context, id, positionID *uuid.UUID, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", s.uri, id, positionID)
	return NewRequestBuilder[T](s.client, path).SetParams(params).Get(ctx)
}

// UpdatePosition Обновление позиции.
func (s *endpointPositions[T]) UpdatePosition(ctx context.Context, id, positionID *uuid.UUID, position *T, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", s.uri, id, positionID)
	return NewRequestBuilder[T](s.client, path).SetParams(params).Put(ctx, position)
}

// CreatePosition Создание позиции документа.
func (s *endpointPositions[T]) CreatePosition(ctx context.Context, id *uuid.UUID, position *T) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", s.uri, id)
	return NewRequestBuilder[T](s.client, path).Post(ctx, position)
}

// CreatePositions Массово создаёт позиции документа.
func (s *endpointPositions[T]) CreatePositions(ctx context.Context, id *uuid.UUID, positions []*T) (*[]T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", s.uri, id)
	return NewRequestBuilder[[]T](s.client, path).Post(ctx, positions)
}

// DeletePosition Удаляет позицию документа.
func (s *endpointPositions[T]) DeletePosition(ctx context.Context, id, positionID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", s.uri, id, positionID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetPositionTrackingCodes Получить Коды маркировки позиции документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-poluchit-kody-markirowki-pozicii-dokumenta
func (s *endpointPositions[T]) GetPositionTrackingCodes(ctx context.Context, id, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes", s.uri, id, positionID)
	return NewRequestBuilder[MetaArray[TrackingCode]](s.client, path).Get(ctx)
}

// CreateOrUpdatePositionTrackingCodes Массовое создание и обновление Кодов маркировки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-sozdanie-i-obnowlenie-kodow-markirowki
func (s *endpointPositions[T]) CreateOrUpdatePositionTrackingCodes(ctx context.Context, id, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes", s.uri, id, positionID)
	return NewRequestBuilder[[]TrackingCode](s.client, path).Post(ctx, trackingCodes)
}

// DeletePositionTrackingCodes Массовое удаление Кодов маркировки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-udalenie-kodow-markirowki
func (s *endpointPositions[T]) DeletePositionTrackingCodes(ctx context.Context, id, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes/delete", s.uri, id, positionID)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, trackingCodes)
}

type endpointPrintDocument struct{ Endpoint }

// PrintDocument Запрос на печать документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-pechat-dokumentow-zapros-na-pechat
func (s *endpointPrintDocument) PrintDocument(ctx context.Context, id *uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/export", s.uri, id)

	_, resp, err := NewRequestBuilder[PrintFile](s.client, path).
		SetHeader(headerGetContent, "true").
		Post(ctx, PrintDocumentArg)

	if err != nil {
		return nil, resp, err
	}

	file, err := GetFileFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return file, resp, err
}

type endpointPrintLabel struct{ Endpoint }

// PrintLabel Запрос на печать этикеток и ценников.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pechat-atiketok-i-cennikow
func (s *endpointPrintLabel) PrintLabel(ctx context.Context, id *uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/export", s.uri, id)

	_, resp, err := NewRequestBuilder[PrintFile](s.client, path).
		SetHeader(headerGetContent, "true").
		Post(ctx, PrintLabelArg)

	if err != nil {
		return nil, resp, err
	}

	file, err := GetFileFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return file, resp, err
}

type endpointPublication struct{ Endpoint }

// GetPublications Запрос на получение списка Публикаций по указанному документу.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikacii
func (s *endpointPublication) GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication", s.uri, id)
	return NewRequestBuilder[MetaArray[Publication]](s.client, path).Get(ctx)
}

// GetPublicationByID Запрос на получение Публикации с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikaciu
func (s *endpointPublication) GetPublicationByID(ctx context.Context, id, publicationID *uuid.UUID) (*Publication, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication/%s", s.uri, id, publicationID)
	return NewRequestBuilder[Publication](s.client, path).Get(ctx)
}

// Publish Запрос на публикацию документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-sozdat-publikaciu
func (s *endpointPublication) Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error) {
	publication := new(Publication).SetTemplate(template)
	path := fmt.Sprintf("%s/%s/publication", s.uri, id)
	return NewRequestBuilder[Publication](s.client, path).Post(ctx, publication)
}

// DeletePublication Запрос на удаление Публикации с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-udalit-publikaciu
func (s *endpointPublication) DeletePublication(ctx context.Context, id, publicationID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication/%s", s.uri, id, publicationID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

type endpointSettings[T any] struct{ Endpoint }

// GetSettings Запрос на получение настроек справочника.
func (s *endpointSettings[T]) GetSettings(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", s.uri)
	return NewRequestBuilder[T](s.client, path).Get(ctx)
}

// UpdateSettings Изменить настройки справочника.
func (s *endpointSettings[T]) UpdateSettings(ctx context.Context, settings *T) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", s.uri)
	return NewRequestBuilder[T](s.client, path).Put(ctx, settings)
}

type endpointStates struct{ Endpoint }

// GetStateByID Запрос на получение статуса по id.
func (s *endpointStates) GetStateByID(ctx context.Context, id *uuid.UUID) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", s.uri, id)
	return NewRequestBuilder[State](s.client, path).Get(ctx)
}

// CreateState Создать новый статус.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-sozdat-status
func (s *endpointStates) CreateState(ctx context.Context, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states", s.uri)
	return NewRequestBuilder[State](s.client, path).Post(ctx, state)
}

// UpdateState Изменить существующий статус.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-izmenit-status
func (s *endpointStates) UpdateState(ctx context.Context, id *uuid.UUID, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", s.uri, id)
	return NewRequestBuilder[State](s.client, path).Put(ctx, state)
}

// CreateOrUpdateStates Массовое создание и обновление Статусов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-massowoe-sozdanie-i-obnowlenie-statusow
func (s *endpointStates) CreateOrUpdateStates(ctx context.Context, id *uuid.UUID, states []*State) (*[]State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", s.uri, id)
	return NewRequestBuilder[[]State](s.client, path).Post(ctx, states)
}

// DeleteState Запрос на удаление Статуса с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-udalit-status
func (s *endpointStates) DeleteState(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", s.uri, id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

type endpointSyncID[T any] struct{ Endpoint }

// GetBySyncID Запрос на получение объекта по syncID.
func (s *endpointSyncID[T]) GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/syncid/%s", s.uri, syncID)
	return NewRequestBuilder[T](s.client, path).Get(ctx)
}

// DeleteBySyncID Запрос на удаление объекта по syncID.
func (s *endpointSyncID[T]) DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/syncid/%s", s.uri, syncID)
	return NewRequestBuilder[T](s.client, path).Delete(ctx)
}

type endpointPrintTemplates struct{ Endpoint }

// GetEmbeddedTemplates Запрос на получение информации о стандартных шаблонах печатных форм для указанного типа сущности.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-standartnyh-shablonow
func (s *endpointPrintTemplates) GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate", s.uri)
	return NewRequestBuilder[List[EmbeddedTemplate]](s.client, path).Get(ctx)
}

// GetEmbeddedTemplateByID Запрос на получение информации об отдельном стандартном шаблоне печатной формы для указанного типа сущности по его id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-standartnyj-shablon
func (s *endpointPrintTemplates) GetEmbeddedTemplateByID(ctx context.Context, id *uuid.UUID) (*EmbeddedTemplate, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate/%s", s.uri, id)
	return NewRequestBuilder[EmbeddedTemplate](s.client, path).Get(ctx)
}

// GetCustomTemplates Запрос на получение информации о пользовательских шаблонах печатных форм для указанного типа сущности.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-pol-zowatel-skih-shablonow
func (s *endpointPrintTemplates) GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/customtemplate", s.uri)
	return NewRequestBuilder[List[CustomTemplate]](s.client, path).Get(ctx)
}

// GetCustomTemplateByID Запрос на получение информации об отдельном пользовательском шаблоне печатной формы для указанного типа сущности по его id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-pol-zowatel-skij-shablon
func (s *endpointPrintTemplates) GetCustomTemplateByID(ctx context.Context, id *uuid.UUID) (*CustomTemplate, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate/%s", s.uri, id)
	return NewRequestBuilder[CustomTemplate](s.client, path).Get(ctx)
}

type endpointRemove struct{ Endpoint }

// Remove Запрос на перемещение документа с указанным id в корзину.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-udalenie-w-korzinu
func (s *endpointRemove) Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/trash", s.uri, id)
	_, resp, err := NewRequestBuilder[any](s.client, path).Post(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}
