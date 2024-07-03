package moysklad

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type mainService[E MetaOwner, P any, M any, S any] struct {
	endpointGetList[E]
	endpointCreate[E]
	endpointCreateUpdateMany[E]
	endpointDeleteMany[E]
	endpointDelete
	endpointGetByID[E]
	endpointUpdate[E]
	endpointMetadata[M]
	endpointAttributes
	endpointNamedFilter
	endpointImages
	endpointSyncID[E]
	endpointAudit
	endpointPrintLabel
	endpointPositions[P]
	endpointTemplate[E]
	endpointPublication
	endpointSettings[S]
	endpointPrintTemplates
	endpointTrash
	endpointPrintDocument
	endpointAccounts
	endpointStates
	endpointFiles
	endpointTemplateBased[E]
	endpointEvaluate[E]
}

func newMainService[E MetaOwner, P any, M any, S any](e Endpoint) *mainService[E, P, M, S] {
	return &mainService[E, P, M, S]{
		endpointGetList:          endpointGetList[E]{e},
		endpointCreate:           endpointCreate[E]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[E]{e},
		endpointDeleteMany:       endpointDeleteMany[E]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[E]{e},
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
		endpointPrintTemplates:   endpointPrintTemplates{e},
		endpointTrash:            endpointTrash{e},
		endpointPrintDocument:    endpointPrintDocument{e},
		endpointAccounts:         endpointAccounts{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
		endpointTemplate:         endpointTemplate[E]{e},
		endpointTemplateBased:    endpointTemplateBased[E]{e},
		endpointEvaluate:         endpointEvaluate[E]{e},
	}
}

type endpointGetList[T any] struct{ Endpoint }

// GetList выполняет запрос на получение объектов в виде списка.
func (endpoint *endpointGetList[T]) GetList(ctx context.Context, params ...*Params) (*List[T], *resty.Response, error) {
	return NewRequestBuilder[List[T]](endpoint.client, endpoint.uri).SetParams(params...).Get(ctx)
}

type endpointDelete struct{ Endpoint }

// Delete выполняет запрос на удаление объекта по ID.
func (endpoint *endpointDelete) Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointGetByID[T any] struct{ Endpoint }

// GetByID выполняет запрос на получение отдельного объекта по ID.
func (endpoint *endpointGetByID[T]) GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", endpoint.uri, id)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Get(ctx)
}

type endpointMetadata[T any] struct{ Endpoint }

// GetMetadata выполняет запрос на получение метаданных объекта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-suschnosti
func (endpoint *endpointMetadata[T]) GetMetadata(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

type endpointTemplate[T any] struct{ Endpoint }

// Template выполняет запрос на получение предзаполненного стандартными полями объект.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-shablony-dokumentow
func (endpoint *endpointTemplate[T]) Template(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/new", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Put(ctx, nil)
}

type endpointTemplateBased[T any] struct{ Endpoint }

func templateBasedPrepare(based []MetaOwner) map[string]MetaWrapper {
	if len(based) == 0 {
		return nil
	}

	var body = make(map[string]MetaWrapper)
	for _, el := range based {
		metaType := el.GetMeta().GetType()
		if metaType == MetaTypeUnknown {
			continue
		}
		var metaTypeTemplate string
		switch metaType {
		case MetaTypeCustomerOrder:
			metaTypeTemplate = "customerOrder"
		case MetaTypeRetailDemand:
			metaTypeTemplate = "retailDemand"
		case MetaTypePurchaseReturn:
			metaTypeTemplate = "purchaseReturn"
		case MetaTypeInvoiceOut:
			metaTypeTemplate = "invoiceOut"
		case MetaTypeInvoiceIn:
			metaTypeTemplate = "invoiceIn"
		case MetaTypeCommissionReportOut:
			metaTypeTemplate = "commissionReportOut"
		case MetaTypeCommissionReportIn:
			metaTypeTemplate = "commissionReportIn"
		case MetaTypeProcessingPlan:
			metaTypeTemplate = "processingPlan"
		case MetaTypeProcessingOrder:
			metaTypeTemplate = "processingOrder"
		case MetaTypeInternalOrder:
			metaTypeTemplate = "internalOrder"
		case MetaTypeSalesReturn:
			metaTypeTemplate = "salesReturn"
		case MetaTypeRetailShift:
			metaTypeTemplate = "retailShift"
		case MetaTypePaymentOut:
			metaTypeTemplate = "paymentOut"
		case MetaTypePaymentIn:
			metaTypeTemplate = "paymentIn"

		default:
			metaTypeTemplate = metaType.String()
		}

		body[metaTypeTemplate] = el.GetMeta().Wrap()
	}

	return body
}

// TemplateBased выполняет запрос на получение предзаполненного стандартными полями объект на основе переданных документов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-shablony-dokumentow
func (endpoint *endpointTemplateBased[T]) TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/new", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Put(ctx, templateBasedPrepare(basedOn))
}

