package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CommissionReportIn Полученный отчет комиссионера.
// Ключевое слово: commissionreportin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera
type CommissionReportIn struct {
	VatSum                        *Decimal                                     `json:"vatSum,omitempty"`
	Organization                  *Organization                                `json:"organization,omitempty"`
	AgentAccount                  *AgentAccount                                `json:"agentAccount,omitempty"`
	Agent                         *Counterparty                                `json:"agent,omitempty"`
	Name                          *string                                      `json:"name,omitempty"`
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
	Files                         *Files                                       `json:"files,omitempty"`
	Group                         *Group                                       `json:"group,omitempty"`
	ID                            *uuid.UUID                                   `json:"id,omitempty"`
	Meta                          *Meta                                        `json:"meta,omitempty"`
	Moment                        *Timestamp                                   `json:"moment,omitempty"`
	AccountID                     *uuid.UUID                                   `json:"accountId,omitempty"`
	Applicable                    *bool                                        `json:"applicable,omitempty"`
	OrganizationAccount           *AgentAccount                                `json:"organizationAccount,omitempty"`
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
	RewardType                    RewardType                                   `json:"rewardType,omitempty"`
	Attributes                    Attributes                                   `json:"attributes,omitempty"`
}

func (commissionReportIn CommissionReportIn) GetVatSum() Decimal {
	return Deref(commissionReportIn.VatSum)
}

func (commissionReportIn CommissionReportIn) GetOrganization() Organization {
	return Deref(commissionReportIn.Organization)
}

func (commissionReportIn CommissionReportIn) GetAgentAccount() AgentAccount {
	return Deref(commissionReportIn.AgentAccount)
}

func (commissionReportIn CommissionReportIn) GetAgent() Counterparty {
	return Deref(commissionReportIn.Agent)
}

func (commissionReportIn CommissionReportIn) GetName() string {
	return Deref(commissionReportIn.Name)
}

func (commissionReportIn CommissionReportIn) GetCode() string {
	return Deref(commissionReportIn.Code)
}

func (commissionReportIn CommissionReportIn) GetCommissionOverhead() CommissionOverhead {
	return Deref(commissionReportIn.CommissionOverhead)
}

func (commissionReportIn CommissionReportIn) GetCommissionPeriodEnd() Timestamp {
	return Deref(commissionReportIn.CommissionPeriodEnd)
}

func (commissionReportIn CommissionReportIn) GetCommissionPeriodStart() Timestamp {
	return Deref(commissionReportIn.CommissionPeriodStart)
}

func (commissionReportIn CommissionReportIn) GetCommitentSum() Decimal {
	return Deref(commissionReportIn.CommitentSum)
}

func (commissionReportIn CommissionReportIn) GetContract() Contract {
	return Deref(commissionReportIn.Contract)
}

func (commissionReportIn CommissionReportIn) GetCreated() Timestamp {
	return Deref(commissionReportIn.Created)
}

func (commissionReportIn CommissionReportIn) GetDeleted() Timestamp {
	return Deref(commissionReportIn.Deleted)
}

func (commissionReportIn CommissionReportIn) GetDescription() string {
	return Deref(commissionReportIn.Description)
}

func (commissionReportIn CommissionReportIn) GetExternalCode() string {
	return Deref(commissionReportIn.ExternalCode)
}

func (commissionReportIn CommissionReportIn) GetFiles() Files {
	return Deref(commissionReportIn.Files)
}

func (commissionReportIn CommissionReportIn) GetGroup() Group {
	return Deref(commissionReportIn.Group)
}

func (commissionReportIn CommissionReportIn) GetID() uuid.UUID {
	return Deref(commissionReportIn.ID)
}

func (commissionReportIn CommissionReportIn) GetMeta() Meta {
	return Deref(commissionReportIn.Meta)
}

func (commissionReportIn CommissionReportIn) GetMoment() Timestamp {
	return Deref(commissionReportIn.Moment)
}

func (commissionReportIn CommissionReportIn) GetAccountID() uuid.UUID {
	return Deref(commissionReportIn.AccountID)
}

func (commissionReportIn CommissionReportIn) GetApplicable() bool {
	return Deref(commissionReportIn.Applicable)
}

