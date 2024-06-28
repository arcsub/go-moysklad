package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CommissionReportOut Выданный отчёт комиссионера.
//
// Код сущности: commissionreportout
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera
type CommissionReportOut struct {
	Applicable            *bool                                   `json:"applicable,omitempty"`            // Отметка о проведении
	OrganizationAccount   *AgentAccount                           `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	AgentAccount          *AgentAccount                           `json:"agentAccount,omitempty"`          // Метаданные счета контрагента
	Organization          *Organization                           `json:"organization,omitempty"`          // Метаданные юрлица
	VatSum                *float64                                `json:"vatSum,omitempty"`                // Сумма НДС
	Code                  *string                                 `json:"code,omitempty"`                  // Код Выданного отчета комиссионера
	CommissionPeriodEnd   *Timestamp                              `json:"commissionPeriodEnd,omitempty"`   // Конец периода
	Agent                 *Counterparty                           `json:"agent,omitempty"`                 // Метаданные контрагента
	CommitentSum          *float64                                `json:"commitentSum,omitempty"`          // Сумма коммитента в установленной валюте
	Contract              *Contract                               `json:"contract,omitempty"`              // Метаданные договора
	Created               *Timestamp                              `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                              `json:"deleted,omitempty"`               // Момент последнего удаления Выданного отчета комиссионера
	Description           *string                                 `json:"description,omitempty"`           // Комментарий Выданного отчета комиссионера
	ExternalCode          *string                                 `json:"externalCode,omitempty"`          // Внешний код Выданного отчета комиссионера
	Files                 *MetaArray[File]                        `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                                  `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                              `json:"id,omitempty"`                    // ID Выданного отчета комиссионера
	Meta                  *Meta                                   `json:"meta,omitempty"`                  // Метаданные Выданного отчета комиссионера
	Moment                *Timestamp                              `json:"moment,omitempty"`                // Дата документа
	Name                  *string                                 `json:"name,omitempty"`                  // Наименование Выданного отчета комиссионера
	AccountID             *uuid.UUID                              `json:"accountId,omitempty"`             // ID учётной записи
	CommissionPeriodStart *Timestamp                              `json:"commissionPeriodStart,omitempty"` // Начало периода
	Owner                 *Employee                               `json:"owner,omitempty"`                 // Метаданные владельца (Сотрудника)
	PayedSum              *float64                                `json:"payedSum,omitempty"`              // Оплаченная сумма
	Positions             *Positions[CommissionReportOutPosition] `json:"positions,omitempty"`             // Метаданные позиций Выданного отчета
	Printed               *bool                                   `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *NullValue[Project]                     `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                                   `json:"published,omitempty"`             // Опубликован ли документ
	Rate                  *NullValue[Rate]                        `json:"rate,omitempty"`                  // Валюта
	RewardPercent         *float64                                `json:"rewardPercent,omitempty"`         // Процент вознаграждения (всегда 0 если вознаграждение не рассчитывается)
	Payments              Slice[Payment]                          `json:"payments,omitempty"`              // Массив ссылок на связанные платежи
	SalesChannel          *NullValue[SalesChannel]                `json:"salesChannel,omitempty"`          // Метаданные канала продаж
	Shared                *bool                                   `json:"shared,omitempty"`                // Общий доступ
	State                 *NullValue[State]                       `json:"state,omitempty"`                 // Метаданные статуса Выданного отчета комиссионера
	Sum                   *float64                                `json:"sum,omitempty"`                   // Сумма Выданного отчета комиссионера в копейках
	SyncID                *uuid.UUID                              `json:"syncId,omitempty"`                // ID синхронизации
	Updated               *Timestamp                              `json:"updated,omitempty"`               // Момент последнего обновления Выданного отчета комиссионера
	VatEnabled            *bool                                   `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                                   `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	RewardType            RewardType                              `json:"rewardType,omitempty"`            // Тип вознаграждения
	Attributes            Slice[Attribute]                        `json:"attributes,omitempty"`            // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (commissionReportOut CommissionReportOut) Clean() *CommissionReportOut {
	if commissionReportOut.Meta == nil {
		return nil
	}
	return &CommissionReportOut{Meta: commissionReportOut.Meta}
}