type endpointCreate[T any] struct{ Endpoint }

// Create выполняет запрос на создание объекта.
func (endpoint *endpointCreate[T]) Create(ctx context.Context, entity *T, params ...*Params) (*T, *resty.Response, error) {
	return NewRequestBuilder[T](endpoint.client, endpoint.uri).SetParams(params...).Post(ctx, entity)
}

// DeleteManyResponse объект ответа на запрос удаления нескольких объектов.
type DeleteManyResponse []struct {
	Info      string    `json:"info"`
	ApiErrors ApiErrors `json:"errors"`
}

type endpointDeleteMany[T MetaOwner] struct{ Endpoint }

// DeleteMany выполняет запрос на удаление нескольких объектов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (endpoint *endpointDeleteMany[T]) DeleteMany(ctx context.Context, entities ...*T) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/delete", endpoint.uri)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(entities))
}

type endpointCreateUpdateMany[T any] struct{ Endpoint }

// CreateUpdateMany выполняет запрос на создание и/или изменение нескольких объектов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (endpoint *endpointCreateUpdateMany[T]) CreateUpdateMany(ctx context.Context, entities Slice[T], params ...*Params) (*Slice[T], *resty.Response, error) {
	return NewRequestBuilder[Slice[T]](endpoint.client, endpoint.uri).SetParams(params...).Post(ctx, entities)
}

type endpointUpdate[T any] struct{ Endpoint }

// Update выполняет запрос на изменение объекта.
func (endpoint *endpointUpdate[T]) Update(ctx context.Context, id uuid.UUID, entity *T, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", endpoint.uri, id)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Put(ctx, entity)
}

type endpointAccounts struct{ Endpoint }

// GetAccountList выполняет запрос на получение всех счетов в виде списка.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-scheta-kontragenta
func (endpoint *endpointAccounts) GetAccountList(ctx context.Context, id uuid.UUID) (*List[AgentAccount], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts", endpoint.uri, id)
	return NewRequestBuilder[List[AgentAccount]](endpoint.client, path).Get(ctx)
}

// GetAccountByID выполняет запрос на получение отдельного счёта по ID.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-schet-kontragenta
func (endpoint *endpointAccounts) GetAccountByID(ctx context.Context, id, accountID uuid.UUID) (*AgentAccount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts/%s", endpoint.uri, id, accountID)
	return NewRequestBuilder[AgentAccount](endpoint.client, path).Get(ctx)
}

// UpdateAccountMany выполняет запрос на изменение списка счетов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico-izmenit-scheta-urlica
func (endpoint *endpointAccounts) UpdateAccountMany(ctx context.Context, id uuid.UUID, accounts ...*AgentAccount) (*MetaArray[AgentAccount], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts", endpoint.uri, id)
	return NewRequestBuilder[MetaArray[AgentAccount]](endpoint.client, path).Post(ctx, accounts)
}

type endpointAttributes struct{ Endpoint }

// GetAttributeList выполняет запрос на получение всех дополнительных полей объекта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-poluchit-wse-dopolnitel-nye-polq-dlq-ukazannogo-tipa
func (endpoint *endpointAttributes) GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", endpoint.uri)
	return NewRequestBuilder[List[Attribute]](endpoint.client, path).Get(ctx)
}