func (commissionReportIn CommissionReportIn) GetOrganizationAccount() AgentAccount {
	return Deref(commissionReportIn.OrganizationAccount)
}

func (commissionReportIn CommissionReportIn) GetOwner() Employee {
	return Deref(commissionReportIn.Owner)
}

func (commissionReportIn CommissionReportIn) GetPayedSum() float64 {
	return Deref(commissionReportIn.PayedSum)
}

func (commissionReportIn CommissionReportIn) GetPositions() Positions[CommissionReportInPosition] {
	return Deref(commissionReportIn.Positions)
}

func (commissionReportIn CommissionReportIn) GetPrinted() bool {
	return Deref(commissionReportIn.Printed)
}

func (commissionReportIn CommissionReportIn) GetProject() Project {
	return Deref(commissionReportIn.Project)
}

func (commissionReportIn CommissionReportIn) GetPublished() bool {
	return Deref(commissionReportIn.Published)
}

func (commissionReportIn CommissionReportIn) GetRate() Rate {
	return Deref(commissionReportIn.Rate)
}

func (commissionReportIn CommissionReportIn) GetReturnToCommissionerPositions() Positions[CommissionReportInReturnPosition] {
	return Deref(commissionReportIn.ReturnToCommissionerPositions)
}

func (commissionReportIn CommissionReportIn) GetRewardPercent() float64 {
	return Deref(commissionReportIn.RewardPercent)
}

func (commissionReportIn CommissionReportIn) GetPayments() Payments {
	return Deref(commissionReportIn.Payments)
}

func (commissionReportIn CommissionReportIn) GetSalesChannel() SalesChannel {
	return Deref(commissionReportIn.SalesChannel)
}

func (commissionReportIn CommissionReportIn) GetShared() bool {
	return Deref(commissionReportIn.Shared)
}

func (commissionReportIn CommissionReportIn) GetState() State {
	return Deref(commissionReportIn.State)
}

func (commissionReportIn CommissionReportIn) GetSum() float64 {
	return Deref(commissionReportIn.Sum)
}

func (commissionReportIn CommissionReportIn) GetSyncID() uuid.UUID {
	return Deref(commissionReportIn.SyncID)
}

func (commissionReportIn CommissionReportIn) GetUpdated() Timestamp {
	return Deref(commissionReportIn.Updated)
}

func (commissionReportIn CommissionReportIn) GetVatEnabled() bool {
	return Deref(commissionReportIn.VatEnabled)
}

func (commissionReportIn CommissionReportIn) GetVatIncluded() bool {
	return Deref(commissionReportIn.VatIncluded)
}

func (commissionReportIn CommissionReportIn) GetRewardType() RewardType {
	return commissionReportIn.RewardType
}

func (commissionReportIn CommissionReportIn) GetAttributes() Attributes {
	return commissionReportIn.Attributes
}