// AsOperation возвращает объект [Operation] c полями meta и linkedSum.
// Значение поля linkedSum заполняется из поля sum.
func (commissionReportOut CommissionReportOut) AsOperation() *Operation {
	return &Operation{Meta: commissionReportOut.GetMeta()}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface.
func (commissionReportOut CommissionReportOut) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: commissionReportOut.Meta}
}

// GetApplicable возвращает Отметку о проведении.
func (commissionReportOut CommissionReportOut) GetApplicable() bool {
	return Deref(commissionReportOut.Applicable)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (commissionReportOut CommissionReportOut) GetOrganizationAccount() AgentAccount {
	return Deref(commissionReportOut.OrganizationAccount)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (commissionReportOut CommissionReportOut) GetAgentAccount() AgentAccount {
	return Deref(commissionReportOut.AgentAccount)
}

// GetOrganization возвращает Метаданные юрлица.
func (commissionReportOut CommissionReportOut) GetOrganization() Organization {
	return Deref(commissionReportOut.Organization)
}

// GetVatSum возвращает Сумму НДС.
func (commissionReportOut CommissionReportOut) GetVatSum() float64 {
	return Deref(commissionReportOut.VatSum)
}

// GetCode возвращает Код Выданного отчета комиссионера.
func (commissionReportOut CommissionReportOut) GetCode() string {
	return Deref(commissionReportOut.Code)
}

// GetCommissionPeriodEnd возвращает Конец периода.
func (commissionReportOut CommissionReportOut) GetCommissionPeriodEnd() Timestamp {
	return Deref(commissionReportOut.CommissionPeriodEnd)
}

// GetAgent возвращает Метаданные контрагента.
func (commissionReportOut CommissionReportOut) GetAgent() Counterparty {
	return Deref(commissionReportOut.Agent)
}

// GetCommitentSum возвращает Сумму комитента в установленной валюте.
func (commissionReportOut CommissionReportOut) GetCommitentSum() float64 {
	return Deref(commissionReportOut.CommitentSum)
}

// GetContract возвращает Метаданные договора.
func (commissionReportOut CommissionReportOut) GetContract() Contract {
	return Deref(commissionReportOut.Contract)
}

// GetCreated возвращает Дату создания.
func (commissionReportOut CommissionReportOut) GetCreated() Timestamp {
	return Deref(commissionReportOut.Created)
}

// GetDeleted возвращает Момент последнего удаления Выданного отчёта комиссионера.
func (commissionReportOut CommissionReportOut) GetDeleted() Timestamp {
	return Deref(commissionReportOut.Deleted)
}

// GetDescription возвращает Комментарий Выданного отчёта комиссионера.
func (commissionReportOut CommissionReportOut) GetDescription() string {
	return Deref(commissionReportOut.Description)
}

// GetExternalCode возвращает Внешний код Выданного отчёта комиссионера.
func (commissionReportOut CommissionReportOut) GetExternalCode() string {
	return Deref(commissionReportOut.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (commissionReportOut CommissionReportOut) GetFiles() MetaArray[File] {
	return Deref(commissionReportOut.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (commissionReportOut CommissionReportOut) GetGroup() Group {
	return Deref(commissionReportOut.Group)
}

// GetID возвращает ID Выданного отчёта комиссионера.
func (commissionReportOut CommissionReportOut) GetID() uuid.UUID {
	return Deref(commissionReportOut.ID)
}

// GetMeta возвращает Метаданные Выданного отчёта комиссионера.
func (commissionReportOut CommissionReportOut) GetMeta() Meta {
	return Deref(commissionReportOut.Meta)
}

// GetMoment возвращает Дату документа.
func (commissionReportOut CommissionReportOut) GetMoment() Timestamp {
	return Deref(commissionReportOut.Moment)
}

// GetName возвращает Наименование Выданного отчёта комиссионера.
func (commissionReportOut CommissionReportOut) GetName() string {
	return Deref(commissionReportOut.Name)
}

// GetAccountID возвращает ID учётной записи.
func (commissionReportOut CommissionReportOut) GetAccountID() uuid.UUID {
	return Deref(commissionReportOut.AccountID)
}

// GetCommissionPeriodStart возвращает Начало периода.
func (commissionReportOut CommissionReportOut) GetCommissionPeriodStart() Timestamp {
	return Deref(commissionReportOut.CommissionPeriodStart)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (commissionReportOut CommissionReportOut) GetOwner() Employee {
	return Deref(commissionReportOut.Owner)
}

// GetPayedSum возвращает Оплаченную сумму.
func (commissionReportOut CommissionReportOut) GetPayedSum() float64 {
	return Deref(commissionReportOut.PayedSum)
}

// GetPositions возвращает Метаданные позиций Выданного отчёта.
func (commissionReportOut CommissionReportOut) GetPositions() Positions[CommissionReportOutPosition] {
	return Deref(commissionReportOut.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (commissionReportOut CommissionReportOut) GetPrinted() bool {
	return Deref(commissionReportOut.Printed)
}

// GetProject возвращает Метаданные проекта.
func (commissionReportOut CommissionReportOut) GetProject() Project {
	return commissionReportOut.Project.Get()
}

// GetPublished возвращает true, если документ опубликован.
func (commissionReportOut CommissionReportOut) GetPublished() bool {
	return Deref(commissionReportOut.Published)
}

// GetRate возвращает Валюту.
func (commissionReportOut CommissionReportOut) GetRate() Rate {
	return commissionReportOut.Rate.Get()
}

// GetRewardPercent возвращает Процент вознаграждения (всегда 0 если вознаграждение не рассчитывается).
func (commissionReportOut CommissionReportOut) GetRewardPercent() float64 {
	return Deref(commissionReportOut.RewardPercent)
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (commissionReportOut CommissionReportOut) GetPayments() Slice[Payment] {
	return commissionReportOut.Payments
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (commissionReportOut CommissionReportOut) GetSalesChannel() SalesChannel {
	return commissionReportOut.SalesChannel.Get()
}

// GetShared возвращает флаг общего доступа.
func (commissionReportOut CommissionReportOut) GetShared() bool {
	return Deref(commissionReportOut.Shared)
}

// GetState возвращает Метаданные статуса.
func (commissionReportOut CommissionReportOut) GetState() State {
	return commissionReportOut.State.Get()
}

// GetSum возвращает Сумму Выданного отчёта комиссионера в копейках.
func (commissionReportOut CommissionReportOut) GetSum() float64 {
	return Deref(commissionReportOut.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (commissionReportOut CommissionReportOut) GetSyncID() uuid.UUID {
	return Deref(commissionReportOut.SyncID)
}

// GetUpdated возвращает Момент последнего обновления.
func (commissionReportOut CommissionReportOut) GetUpdated() Timestamp {
	return Deref(commissionReportOut.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (commissionReportOut CommissionReportOut) GetVatEnabled() bool {
	return Deref(commissionReportOut.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (commissionReportOut CommissionReportOut) GetVatIncluded() bool {
	return Deref(commissionReportOut.VatIncluded)
}

// GetRewardType возвращает Тип вознаграждения.
func (commissionReportOut CommissionReportOut) GetRewardType() RewardType {
	return commissionReportOut.RewardType
}

// GetAttributes возвращает Список метаданных доп. полей.
func (commissionReportOut CommissionReportOut) GetAttributes() Slice[Attribute] {
	return commissionReportOut.Attributes
}

// SetApplicable устанавливает Отметку о проведении.
func (commissionReportOut *CommissionReportOut) SetApplicable(applicable bool) *CommissionReportOut {
	commissionReportOut.Applicable = &applicable
	return commissionReportOut
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (commissionReportOut *CommissionReportOut) SetOrganizationAccount(organizationAccount *AgentAccount) *CommissionReportOut {
	if organizationAccount != nil {
		commissionReportOut.OrganizationAccount = organizationAccount.Clean()
	}
	return commissionReportOut
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (commissionReportOut *CommissionReportOut) SetAgentAccount(agentAccount *AgentAccount) *CommissionReportOut {
	if agentAccount != nil {
		commissionReportOut.AgentAccount = agentAccount.Clean()
	}
	return commissionReportOut
}

// SetOrganization устанавливает Метаданные юрлица.
func (commissionReportOut *CommissionReportOut) SetOrganization(organization *Organization) *CommissionReportOut {
	if organization != nil {
		commissionReportOut.Organization = organization.Clean()
	}
	return commissionReportOut
}

// SetCode устанавливает Код Выданного отчёта комиссионера.
func (commissionReportOut *CommissionReportOut) SetCode(code string) *CommissionReportOut {
	commissionReportOut.Code = &code
	return commissionReportOut
}

// SetCommissionPeriodEnd устанавливает Конец периода.
func (commissionReportOut *CommissionReportOut) SetCommissionPeriodEnd(commissionPeriodEnd *Timestamp) *CommissionReportOut {
	commissionReportOut.CommissionPeriodEnd = commissionPeriodEnd
	return commissionReportOut
}

// SetAgent устанавливает Метаданные контрагента.
func (commissionReportOut *CommissionReportOut) SetAgent(agent *Counterparty) *CommissionReportOut {
	if agent != nil {
		commissionReportOut.Agent = agent.Clean()
	}
	return commissionReportOut
}

// SetContract устанавливает Метаданные договора.
func (commissionReportOut *CommissionReportOut) SetContract(contract *Contract) *CommissionReportOut {
	if contract != nil {
		commissionReportOut.Contract = contract.Clean()
	}
	return commissionReportOut
}

// SetDescription устанавливает Комментарий Выданного отчёта комиссионера.
func (commissionReportOut *CommissionReportOut) SetDescription(description string) *CommissionReportOut {
	commissionReportOut.Description = &description
	return commissionReportOut
}

// SetExternalCode устанавливает Внешний код Выданного отчёта комиссионера.
func (commissionReportOut *CommissionReportOut) SetExternalCode(externalCode string) *CommissionReportOut {
	commissionReportOut.ExternalCode = &externalCode
	return commissionReportOut
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (commissionReportOut *CommissionReportOut) SetFiles(files ...*File) *CommissionReportOut {
	commissionReportOut.Files = NewMetaArrayFrom(files)
	return commissionReportOut
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (commissionReportOut *CommissionReportOut) SetGroup(group *Group) *CommissionReportOut {
	if group != nil {
		commissionReportOut.Group = group.Clean()
	}
	return commissionReportOut
}

// SetMeta устанавливает Метаданные Выданного отчёта комиссионера.
func (commissionReportOut *CommissionReportOut) SetMeta(meta *Meta) *CommissionReportOut {
	commissionReportOut.Meta = meta
	return commissionReportOut
}

// SetMoment устанавливает Дату документа.
func (commissionReportOut *CommissionReportOut) SetMoment(moment *Timestamp) *CommissionReportOut {
	commissionReportOut.Moment = moment
	return commissionReportOut
}

// SetName устанавливает Наименование Выданного отчёта комиссионера.
func (commissionReportOut *CommissionReportOut) SetName(name string) *CommissionReportOut {
	commissionReportOut.Name = &name
	return commissionReportOut
}

// SetCommissionPeriodStart устанавливает Начало периода.
func (commissionReportOut *CommissionReportOut) SetCommissionPeriodStart(commissionPeriodStart *Timestamp) *CommissionReportOut {
	commissionReportOut.CommissionPeriodStart = commissionPeriodStart
	return commissionReportOut
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (commissionReportOut *CommissionReportOut) SetOwner(owner *Employee) *CommissionReportOut {
	if owner != nil {
		commissionReportOut.Owner = owner.Clean()
	}
	return commissionReportOut
}

// SetPositions устанавливает Метаданные позиций реализовано комиссионером Полученного отчёта комиссионера.
//
// Принимает множество объектов [CommissionReportInPosition].
func (commissionReportOut *CommissionReportOut) SetPositions(positions ...*CommissionReportOutPosition) *CommissionReportOut {
	commissionReportOut.Positions = NewPositionsFrom(positions)
	return commissionReportOut
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (commissionReportOut *CommissionReportOut) SetProject(project *Project) *CommissionReportOut {
	if project == nil {
		commissionReportOut.Project = NewNullValue[Project]()
	} else {
		commissionReportOut.Project = NewNullValueFrom(project.Clean())
	}
	return commissionReportOut
}

// SetRate устанавливает Валюту.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (commissionReportOut *CommissionReportOut) SetRate(rate *Rate) *CommissionReportOut {
	if rate == nil {
		commissionReportOut.Rate = NewNullValue[Rate]()
	} else {
		commissionReportOut.Rate = NewNullValueFrom(rate)
	}
	return commissionReportOut
}

// SetRewardPercent устанавливает Процент вознаграждения.
func (commissionReportOut *CommissionReportOut) SetRewardPercent(rewardPercent float64) *CommissionReportOut {
	commissionReportOut.RewardPercent = &rewardPercent
	return commissionReportOut
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [AsPaymentInterface].
func (commissionReportOut *CommissionReportOut) SetPayments(payments ...AsPaymentInterface) *CommissionReportOut {
	commissionReportOut.Payments = NewPaymentsFrom(payments)
	return commissionReportOut
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (commissionReportOut *CommissionReportOut) SetSalesChannel(salesChannel *SalesChannel) *CommissionReportOut {
	if salesChannel == nil {
		commissionReportOut.SalesChannel = NewNullValue[SalesChannel]()
	} else {
		commissionReportOut.SalesChannel = NewNullValueFrom(salesChannel.Clean())
	}
	return commissionReportOut
}

// SetShared устанавливает флаг общего доступа.
func (commissionReportOut *CommissionReportOut) SetShared(shared bool) *CommissionReportOut {
	commissionReportOut.Shared = &shared
	return commissionReportOut
}

// SetState устанавливает Метаданные статуса Выданного отчёта комиссионера.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (commissionReportOut *CommissionReportOut) SetState(state *State) *CommissionReportOut {
	if state == nil {
		commissionReportOut.State = NewNullValue[State]()
	} else {
		commissionReportOut.State = NewNullValueFrom(state.Clean())
	}
	return commissionReportOut
}

// SetSyncID устанавливает ID синхронизации.
func (commissionReportOut *CommissionReportOut) SetSyncID(syncID uuid.UUID) *CommissionReportOut {
	commissionReportOut.SyncID = &syncID
	return commissionReportOut
}

// SetVatEnabled устанавливает флаг включения НДС для Выданного отчёта комиссионера.
func (commissionReportOut *CommissionReportOut) SetVatEnabled(vatEnabled bool) *CommissionReportOut {
	commissionReportOut.VatEnabled = &vatEnabled
	return commissionReportOut
}

// SetVatIncluded устанавливает флаг включения НДС.
func (commissionReportOut *CommissionReportOut) SetVatIncluded(vatIncluded bool) *CommissionReportOut {
	commissionReportOut.VatIncluded = &vatIncluded
	return commissionReportOut
}

// SetRewardType устанавливает Тип вознаграждения.
func (commissionReportOut *CommissionReportOut) SetRewardType(rewardType RewardType) *CommissionReportOut {
	commissionReportOut.RewardType = rewardType
	return commissionReportOut
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (commissionReportOut *CommissionReportOut) SetAttributes(attributes ...*Attribute) *CommissionReportOut {
	commissionReportOut.Attributes = attributes
	return commissionReportOut
}

// String реализует интерфейс [fmt.Stringer].
func (commissionReportOut CommissionReportOut) String() string {
	return Stringify(commissionReportOut)
}

// MetaType возвращает код сущности.
func (CommissionReportOut) MetaType() MetaType {
	return MetaTypeCommissionReportOut
}

// Update shortcut
func (commissionReportOut CommissionReportOut) Update(ctx context.Context, client *Client, params ...*Params) (*CommissionReportOut, *resty.Response, error) {
	return NewCommissionReportOutService(client).Update(ctx, commissionReportOut.GetID(), &commissionReportOut, params...)
}

// Create shortcut
func (commissionReportOut CommissionReportOut) Create(ctx context.Context, client *Client, params ...*Params) (*CommissionReportOut, *resty.Response, error) {
	return NewCommissionReportOutService(client).Create(ctx, &commissionReportOut, params...)
}

// Delete shortcut
func (commissionReportOut CommissionReportOut) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCommissionReportOutService(client).Delete(ctx, commissionReportOut.GetID())
}

// CommissionReportOutPosition Позиция Выданного отчёта комиссионера.
//
// Код сущности: commissionreportoutposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera-vydannye-otchety-komissionera-pozicii-vydannogo-otcheta-komissionera
type CommissionReportOutPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

// GetAccountID возвращает ID учётной записи.
func (commissionReportOutPosition CommissionReportOutPosition) GetAccountID() uuid.UUID {
	return Deref(commissionReportOutPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (commissionReportOutPosition CommissionReportOutPosition) GetAssortment() AssortmentPosition {
	return Deref(commissionReportOutPosition.Assortment)
}

// GetID возвращает ID позиции.
func (commissionReportOutPosition CommissionReportOutPosition) GetID() uuid.UUID {
	return Deref(commissionReportOutPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (commissionReportOutPosition CommissionReportOutPosition) GetPack() Pack {
	return Deref(commissionReportOutPosition.Pack)
}

// GetPrice возвращает Цена товара/услуги в копейках.
func (commissionReportOutPosition CommissionReportOutPosition) GetPrice() float64 {
	return Deref(commissionReportOutPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (commissionReportOutPosition CommissionReportOutPosition) GetQuantity() float64 {
	return Deref(commissionReportOutPosition.Quantity)
}

// GetReward возвращает Вознаграждение.
func (commissionReportOutPosition CommissionReportOutPosition) GetReward() float64 {
	return Deref(commissionReportOutPosition.Reward)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (commissionReportOutPosition CommissionReportOutPosition) GetVat() int {
	return Deref(commissionReportOutPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (commissionReportOutPosition CommissionReportOutPosition) GetVatEnabled() bool {
	return Deref(commissionReportOutPosition.VatEnabled)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AsAssortmentInterface].
func (commissionReportOutPosition *CommissionReportOutPosition) SetAssortment(assortment AsAssortmentInterface) *CommissionReportOutPosition {
	commissionReportOutPosition.Assortment = assortment.AsAssortment()
	return commissionReportOutPosition
}

// SetPack устанавливает Упаковку Товара.
func (commissionReportOutPosition *CommissionReportOutPosition) SetPack(pack *Pack) *CommissionReportOutPosition {
	commissionReportOutPosition.Pack = pack
	return commissionReportOutPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (commissionReportOutPosition *CommissionReportOutPosition) SetPrice(price *float64) *CommissionReportOutPosition {
	commissionReportOutPosition.Price = price
	return commissionReportOutPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (commissionReportOutPosition *CommissionReportOutPosition) SetQuantity(quantity float64) *CommissionReportOutPosition {
	commissionReportOutPosition.Quantity = &quantity
	return commissionReportOutPosition
}

// SetReward устанавливает Вознаграждение.
func (commissionReportOutPosition *CommissionReportOutPosition) SetReward(reward *float64) *CommissionReportOutPosition {
	commissionReportOutPosition.Reward = reward
	return commissionReportOutPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (commissionReportOutPosition *CommissionReportOutPosition) SetVat(vat int) *CommissionReportOutPosition {
	commissionReportOutPosition.Vat = &vat
	return commissionReportOutPosition
}

// SetVatEnabled устанавливает флаг включения НДС для текущей позиции.
func (commissionReportOutPosition *CommissionReportOutPosition) SetVatEnabled(vatEnabled bool) *CommissionReportOutPosition {
	commissionReportOutPosition.VatEnabled = &vatEnabled
	return commissionReportOutPosition
}

// String реализует интерфейс [fmt.Stringer].
func (commissionReportOutPosition CommissionReportOutPosition) String() string {
	return Stringify(commissionReportOutPosition)
}

// MetaType возвращает код сущности.
func (CommissionReportOutPosition) MetaType() MetaType {
	return MetaTypeCommissionReportOutPosition
}

// CommissionReportOutService описывает методы сервиса для работы с выданными отчётами комиссионера.
type CommissionReportOutService interface {
	// GetList выполняет запрос на получение списка выданных отчётов комиссионера.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[CommissionReportOut], *resty.Response, error)

	// Create выполняет запрос на создание выданного отчёта комиссионера.
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
	Create(ctx context.Context, commissionReportOut *CommissionReportOut, params ...*Params) (*CommissionReportOut, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение выданных отчётов комиссионера.
	// Изменяемые выданные отчёты комиссионера должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список выданных отчётов комиссионера и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых выданных отчётов комиссионера.
	CreateUpdateMany(ctx context.Context, commissionReportOutList Slice[CommissionReportOut], params ...*Params) (*Slice[CommissionReportOut], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление выданных отчётов комиссионера.
	// Принимает контекст и множество выданных отчётов комиссионера.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*CommissionReportOut) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление выданного отчёта комиссионера.
	// Принимает контекст и ID выданного отчёта комиссионера.
	// Возвращает true в случае успешного удаления выданного отчёта комиссионера.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного выданного отчёта комиссионера по ID.
	// Принимает контекст, ID выданного отчёта комиссионера и опционально объект параметров запроса Params.
	// Возвращает найденный выданный отчёт комиссионера.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*CommissionReportOut, *resty.Response, error)

	// Update выполняет запрос на изменение выданного отчёта комиссионера.
	// Принимает контекст, выданный отчёт комиссионера и опционально объект параметров запроса Params.
	// Возвращает изменённый выданный отчёт комиссионера.
	Update(ctx context.Context, id uuid.UUID, commissionReportOut *CommissionReportOut, params ...*Params) (*CommissionReportOut, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных выданных отчётов комиссионера.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesSharedStatesWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[CommissionReportOutPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*CommissionReportOutPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *CommissionReportOutPosition, params ...*Params) (*CommissionReportOutPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *CommissionReportOutPosition, params ...*Params) (*CommissionReportOutPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*CommissionReportOutPosition) (*Slice[CommissionReportOutPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает true в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*CommissionReportOutPosition) (*DeleteManyResponse, *resty.Response, error)

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

	// GetBySyncID выполняет запрос на получение документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*CommissionReportOut, *resty.Response, error)

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

	// GetStateByID выполняет запрос на получение статуса документа по ID.
	// Принимает контекст и ID статуса.
	// Возвращает найденный статус.
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)

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
	// Возвращает объект MetaArray.
	GetFileList(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *CommissionReportOut, evaluate ...Evaluate) (*CommissionReportOut, *resty.Response, error)
}

// NewCommissionReportOutService возвращает сервис для работы с выданными отчётами комиссионера.
func NewCommissionReportOutService(client *Client) CommissionReportOutService {
	return newMainService[CommissionReportOut, CommissionReportOutPosition, MetaAttributesSharedStatesWrapper, any](NewEndpoint(client, "entity/commissionreportout"))
}