// GetAttributeByID выполняет запрос на получение отдельного дополнительного поле по ID.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-poluchit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%s", endpoint.uri, id)
	return NewRequestBuilder[Attribute](endpoint.client, path).Get(ctx)
}

// CreateAttribute выполняет запрос на создание дополнительного поля.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (endpoint *endpointAttributes) CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", endpoint.uri)
	return NewRequestBuilder[Attribute](endpoint.client, path).Post(ctx, attribute)
}

// CreateUpdateAttributeMany выполняет запрос на создание нескольких дополнительных полей.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (endpoint *endpointAttributes) CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", endpoint.uri)
	// при передаче массива из 1-го доп поля сервис возвращает 1 доп поле, а не массив доп полей.
	// если количество передаваемых доп полей равняется 1, то дополнительно оборачиваем в срез.
	if len(attributes) == 1 {
		attribute, resp, err := NewRequestBuilder[Attribute](endpoint.client, path).Post(ctx, attributes[0])
		return (&Slice[Attribute]{}).Push(attribute), resp, err
	}
	return NewRequestBuilder[Slice[Attribute]](endpoint.client, path).Post(ctx, attributes)
}

// UpdateAttribute выполняет запрос на изменение дополнительного поля.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-izmenit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%s", endpoint.uri, id)
	return NewRequestBuilder[Attribute](endpoint.client, path).Put(ctx, attribute)
}

// DeleteAttribute выполняет запрос на удаление дополнительного поля.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-udalit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%s", endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeleteAttributeMany выполняет запрос на удаление нескольких дополнительных полей.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-udalit-dopolnitel-nye-polq
func (endpoint *endpointAttributes) DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/delete", endpoint.uri)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(attributes))
}

type endpointAudit struct{ Endpoint }

// GetAudit выполняет запрос на получение событий сущности с указанным ID.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-sobytiq-po-suschnosti
func (endpoint *endpointAudit) GetAudit(ctx context.Context, id uuid.UUID, params ...*Params) (*List[AuditEvent], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/audit", endpoint.uri, id)
	return NewRequestBuilder[List[AuditEvent]](endpoint.client, path).SetParams(params...).Get(ctx)
}

type endpointFiles struct{ Endpoint }

// GetFileList выполняет запрос на получение файлов в виде списка.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-poluchit-spisok-fajlow-operacii-nomenklatury-zadachi-ili-kontragenta
func (endpoint *endpointFiles) GetFileList(ctx context.Context, id uuid.UUID) (*List[File], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files", endpoint.uri, id)
	return NewRequestBuilder[List[File]](endpoint.client, path).Get(ctx)
}

// CreateFile выполняет запрос на добавление файла.
func (endpoint *endpointFiles) CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files", endpoint.uri, id)
	return NewRequestBuilder[Slice[File]](endpoint.client, path).Post(ctx, file)
}

// UpdateFileMany выполняет запрос на массовое добавление/обновление файлов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-dobawit-fajly-k-operacii-nomenklature-ili-kontragentu
func (endpoint *endpointFiles) UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files", endpoint.uri, id)
	return NewRequestBuilder[Slice[File]](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(files))
}

// DeleteFile выполняет запрос на удаление файла.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-udalit-fajl
func (endpoint *endpointFiles) DeleteFile(ctx context.Context, id, fileID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files/%s", endpoint.uri, id, fileID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeleteFileMany выполняет запрос на удаление нескольких файлов.
func (endpoint *endpointFiles) DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files/delete", endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(files))
}

type endpointImages struct{ Endpoint }

// GetImageList выполняет запрос на получение изображений в виде списка.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-poluchit-spisok-izobrazhenij-towara-komplekta-i-modifikacii
func (endpoint *endpointImages) GetImageList(ctx context.Context, id uuid.UUID) (*List[Image], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images", endpoint.uri, id)
	return NewRequestBuilder[List[Image]](endpoint.client, path).Get(ctx)
}

