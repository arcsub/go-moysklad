package moysklad

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

const (
	EndpointEntity   = "entity/"
	EndpointContext  = "context/"
	EndpointReport   = "report/"
	EndpointSecurity = "security/"
	EndpointDelete   = "/delete"
	EndpointSettings = "%s/settings"

	EndpointExport = "%s/%s/export"

	EndpointSyncID = "%s/syncid/%s"

	EndpointTrash = "%s/%s/trash"

	EndpointPublication   = "%s/%s/publication"
	EndpointPublicationID = EndpointPublication + "/%s"

	EndpointEmbeddedTemplate   = "%s/metadata/embeddedtemplate"
	EndpointEmbeddedTemplateID = EndpointEmbeddedTemplate + "/%s"

	EndpointCustomTemplate   = "%s/metadata/customtemplate"
	EndpointCustomTemplateID = EndpointCustomTemplate + "/%s"

	EndpointAccounts   = "%s/%s/accounts"
	EndpointAccountsID = EndpointAccounts + "/%s"

	EndpointStates   = "%s/metadata/states"
	EndpointStatesID = EndpointStates + "/%s"

	EndpointAttributes       = "%s/metadata/attributes"
	EndpointAttributesID     = EndpointAttributes + "/%s"
	EndpointAttributesDelete = EndpointAttributes + EndpointDelete

	EndpointFiles       = "%s/%s/files"
	EndpointFilesID     = EndpointFiles + "/%s"
	EndpointFilesDelete = EndpointFiles + EndpointDelete

	EndpointImages       = "%s/%s/images"
	EndpointImagesID     = EndpointFiles + "/%s"
	EndpointImagesDelete = EndpointFiles + EndpointDelete

	EndpointNamedFilter   = "%s/namedfilter"
	EndpointNamedFilterID = EndpointNamedFilter + "/%s"

	EndpointPositions       = "/%s/%s/positions"
	EndpointPositionsID     = EndpointPositions + "/%s"
	EndpointPositionsDelete = EndpointPositions + EndpointDelete

	EndpointTrackingCodes       = EndpointPositions + "/%s/trackingCodes"
	EndpointTrackingCodesDelete = EndpointTrackingCodes + EndpointDelete
)

