package moysklad

import (
	"github.com/google/uuid"
)

// CommissionReportIn Полученный отчет комиссионера.
// Ключевое слово: commissionreportin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera
type CommissionReportIn struct {
	Name                          *string                                      `json:"name,omitempty"`
	Applicable                    *bool                                        `json:"applicable,omitempty"`
	AgentAccount                  *AgentAccount                                `json:"agentAccount,omitempty"`
	Organization                  *Organization                                `json:"organization,omitempty"`
	Attributes                    *Attributes                                  `json:"attributes,omitempty"`
	Code                          *string                                      `json:"code,omitempty"`
	CommissionOverhead            *CommissionOverhead                          `json:"commissionOverhead,omitempty"`
	CommissionPeriodEnd           *Timestamp                                   `json:"commissionPeriodEnd,omitempty"`
	CommissionPeriodStart         *Timestamp                                   `json:"commissionPeriodStart,omitempty"`
	CommitentSum                  *Decimal                                     `json:"commitentSum,omitempty"`
	Contract                      *Contract                                    `json:"contract,omitempty"`
	Created                       *Timestamp                                   `json:"created,omitempty"`
	Deleted                       *Timestamp                                   `json:"deleted,omitempty"`
	Description                   *string                                      `json:"description,omitempty"`
	ExternalCode                  *string                                      `json:"externalCode,omitempty"`
	OrganizationAccount           *AgentAccount                                `json:"organizationAccount,omitempty"`
	Group                         *Group                                       `json:"group,omitempty"`
	ID                            *uuid.UUID                                   `json:"id,omitempty"`
	Meta                          *Meta                                        `json:"meta,omitempty"`
	Moment                        *Timestamp                                   `json:"moment,omitempty"`
	VatSum                        *Decimal                                     `json:"vatSum,omitempty"`
	Agent                         *Counterparty                                `json:"agent,omitempty"`
	Files                         *Files                                       `json:"files,omitempty"`
	Owner                         *Employee                                    `json:"owner,omitempty"`
	PayedSum                      *float64                                     `json:"payedSum,omitempty"`
	Positions                     *Positions[CommissionReportInPosition]       `json:"positions,omitempty"`
	Printed                       *bool                                        `json:"printed,omitempty"`
	Project                       *Project                                     `json:"project,omitempty"`
	Published                     *bool                                        `json:"published,omitempty"`
	Rate                          *Rate                                        `json:"rate,omitempty"`
	ReturnToCommissionerPositions *Positions[CommissionReportInReturnPosition] `json:"returnToCommissionerPositions,omitempty"`
	RewardPercent                 *float64                                     `json:"rewardPercent,omitempty"`
	Payments                      *Payments                                    `json:"payments,omitempty"`
	SalesChannel                  *SalesChannel                                `json:"salesChannel,omitempty"`
	Shared                        *bool                                        `json:"shared,omitempty"`
	State                         *State                                       `json:"state,omitempty"`
	Sum                           *float64                                     `json:"sum,omitempty"`
	SyncID                        *uuid.UUID                                   `json:"syncId,omitempty"`
	Updated                       *Timestamp                                   `json:"updated,omitempty"`
	VatEnabled                    *bool                                        `json:"vatEnabled,omitempty"`
	VatIncluded                   *bool                                        `json:"vatIncluded,omitempty"`
	AccountID                     *uuid.UUID                                   `json:"accountId,omitempty"`
	RewardType                    RewardType                                   `json:"rewardType,omitempty"`
}

func (c CommissionReportIn) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CommissionReportIn) GetMeta() *Meta {
	return c.Meta
}

func (c CommissionReportIn) MetaType() MetaType {
	return MetaTypeCommissionReportIn
}

// CommissionOverhead Прочие расходы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-prochie-rashody
type CommissionOverhead struct {
	Sum *float64 `json:"sum,omitempty"` // Сумма в копейках
}

func (c CommissionOverhead) String() string {
	return Stringify(c)
}

// CommissionReportInPosition Позиция Полученного отчета комиссионера.
// Ключевое слово: commissionreportinposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-pozicii-poluchennogo-otcheta-komissionera
type CommissionReportInPosition struct {
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

func (c CommissionReportInPosition) String() string {
	return Stringify(c)
}

func (c CommissionReportInPosition) MetaType() MetaType {
	return MetaTypeCommissionReportInPosition
}

// CommissionReportInReturnPosition Позиция возврата на склад комиссионера.
// Ключевое слово: commissionreportinreturnedposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-pozicii-poluchennogo-otcheta-komissionera-ob-ekt-pozicii-wozwrata-na-sklad-komissionera-soderzhit-sleduuschie-polq
type CommissionReportInReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Meta       *Meta               `json:"meta,omitempty"`       // Метаданные
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (c CommissionReportInReturnPosition) String() string {
	return Stringify(c)
}

func (c CommissionReportInReturnPosition) MetaType() MetaType {
	return MetaTypeCommissionReportInReturnPosition
}
