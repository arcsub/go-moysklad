package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CommissionReportIn Полученный отчёт комиссионера.
//
// Код сущности: commissionreportin
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera
type CommissionReportIn struct {
	VatSum                        *float64                                     `json:"vatSum,omitempty"`                        // Сумма НДС
	Organization                  *Organization                                `json:"organization,omitempty"`                  // Метаданные юрлица
	AgentAccount                  *AgentAccount                                `json:"agentAccount,omitempty"`                  // Метаданные счета контрагента
	Agent                         *Counterparty                                `json:"agent,omitempty"`                         // Метаданные контрагента
	Name                          *string                                      `json:"name,omitempty"`                          // Наименование Полученного отчёта комиссионера
	Code                          *string                                      `json:"code,omitempty"`                          // Код Полученного отчёта комиссионера
	CommissionOverhead            *CommissionOverhead                          `json:"commissionOverhead,omitempty"`            // Прочие расходы. Если Позиции отчёта комиссионера не заданы, то расходы нельзя задать
	CommissionPeriodEnd           *Timestamp                                   `json:"commissionPeriodEnd,omitempty"`           // Конец периода
	CommissionPeriodStart         *Timestamp                                   `json:"commissionPeriodStart,omitempty"`         // Начало периода
	CommitentSum                  *float64                                     `json:"commitentSum,omitempty"`                  // Сумма комитента в установленной валюте
	Contract                      *Contract                                    `json:"contract,omitempty"`                      // Метаданные договора
	Created                       *Timestamp                                   `json:"created,omitempty"`                       // Дата создания
	Deleted                       *Timestamp                                   `json:"deleted,omitempty"`                       // Момент последнего удаления Полученного отчёта комиссионера
	Description                   *string                                      `json:"description,omitempty"`                   // Комментарий Полученного отчёта комиссионера
	ExternalCode                  *string                                      `json:"externalCode,omitempty"`                  // Внешний код Полученного отчёта комиссионера
	Files                         *MetaArray[File]                             `json:"files,omitempty"`                         // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                         *Group                                       `json:"group,omitempty"`                         // Отдел сотрудника
	ID                            *uuid.UUID                                   `json:"id,omitempty"`                            // ID Полученного отчёта комиссионера
	Meta                          *Meta                                        `json:"meta,omitempty"`                          // Метаданные Полученного отчёта комиссионера
	Moment                        *Timestamp                                   `json:"moment,omitempty"`                        // Дата документа
	AccountID                     *uuid.UUID                                   `json:"accountId,omitempty"`                     // ID учётной записи
	Applicable                    *bool                                        `json:"applicable,omitempty"`                    // Отметка о проведении
	OrganizationAccount           *AgentAccount                                `json:"organizationAccount,omitempty"`           // Метаданные счета юрлица
	Owner                         *Employee                                    `json:"owner,omitempty"`                         // Метаданные владельца (Сотрудника)
	PayedSum                      *float64                                     `json:"payedSum,omitempty"`                      // Оплаченная сумма
	Positions                     *MetaArray[CommissionReportInPosition]       `json:"positions,omitempty"`                     // Метаданные позиций реализовано комиссионером Полученного отчёта комиссионера
	Printed                       *bool                                        `json:"printed,omitempty"`                       // Напечатан ли документ
	Project                       *NullValue[Project]                          `json:"project,omitempty"`                       // Метаданные проекта
	Published                     *bool                                        `json:"published,omitempty"`                     // Опубликован ли документ
	Rate                          *NullValue[Rate]                             `json:"rate,omitempty"`                          // Валюта
	ReturnToCommissionerPositions *MetaArray[CommissionReportInReturnPosition] `json:"returnToCommissionerPositions,omitempty"` // Метаданные позиций возврата на склад комиссионера Полученного отчёта комиссионера
	RewardPercent                 *float64                                     `json:"rewardPercent,omitempty"`                 // Процент вознаграждения (всегда 0 если вознаграждение не рассчитывается)
	Payments                      Slice[Payment]                               `json:"payments,omitempty"`                      // Массив ссылок на связанные платежи
	SalesChannel                  *NullValue[SalesChannel]                     `json:"salesChannel,omitempty"`                  // Метаданные канала продаж
	Shared                        *bool                                        `json:"shared,omitempty"`                        // Общий доступ
	State                         *NullValue[State]                            `json:"state,omitempty"`                         // Метаданные статуса Полученного отчёта комиссионера
	Sum                           *float64                                     `json:"sum,omitempty"`                           // Сумма Полученного отчёта комиссионера в копейках
	SyncID                        *uuid.UUID                                   `json:"syncId,omitempty"`                        // ID синхронизации
	Updated                       *Timestamp                                   `json:"updated,omitempty"`                       // Момент последнего обновления Полученного отчёта комиссионера
	VatEnabled                    *bool                                        `json:"vatEnabled,omitempty"`                    // Учитывается ли НДС
	VatIncluded                   *bool                                        `json:"vatIncluded,omitempty"`                   // Включен ли НДС в цену
	RewardType                    RewardType                                   `json:"rewardType,omitempty"`                    // Тип вознаграждения
	Attributes                    Slice[Attribute]                             `json:"attributes,omitempty"`                    // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (commissionReportIn CommissionReportIn) Clean() *CommissionReportIn {
	if commissionReportIn.Meta == nil {
		return nil
	}
	return &CommissionReportIn{Meta: commissionReportIn.Meta}
}

