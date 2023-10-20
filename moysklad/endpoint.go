package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type endpointDelete struct{ Endpoint }

func (s *endpointDelete) Delete(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	path := id.String()
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

type endpointGetList[T any] struct{ Endpoint }

func (s *endpointGetList[T]) Get(ctx context.Context, params *Params) (*List[T], *Response, error) {
	return NewRequestBuilder[List[T]](s.Endpoint, ctx).WithParams(params).Get()
}

type endpointGetById[T any] struct{ Endpoint }

func (s *endpointGetById[T]) GetById(ctx context.Context, id *uuid.UUID, params *Params) (*T, *Response, error) {
	path := id.String()
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

type endpointGetOne[T any] struct{ Endpoint }

func (s *endpointGetOne[T]) Get(ctx context.Context, params *Params) (*T, *Response, error) {
	return NewRequestBuilder[T](s.Endpoint, ctx).WithParams(params).Get()
}

type endpointGetOneAsync[T any] struct{ Endpoint }

func (s *endpointGetOneAsync[T]) GetAsync(ctx context.Context) (*AsyncResultService[T], *Response, error) {
	return NewRequestBuilder[T](s.Endpoint, ctx).Async()
}

type endpointMetadata[T any] struct{ Endpoint }

// GetMetadata Получить метаданные сущности
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-suschnosti
func (s *endpointMetadata[T]) GetMetadata(ctx context.Context) (*T, *Response, error) {
	path := "metadata"
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).Get()
}

type endpointTemplate[T MetaTyper] struct{ Endpoint }

// Template Получить предзаполненный стандартными полями JSON-объект
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-shablony-dokumentow
func (s *endpointTemplate[T]) Template(ctx context.Context) (*T, *Response, error) {
	path := "new"
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).Put()
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

// TemplateBasedOn Получить предзаполненный стандартными полями JSON-объект на основании документа(-ов)
func (s *endpointTemplateBasedOn[T, A]) TemplateBasedOn(ctx context.Context, arg *A) (*T, *Response, error) {
	path := "new"
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).WithBody(arg).Put()
}

type endpointCreate[T any] struct{ Endpoint }

// Create создать элемент
func (s *endpointCreate[T]) Create(ctx context.Context, entity *T, params *Params) (*T, *Response, error) {
	return NewRequestBuilder[T](s.Endpoint, ctx).WithParams(params).WithBody(entity).Post()
}

// DeleteManyResponse объект ответа на запрос удаления нескольких элементов
type DeleteManyResponse []struct {
	Info      string    `json:"info"`
	ApiErrors ApiErrors `json:"errors"`
}

type endpointDeleteMany[T any] struct{ Endpoint }

// DeleteMany Удаление нескольких объектов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (s *endpointDeleteMany[T]) DeleteMany(ctx context.Context, entities []*T) (*DeleteManyResponse, *Response, error) {
	path := "delete"
	return NewRequestBuilder[DeleteManyResponse](s.Endpoint, ctx).WithPath(path).WithBody(entities).Post()
}

type endpointCreateUpdateDeleteMany[T any] struct{ Endpoint }

// CreateUpdateMany Создание и обновление нескольких объектов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (s *endpointCreateUpdateDeleteMany[T]) CreateUpdateMany(ctx context.Context, entities []*T, params *Params) (*[]T, *Response, error) {
	return NewRequestBuilder[[]T](s.Endpoint, ctx).WithParams(params).WithBody(entities).Post()
}

// DeleteMany Удаление нескольких объектов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (s *endpointCreateUpdateDeleteMany[T]) DeleteMany(ctx context.Context, entities []*T) (*DeleteManyResponse, *Response, error) {
	path := "delete"
	return NewRequestBuilder[DeleteManyResponse](s.Endpoint, ctx).WithPath(path).WithBody(entities).Post()
}

type endpointUpdate[T any] struct{ Endpoint }