// CreateImage выполняет запрос на добавление изображения.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-dobawit-izobrazhenie-k-towaru-komplektu-ili-modifikacii
func (endpoint *endpointImages) CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*Slice[Image], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images", endpoint.uri, id)
	return NewRequestBuilder[Slice[Image]](endpoint.client, path).Post(ctx, image)
}

// UpdateImageMany выполняет запрос на изменение списка изображений.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-izmenenie-spiska-izobrazhenij-u-towara-komplekta-ili-modifikacii
func (endpoint *endpointImages) UpdateImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*Slice[Image], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images", endpoint.uri, id)
	return NewRequestBuilder[Slice[Image]](endpoint.client, path).Post(ctx, images)
}

// DeleteImage выполняет запрос на удаление изображения.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-izobrazhenie
func (endpoint *endpointImages) DeleteImage(ctx context.Context, id, imageID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images/%s", endpoint.uri, id, imageID)
	return NewRequestBuilder[[]Image](endpoint.client, path).Delete(ctx)
}

// DeleteImageMany выполняет запрос на удаление нескольких изображений.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-gruppu-izobrazhenij
func (endpoint *endpointImages) DeleteImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images/delete", endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, images)
}

type endpointNamedFilter struct{ Endpoint }

// GetNamedFilterList выполняет запрос на получение фильтров в виде списка.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-spisok-fil-trow
func (endpoint *endpointNamedFilter) GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error) {
	path := fmt.Sprintf("%s/namedfilter", endpoint.uri)
	return NewRequestBuilder[List[NamedFilter]](endpoint.client, path).SetParams(params...).Get(ctx)
}

// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-fil-tr-po-id
func (endpoint *endpointNamedFilter) GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error) {
	path := fmt.Sprintf("%s/namedfilter/%s", endpoint.uri, id)
	return NewRequestBuilder[NamedFilter](endpoint.client, path).Get(ctx)
}

type endpointPositions[T any] struct{ Endpoint }

// GetPositionList выполняет запрос на получение всех позиций документа.
func (endpoint *endpointPositions[T]) GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[T], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", endpoint.uri, id)
	return NewRequestBuilder[List[T]](endpoint.client, path).SetParams(params...).Get(ctx)
}

// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
func (endpoint *endpointPositions[T]) GetPositionByID(ctx context.Context, id, positionID uuid.UUID, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", endpoint.uri, id, positionID)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Get(ctx)
}

// UpdatePosition выполняет запрос на изменение позиции документа.
func (endpoint *endpointPositions[T]) UpdatePosition(ctx context.Context, id, positionID uuid.UUID, position *T, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", endpoint.uri, id, positionID)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Put(ctx, position)
}

// CreatePosition выполняет запрос на создание позиции документа.
func (endpoint *endpointPositions[T]) CreatePosition(ctx context.Context, id uuid.UUID, position *T, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", endpoint.uri, id)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Post(ctx, position)
}

// CreatePositionMany выполняет запрос на массовое создание позиций документа.
func (endpoint *endpointPositions[T]) CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*T) (*Slice[T], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", endpoint.uri, id)
	return NewRequestBuilder[Slice[T]](endpoint.client, path).Post(ctx, positions)
}

// DeletePosition выполняет запрос на удаление позиции документа.
func (endpoint *endpointPositions[T]) DeletePosition(ctx context.Context, id, positionID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", endpoint.uri, id, positionID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeletePositionMany выполняет запрос на удаление нескольких позиций документа.
func (endpoint *endpointPositions[T]) DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*T) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/delete", endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, entities)
}

// GetPositionTrackingCodeList выполняет запрос на получение коды маркировки позиции документа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-poluchit-kody-markirowki-pozicii-dokumenta
func (endpoint *endpointPositions[T]) GetPositionTrackingCodeList(ctx context.Context, id, positionID uuid.UUID) (*List[TrackingCode], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes", endpoint.uri, id, positionID)
	return NewRequestBuilder[List[TrackingCode]](endpoint.client, path).Get(ctx)
}

// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание и/или изменение Кодов маркировки.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-sozdanie-i-obnowlenie-kodow-markirowki
func (endpoint *endpointPositions[T]) CreateUpdatePositionTrackingCodeMany(ctx context.Context, id, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes", endpoint.uri, id, positionID)
	return NewRequestBuilder[Slice[TrackingCode]](endpoint.client, path).Post(ctx, trackingCodes)
}

// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление Кодов маркировки.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-udalenie-kodow-markirowki
func (endpoint *endpointPositions[T]) DeletePositionTrackingCodeMany(ctx context.Context, id, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes/delete", endpoint.uri, id, positionID)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, trackingCodes)
}

type endpointPrintDocument struct{ Endpoint }

var reContentDisposition = regexp.MustCompile(`filename="(.*)"`)

func getFileFromResponse(resp *resty.Response) (*PrintFile, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.RawBody()); err != nil {
		return nil, err
	}

	var fileName string
	headerStr := resp.Header().Get(headerContentDisposition)
	if match := reContentDisposition.FindStringSubmatch(headerStr); len(match) > 1 {
		fileName = match[1]
	}
	file := &PrintFile{buf, fileName}
	return file, nil
}

// PrintDocument выполняет запрос на печать документа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-pechat-dokumentow-zapros-na-pechat
func (endpoint *endpointPrintDocument) PrintDocument(ctx context.Context, id uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/export", endpoint.uri, id)
	_, resp, err := NewRequestBuilder[PrintFile](endpoint.client, path).SetHeader(headerGetContent, "true").Post(ctx, PrintDocumentArg)
	if err != nil {
		return nil, resp, err
	}

	file, err := getFileFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return file, resp, err
}

type endpointPrintLabel struct{ Endpoint }

// PrintLabel выполняет запрос на печать этикеток и ценников.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pechat-atiketok-i-cennikow
func (endpoint *endpointPrintLabel) PrintLabel(ctx context.Context, id uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/export", endpoint.uri, id)

	_, resp, err := NewRequestBuilder[PrintFile](endpoint.client, path).
		SetHeader(headerGetContent, "true").
		Post(ctx, PrintLabelArg)

	if err != nil {
		return nil, resp, err
	}

	file, err := getFileFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return file, resp, err
}

type endpointPublication struct{ Endpoint }

// GetPublicationList выполняет запрос на получение списка Публикаций по указанному документу.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikacii
func (endpoint *endpointPublication) GetPublicationList(ctx context.Context, id uuid.UUID) (*List[Publication], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication", endpoint.uri, id)
	return NewRequestBuilder[List[Publication]](endpoint.client, path).Get(ctx)
}

// GetPublicationByID выполняет запрос на получение Публикации с указанным id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikaciu
func (endpoint *endpointPublication) GetPublicationByID(ctx context.Context, id, publicationID uuid.UUID) (*Publication, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication/%s", endpoint.uri, id, publicationID)
	return NewRequestBuilder[Publication](endpoint.client, path).Get(ctx)
}

// Publish выполняет запрос на публикацию документа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-sozdat-publikaciu
func (endpoint *endpointPublication) Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error) {
	publication := new(Publication).SetTemplate(template)
	path := fmt.Sprintf("%s/%s/publication", endpoint.uri, id)
	return NewRequestBuilder[Publication](endpoint.client, path).Post(ctx, publication)
}

// DeletePublication выполняет запрос на удаление Публикации с указанным id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-udalit-publikaciu
func (endpoint *endpointPublication) DeletePublication(ctx context.Context, id, publicationID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication/%s", endpoint.uri, id, publicationID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointSettings[T any] struct{ Endpoint }

// GetSettings выполняет запрос на получение настроек справочника.
func (endpoint *endpointSettings[T]) GetSettings(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

// UpdateSettings выполняет запрос на изменение настроек справочника.
func (endpoint *endpointSettings[T]) UpdateSettings(ctx context.Context, settings *T) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Put(ctx, settings)
}

type endpointStates struct{ Endpoint }

// GetStateByID выполняет запрос на получение статуса по ID.
func (endpoint *endpointStates) GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", endpoint.uri, id)
	return NewRequestBuilder[State](endpoint.client, path).Get(ctx)
}