// operation возвращает объект [Operation] c полями meta и linkedSum.
//
// Значение поля linkedSum заполняется из поля sum.
func (commissionReportIn CommissionReportIn) operation() *Operation {
	return &Operation{Meta: commissionReportIn.GetMeta(), LinkedSum: commissionReportIn.GetSum()}
}

// asTaskOperation реализует интерфейс [TaskOperationInterface].
func (commissionReportIn CommissionReportIn) asTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: commissionReportIn.Meta}
}

// GetVatSum возвращает Сумму НДС.
func (commissionReportIn CommissionReportIn) GetVatSum() float64 {
	return Deref(commissionReportIn.VatSum)
}

// GetOrganization возвращает Метаданные юрлица.
func (commissionReportIn CommissionReportIn) GetOrganization() Organization {
	return Deref(commissionReportIn.Organization)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (commissionReportIn CommissionReportIn) GetAgentAccount() AgentAccount {
	return Deref(commissionReportIn.AgentAccount)
}

// GetAgent возвращает Метаданные контрагента.
func (commissionReportIn CommissionReportIn) GetAgent() Counterparty {
	return Deref(commissionReportIn.Agent)
}

// GetName возвращает Наименование Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetName() string {
	return Deref(commissionReportIn.Name)
}

// GetCode возвращает Код Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetCode() string {
	return Deref(commissionReportIn.Code)
}

// GetCommissionOverhead возвращает Прочие расходы.
func (commissionReportIn CommissionReportIn) GetCommissionOverhead() CommissionOverhead {
	return Deref(commissionReportIn.CommissionOverhead)
}

// GetCommissionOverheadSum возвращает Сумму в копейках Прочих расходов.
func (commissionReportIn CommissionReportIn) GetCommissionOverheadSum() float64 {
	return Deref(commissionReportIn.CommissionOverhead).GetSum()
}

// GetCommissionPeriodEnd возвращает Конец периода.
func (commissionReportIn CommissionReportIn) GetCommissionPeriodEnd() Timestamp {
	return Deref(commissionReportIn.CommissionPeriodEnd)
}

// GetCommissionPeriodStart возвращает Начало периода.
func (commissionReportIn CommissionReportIn) GetCommissionPeriodStart() Timestamp {
	return Deref(commissionReportIn.CommissionPeriodStart)
}

// GetCommitentSum возвращает Сумму комитента в установленной валюте.
func (commissionReportIn CommissionReportIn) GetCommitentSum() float64 {
	return Deref(commissionReportIn.CommitentSum)
}

// GetContract возвращает Метаданные договора.
func (commissionReportIn CommissionReportIn) GetContract() Contract {
	return Deref(commissionReportIn.Contract)
}

// GetCreated возвращает Дату создания.
func (commissionReportIn CommissionReportIn) GetCreated() Timestamp {
	return Deref(commissionReportIn.Created)
}

// GetDeleted возвращает Момент последнего удаления Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetDeleted() Timestamp {
	return Deref(commissionReportIn.Deleted)
}

// GetDescription возвращает Комментарий Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetDescription() string {
	return Deref(commissionReportIn.Description)
}