type mainService[E MetaIDOwner, P any, M any, S any] struct {
	endpointGetList[E]
	endpointCreate[E]
	endpointCreateUpdateMany[E]
	endpointDeleteMany[E]
	endpointDeleteByID
	endpointDelete[E]
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

func newMainService[E MetaIDOwner, P any, M any, S any](client *Client, path string) *mainService[E, P, M, S] {
	endpoint := NewEndpoint(client, path)

	return &mainService[E, P, M, S]{
		endpointGetList:          endpointGetList[E]{endpoint},
		endpointCreate:           endpointCreate[E]{endpoint},
		endpointCreateUpdateMany: endpointCreateUpdateMany[E]{endpoint},
		endpointDeleteMany:       endpointDeleteMany[E]{endpoint},
		endpointDeleteByID:       endpointDeleteByID{endpoint},
		endpointDelete:           endpointDelete[E]{endpoint},
		endpointGetByID:          endpointGetByID[E]{endpoint},
		endpointUpdate:           endpointUpdate[E]{endpoint},
		endpointMetadata:         endpointMetadata[M]{endpoint},
		endpointAttributes:       endpointAttributes{endpoint},
		endpointNamedFilter:      endpointNamedFilter{endpoint},
		endpointImages:           endpointImages{endpoint},
		endpointSyncID:           endpointSyncID[E]{endpoint},
		endpointAudit:            endpointAudit{endpoint},
		endpointPrintLabel:       endpointPrintLabel{endpoint},
		endpointPositions:        endpointPositions[P]{endpoint},
		endpointPublication:      endpointPublication{endpoint},
		endpointSettings:         endpointSettings[S]{endpoint},
		endpointPrintTemplates:   endpointPrintTemplates{endpoint},
		endpointTrash:            endpointTrash{endpoint},
		endpointPrintDocument:    endpointPrintDocument{endpoint},
		endpointAccounts:         endpointAccounts{endpoint},
		endpointStates:           endpointStates{endpoint},
		endpointFiles:            endpointFiles{endpoint},
		endpointTemplate:         endpointTemplate[E]{endpoint},
		endpointTemplateBased:    endpointTemplateBased[E]{endpoint},
		endpointEvaluate:         endpointEvaluate[E]{endpoint},
	}
}

func getAll[T any](ctx context.Context, client *Client, path string, params []*Params) (*Slice[T], *resty.Response, error) {
	var offset = 1
	var perPage = MaxPositions
	var data Slice[T]
	var mu sync.Mutex
	var wg sync.WaitGroup

	paramsCpy := GetParamsFromSliceOrNew(params).WithLimit(offset).WithOffset(0)

	list, resp, err := NewRequestBuilder[List[T]](client, path).SetParams(paramsCpy).Get(ctx)
	if err != nil {
		return nil, resp, err
	}

	data = append(data, list.Rows...)
	size := list.Meta.Size

	if len(paramsCpy.Expand) > 0 {
		perPage = 100
	}

	for i := offset; i < size; i += perPage {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			paramsCpy := paramsCpy.Clone().WithLimit(perPage).WithOffset(i)

			list, resResp, err := NewRequestBuilder[List[T]](client, path).SetParams(paramsCpy).Get(ctx)

			mu.Lock()
			resp = resResp
			mu.Unlock()

			if err != nil {
				log.Println("getAll error:", err)
				return
			}

			mu.Lock()
			data = append(data, list.Rows...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	return &data, resp, nil
}

func posAll[T any](ctx context.Context, client *Client, path string, entities Slice[T], params []*Params) (*Slice[T], *resty.Response, error) {
	paramsCpy := GetParamsFromSliceOrNew(params)

	if entities.Len() > MaxPositions {
		var data Slice[T]
		var resp *resty.Response
		var mu sync.Mutex
		var wg sync.WaitGroup

		for _, chunk := range entities.IntoChunks(MaxPositions) {
			wg.Add(1)

			go func(chunk Slice[T]) {
				defer wg.Done()

				list, resResp, err := NewRequestBuilder[Slice[T]](client, path).SetParams(paramsCpy).Post(ctx, chunk)

				mu.Lock()
				resp = resResp
				mu.Unlock()

				if err != nil {
					log.Println("postAll error:", err)
					return
				}

				if list.Len() > 0 {
					mu.Lock()
					data.Push(list.S()...)
					mu.Unlock()
				}
			}(chunk)
		}

		wg.Wait()

		return &data, resp, nil
	}

	return NewRequestBuilder[Slice[T]](client, path).SetParams(paramsCpy).Post(ctx, entities)
}

func deleteAll[T MetaOwner](ctx context.Context, client *Client, path string, entities Slice[T]) (*DeleteManyResponse, *resty.Response, error) {
	if entities.Len() > MaxPositions {
		var data DeleteManyResponse
		var resp *resty.Response
		var mu sync.Mutex
		var wg sync.WaitGroup

		for _, chunk := range entities.IntoChunks(MaxPositions) {
			wg.Add(1)

			go func(chunk Slice[T], resp *resty.Response) {
				defer wg.Done()

				list, resResp, err := NewRequestBuilder[DeleteManyResponse](client, path).Post(ctx, AsMetaWrapperSlice(chunk))

				mu.Lock()
				resp = resResp
				mu.Unlock()

				if err != nil {
					log.Println("deleteAll error:", err)
					return
				}

				mu.Lock()
				data = append(data, Deref(list)...)
				mu.Unlock()
			}(chunk, resp)
		}

		wg.Wait()

		return &data, resp, nil
	}

	return NewRequestBuilder[DeleteManyResponse](client, path).Post(ctx, AsMetaWrapperSlice(entities))
}

type endpointGetList[T any] struct{ Endpoint }

// GetList выполняет запрос на получение объектов в виде списка.
func (endpoint *endpointGetList[T]) GetList(ctx context.Context, params ...*Params) (*List[T], *resty.Response, error) {
	return NewRequestBuilder[List[T]](endpoint.client, endpoint.uri).SetParams(params...).Get(ctx)
}

// GetListAll выполняет запрос на получение всех объектов в виде списка.
func (endpoint *endpointGetList[T]) GetListAll(ctx context.Context, params ...*Params) (*Slice[T], *resty.Response, error) {
	return getAll[T](ctx, endpoint.client, endpoint.uri, params)
}

type endpointDeleteByID struct{ Endpoint }

// DeleteByID выполняет запрос на удаление объекта по ID.
func (endpoint *endpointDeleteByID) DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointDelete[T MetaIDOwner] struct{ Endpoint }

// Delete выполняет запрос на удаление объекта.
func (endpoint *endpointDelete[T]) Delete(ctx context.Context, entity *T) (bool, *resty.Response, error) {
	id := GetUUIDFromEntity(entity)
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
type DeleteManyResponse []DeleteManyRow

type DeleteManyRow struct {
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
	return deleteAll[T](ctx, endpoint.client, path, entities)
}

type endpointCreateUpdateMany[T any] struct{ Endpoint }

// CreateUpdateMany выполняет запрос на создание и/или изменение нескольких объектов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (endpoint *endpointCreateUpdateMany[T]) CreateUpdateMany(ctx context.Context, entities Slice[T], params ...*Params) (*Slice[T], *resty.Response, error) {
	return posAll[T](ctx, endpoint.client, endpoint.uri, entities, params)
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
	path := fmt.Sprintf(EndpointAccounts, endpoint.uri, id)
	return NewRequestBuilder[List[AgentAccount]](endpoint.client, path).Get(ctx)
}

// GetAccountByID выполняет запрос на получение отдельного счёта по ID.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-schet-kontragenta
func (endpoint *endpointAccounts) GetAccountByID(ctx context.Context, id, accountID uuid.UUID) (*AgentAccount, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAccountsID, endpoint.uri, id, accountID)
	return NewRequestBuilder[AgentAccount](endpoint.client, path).Get(ctx)
}

// UpdateAccountMany выполняет запрос на изменение списка счетов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico-izmenit-scheta-urlica
func (endpoint *endpointAccounts) UpdateAccountMany(ctx context.Context, id uuid.UUID, accounts ...*AgentAccount) (*MetaArray[AgentAccount], *resty.Response, error) {
	path := fmt.Sprintf(EndpointAccounts, endpoint.uri, id)
	return NewRequestBuilder[MetaArray[AgentAccount]](endpoint.client, path).Post(ctx, accounts)
}

type endpointAttributes struct{ Endpoint }

// GetAttributeList выполняет запрос на получение всех дополнительных полей объекта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-poluchit-wse-dopolnitel-nye-polq-dlq-ukazannogo-tipa
func (endpoint *endpointAttributes) GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error) {
	path := fmt.Sprintf(EndpointAttributes, endpoint.uri)
	return NewRequestBuilder[List[Attribute]](endpoint.client, path).Get(ctx)
}

// GetAttributeByID выполняет запрос на получение отдельного дополнительного поле по ID.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-poluchit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAttributesID, endpoint.uri, id)
	return NewRequestBuilder[Attribute](endpoint.client, path).Get(ctx)
}