func (s *endpointUpdate[T]) Update(ctx context.Context, id *uuid.UUID, entity *T, params *Params) (*T, *Response, error) {
	path := id.String()
	return NewRequestBuilder[T](s.Endpoint, ctx).WithParams(params).WithPath(path).WithBody(entity).Put()
}

type endpointAccounts struct{ Endpoint }

// GetAccounts Получить все счета
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-scheta-kontragenta
func (s *endpointAccounts) GetAccounts(ctx context.Context, id *uuid.UUID) (*List[AgentAccount], *Response, error) {
	path := fmt.Sprintf("%s/accounts", id)
	return NewRequestBuilder[List[AgentAccount]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetAccountById Получить отдельный счёт
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-schet-kontragenta
func (s *endpointAccounts) GetAccountById(ctx context.Context, id, accountId uuid.UUID) (*AgentAccount, *Response, error) {
	path := fmt.Sprintf("%s/accounts/%s", id, accountId)
	return NewRequestBuilder[AgentAccount](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateAccounts Изменить счета
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico-izmenit-scheta-urlica
func (s *endpointAccounts) UpdateAccounts(ctx context.Context, id *uuid.UUID, accounts []*AgentAccount) (*Slice[AgentAccount], *Response, error) {
	path := fmt.Sprintf("%s/accounts", id)
	return NewRequestBuilder[Slice[AgentAccount]](s.Endpoint, ctx).WithPath(path).WithBody(accounts).Post()
}

type endpointAttributes struct{ Endpoint }

// GetAttributes Получить все дополнительные поля для указанного типа
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-poluchit-wse-dopolnitel-nye-polq-dlq-ukazannogo-tipa
func (s *endpointAttributes) GetAttributes(ctx context.Context) (*MetaArray[Attribute], *Response, error) {
	path := "metadata/attributes"
	return NewRequestBuilder[MetaArray[Attribute]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetAttributeById Получить дополнительное поле
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-poluchit-dopolnitel-noe-pole
func (s *endpointAttributes) GetAttributeById(ctx context.Context, id *uuid.UUID) (*Attribute, *Response, error) {
	path := fmt.Sprintf("metadata/attributes/%s", id)
	return NewRequestBuilder[Attribute](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateAttribute Создать дополнительное поле
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (s *endpointAttributes) CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *Response, error) {
	path := "metadata/attributes"
	return NewRequestBuilder[Attribute](s.Endpoint, ctx).WithPath(path).WithBody(attribute).Post()
}

// CreateAttributes Создать дополнительные поля
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (s *endpointAttributes) CreateAttributes(ctx context.Context, attributes []*Attribute) (*Slice[Attribute], *Response, error) {
	path := "metadata/attributes"

	// при передаче массива из 1-го доп поля сервис возвращает 1 доп поле, а не массив доп полей.
	// если количество передаваемых доп полей равняется 1, то дополнительно оборачиваем в срез.
	if len(attributes) == 1 {
		attribute, response, err := NewRequestBuilder[Attribute](s.Endpoint, ctx).
			WithPath(path).WithBody(attributes[0]).Post()
		return &Slice[Attribute]{attribute}, response, err
	}
	return NewRequestBuilder[Slice[Attribute]](s.Endpoint, ctx).WithPath(path).WithBody(attributes).Post()
}

// UpdateAttribute Изменить дополнительное поле
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-izmenit-dopolnitel-noe-pole
func (s *endpointAttributes) UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *Response, error) {
	path := fmt.Sprintf("metadata/attributes/%s", id)
	return NewRequestBuilder[Attribute](s.Endpoint, ctx).WithPath(path).WithBody(attribute).Put()
}

// DeleteAttribute Удалить дополнительное поле
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-udalit-dopolnitel-noe-pole
func (s *endpointAttributes) DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("metadata/attributes/%s", id)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

// DeleteAttributes Удалить дополнительные поля
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-udalit-dopolnitel-nye-polq
func (s *endpointAttributes) DeleteAttributes(ctx context.Context, attributes []*Attribute) (*DeleteManyResponse, *Response, error) {
	path := "metadata/attributes/delete"
	return NewRequestBuilder[DeleteManyResponse](s.Endpoint, ctx).WithPath(path).WithBody(attributes).Post()
}

type endpointAudit struct{ Endpoint }

// GetAudit Запрос на получение событий по сущности с указанным id
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-sobytiq-po-suschnosti
func (s *endpointAudit) GetAudit(ctx context.Context, id *uuid.UUID, params *Params) (*List[AuditEvent], *Response, error) {
	path := fmt.Sprintf("%s/audit", id)
	return NewRequestBuilder[List[AuditEvent]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

type endpointFiles struct{ Endpoint }

// GetFiles Получить список Файлов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-poluchit-spisok-fajlow-operacii-nomenklatury-zadachi-ili-kontragenta
func (s *endpointFiles) GetFiles(ctx context.Context, id *uuid.UUID) (*MetaArray[File], *Response, error) {
	path := fmt.Sprintf("%s/files", id)
	return NewRequestBuilder[MetaArray[File]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateFile Добавить Файл
func (s *endpointFiles) CreateFile(ctx context.Context, id *uuid.UUID, file *File) (*Slice[File], *Response, error) {
	path := fmt.Sprintf("%s/files", id)
	return NewRequestBuilder[Slice[File]](s.Endpoint, ctx).WithPath(path).WithBody(file).Post()
}

// UpdateFiles Добавить/обновить Файлы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-dobawit-fajly-k-operacii-nomenklature-ili-kontragentu
func (s *endpointFiles) UpdateFiles(ctx context.Context, id *uuid.UUID, files []File) (*Slice[File], *Response, error) {
	path := fmt.Sprintf("%s/files", id)
	return NewRequestBuilder[Slice[File]](s.Endpoint, ctx).WithPath(path).WithBody(files).Post()
}

// DeleteFile Удалить Файл
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-udalit-fajl
func (s *endpointFiles) DeleteFile(ctx context.Context, id, fileId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/files/%s", id, fileId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

// DeleteFiles Удалить Файлы
func (s *endpointFiles) DeleteFiles(ctx context.Context, id *uuid.UUID, files []*File) (*DeleteManyResponse, *Response, error) {
	path := fmt.Sprintf("%s/files/delete", id)
	return NewRequestBuilder[DeleteManyResponse](s.Endpoint, ctx).WithPath(path).WithBody(files).Post()
}

type endpointImages struct{ Endpoint }

// GetImages Получить список Изображений
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-poluchit-spisok-izobrazhenij-towara-komplekta-i-modifikacii
func (s *endpointImages) GetImages(ctx context.Context, id *uuid.UUID) (*MetaArray[Image], *Response, error) {
	path := fmt.Sprintf("%s/images", id)
	return NewRequestBuilder[MetaArray[Image]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateImage Добавить Изображение
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-dobawit-izobrazhenie-k-towaru-komplektu-ili-modifikacii
func (s *endpointImages) CreateImage(ctx context.Context, id *uuid.UUID, image *Image) (*Slice[Image], *Response, error) {
	path := fmt.Sprintf("%s/images", id)
	return NewRequestBuilder[Slice[Image]](s.Endpoint, ctx).WithPath(path).WithBody(image).Post()
}

// UpdateImages Изменение списка Изображений
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-izmenenie-spiska-izobrazhenij-u-towara-komplekta-ili-modifikacii
func (s *endpointImages) UpdateImages(ctx context.Context, id *uuid.UUID, images []*Image) (*Slice[Image], *Response, error) {
	path := fmt.Sprintf("%s/images", id)
	return NewRequestBuilder[Slice[Image]](s.Endpoint, ctx).WithPath(path).WithBody(images).Post()
}

// DeleteImage Удалить Изображение
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-izobrazhenie
func (s *endpointImages) DeleteImage(ctx context.Context, id, imageId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/images/%s", id, imageId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

// DeleteImages Удалить несколько Изображений
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-gruppu-izobrazhenij
func (s *endpointImages) DeleteImages(ctx context.Context, id *uuid.UUID, images []*Image) (*DeleteManyResponse, *Response, error) {
	path := fmt.Sprintf("%s/images/delete", id)
	return NewRequestBuilder[DeleteManyResponse](s.Endpoint, ctx).WithPath(path).WithBody(images).Post()
}

type endpointNamedFilter struct{ Endpoint }

// GetNamedFilters Получить список фильтров
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-spisok-fil-trow
func (s *endpointNamedFilter) GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *Response, error) {
	path := "namedfilter"
	return NewRequestBuilder[List[NamedFilter]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// GetNamedFilterById Получить фильтр по id
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-fil-tr-po-id
func (s *endpointNamedFilter) GetNamedFilterById(ctx context.Context, id *uuid.UUID) (*NamedFilter, *Response, error) {
	path := fmt.Sprintf("namedfilter/%s", id)
	return NewRequestBuilder[NamedFilter](s.Endpoint, ctx).WithPath(path).Get()
}

type endpointPositions[T any] struct{ Endpoint }

// GetPositions Получить все позиции документа
// Флаг withStockPositions - получить остатки и себестоимость в позициях (поле Stock)
func (s *endpointPositions[T]) GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[T], *Response, error) {
	path := fmt.Sprintf("%s/positions", id)
	return NewRequestBuilder[MetaArray[T]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// GetPositionById Получение отдельной позиции
// Флаг withStockPositions - получить остатки и себестоимость в позициях (поле Stock)
func (s *endpointPositions[T]) GetPositionById(ctx context.Context, id, positionId uuid.UUID, params *Params) (*T, *Response, error) {
	path := fmt.Sprintf("%s/positions/%s", id, positionId)
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// UpdatePosition Обновление позиции
// Флаг withStockPositions - получить остатки и себестоимость в позициях (поле Stock)
func (s *endpointPositions[T]) UpdatePosition(ctx context.Context, id, positionId uuid.UUID, position *T, params *Params) (*T, *Response, error) {
	path := fmt.Sprintf("%s/positions/%s", id, positionId)
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).WithParams(params).WithBody(position).Put()
}

// CreatePosition Создание позиции документа
func (s *endpointPositions[T]) CreatePosition(ctx context.Context, id *uuid.UUID, position *T) (*T, *Response, error) {
	path := fmt.Sprintf("%s/positions", id)
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).WithBody(position).Post()
}

// CreatePositions Массово создаёт позиции документа.
func (s *endpointPositions[T]) CreatePositions(ctx context.Context, id *uuid.UUID, positions []*T) (*Slice[T], *Response, error) {
	path := fmt.Sprintf("%s/positions", id)
	return NewRequestBuilder[Slice[T]](s.Endpoint, ctx).WithPath(path).WithBody(positions).Post()
}

// DeletePosition Удаляет позицию документа.
func (s *endpointPositions[T]) DeletePosition(ctx context.Context, id, positionId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/positions/%s", id, positionId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

// GetPositionTrackingCodes Получить Коды маркировки позиции документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-poluchit-kody-markirowki-pozicii-dokumenta
func (s *endpointPositions[T]) GetPositionTrackingCodes(ctx context.Context, id, positionId uuid.UUID) (*MetaArray[TrackingCode], *Response, error) {
	path := fmt.Sprintf("%s/positions/%s/trackingCodes", id, positionId)
	return NewRequestBuilder[MetaArray[TrackingCode]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateOrUpdatePositionTrackingCodes Массовое создание и обновление Кодов маркировки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-sozdanie-i-obnowlenie-kodow-markirowki
func (s *endpointPositions[T]) CreateOrUpdatePositionTrackingCodes(ctx context.Context, id, positionId uuid.UUID, trackingCodes TrackingCodes) (*Slice[TrackingCode], *Response, error) {
	path := fmt.Sprintf("%s/positions/%s/trackingCodes", id, positionId)
	return NewRequestBuilder[Slice[TrackingCode]](s.Endpoint, ctx).WithPath(path).WithBody(trackingCodes).Post()
}

// DeletePositionTrackingCodes Массовое удаление Кодов маркировки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-udalenie-kodow-markirowki
func (s *endpointPositions[T]) DeletePositionTrackingCodes(ctx context.Context, id, positionId uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *Response, error) {
	path := fmt.Sprintf("%s/positions/%s/trackingCodes/delete", id, positionId)
	return NewRequestBuilder[DeleteManyResponse](s.Endpoint, ctx).WithPath(path).WithBody(trackingCodes).Post()
}

type endpointPrintDoc struct{ Endpoint }

// PrintDoc Запрос на печать документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-pechat-dokumentow-zapros-na-pechat
func (s *endpointPrintDoc) PrintDoc(ctx context.Context, id *uuid.UUID, printDocRequest *PrintDocRequest) (*PrintFile, *Response, error) {
	path := fmt.Sprintf("%s/export", id)
	rb := NewRequestBuilder[CustomTemplate](s.Endpoint, ctx).WithPath(path).WithBody(printDocRequest).setContentHeader()
	resp, err := rb.do(http.MethodPost)
	if err != nil {
		return nil, resp, err
	}

	file, err := GetFileFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return file, resp, err
}

type endpointPrintPrice struct{ Endpoint }

// PrintPrice Запрос на печать этикеток и ценников.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pechat-atiketok-i-cennikow
func (s *endpointPrintPrice) PrintPrice(ctx context.Context, id *uuid.UUID, printPriceArg *PrintPriceArg) (*PrintFile, *Response, error) {
	path := fmt.Sprintf("%s/export", id)
	rb := NewRequestBuilder[CustomTemplate](s.Endpoint, ctx).WithPath(path).WithBody(printPriceArg).setContentHeader()
	resp, err := rb.do(http.MethodPost)
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
func (s *endpointPublication) GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *Response, error) {
	path := fmt.Sprintf("%s/publication", id)
	return NewRequestBuilder[MetaArray[Publication]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetPublicationById Запрос на получение Публикации с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikaciu
func (s *endpointPublication) GetPublicationById(ctx context.Context, id, publicationId uuid.UUID) (*Publication, *Response, error) {
	path := fmt.Sprintf("%s/publication/%s", id, publicationId)
	return NewRequestBuilder[Publication](s.Endpoint, ctx).WithPath(path).Get()
}

// Publish Запрос на публикацию документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-sozdat-publikaciu
func (s *endpointPublication) Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *Response, error) {
	path := fmt.Sprintf("%s/publication", id)
	publication := new(Publication).SetTemplate(template)
	return NewRequestBuilder[Publication](s.Endpoint, ctx).WithPath(path).WithBody(publication).Post()
}

// DeletePublication Запрос на удаление Публикации с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-udalit-publikaciu
func (s *endpointPublication) DeletePublication(ctx context.Context, id, publicationId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/publication/%s", id, publicationId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

type endpointSettings[T any] struct{ Endpoint }

// GetSettings Запрос на получение настроек справочника.
func (s *endpointSettings[T]) GetSettings(ctx context.Context) (*T, *Response, error) {
	path := "settings"
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateSettings Изменить настройки справочника.
func (s *endpointSettings[T]) UpdateSettings(ctx context.Context, settings *T) (*T, *Response, error) {
	path := "settings"
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).WithBody(settings).Put()
}

type endpointStates struct{ Endpoint }

// GetStateById Запрос на получение статуса по id.
func (s *endpointStates) GetStateById(ctx context.Context, id *uuid.UUID) (*State, *Response, error) {
	path := fmt.Sprintf("metadata/states/%s", id)
	return NewRequestBuilder[State](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateState Создать новый статус.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-sozdat-status
func (s *endpointStates) CreateState(ctx context.Context, state *State) (*State, *Response, error) {
	path := "metadata/states"
	return NewRequestBuilder[State](s.Endpoint, ctx).WithPath(path).WithBody(state).Post()
}

// UpdateState Изменить существующий статус.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-izmenit-status
func (s *endpointStates) UpdateState(ctx context.Context, id *uuid.UUID, state *State) (*State, *Response, error) {
	path := fmt.Sprintf("metadata/states/%s", id)
	return NewRequestBuilder[State](s.Endpoint, ctx).WithPath(path).WithBody(state).Put()
}

// CreateOrUpdateStates Массовое создание и обновление Статусов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-massowoe-sozdanie-i-obnowlenie-statusow
func (s *endpointStates) CreateOrUpdateStates(ctx context.Context, id *uuid.UUID, states []*State) (*Slice[State], *Response, error) {
	path := fmt.Sprintf("metadata/states/%s", id)
	return NewRequestBuilder[Slice[State]](s.Endpoint, ctx).WithPath(path).WithBody(states).Post()
}

// DeleteState Запрос на удаление Статуса с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-udalit-status
func (s *endpointStates) DeleteState(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("metadata/states/%s", id)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

type endpointSyncId[T any] struct{ Endpoint }

func (s *endpointSyncId[T]) GetBySyncId(ctx context.Context, syncId uuid.UUID) (*T, *Response, error) {
	path := fmt.Sprintf("syncid/%s", syncId)
	return NewRequestBuilder[T](s.Endpoint, ctx).WithPath(path).Get()
}

func (s *endpointSyncId[T]) DeleteBySyncId(ctx context.Context, syncId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("syncid/%s", syncId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

type endpointTemplates struct{ Endpoint }

// GetEmbeddedTemplates Запрос на получение информации о стандартных шаблонах печатных форм для указанного типа сущности.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-standartnyh-shablonow
func (s *endpointTemplates) GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *Response, error) {
	path := "metadata/embeddedtemplate"
	return NewRequestBuilder[List[EmbeddedTemplate]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetEmbeddedTemplateById Запрос на получение информации об отдельном стандартном шаблоне печатной формы для указанного типа сущности по его id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-standartnyj-shablon
func (s *endpointTemplates) GetEmbeddedTemplateById(ctx context.Context, id *uuid.UUID) (*EmbeddedTemplate, *Response, error) {
	path := fmt.Sprintf("metadata/embeddedtemplate/%s", id)
	return NewRequestBuilder[EmbeddedTemplate](s.Endpoint, ctx).WithPath(path).Get()
}

// GetCustomTemplates Запрос на получение информации о пользовательских шаблонах печатных форм для указанного типа сущности.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-pol-zowatel-skih-shablonow
func (s *endpointTemplates) GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *Response, error) {
	path := "metadata/customtemplate"
	return NewRequestBuilder[List[CustomTemplate]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetCustomTemplateById Запрос на получение информации об отдельном пользовательском шаблоне печатной формы для указанного типа сущности по его id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-pol-zowatel-skij-shablon
func (s *endpointTemplates) GetCustomTemplateById(ctx context.Context, id *uuid.UUID) (*CustomTemplate, *Response, error) {
	path := fmt.Sprintf("metadata/customtemplate/%s", id)
	return NewRequestBuilder[CustomTemplate](s.Endpoint, ctx).WithPath(path).Get()
}

type endpointRemove struct{ Endpoint }

// Remove Запрос на перемещение документа с указанным id в корзину.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-udalenie-w-korzinu
func (s *endpointRemove) Remove(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/trash", id)
	resp, err := NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).do(http.MethodPost)
	if err != nil {
		return false, resp, err
	}
	ok := resp.StatusCode == http.StatusOK
	return ok, resp, nil
}
