package moysklad

import (
	"github.com/google/uuid"
)

// CommissionReportOut Выданный отчет комиссионера.
// Ключевое слово: commissionreportout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera
type CommissionReportOut struct {
	Organization          *Organization                           `json:"organization,omitempty"`
	CommissionPeriodStart *Timestamp                              `json:"commissionPeriodStart,omitempty"`
	AgentAccount          *AgentAccount                           `json:"agentAccount,omitempty"`
	AccountID             *uuid.UUID                              `json:"accountId,omitempty"`
	Attributes            *Attributes                             `json:"attributes,omitempty"`
	Owner                 *Employee                               `json:"owner,omitempty"`
	CommissionPeriodEnd   *Timestamp                              `json:"commissionPeriodEnd,omitempty"`
	OrganizationAccount   *AgentAccount                           `json:"organizationAccount,omitempty"`
	CommitentSum          *float64                                `json:"commitentSum,omitempty"`
	Contract              *Contract                               `json:"contract,omitempty"`
	Created               *Timestamp                              `json:"created,omitempty"`
	Deleted               *Timestamp                              `json:"deleted,omitempty"`
	Description           *string                                 `json:"description,omitempty"`
	ExternalCode          *string                                 `json:"externalCode,omitempty"`
	Files                 *Files                                  `json:"files,omitempty"`
	Group                 *Group                                  `json:"group,omitempty"`
	ID                    *uuid.UUID                              `json:"id,omitempty"`
	Meta                  *Meta                                   `json:"meta,omitempty"`
	Moment                *Timestamp                              `json:"moment,omitempty"`
	Name                  *string                                 `json:"name,omitempty"`
	Applicable            *bool                                   `json:"applicable,omitempty"`
	Agent                 *Counterparty                           `json:"agent,omitempty"`
	Code                  *string                                 `json:"code,omitempty"`
	PayedSum              *float64                                `json:"payedSum,omitempty"`
	Positions             *Positions[CommissionReportOutPosition] `json:"positions,omitempty"`
	Printed               *bool                                   `json:"printed,omitempty"`
	Project               *Project                                `json:"project,omitempty"`
	Published             *bool                                   `json:"published,omitempty"`
	Rate                  *Rate                                   `json:"rate,omitempty"`
	RewardPercent         *float64                                `json:"rewardPercent,omitempty"`
	Payments              *Payments                               `json:"payments,omitempty"`
	SalesChannel          *SalesChannel                           `json:"salesChannel,omitempty"`
	Shared                *bool                                   `json:"shared,omitempty"`
	State                 *State                                  `json:"state,omitempty"`
	Sum                   *float64                                `json:"sum,omitempty"`
	SyncID                *uuid.UUID                              `json:"syncId,omitempty"`
	Updated               *Timestamp                              `json:"updated,omitempty"`
	VatEnabled            *bool                                   `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                                   `json:"vatIncluded,omitempty"`
	VatSum                *float64                                `json:"vatSum,omitempty"`
	RewardType            RewardType                              `json:"rewardType,omitempty"`
}

func (c CommissionReportOut) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CommissionReportOut) GetMeta() *Meta {
	return c.Meta
}

func (c CommissionReportOut) MetaType() MetaType {
	return MetaTypeCommissionReportOut
}

// CommissionReportOutPosition Позиция Выданного отчета комиссионера.
// Ключевое слово: commissionreportoutposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera-vydannye-otchety-komissionera-pozicii-vydannogo-otcheta-komissionera
type CommissionReportOutPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Meta       *Meta               `json:"meta,omitempty"`       // Метаданные
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (c CommissionReportOutPosition) String() string {
	return Stringify(c)
}

func (c CommissionReportOutPosition) MetaType() MetaType {
	return MetaTypeCommissionReportOutPosition
}