// CreateAttribute выполняет запрос на создание дополнительного поля.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (endpoint *endpointAttributes) CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAttributes, endpoint.uri)
	return NewRequestBuilder[Attribute](endpoint.client, path).Post(ctx, attribute)
}

// CreateUpdateAttributeMany выполняет запрос на создание нескольких дополнительных полей.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (endpoint *endpointAttributes) CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error) {
	path := fmt.Sprintf(EndpointAttributes, endpoint.uri)
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
	path := fmt.Sprintf(EndpointAttributesID, endpoint.uri, id)
	return NewRequestBuilder[Attribute](endpoint.client, path).Put(ctx, attribute)
}

// DeleteAttribute выполняет запрос на удаление дополнительного поля.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-udalit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAttributesID, endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeleteAttributeMany выполняет запрос на удаление нескольких дополнительных полей.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-udalit-dopolnitel-nye-polq
func (endpoint *endpointAttributes) DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAttributesDelete, endpoint.uri)
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
	path := fmt.Sprintf(EndpointFiles, endpoint.uri, id)
	return NewRequestBuilder[List[File]](endpoint.client, path).Get(ctx)
}

// CreateFile выполняет запрос на добавление файла.
func (endpoint *endpointFiles) CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error) {
	path := fmt.Sprintf(EndpointFiles, endpoint.uri, id)
	return NewRequestBuilder[Slice[File]](endpoint.client, path).Post(ctx, file)
}

