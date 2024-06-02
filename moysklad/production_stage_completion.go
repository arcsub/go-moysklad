package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionStageCompletion Выполнение этапа производства
// Ключевое слово: productionstagecompletion
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa
type ProductionStageCompletion struct {
	AccountId          *uuid.UUID                                    `json:"accountId,omitempty"`          // ID учетной записи
	Created            *Timestamp                                    `json:"created,omitempty"`            // Дата создания
	ExternalCode       *string                                       `json:"externalCode,omitempty"`       // Внешний код Выполнения этапа производства
	Group              *Group                                        `json:"group,omitempty"`              // Отдел сотрудника
	ID                 *uuid.UUID                                    `json:"id,omitempty"`                 // ID Выполнения этапа производства
	LabourUnitCost     *float64                                      `json:"labourUnitCost,omitempty"`     // Оплата труда за единицу объема производства
	Materials          *Positions[ProductionStageCompletionMaterial] `json:"materials,omitempty"`          // Метаданные Материалов выполнения этапа производства
	Meta               *Meta                                         `json:"meta,omitempty"`               // Метаданные Выполнения этапа производства
	Moment             *Timestamp                                    `json:"moment,omitempty"`             // Дата документа
	Name               *string                                       `json:"name,omitempty"`               // Наименование Выполнения этапа производства
	Owner              *Employee                                     `json:"owner,omitempty"`              // Владелец (Сотрудник)
	Performer          *Employee                                     `json:"performer,omitempty"`          // Исполнитель (Сотрудник)
	ProcessingUnitCost *float64                                      `json:"processingUnitCost,omitempty"` // Затраты на единицу объема производства
	ProductionStage    *ProductionStage                              `json:"productionStage,omitempty"`    // Производственный этап
	ProductionVolume   *float64                                      `json:"productionVolume,omitempty"`   // Объем производства
	Products           *Positions[ProductionStageCompletionResult]   `json:"products,omitempty"`           // Метаданные Продуктов выполнения этапа производства. Есть только у последнего этапа
	Shared             *bool                                         `json:"shared,omitempty"`             // Общий доступ
	Updated            *Timestamp                                    `json:"updated,omitempty"`            // Момент последнего обновления Выполнения этапа производства
}

// ProductionStageCompletionMaterial Материалы Выполнения этапа производства
// Ключевое слово: productionstagecompletionmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-vypolnenie-atapa-proizwodstwa-materialy-vypolneniq-atapa-proizwodstwa
type ProductionStageCompletionMaterial struct {
	AccountId        *uuid.UUID          `json:"accountId,omitempty"`        // ID учетной записи
	Assortment       *AssortmentPosition `json:"assortment,omitempty"`       // Метаданные товара/модификации/серии, которую представляет собой позиция
	ConsumedQuantity *float64            `json:"consumedQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	ID               *uuid.UUID          `json:"id,omitempty"`               // ID позиции
	Things           *Things             `json:"things,omitempty"`           // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

// ProductionStageCompletionResult Продукт Выполнения этапа производства
// Ключевое слово: productionstagecompletionresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-vypolnenie-atapa-proizwodstwa-produkty-vypolneniq-atapa-proizwodstwa
type ProductionStageCompletionResult struct {
	AccountId        *uuid.UUID          `json:"accountId,omitempty"`        // ID учетной записи
	Assortment       *AssortmentPosition `json:"assortment,omitempty"`       // Метаданные товара/модификации/серии, которую представляет собой позиция
	ID               *uuid.UUID          `json:"id,omitempty"`               // ID позиции
	ProducedQuantity *float64            `json:"producedQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	Things           *Things             `json:"things,omitempty"`           // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

// ProductionStageCompletionService
// Сервис для работы с выполнениями этапов производства
type ProductionStageCompletionService interface {
	GetList(ctx context.Context, params *Params) (*List[ProductionStageCompletion], *resty.Response, error)
	Create(ctx context.Context, productionStageCompletion *ProductionStageCompletion, params *Params) (*ProductionStageCompletion, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productionStageCompletionList []*ProductionStageCompletion, params *Params) (*[]ProductionStageCompletion, *resty.Response, error)
	DeleteMany(ctx context.Context, productionStageCompletionList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProductionStageCompletion, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, productionStageCompletion *ProductionStageCompletion, params *Params) (*ProductionStageCompletion, *resty.Response, error)
	GetMaterials(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionStageCompletionMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id *uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params *Params) (*ProductionStageCompletionMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id *uuid.UUID, materialID *uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params *Params) (*ProductionStageCompletionMaterial, *resty.Response, error)
	GetProducts(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionStageCompletionResult], *resty.Response, error)
	UpdateProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, productionStageCompletionResult *ProductionStageCompletionResult, params *Params) (*ProductionStageCompletionResult, *resty.Response, error)
}

type productionStageCompletionService struct {
	Endpoint
	endpointGetList[ProductionStageCompletion]
	endpointCreate[ProductionStageCompletion]
	endpointCreateUpdateMany[ProductionStageCompletion]
	endpointDeleteMany[ProductionStageCompletion]
	endpointDelete
	endpointGetById[ProductionStageCompletion]
	endpointUpdate[ProductionStageCompletion]
}

func NewProductionStageCompletionService(client *Client) ProductionStageCompletionService {
	e := NewEndpoint(client, "entity/productionstagecompletion")
	return &productionStageCompletionService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[ProductionStageCompletion]{e},
		endpointCreate:           endpointCreate[ProductionStageCompletion]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[ProductionStageCompletion]{e},
		endpointDeleteMany:       endpointDeleteMany[ProductionStageCompletion]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[ProductionStageCompletion]{e},
		endpointUpdate:           endpointUpdate[ProductionStageCompletion]{e},
	}
}

// GetMaterials Получить Материалы выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-poluchit-materialy-wypolneniq-atapa-proizwodstwa
func (s *productionStageCompletionService) GetMaterials(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionStageCompletionMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[MetaArray[ProductionStageCompletionMaterial]](s.client, path).SetParams(params).Get(ctx)
}

// CreateMaterial Добавить Материал выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-dobawit-material-wypolneniq-atapa-proizwodstwa
func (s *productionStageCompletionService) CreateMaterial(ctx context.Context, id *uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params *Params) (*ProductionStageCompletionMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[ProductionStageCompletionMaterial](s.client, path).SetParams(params).Post(ctx, productionStageCompletionMaterial)
}

// UpdateMaterial Изменить Материал выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-material-wypolneniq-atapa-proizwodstwa
func (s *productionStageCompletionService) UpdateMaterial(ctx context.Context, id *uuid.UUID, materialID *uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params *Params) (*ProductionStageCompletionMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialID)
	return NewRequestBuilder[ProductionStageCompletionMaterial](s.client, path).SetParams(params).Put(ctx, productionStageCompletionMaterial)
}

// GetProducts Получить Продукты выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-poluchit-produkty-wypolneniq-atapa-proizwodstwa
func (s *productionStageCompletionService) GetProducts(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionStageCompletionResult], *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[MetaArray[ProductionStageCompletionResult]](s.client, path).SetParams(params).Get(ctx)
}

// UpdateProduct Изменить Продукт выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-produkt-wypolneniq-atapa-proizwodstwa
func (s *productionStageCompletionService) UpdateProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, productionStageCompletionResult *ProductionStageCompletionResult, params *Params) (*ProductionStageCompletionResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionStageCompletionResult](s.client, path).SetParams(params).Put(ctx, productionStageCompletionResult)
}