func (commissionReportIn *CommissionReportIn) SetOrganization(organization *Organization) *CommissionReportIn {
	commissionReportIn.Organization = organization
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetAgentAccount(agentAccount *AgentAccount) *CommissionReportIn {
	commissionReportIn.AgentAccount = agentAccount
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetAgent(agent *Counterparty) *CommissionReportIn {
	commissionReportIn.Agent = agent
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetName(name string) *CommissionReportIn {
	commissionReportIn.Name = &name
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetCode(code string) *CommissionReportIn {
	commissionReportIn.Code = &code
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetCommissionOverhead(commissionOverhead *CommissionOverhead) *CommissionReportIn {
	commissionReportIn.CommissionOverhead = commissionOverhead
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetCommissionPeriodEnd(commissionPeriodEnd *Timestamp) *CommissionReportIn {
	commissionReportIn.CommissionPeriodEnd = commissionPeriodEnd
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetCommissionPeriodStart(commissionPeriodStart *Timestamp) *CommissionReportIn {
	commissionReportIn.CommissionPeriodStart = commissionPeriodStart
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetContract(contract *Contract) *CommissionReportIn {
	commissionReportIn.Contract = contract
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetDescription(description string) *CommissionReportIn {
	commissionReportIn.Description = &description
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetExternalCode(externalCode string) *CommissionReportIn {
	commissionReportIn.ExternalCode = &externalCode
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetFiles(files *Files) *CommissionReportIn {
	commissionReportIn.Files = files
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetGroup(group *Group) *CommissionReportIn {
	commissionReportIn.Group = group
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetMeta(meta *Meta) *CommissionReportIn {
	commissionReportIn.Meta = meta
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetMoment(moment *Timestamp) *CommissionReportIn {
	commissionReportIn.Moment = moment
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetApplicable(applicable bool) *CommissionReportIn {
	commissionReportIn.Applicable = &applicable
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetOrganizationAccount(organizationAccount *AgentAccount) *CommissionReportIn {
	commissionReportIn.OrganizationAccount = organizationAccount
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetOwner(owner *Employee) *CommissionReportIn {
	commissionReportIn.Owner = owner
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetPositions(positions *Positions[CommissionReportInPosition]) *CommissionReportIn {
	commissionReportIn.Positions = positions
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetProject(project *Project) *CommissionReportIn {
	commissionReportIn.Project = project
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetRate(rate *Rate) *CommissionReportIn {
	commissionReportIn.Rate = rate
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetReturnToCommissionerPositions(returnToCommissionerPositions *Positions[CommissionReportInReturnPosition]) *CommissionReportIn {
	commissionReportIn.ReturnToCommissionerPositions = returnToCommissionerPositions
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetRewardPercent(rewardPercent float64) *CommissionReportIn {
	commissionReportIn.RewardPercent = &rewardPercent
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetPayments(payments *Payments) *CommissionReportIn {
	commissionReportIn.Payments = payments
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetSalesChannel(salesChannel *SalesChannel) *CommissionReportIn {
	commissionReportIn.SalesChannel = salesChannel
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetShared(shared bool) *CommissionReportIn {
	commissionReportIn.Shared = &shared
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetState(state *State) *CommissionReportIn {
	commissionReportIn.State = state
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetSyncID(syncID *uuid.UUID) *CommissionReportIn {
	commissionReportIn.SyncID = syncID
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetVatEnabled(vatEnabled bool) *CommissionReportIn {
	commissionReportIn.VatEnabled = &vatEnabled
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetVatIncluded(vatIncluded bool) *CommissionReportIn {
	commissionReportIn.VatIncluded = &vatIncluded
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetRewardType(rewardType RewardType) *CommissionReportIn {
	commissionReportIn.RewardType = rewardType
	return commissionReportIn
}

func (commissionReportIn *CommissionReportIn) SetAttributes(attributes Attributes) *CommissionReportIn {
	commissionReportIn.Attributes = attributes
	return commissionReportIn
}

func (commissionReportIn CommissionReportIn) String() string {
	return Stringify(commissionReportIn)
}

func (commissionReportIn CommissionReportIn) MetaType() MetaType {
	return MetaTypeCommissionReportIn
}

// CommissionOverhead Прочие расходы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-prochie-rashody
type CommissionOverhead struct {
	Sum *float64 `json:"sum,omitempty"` // Сумма в копейках
}

func (commissionOverhead CommissionOverhead) GetSum() float64 {
	return Deref(commissionOverhead.Sum)
}

func (commissionOverhead *CommissionOverhead) SetSum(sum float64) *CommissionOverhead {
	commissionOverhead.Sum = &sum
	return commissionOverhead
}

func (commissionOverhead CommissionOverhead) String() string {
	return Stringify(commissionOverhead)
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
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *Decimal            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (commissionReportInPosition CommissionReportInPosition) GetAccountID() uuid.UUID {
	return Deref(commissionReportInPosition.AccountID)
}

func (commissionReportInPosition CommissionReportInPosition) GetAssortment() AssortmentPosition {
	return Deref(commissionReportInPosition.Assortment)
}

func (commissionReportInPosition CommissionReportInPosition) GetID() uuid.UUID {
	return Deref(commissionReportInPosition.ID)
}

func (commissionReportInPosition CommissionReportInPosition) GetPack() Pack {
	return Deref(commissionReportInPosition.Pack)
}

func (commissionReportInPosition CommissionReportInPosition) GetMeta() Meta {
	return Deref(commissionReportInPosition.Meta)
}

func (commissionReportInPosition CommissionReportInPosition) GetPrice() Decimal {
	return Deref(commissionReportInPosition.Price)
}

func (commissionReportInPosition CommissionReportInPosition) GetQuantity() float64 {
	return Deref(commissionReportInPosition.Quantity)
}

func (commissionReportInPosition CommissionReportInPosition) GetReward() Decimal {
	return Deref(commissionReportInPosition.Reward)
}

func (commissionReportInPosition CommissionReportInPosition) GetVat() int {
	return Deref(commissionReportInPosition.Vat)
}

func (commissionReportInPosition CommissionReportInPosition) GetVatEnabled() bool {
	return Deref(commissionReportInPosition.VatEnabled)
}

func (commissionReportInPosition *CommissionReportInPosition) SetAssortment(assortment *AssortmentPosition) *CommissionReportInPosition {
	commissionReportInPosition.Assortment = assortment
	return commissionReportInPosition
}

func (commissionReportInPosition *CommissionReportInPosition) SetPack(pack *Pack) *CommissionReportInPosition {
	commissionReportInPosition.Pack = pack
	return commissionReportInPosition
}

func (commissionReportInPosition *CommissionReportInPosition) SetMeta(meta *Meta) *CommissionReportInPosition {
	commissionReportInPosition.Meta = meta
	return commissionReportInPosition
}

func (commissionReportInPosition *CommissionReportInPosition) SetPrice(price *Decimal) *CommissionReportInPosition {
	commissionReportInPosition.Price = price
	return commissionReportInPosition
}

func (commissionReportInPosition *CommissionReportInPosition) SetQuantity(quantity float64) *CommissionReportInPosition {
	commissionReportInPosition.Quantity = &quantity
	return commissionReportInPosition
}

func (commissionReportInPosition *CommissionReportInPosition) SetReward(reward *Decimal) *CommissionReportInPosition {
	commissionReportInPosition.Reward = reward
	return commissionReportInPosition
}

func (commissionReportInPosition *CommissionReportInPosition) SetVat(vat int) *CommissionReportInPosition {
	commissionReportInPosition.Vat = &vat
	return commissionReportInPosition
}

func (commissionReportInPosition *CommissionReportInPosition) SetVatEnabled(vatEnabled bool) *CommissionReportInPosition {
	commissionReportInPosition.VatEnabled = &vatEnabled
	return commissionReportInPosition
}

func (commissionReportInPosition CommissionReportInPosition) String() string {
	return Stringify(commissionReportInPosition)
}

func (commissionReportInPosition CommissionReportInPosition) MetaType() MetaType {
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

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetAccountID() uuid.UUID {
	return Deref(commissionReportInReturnPosition.AccountID)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(commissionReportInReturnPosition.Assortment)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetID() uuid.UUID {
	return Deref(commissionReportInReturnPosition.ID)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetMeta() Meta {
	return Deref(commissionReportInReturnPosition.Meta)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetPrice() float64 {
	return Deref(commissionReportInReturnPosition.Price)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetQuantity() float64 {
	return Deref(commissionReportInReturnPosition.Quantity)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetReward() float64 {
	return Deref(commissionReportInReturnPosition.Reward)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetVat() int {
	return Deref(commissionReportInReturnPosition.Vat)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetVatEnabled() bool {
	return Deref(commissionReportInReturnPosition.VatEnabled)
}

func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetAssortment(assortment *AssortmentPosition) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Assortment = assortment
	return commissionReportInReturnPosition
}

func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetMeta(meta *Meta) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Meta = meta
	return commissionReportInReturnPosition
}

func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetPrice(price float64) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Price = &price
	return commissionReportInReturnPosition
}

func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetQuantity(quantity float64) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Quantity = &quantity
	return commissionReportInReturnPosition
}

func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetReward(reward float64) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Reward = &reward
	return commissionReportInReturnPosition
}

func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetVat(vat int) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Vat = &vat
	return commissionReportInReturnPosition
}

func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetVatEnabled(vatEnabled bool) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.VatEnabled = &vatEnabled
	return commissionReportInReturnPosition
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) String() string {
	return Stringify(commissionReportInReturnPosition)
}

func (commissionReportInReturnPosition CommissionReportInReturnPosition) MetaType() MetaType {
	return MetaTypeCommissionReportInReturnPosition
}

// CommissionReportInService
// Сервис для работы с полученными отчётами комиссионера.
type CommissionReportInService interface {
	GetList(ctx context.Context, params *Params) (*List[CommissionReportIn], *resty.Response, error)
	Create(ctx context.Context, commissionReportIn *CommissionReportIn, params *Params) (*CommissionReportIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, commissionReportInList []*CommissionReportIn, params *Params) (*[]CommissionReportIn, *resty.Response, error)
	DeleteMany(ctx context.Context, commissionReportInList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CommissionReportIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, commissionReportIn *CommissionReportIn, params *Params) (*CommissionReportIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportInPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*CommissionReportInPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *CommissionReportInPosition, params *Params) (*CommissionReportInPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *CommissionReportInPosition) (*CommissionReportInPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*CommissionReportInPosition) (*[]CommissionReportInPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CommissionReportIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	// Template(ctx context.Context) (*CommissionReportIn, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetReturnPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportInReturnPosition], *resty.Response, error)
	GetReturnPositionByID(ctx context.Context, id, positionID *uuid.UUID, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error)
	CreateReturnPosition(ctx context.Context, id *uuid.UUID, position *CommissionReportInReturnPosition) (*CommissionReportInReturnPosition, *resty.Response, error)
	UpdateReturnPosition(ctx context.Context, id, positionID *uuid.UUID, position *CommissionReportInReturnPosition, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error)
	DeleteReturnPosition(ctx context.Context, id, positionID *uuid.UUID) (bool, *resty.Response, error)
}

type commissionReportInService struct {
	Endpoint
	endpointGetList[CommissionReportIn]
	endpointCreate[CommissionReportIn]
	endpointCreateUpdateMany[CommissionReportIn]
	endpointDeleteMany[CommissionReportIn]
	endpointDelete
	endpointGetById[CommissionReportIn]
	endpointUpdate[CommissionReportIn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[CommissionReportInPosition]
	endpointAttributes
	endpointSyncID[CommissionReportIn]
	endpointNamedFilter
	endpointPublication
	endpointRemove
}

func NewCommissionReportInService(client *Client) CommissionReportInService {
	e := NewEndpoint(client, "entity/commissionreportin")
	return &commissionReportInService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[CommissionReportIn]{e},
		endpointCreate:           endpointCreate[CommissionReportIn]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[CommissionReportIn]{e},
		endpointDeleteMany:       endpointDeleteMany[CommissionReportIn]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[CommissionReportIn]{e},
		endpointUpdate:           endpointUpdate[CommissionReportIn]{e},
		endpointMetadata:         endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:        endpointPositions[CommissionReportInPosition]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[CommissionReportIn]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		//endpointTemplate:         endpointTemplate[CommissionReportIn]{e},
		endpointPublication: endpointPublication{e},
		endpointRemove:      endpointRemove{e},
	}
}

// GetReturnPositions Получить позиции возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchit-pozicii-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) GetReturnPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportInReturnPosition], *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[MetaArray[CommissionReportInReturnPosition]](s.client, path).SetParams(params).Get(ctx)
}

// GetReturnPositionByID Получить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchit-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) GetReturnPositionByID(ctx context.Context, id, positionID *uuid.UUID, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionID)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).SetParams(params).Get(ctx)
}

// CreateReturnPosition Создать позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-sozdat-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) CreateReturnPosition(ctx context.Context, id *uuid.UUID, position *CommissionReportInReturnPosition) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).Post(ctx, position)
}

// UpdateReturnPosition Изменить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-izmenit-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) UpdateReturnPosition(ctx context.Context, id, positionID *uuid.UUID, position *CommissionReportInReturnPosition, params *Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionID)
	return NewRequestBuilder[CommissionReportInReturnPosition](s.client, path).SetParams(params).Put(ctx, position)
}

// DeleteReturnPosition Удалить позицию возврата на склад комиссионера.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-udalit-poziciu-wozwrata-na-sklad-komissionera
func (s *commissionReportInService) DeleteReturnPosition(ctx context.Context, id, positionID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/positions/%s", id, positionID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