// UpdateFileMany выполняет запрос на массовое добавление/обновление файлов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-dobawit-fajly-k-operacii-nomenklature-ili-kontragentu
func (endpoint *endpointFiles) UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error) {
	path := fmt.Sprintf(EndpointFiles, endpoint.uri, id)
	return NewRequestBuilder[Slice[File]](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(files))
}

// DeleteFile выполняет запрос на удаление файла.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-udalit-fajl
func (endpoint *endpointFiles) DeleteFile(ctx context.Context, id, fileID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointFilesID, endpoint.uri, id, fileID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeleteFileMany выполняет запрос на удаление нескольких файлов.
func (endpoint *endpointFiles) DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointFilesDelete, endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(files))
}

type endpointImages struct{ Endpoint }

// GetImageList выполняет запрос на получение изображений в виде списка.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-poluchit-spisok-izobrazhenij-towara-komplekta-i-modifikacii
func (endpoint *endpointImages) GetImageList(ctx context.Context, id uuid.UUID) (*List[Image], *resty.Response, error) {
	path := fmt.Sprintf(EndpointImages, endpoint.uri, id)
	return NewRequestBuilder[List[Image]](endpoint.client, path).Get(ctx)
}

// CreateImage выполняет запрос на добавление изображения.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-dobawit-izobrazhenie-k-towaru-komplektu-ili-modifikacii
func (endpoint *endpointImages) CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*Slice[Image], *resty.Response, error) {
	path := fmt.Sprintf(EndpointImages, endpoint.uri, id)
	return NewRequestBuilder[Slice[Image]](endpoint.client, path).Post(ctx, image)
}

// UpdateImageMany выполняет запрос на изменение списка изображений.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-izmenenie-spiska-izobrazhenij-u-towara-komplekta-ili-modifikacii
func (endpoint *endpointImages) UpdateImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*Slice[Image], *resty.Response, error) {
	path := fmt.Sprintf(EndpointImages, endpoint.uri, id)
	return NewRequestBuilder[Slice[Image]](endpoint.client, path).Post(ctx, images)
}

// DeleteImage выполняет запрос на удаление изображения.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-izobrazhenie
func (endpoint *endpointImages) DeleteImage(ctx context.Context, id, imageID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointImagesID, endpoint.uri, id, imageID)
	return NewRequestBuilder[[]Image](endpoint.client, path).Delete(ctx)
}

// DeleteImageMany выполняет запрос на удаление нескольких изображений.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-gruppu-izobrazhenij
func (endpoint *endpointImages) DeleteImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointImagesDelete, endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, images)
}

type endpointNamedFilter struct{ Endpoint }