// GetExternalCode возвращает Внешний код Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetExternalCode() string {
	return Deref(commissionReportIn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (commissionReportIn CommissionReportIn) GetFiles() MetaArray[File] {
	return Deref(commissionReportIn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (commissionReportIn CommissionReportIn) GetGroup() Group {
	return Deref(commissionReportIn.Group)
}

// GetID возвращает ID Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetID() uuid.UUID {
	return Deref(commissionReportIn.ID)
}

// GetMeta возвращает Метаданные Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetMeta() Meta {
	return Deref(commissionReportIn.Meta)
}

// GetMoment возвращает Дату документа.
func (commissionReportIn CommissionReportIn) GetMoment() Timestamp {
	return Deref(commissionReportIn.Moment)
}

// GetAccountID возвращает ID учётной записи.
func (commissionReportIn CommissionReportIn) GetAccountID() uuid.UUID {
	return Deref(commissionReportIn.AccountID)
}

// GetApplicable возвращает Отметку о проведении.
func (commissionReportIn CommissionReportIn) GetApplicable() bool {
	return Deref(commissionReportIn.Applicable)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (commissionReportIn CommissionReportIn) GetOrganizationAccount() AgentAccount {
	return Deref(commissionReportIn.OrganizationAccount)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (commissionReportIn CommissionReportIn) GetOwner() Employee {
	return Deref(commissionReportIn.Owner)
}

// GetPayedSum возвращает Оплаченную сумму.
func (commissionReportIn CommissionReportIn) GetPayedSum() float64 {
	return Deref(commissionReportIn.PayedSum)
}

// GetPositions возвращает Метаданные позиций реализовано комиссионером Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetPositions() MetaArray[CommissionReportInPosition] {
	return Deref(commissionReportIn.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (commissionReportIn CommissionReportIn) GetPrinted() bool {
	return Deref(commissionReportIn.Printed)
}

// GetProject возвращает Метаданные проекта.
func (commissionReportIn CommissionReportIn) GetProject() Project {
	return commissionReportIn.Project.GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (commissionReportIn CommissionReportIn) GetPublished() bool {
	return Deref(commissionReportIn.Published)
}

// GetRate возвращает Валюту.
func (commissionReportIn CommissionReportIn) GetRate() Rate {
	return commissionReportIn.Rate.GetValue()
}

// GetReturnToCommissionerPositions возвращает Метаданные позиций возврата на склад комиссионера Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetReturnToCommissionerPositions() MetaArray[CommissionReportInReturnPosition] {
	return Deref(commissionReportIn.ReturnToCommissionerPositions)
}

// GetRewardPercent возвращает Процент вознаграждения (всегда 0 если вознаграждение не рассчитывается).
func (commissionReportIn CommissionReportIn) GetRewardPercent() float64 {
	return Deref(commissionReportIn.RewardPercent)
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (commissionReportIn CommissionReportIn) GetPayments() Slice[Payment] {
	return commissionReportIn.Payments
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (commissionReportIn CommissionReportIn) GetSalesChannel() SalesChannel {
	return commissionReportIn.SalesChannel.GetValue()
}

// GetShared возвращает флаг общего доступа.
func (commissionReportIn CommissionReportIn) GetShared() bool {
	return Deref(commissionReportIn.Shared)
}

// GetState возвращает Метаданные статуса Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetState() State {
	return commissionReportIn.State.GetValue()
}

// GetSum возвращает Сумму Полученного отчёта комиссионера в копейках.
func (commissionReportIn CommissionReportIn) GetSum() float64 {
	return Deref(commissionReportIn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (commissionReportIn CommissionReportIn) GetSyncID() uuid.UUID {
	return Deref(commissionReportIn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Полученного отчёта комиссионера.
func (commissionReportIn CommissionReportIn) GetUpdated() Timestamp {
	return Deref(commissionReportIn.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (commissionReportIn CommissionReportIn) GetVatEnabled() bool {
	return Deref(commissionReportIn.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (commissionReportIn CommissionReportIn) GetVatIncluded() bool {
	return Deref(commissionReportIn.VatIncluded)
}

// GetRewardType возвращает Тип вознаграждения.
func (commissionReportIn CommissionReportIn) GetRewardType() RewardType {
	return commissionReportIn.RewardType
}

// GetAttributes возвращает Список метаданных доп. полей.
func (commissionReportIn CommissionReportIn) GetAttributes() Slice[Attribute] {
	return commissionReportIn.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (commissionReportIn *CommissionReportIn) SetOrganization(organization *Organization) *CommissionReportIn {
	if organization != nil {
		commissionReportIn.Organization = organization.Clean()
	}
	return commissionReportIn
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (commissionReportIn *CommissionReportIn) SetAgentAccount(agentAccount *AgentAccount) *CommissionReportIn {
	if agentAccount != nil {
		commissionReportIn.AgentAccount = agentAccount.Clean()
	}
	return commissionReportIn
}

// SetAgent устанавливает Метаданные контрагента.
func (commissionReportIn *CommissionReportIn) SetAgent(agent *Counterparty) *CommissionReportIn {
	if agent != nil {
		commissionReportIn.Agent = agent.Clean()
	}
	return commissionReportIn
}

// SetName устанавливает Наименование Полученного отчёта комиссионера.
func (commissionReportIn *CommissionReportIn) SetName(name string) *CommissionReportIn {
	commissionReportIn.Name = &name
	return commissionReportIn
}

// SetCode устанавливает Код Полученного отчёта комиссионера.
func (commissionReportIn *CommissionReportIn) SetCode(code string) *CommissionReportIn {
	commissionReportIn.Code = &code
	return commissionReportIn
}

// SetCommissionOverhead устанавливает Прочие расходы.
//
// Если Позиции отчёта комиссионера не заданы, то расходы нельзя задать.
func (commissionReportIn *CommissionReportIn) SetCommissionOverhead(commissionOverhead *CommissionOverhead) *CommissionReportIn {
	if commissionOverhead != nil {
		commissionReportIn.CommissionOverhead = commissionOverhead
	}
	return commissionReportIn
}

// SetCommissionOverheadSum устанавливает сумму в копейках Прочих расходов.
//
// Если Позиции отчёта комиссионера не заданы, то расходы нельзя задать.
func (commissionReportIn *CommissionReportIn) SetCommissionOverheadSum(sum float64) *CommissionReportIn {
	commissionReportIn.CommissionOverhead = &CommissionOverhead{&sum}
	return commissionReportIn
}

// SetCommissionPeriodEnd устанавливает Конец периода.
func (commissionReportIn *CommissionReportIn) SetCommissionPeriodEnd(commissionPeriodEnd *Timestamp) *CommissionReportIn {
	commissionReportIn.CommissionPeriodEnd = commissionPeriodEnd
	return commissionReportIn
}

// SetCommissionPeriodStart устанавливает Начало периода.
func (commissionReportIn *CommissionReportIn) SetCommissionPeriodStart(commissionPeriodStart *Timestamp) *CommissionReportIn {
	commissionReportIn.CommissionPeriodStart = commissionPeriodStart
	return commissionReportIn
}

// SetContract устанавливает Метаданные договора.
func (commissionReportIn *CommissionReportIn) SetContract(contract *Contract) *CommissionReportIn {
	if contract != nil {
		commissionReportIn.Contract = contract.Clean()
	}
	return commissionReportIn
}

// SetDescription устанавливает Комментарий Полученного отчёта комиссионера.
func (commissionReportIn *CommissionReportIn) SetDescription(description string) *CommissionReportIn {
	commissionReportIn.Description = &description
	return commissionReportIn
}

// SetExternalCode устанавливает Внешний код Полученного отчёта комиссионера.
func (commissionReportIn *CommissionReportIn) SetExternalCode(externalCode string) *CommissionReportIn {
	commissionReportIn.ExternalCode = &externalCode
	return commissionReportIn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (commissionReportIn *CommissionReportIn) SetFiles(files ...*File) *CommissionReportIn {
	commissionReportIn.Files = NewMetaArrayFrom(files)
	return commissionReportIn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (commissionReportIn *CommissionReportIn) SetGroup(group *Group) *CommissionReportIn {
	if group != nil {
		commissionReportIn.Group = group.Clean()
	}
	return commissionReportIn
}

// SetMeta устанавливает Метаданные Полученного отчёта комиссионера.
func (commissionReportIn *CommissionReportIn) SetMeta(meta *Meta) *CommissionReportIn {
	commissionReportIn.Meta = meta
	return commissionReportIn
}

// SetMoment устанавливает Дату документа.
func (commissionReportIn *CommissionReportIn) SetMoment(moment *Timestamp) *CommissionReportIn {
	commissionReportIn.Moment = moment
	return commissionReportIn
}

// SetApplicable устанавливает Отметку о проведении.
func (commissionReportIn *CommissionReportIn) SetApplicable(applicable bool) *CommissionReportIn {
	commissionReportIn.Applicable = &applicable
	return commissionReportIn
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (commissionReportIn *CommissionReportIn) SetOrganizationAccount(organizationAccount *AgentAccount) *CommissionReportIn {
	if organizationAccount != nil {
		commissionReportIn.OrganizationAccount = organizationAccount.Clean()
	}
	return commissionReportIn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (commissionReportIn *CommissionReportIn) SetOwner(owner *Employee) *CommissionReportIn {
	if owner != nil {
		commissionReportIn.Owner = owner.Clean()
	}
	return commissionReportIn
}

// SetPositions устанавливает Метаданные позиций реализовано комиссионером Полученного отчёта комиссионера.
//
// Принимает множество объектов [CommissionReportInPosition].
func (commissionReportIn *CommissionReportIn) SetPositions(positions ...*CommissionReportInPosition) *CommissionReportIn {
	commissionReportIn.Positions = NewMetaArrayFrom(positions)
	return commissionReportIn
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (commissionReportIn *CommissionReportIn) SetProject(project *Project) *CommissionReportIn {
	commissionReportIn.Project = NewNullValue(project)
	return commissionReportIn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (commissionReportIn *CommissionReportIn) SetRate(rate *Rate) *CommissionReportIn {
	commissionReportIn.Rate = NewNullValue(rate)
	return commissionReportIn
}

// SetReturnToCommissionerPositions устанавливает Метаданные позиций возврата на склад комиссионера Полученного отчёта комиссионера.
//
// Принимает множество объектов [CommissionReportInReturnPosition].
func (commissionReportIn *CommissionReportIn) SetReturnToCommissionerPositions(returnToCommissionerPositions ...*CommissionReportInReturnPosition) *CommissionReportIn {
	commissionReportIn.ReturnToCommissionerPositions = NewMetaArrayFrom(returnToCommissionerPositions)
	return commissionReportIn
}

// SetRewardPercent устанавливает Процент вознаграждения.
func (commissionReportIn *CommissionReportIn) SetRewardPercent(rewardPercent float64) *CommissionReportIn {
	commissionReportIn.RewardPercent = &rewardPercent
	return commissionReportIn
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [AsPaymentInterface].
func (commissionReportIn *CommissionReportIn) SetPayments(payments ...AsPaymentInterface) *CommissionReportIn {
	commissionReportIn.Payments = NewPaymentsFrom(payments)
	return commissionReportIn
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil передаёт сброс значения (null).
func (commissionReportIn *CommissionReportIn) SetSalesChannel(salesChannel *SalesChannel) *CommissionReportIn {
	commissionReportIn.SalesChannel = NewNullValue(salesChannel)
	return commissionReportIn
}

// SetShared устанавливает флаг общего доступа.
func (commissionReportIn *CommissionReportIn) SetShared(shared bool) *CommissionReportIn {
	commissionReportIn.Shared = &shared
	return commissionReportIn
}

// SetState устанавливает Метаданные статуса Приходного ордера.
//
// Передача nil передаёт сброс значения (null).
func (commissionReportIn *CommissionReportIn) SetState(state *State) *CommissionReportIn {
	commissionReportIn.State = NewNullValue(state)
	return commissionReportIn
}

// SetSyncID устанавливает ID синхронизации.
func (commissionReportIn *CommissionReportIn) SetSyncID(syncID uuid.UUID) *CommissionReportIn {
	commissionReportIn.SyncID = &syncID
	return commissionReportIn
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Полученного отчёта комиссионера.
func (commissionReportIn *CommissionReportIn) SetVatEnabled(vatEnabled bool) *CommissionReportIn {
	commissionReportIn.VatEnabled = &vatEnabled
	return commissionReportIn
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (commissionReportIn *CommissionReportIn) SetVatIncluded(vatIncluded bool) *CommissionReportIn {
	commissionReportIn.VatIncluded = &vatIncluded
	return commissionReportIn
}

// SetRewardType устанавливает Тип вознаграждения.
func (commissionReportIn *CommissionReportIn) SetRewardType(rewardType RewardType) *CommissionReportIn {
	commissionReportIn.RewardType = rewardType
	return commissionReportIn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (commissionReportIn *CommissionReportIn) SetAttributes(attributes ...*Attribute) *CommissionReportIn {
	commissionReportIn.Attributes.Push(attributes...)
	return commissionReportIn
}

// String реализует интерфейс [fmt.Stringer].
func (commissionReportIn CommissionReportIn) String() string {
	return Stringify(commissionReportIn)
}

// MetaType возвращает код сущности.
func (CommissionReportIn) MetaType() MetaType {
	return MetaTypeCommissionReportIn
}

// Update shortcut
func (commissionReportIn CommissionReportIn) Update(ctx context.Context, client *Client, params ...*Params) (*CommissionReportIn, *resty.Response, error) {
	return NewCommissionReportInService(client).Update(ctx, commissionReportIn.GetID(), &commissionReportIn, params...)
}

// Create shortcut
func (commissionReportIn CommissionReportIn) Create(ctx context.Context, client *Client, params ...*Params) (*CommissionReportIn, *resty.Response, error) {
	return NewCommissionReportInService(client).Create(ctx, &commissionReportIn, params...)
}

// Delete shortcut
func (commissionReportIn CommissionReportIn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCommissionReportInService(client).Delete(ctx, commissionReportIn.GetID())
}

// CommissionOverhead Прочие расходы.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-prochie-rashody
type CommissionOverhead struct {
	Sum *float64 `json:"sum,omitempty"` // Сумма в копейках
}

// GetSum возвращает сумму в копейках.
func (commissionOverhead CommissionOverhead) GetSum() float64 {
	return Deref(commissionOverhead.Sum)
}

// SetSum устанавливает сумму в копейках.
func (commissionOverhead *CommissionOverhead) SetSum(sum float64) *CommissionOverhead {
	commissionOverhead.Sum = &sum
	return commissionOverhead
}

// String реализует интерфейс [fmt.Stringer].
func (commissionOverhead CommissionOverhead) String() string {
	return Stringify(commissionOverhead)
}

// CommissionReportInPosition Позиция Полученного отчёта комиссионера.
//
// Код сущности: commissionreportinposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-pozicii-poluchennogo-otcheta-komissionera
type CommissionReportInPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции.
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

// GetAccountID возвращает ID учётной записи.
func (commissionReportInPosition CommissionReportInPosition) GetAccountID() uuid.UUID {
	return Deref(commissionReportInPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (commissionReportInPosition CommissionReportInPosition) GetAssortment() AssortmentPosition {
	return Deref(commissionReportInPosition.Assortment)
}

// GetID возвращает ID Позиции Полученного отчёта комиссионера.
func (commissionReportInPosition CommissionReportInPosition) GetID() uuid.UUID {
	return Deref(commissionReportInPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (commissionReportInPosition CommissionReportInPosition) GetPack() Pack {
	return Deref(commissionReportInPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (commissionReportInPosition CommissionReportInPosition) GetPrice() float64 {
	return Deref(commissionReportInPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (commissionReportInPosition CommissionReportInPosition) GetQuantity() float64 {
	return Deref(commissionReportInPosition.Quantity)
}

// GetReward возвращает Вознаграждение.
func (commissionReportInPosition CommissionReportInPosition) GetReward() float64 {
	return Deref(commissionReportInPosition.Reward)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (commissionReportInPosition CommissionReportInPosition) GetVat() int {
	return Deref(commissionReportInPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (commissionReportInPosition CommissionReportInPosition) GetVatEnabled() bool {
	return Deref(commissionReportInPosition.VatEnabled)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (commissionReportInPosition *CommissionReportInPosition) SetAssortment(assortment AssortmentInterface) *CommissionReportInPosition {
	if assortment != nil {
		commissionReportInPosition.Assortment = assortment.asAssortment()
	}
	return commissionReportInPosition
}

// SetPack устанавливает Упаковку Товара.
func (commissionReportInPosition *CommissionReportInPosition) SetPack(pack *Pack) *CommissionReportInPosition {
	if pack != nil {
		commissionReportInPosition.Pack = pack
	}
	return commissionReportInPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (commissionReportInPosition *CommissionReportInPosition) SetPrice(price float64) *CommissionReportInPosition {
	commissionReportInPosition.Price = &price
	return commissionReportInPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (commissionReportInPosition *CommissionReportInPosition) SetQuantity(quantity float64) *CommissionReportInPosition {
	commissionReportInPosition.Quantity = &quantity
	return commissionReportInPosition
}

// SetReward устанавливает Вознаграждение.
func (commissionReportInPosition *CommissionReportInPosition) SetReward(reward float64) *CommissionReportInPosition {
	commissionReportInPosition.Reward = &reward
	return commissionReportInPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (commissionReportInPosition *CommissionReportInPosition) SetVat(vat int) *CommissionReportInPosition {
	commissionReportInPosition.Vat = &vat
	return commissionReportInPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (commissionReportInPosition *CommissionReportInPosition) SetVatEnabled(vatEnabled bool) *CommissionReportInPosition {
	commissionReportInPosition.VatEnabled = &vatEnabled
	return commissionReportInPosition
}

// String реализует интерфейс [fmt.Stringer].
func (commissionReportInPosition CommissionReportInPosition) String() string {
	return Stringify(commissionReportInPosition)
}

// MetaType возвращает код сущности.
func (CommissionReportInPosition) MetaType() MetaType {
	return MetaTypeCommissionReportInPosition
}

// CommissionReportInReturnPosition Позиция возврата на склад комиссионера.
//
// Код сущности: commissionreportinreturnedposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-pozicii-poluchennogo-otcheta-komissionera-ob-ekt-pozicii-wozwrata-na-sklad-komissionera-soderzhit-sleduuschie-polq
type CommissionReportInReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

// GetAccountID возвращает ID учётной записи.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetAccountID() uuid.UUID {
	return Deref(commissionReportInReturnPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(commissionReportInReturnPosition.Assortment)
}

// GetID возвращает ID позиции.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetID() uuid.UUID {
	return Deref(commissionReportInReturnPosition.ID)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetPrice() float64 {
	return Deref(commissionReportInReturnPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetQuantity() float64 {
	return Deref(commissionReportInReturnPosition.Quantity)
}

// GetReward возвращает Вознаграждение.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetReward() float64 {
	return Deref(commissionReportInReturnPosition.Reward)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetVat() int {
	return Deref(commissionReportInReturnPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (commissionReportInReturnPosition CommissionReportInReturnPosition) GetVatEnabled() bool {
	return Deref(commissionReportInReturnPosition.VatEnabled)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetAssortment(assortment AssortmentInterface) *CommissionReportInReturnPosition {
	if assortment != nil {
		commissionReportInReturnPosition.Assortment = assortment.asAssortment()
	}
	return commissionReportInReturnPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetPrice(price float64) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Price = &price
	return commissionReportInReturnPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetQuantity(quantity float64) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Quantity = &quantity
	return commissionReportInReturnPosition
}

// SetReward устанавливает Вознаграждение.
func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetReward(reward float64) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Reward = &reward
	return commissionReportInReturnPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetVat(vat int) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.Vat = &vat
	return commissionReportInReturnPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (commissionReportInReturnPosition *CommissionReportInReturnPosition) SetVatEnabled(vatEnabled bool) *CommissionReportInReturnPosition {
	commissionReportInReturnPosition.VatEnabled = &vatEnabled
	return commissionReportInReturnPosition
}

// String реализует интерфейс [fmt.Stringer].
func (commissionReportInReturnPosition CommissionReportInReturnPosition) String() string {
	return Stringify(commissionReportInReturnPosition)
}

// MetaType возвращает код сущности.
func (CommissionReportInReturnPosition) MetaType() MetaType {
	return MetaTypeCommissionReportInReturnPosition
}

// CommissionReportInService методы сервиса для работы с полученными отчётами комиссионера.
type CommissionReportInService interface {
	// GetList выполняет запрос на получение списка полученных отчётов комиссионера.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[CommissionReportIn], *resty.Response, error)

	// Create выполняет запрос на создание полученного отчёта комиссионера.
	// Обязательные поля для заполнения:
	//	- agent (Контрагент)
	//	- contract (Договор)
	//	- organization (Юрлицо)
	//	- commissionPeriodStart (Начало периода)
	//	- commissionPeriodEnd (Конец периода)
	//	- organizationAccount (Счет юрлица, если у юрлица несколько счетов)
	//	- agentAccount (Счет контрагента, если у контрагента несколько счетов)
	// Принимает контекст, расходный ордер и опционально объект параметров запроса Params.
	// Возвращает созданный расходный ордер.
	Create(ctx context.Context, commissionReportIn *CommissionReportIn, params ...*Params) (*CommissionReportIn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение полученных отчётов комиссионера.
	// Изменяемые полученные отчёты комиссионера должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список полученных отчётов комиссионера и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых полученных отчётов комиссионера.
	CreateUpdateMany(ctx context.Context, commissionReportInList Slice[CommissionReportIn], params ...*Params) (*Slice[CommissionReportIn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление полученных отчётов комиссионера.
	// Принимает контекст и множество полученных отчётов комиссионера.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*CommissionReportIn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление полученного отчёта комиссионера.
	// Принимает контекст и ID полученного отчёта комиссионера.
	// Возвращает true в случае успешного удаления полученного отчёта комиссионера.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного полученного отчёта комиссионера по ID.
	// Принимает контекст, ID полученного отчёта комиссионера и опционально объект параметров запроса Params.
	// Возвращает найденный полученный отчёт комиссионера.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*CommissionReportIn, *resty.Response, error)

	// Update выполняет запрос на изменение полученного отчёта комиссионера.
	// Принимает контекст, полученный отчёт комиссионера и опционально объект параметров запроса Params.
	// Возвращает изменённый полученный отчёт комиссионера.
	Update(ctx context.Context, id uuid.UUID, commissionReportIn *CommissionReportIn, params ...*Params) (*CommissionReportIn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных полученных отчётов комиссионера.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[CommissionReportInPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*CommissionReportInPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *CommissionReportInPosition, params ...*Params) (*CommissionReportInPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *CommissionReportInPosition, params ...*Params) (*CommissionReportInPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*CommissionReportInPosition) (*Slice[CommissionReportInPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает true в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*CommissionReportInPosition) (*DeleteManyResponse, *resty.Response, error)

	// GetPositionTrackingCodeList выполняет запрос на получение кодов маркировки позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект List.
	GetPositionTrackingCodeList(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*List[TrackingCode], *resty.Response, error)

	// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание/изменение кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает список созданных и/или изменённых кодов маркировки позиции документа.
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)

	// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает true в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*CommissionReportIn, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает true в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// GetPublicationList выполняет запрос на получение списка публикаций.
	// Принимает контекст и ID документа.
	// Возвращает объект List.
	GetPublicationList(ctx context.Context, id uuid.UUID) (*List[Publication], *resty.Response, error)

	// GetPublicationByID выполняет запрос на получение отдельной публикации по ID.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает найденную публикацию.
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)

	// Publish выполняет запрос на создание публикации.
	// Принимает контекст, ID документа и шаблон.
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает true в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает true в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetReturnPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetReturnPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[CommissionReportInReturnPosition], *resty.Response, error)

	// GetReturnPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetReturnPositionByID(ctx context.Context, id, positionID uuid.UUID, params ...*Params) (*CommissionReportInReturnPosition, *resty.Response, error)

	// CreateReturnPosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreateReturnPosition(ctx context.Context, id uuid.UUID, position *CommissionReportInReturnPosition) (*CommissionReportInReturnPosition, *resty.Response, error)

	// UpdateReturnPosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdateReturnPosition(ctx context.Context, id, positionID uuid.UUID, position *CommissionReportInReturnPosition, params ...*Params) (*CommissionReportInReturnPosition, *resty.Response, error)

	// DeleteReturnPosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает true в случае успешного удаления позиции.
	DeleteReturnPosition(ctx context.Context, id, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeleteReturnPositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteReturnPositionMany(ctx context.Context, id uuid.UUID, positions ...*CommissionReportInReturnPosition) (*DeleteManyResponse, *resty.Response, error)

	// CreateState выполняет запрос на создание статуса документа.
	// Принимает контекст и статус.
	// Возвращает созданный статус.
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)

	// UpdateState выполняет запрос на изменение статуса документа.
	// Принимает контекст, ID статуса и статус.
	// Возвращает изменённый статус.
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)

	// CreateUpdateStateMany выполняет запрос на массовое создание и/или изменение статусов документа.
	// Принимает контекст и множество статусов.
	// Возвращает список созданных и/или изменённых статусов.
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)

	// DeleteState выполняет запрос на удаление статуса документа.
	// Принимает контекст и ID статуса.
	// Возвращает true в случае успешного удаления статуса.
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetFileList выполняет запрос на получение файлов в виде списка.
	// Принимает контекст и ID сущности/документа.
	// Возвращает объект List.
	GetFileList(ctx context.Context, id uuid.UUID) (*List[File], *resty.Response, error)

	// CreateFile выполняет запрос на добавление файла.
	// Принимает контекст, ID сущности/документа и файл.
	// Возвращает список файлов.
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)

	// UpdateFileMany выполняет запрос на массовое создание и/или изменение файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает созданных и/или изменённых файлов.
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)

	// DeleteFile выполняет запрос на удаление файла сущности/документа.
	// Принимает контекст, ID сущности/документа и ID файла.
	// Возвращает true в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)

	// Evaluate выполняет запрос на получение шаблона документа с автозаполнением.
	// Принимает контекст, документ и множество значений Evaluate.
	//
	// Возможные значения типа Evaluate:
	//	- EvaluateDiscount – скидки
	//	- EvaluatePrice    – цены
	//	- EvaluateVat      – ндс
	//	- EvaluateCost     – себестоимость
	// Возвращает шаблон документа с автозаполнением.
	Evaluate(ctx context.Context, commissionReportIn *CommissionReportIn, evaluate ...Evaluate) (*CommissionReportIn, *resty.Response, error)
}

type commissionReportInService struct {
	Endpoint
	endpointGetList[CommissionReportIn]
	endpointCreate[CommissionReportIn]
	endpointCreateUpdateMany[CommissionReportIn]
	endpointDeleteMany[CommissionReportIn]
	endpointDelete
	endpointGetByID[CommissionReportIn]
	endpointUpdate[CommissionReportIn]
	endpointMetadata[MetaAttributesStatesSharedWrapper]
	endpointPositions[CommissionReportInPosition]
	endpointAttributes
	endpointSyncID[CommissionReportIn]
	endpointNamedFilter
	endpointPublication
	endpointTrash
	endpointStates
	endpointFiles
	endpointTemplate[CommissionReportIn]
	endpointEvaluate[CommissionReportIn]
}

// NewCommissionReportInService принимает [Client] и возвращает сервис для работы с полученными отчётами комиссионера.
func NewCommissionReportInService(client *Client) CommissionReportInService {
	e := NewEndpoint(client, "entity/commissionreportin")
	return &commissionReportInService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[CommissionReportIn]{e},
		endpointCreate:           endpointCreate[CommissionReportIn]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[CommissionReportIn]{e},
		endpointDeleteMany:       endpointDeleteMany[CommissionReportIn]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[CommissionReportIn]{e},
		endpointUpdate:           endpointUpdate[CommissionReportIn]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesStatesSharedWrapper]{e},
		endpointPositions:        endpointPositions[CommissionReportInPosition]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[CommissionReportIn]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointTemplate:         endpointTemplate[CommissionReportIn]{e},
		endpointPublication:      endpointPublication{e},
		endpointTrash:            endpointTrash{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
		endpointEvaluate:         endpointEvaluate[CommissionReportIn]{e},
	}
}

func (service *commissionReportInService) GetReturnPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[CommissionReportInReturnPosition], *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[List[CommissionReportInReturnPosition]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *commissionReportInService) GetReturnPositionByID(ctx context.Context, id, positionID uuid.UUID, params ...*Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionID)
	return NewRequestBuilder[CommissionReportInReturnPosition](service.client, path).SetParams(params...).Get(ctx)
}

func (service *commissionReportInService) CreateReturnPosition(ctx context.Context, id uuid.UUID, position *CommissionReportInReturnPosition) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions", id)
	return NewRequestBuilder[CommissionReportInReturnPosition](service.client, path).Post(ctx, position)
}

func (service *commissionReportInService) UpdateReturnPosition(ctx context.Context, id, positionID uuid.UUID, position *CommissionReportInReturnPosition, params ...*Params) (*CommissionReportInReturnPosition, *resty.Response, error) {
	path := fmt.Sprintf("%s/returntocommissionerpositions/%s", id, positionID)
	return NewRequestBuilder[CommissionReportInReturnPosition](service.client, path).SetParams(params...).Put(ctx, position)
}

func (service *commissionReportInService) DeleteReturnPosition(ctx context.Context, id, positionID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/positions/%s", id, positionID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *commissionReportInService) DeleteReturnPositionMany(ctx context.Context, id uuid.UUID, entities ...*CommissionReportInReturnPosition) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/positions/delete", id)
	return NewRequestBuilder[DeleteManyResponse](service.client, path).Post(ctx, entities)
}
