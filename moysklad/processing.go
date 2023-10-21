package moysklad

import (
	"github.com/google/uuid"
)

// Processing Техоперация.
// Ключевое слово: processing
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq
type Processing struct {
	Organization        *Organization                     `json:"organization,omitempty"`
	SyncID              *uuid.UUID                        `json:"syncId,omitempty"`
	Attributes          *Attributes                       `json:"attributes,omitempty"`
	Code                *string                           `json:"code,omitempty"`
	Created             *Timestamp                        `json:"created,omitempty"`
	Deleted             *Timestamp                        `json:"deleted,omitempty"`
	AccountID           *uuid.UUID                        `json:"accountId,omitempty"`
	ExternalCode        *string                           `json:"externalCode,omitempty"`
	Files               *Files                            `json:"files,omitempty"`
	Group               *Group                            `json:"group,omitempty"`
	ID                  *uuid.UUID                        `json:"id,omitempty"`
	Moment              *Timestamp                        `json:"moment,omitempty"`
	MaterialsStore      *Store                            `json:"materialsStore,omitempty"`
	Meta                *Meta                             `json:"meta,omitempty"`
	ProcessingOrder     *ProcessingOrder                  `json:"processingOrder,omitempty"`
	Applicable          *bool                             `json:"applicable,omitempty"`
	Description         *string                           `json:"description,omitempty"`
	OrganizationAccount *AgentAccount                     `json:"organizationAccount,omitempty"`
	Owner               *Employee                         `json:"owner,omitempty"`
	Printed             *bool                             `json:"printed,omitempty"`
	ProcessingPlan      *ProcessingPlan                   `json:"processingPlan,omitempty"`
	ProcessingSum       *float64                          `json:"processingSum,omitempty"`
	Updated             *Timestamp                        `json:"updated,omitempty"`
	ProductsStore       *Store                            `json:"productsStore,omitempty"`
	Project             *Project                          `json:"project,omitempty"`
	Published           *bool                             `json:"published,omitempty"`
	Quantity            *float64                          `json:"quantity,omitempty"`
	Shared              *bool                             `json:"shared,omitempty"`
	State               *State                            `json:"state,omitempty"`
	Name                *string                           `json:"name,omitempty"`
	Products            Slice[ProcessingPositionProduct]  `json:"products,omitempty"`
	Materials           Slice[ProcessingPositionMaterial] `json:"materials,omitempty"`
}

func (p Processing) String() string {
	return Stringify(p)
}

func (p Processing) MetaType() MetaType {
	return MetaTypeProcessing
}

type Processings = Slice[Processing]

// ProcessingPositionMaterial Материал Техоперации.
// Ключевое слово: processingpositionmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-materialy-tehoperacii
type ProcessingPositionMaterial struct {
	ProcessingPosition
}

func (p ProcessingPositionMaterial) String() string {
	return Stringify(p)
}

func (p ProcessingPositionMaterial) MetaType() MetaType {
	return MetaTypeProcessingPositionMaterial
}

// ProcessingPositionProduct Продукт Техоперации.
// Ключевое слово: processingpositionresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-produkty-tehoperacii
type ProcessingPositionProduct struct {
	ProcessingPosition
}

func (p ProcessingPositionProduct) String() string {
	return Stringify(p)
}

func (p ProcessingPositionProduct) MetaType() MetaType {
	return MetaTypeProcessingPositionProduct
}

// ProcessingTemplateArg
// Документ: Техоперация (processing)
// Основание, на котором он может быть создан:
// - Заказ на производство (processingorder)
// - Техкарта (processingplan)
type ProcessingTemplateArg struct {
	ProcessingOrder *MetaWrapper `json:"processingOrder,omitempty"`
	ProcessingPlan  *MetaWrapper `json:"processingPlan,omitempty"`
}