// GetNamedFilterList выполняет запрос на получение фильтров в виде списка.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-spisok-fil-trow
func (endpoint *endpointNamedFilter) GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error) {
	path := fmt.Sprintf(EndpointNamedFilter, endpoint.uri)
	return NewRequestBuilder[List[NamedFilter]](endpoint.client, path).SetParams(params...).Get(ctx)
}

// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-fil-tr-po-id
func (endpoint *endpointNamedFilter) GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error) {
	path := fmt.Sprintf(EndpointNamedFilterID, endpoint.uri, id)
	return NewRequestBuilder[NamedFilter](endpoint.client, path).Get(ctx)
}

type endpointPositions[T any] struct{ Endpoint }

// GetPositionList выполняет запрос на получение всех позиций документа.
func (endpoint *endpointPositions[T]) GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[T], *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositions, endpoint.uri, id)
	return NewRequestBuilder[List[T]](endpoint.client, path).SetParams(params...).Get(ctx)
}

func (endpoint *endpointPositions[T]) GetPositionListAll(ctx context.Context, id uuid.UUID, params ...*Params) (*Slice[T], *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositions, endpoint.uri, id)
	return getAll[T](ctx, endpoint.client, path, params)
}

// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
func (endpoint *endpointPositions[T]) GetPositionByID(ctx context.Context, id, positionID uuid.UUID, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositionsID, endpoint.uri, id, positionID)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Get(ctx)
}

// UpdatePosition выполняет запрос на изменение позиции документа.
func (endpoint *endpointPositions[T]) UpdatePosition(ctx context.Context, id, positionID uuid.UUID, position *T, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositionsID, endpoint.uri, id, positionID)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Put(ctx, position)
}

// CreatePosition выполняет запрос на создание позиции документа.
func (endpoint *endpointPositions[T]) CreatePosition(ctx context.Context, id uuid.UUID, position *T, params ...*Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositions, endpoint.uri, id)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params...).Post(ctx, position)
}

// CreatePositionMany выполняет запрос на массовое создание позиций документа.
func (endpoint *endpointPositions[T]) CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*T) (*Slice[T], *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositions, endpoint.uri, id)
	return posAll[T](ctx, endpoint.client, path, positions, nil)
}

// DeletePosition выполняет запрос на удаление позиции документа.
func (endpoint *endpointPositions[T]) DeletePosition(ctx context.Context, id, positionID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositionsID, endpoint.uri, id, positionID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeletePositionMany выполняет запрос на удаление нескольких позиций документа.
func (endpoint *endpointPositions[T]) DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*T) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPositionsDelete, endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, entities)
}

// GetPositionTrackingCodeList выполняет запрос на получение коды маркировки позиции документа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-poluchit-kody-markirowki-pozicii-dokumenta
func (endpoint *endpointPositions[T]) GetPositionTrackingCodeList(ctx context.Context, id, positionID uuid.UUID) (*List[TrackingCode], *resty.Response, error) {
	path := fmt.Sprintf(EndpointTrackingCodes, endpoint.uri, id, positionID)
	return NewRequestBuilder[List[TrackingCode]](endpoint.client, path).Get(ctx)
}

// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание и/или изменение Кодов маркировки.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-sozdanie-i-obnowlenie-kodow-markirowki
func (endpoint *endpointPositions[T]) CreateUpdatePositionTrackingCodeMany(ctx context.Context, id, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error) {
	path := fmt.Sprintf(EndpointTrackingCodes, endpoint.uri, id, positionID)
	return NewRequestBuilder[Slice[TrackingCode]](endpoint.client, path).Post(ctx, trackingCodes)
}

// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление Кодов маркировки.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-udalenie-kodow-markirowki
func (endpoint *endpointPositions[T]) DeletePositionTrackingCodeMany(ctx context.Context, id, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointTrackingCodesDelete, endpoint.uri, id, positionID)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, trackingCodes)
}

type endpointPrintDocument struct{ Endpoint }