// CreateState выполняет запрос на создание нового статуса.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-sozdat-status
func (endpoint *endpointStates) CreateState(ctx context.Context, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states", endpoint.uri)
	return NewRequestBuilder[State](endpoint.client, path).Post(ctx, state)
}

// UpdateState выполняет запрос на изменение существующего статуса.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-izmenit-status
func (endpoint *endpointStates) UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", endpoint.uri, id)
	return NewRequestBuilder[State](endpoint.client, path).Put(ctx, state)
}

// CreateUpdateStateMany выполняет запрос на массовое создание и/ли изменение Статусов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-massowoe-sozdanie-i-obnowlenie-statusow
func (endpoint *endpointStates) CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states", endpoint.uri)
	return NewRequestBuilder[Slice[State]](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(states))
}

// DeleteState выполняет запрос на удаление Статуса с указанным id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-udalit-status
func (endpoint *endpointStates) DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointSyncID[T any] struct{ Endpoint }

// GetBySyncID выполняет запрос на получение объекта по syncID.
func (endpoint *endpointSyncID[T]) GetBySyncID(ctx context.Context, syncID uuid.UUID) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/syncid/%s", endpoint.uri, syncID)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

// DeleteBySyncID выполняет запрос на удаление объекта по syncID.
func (endpoint *endpointSyncID[T]) DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/syncid/%s", endpoint.uri, syncID)
	return NewRequestBuilder[T](endpoint.client, path).Delete(ctx)
}

type endpointPrintTemplates struct{ Endpoint }

// GetEmbeddedTemplateList выполняет запрос на получение информации о стандартных шаблонах печатных форм для указанного типа сущности.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-standartnyh-shablonow
func (endpoint *endpointPrintTemplates) GetEmbeddedTemplateList(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate", endpoint.uri)
	return NewRequestBuilder[List[EmbeddedTemplate]](endpoint.client, path).Get(ctx)
}

// GetEmbeddedTemplateByID выполняет запрос на получение информации об отдельном стандартном шаблоне печатной формы для указанного типа сущности по его id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-standartnyj-shablon
func (endpoint *endpointPrintTemplates) GetEmbeddedTemplateByID(ctx context.Context, id uuid.UUID) (*EmbeddedTemplate, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate/%s", endpoint.uri, id)
	return NewRequestBuilder[EmbeddedTemplate](endpoint.client, path).Get(ctx)
}

// GetCustomTemplateList выполняет запрос на получение информации о пользовательских шаблонах печатных форм для указанного типа сущности.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-pol-zowatel-skih-shablonow
func (endpoint *endpointPrintTemplates) GetCustomTemplateList(ctx context.Context) (*List[CustomTemplate], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/customtemplate", endpoint.uri)
	return NewRequestBuilder[List[CustomTemplate]](endpoint.client, path).Get(ctx)
}

// GetCustomTemplateByID выполняет запрос на получение информации об отдельном пользовательском шаблоне печатной формы для указанного типа сущности по его id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-pol-zowatel-skij-shablon
func (endpoint *endpointPrintTemplates) GetCustomTemplateByID(ctx context.Context, id uuid.UUID) (*CustomTemplate, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate/%s", endpoint.uri, id)
	return NewRequestBuilder[CustomTemplate](endpoint.client, path).Get(ctx)
}

type endpointTrash struct{ Endpoint }

// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-udalenie-w-korzinu
func (endpoint *endpointTrash) MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/trash", endpoint.uri, id)
	_, resp, err := NewRequestBuilder[any](endpoint.client, path).Post(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}

type endpointEvaluate[T any] struct{ Endpoint }

// Evaluate выполняет запрос на автозаполнение.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-awtozapolnenie
func (endpoint *endpointEvaluate[T]) Evaluate(ctx context.Context, entity *T, evaluate ...Evaluate) (*T, *resty.Response, error) {
	uriParts := strings.Split(endpoint.uri, "/")
	path := fmt.Sprintf("wizard/%s", uriParts[len(uriParts)-1])
	params := NewParams().WithEvaluate(evaluate...)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params).Post(ctx, entity)
}