var reContentDisposition = regexp.MustCompile(`filename="(.*)"`)

func printFileFromResp(resp *resty.Response) (*PrintFile, *resty.Response, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.RawBody()); err != nil {
		return nil, resp, err
	}

	var fileName string
	headerStr := resp.Header().Get(headerContentDisposition)
	if match := reContentDisposition.FindStringSubmatch(headerStr); len(match) > 1 {
		fileName = match[1]
	}
	file := &PrintFile{buf, fileName}
	return file, resp, nil
}

// PrintDocument выполняет запрос на печать документа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-pechat-dokumentow-zapros-na-pechat
func (endpoint *endpointPrintDocument) PrintDocument(ctx context.Context, id uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf(EndpointExport, endpoint.uri, id)
	_, resp, err := NewRequestBuilder[PrintFile](endpoint.client, path).SetHeader(headerGetContent, "true").Post(ctx, PrintDocumentArg)
	if err != nil {
		return nil, resp, err
	}
	return printFileFromResp(resp)
}

type endpointPrintLabel struct{ Endpoint }

// PrintLabel выполняет запрос на печать этикеток и ценников.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pechat-atiketok-i-cennikow
func (endpoint *endpointPrintLabel) PrintLabel(ctx context.Context, id uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf(EndpointExport, endpoint.uri, id)

	_, resp, err := NewRequestBuilder[PrintFile](endpoint.client, path).
		SetHeader(headerGetContent, "true").
		Post(ctx, PrintLabelArg)

	if err != nil {
		return nil, resp, err
	}
	return printFileFromResp(resp)
}

type endpointPublication struct{ Endpoint }

// GetPublicationList выполняет запрос на получение списка Публикаций по указанному документу.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikacii
func (endpoint *endpointPublication) GetPublicationList(ctx context.Context, id uuid.UUID) (*List[Publication], *resty.Response, error) {
	path := fmt.Sprintf(EndpointPublication, endpoint.uri, id)
	return NewRequestBuilder[List[Publication]](endpoint.client, path).Get(ctx)
}

// GetPublicationByID выполняет запрос на получение Публикации с указанным id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikaciu
func (endpoint *endpointPublication) GetPublicationByID(ctx context.Context, id, publicationID uuid.UUID) (*Publication, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPublicationID, endpoint.uri, id, publicationID)
	return NewRequestBuilder[Publication](endpoint.client, path).Get(ctx)
}

// Publish выполняет запрос на публикацию документа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-sozdat-publikaciu
func (endpoint *endpointPublication) Publish(ctx context.Context, id uuid.UUID, template TemplateConverter) (*Publication, *resty.Response, error) {
	if template == nil {
		return nil, nil, fmt.Errorf("publish: template is empty")
	}

	path := fmt.Sprintf(EndpointPublication, endpoint.uri, id)
	return NewRequestBuilder[Publication](endpoint.client, path).Post(ctx, template.AsTemplate())
}

// DeletePublication выполняет запрос на удаление Публикации с указанным id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-udalit-publikaciu
func (endpoint *endpointPublication) DeletePublication(ctx context.Context, id, publicationID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPublicationID, endpoint.uri, id, publicationID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointSettings[T any] struct{ Endpoint }

// GetSettings выполняет запрос на получение настроек справочника.
func (endpoint *endpointSettings[T]) GetSettings(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf(EndpointSettings, endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

// UpdateSettings выполняет запрос на изменение настроек справочника.
func (endpoint *endpointSettings[T]) UpdateSettings(ctx context.Context, settings *T) (*T, *resty.Response, error) {
	path := fmt.Sprintf(EndpointSettings, endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Put(ctx, settings)
}

type endpointStates struct{ Endpoint }

// GetStateByID выполняет запрос на получение статуса по ID.
func (endpoint *endpointStates) GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStatesID, endpoint.uri, id)
	return NewRequestBuilder[State](endpoint.client, path).Get(ctx)
}

// CreateState выполняет запрос на создание нового статуса.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-sozdat-status
func (endpoint *endpointStates) CreateState(ctx context.Context, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStates, endpoint.uri)
	return NewRequestBuilder[State](endpoint.client, path).Post(ctx, state)
}

// UpdateState выполняет запрос на изменение существующего статуса.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-izmenit-status
func (endpoint *endpointStates) UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStatesID, endpoint.uri, id)
	return NewRequestBuilder[State](endpoint.client, path).Put(ctx, state)
}

// CreateUpdateStateMany выполняет запрос на массовое создание и/ли изменение Статусов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-massowoe-sozdanie-i-obnowlenie-statusow
func (endpoint *endpointStates) CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error) {
	path := fmt.Sprintf(EndpointStates, endpoint.uri)
	return NewRequestBuilder[Slice[State]](endpoint.client, path).Post(ctx, AsMetaWrapperSlice(states))
}

// DeleteState выполняет запрос на удаление Статуса с указанным id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-udalit-status
func (endpoint *endpointStates) DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStatesID, endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointSyncID[T any] struct{ Endpoint }

// GetBySyncID выполняет запрос на получение объекта по syncID.
func (endpoint *endpointSyncID[T]) GetBySyncID(ctx context.Context, syncID uuid.UUID) (*T, *resty.Response, error) {
	path := fmt.Sprintf(EndpointSyncID, endpoint.uri, syncID)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

// DeleteBySyncID выполняет запрос на удаление объекта по syncID.
func (endpoint *endpointSyncID[T]) DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointSyncID, endpoint.uri, syncID)
	return NewRequestBuilder[T](endpoint.client, path).Delete(ctx)
}

type endpointPrintTemplates struct{ Endpoint }

// GetEmbeddedTemplateList выполняет запрос на получение информации о стандартных шаблонах печатных форм для указанного типа сущности.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-standartnyh-shablonow
func (endpoint *endpointPrintTemplates) GetEmbeddedTemplateList(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error) {
	path := fmt.Sprintf(EndpointEmbeddedTemplate, endpoint.uri)
	return NewRequestBuilder[List[EmbeddedTemplate]](endpoint.client, path).Get(ctx)
}

// GetEmbeddedTemplateByID выполняет запрос на получение информации об отдельном стандартном шаблоне печатной формы для указанного типа сущности по его id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-standartnyj-shablon
func (endpoint *endpointPrintTemplates) GetEmbeddedTemplateByID(ctx context.Context, id uuid.UUID) (*EmbeddedTemplate, *resty.Response, error) {
	path := fmt.Sprintf(EndpointEmbeddedTemplateID, endpoint.uri, id)
	return NewRequestBuilder[EmbeddedTemplate](endpoint.client, path).Get(ctx)
}

// GetCustomTemplateList выполняет запрос на получение информации о пользовательских шаблонах печатных форм для указанного типа сущности.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-pol-zowatel-skih-shablonow
func (endpoint *endpointPrintTemplates) GetCustomTemplateList(ctx context.Context) (*List[CustomTemplate], *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomTemplate, endpoint.uri)
	return NewRequestBuilder[List[CustomTemplate]](endpoint.client, path).Get(ctx)
}

// GetCustomTemplateByID выполняет запрос на получение информации об отдельном пользовательском шаблоне печатной формы для указанного типа сущности по его id.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-pol-zowatel-skij-shablon
func (endpoint *endpointPrintTemplates) GetCustomTemplateByID(ctx context.Context, id uuid.UUID) (*CustomTemplate, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomTemplateID, endpoint.uri, id)
	return NewRequestBuilder[CustomTemplate](endpoint.client, path).Get(ctx)
}

type endpointTrash struct{ Endpoint }

// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-udalenie-w-korzinu
func (endpoint *endpointTrash) MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointTrash, endpoint.uri, id)
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
