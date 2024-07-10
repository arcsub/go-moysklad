package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"time"
)

// Task Задача.
//
// Код сущности: task
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha
type Task struct {
	AccountID         *uuid.UUID           `json:"accountId,omitempty"`         // ID учётной записи
	Agent             *Agent               `json:"agent,omitempty"`             // Метаданные Контрагента или юрлица, связанного с задачей. Задача может быть привязана либо к контрагенту, либо к юрлицу, либо к документу
	Assignee          *Employee            `json:"assignee,omitempty"`          // Метаданные ответственного за выполнение задачи
	Author            *Employee            `json:"author,omitempty"`            // Метаданные Сотрудника, создавшего задачу (администратор аккаунта, если автор - Приложение)
	AuthorApplication *Application         `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего задачу
	Completed         *Timestamp           `json:"completed,omitempty"`         // Время выполнения задачи
	Created           *Timestamp           `json:"created,omitempty"`           // Момент создания
	Description       *string              `json:"description,omitempty"`       // Текст задачи
	Done              *bool                `json:"done,omitempty"`              // Отметка о выполнении задачи
	DueToDate         *Timestamp           `json:"dueToDate,omitempty"`         // Срок задачи
	Files             *MetaArray[File]     `json:"files,omitempty"`             // Метаданные массива Файлов (Максимальное количество файлов - 100)
	ID                *uuid.UUID           `json:"id,omitempty"`                // ID Задачи
	Implementer       *Employee            `json:"implementer,omitempty"`       // Метаданные Сотрудника, выполнившего задачу
	Meta              *Meta                `json:"meta,omitempty"`              // Метаданные Задачи
	State             *State               `json:"state,omitempty"`             // Метаданные Типа задачи
	Notes             *MetaArray[TaskNote] `json:"notes,omitempty"`             // Метаданные комментариев к задаче
	Operation         *TaskOperation       `json:"operation,omitempty"`         // Метаданные Документа, связанного с задачей. Задача может быть привязана либо к контрагенту, либо к юрлицу, либо к документу
	Updated           *Timestamp           `json:"updated,omitempty"`           // Момент последнего обновления Задачи
}

// TaskOperationConverter описывает метод, который возвращает объект [TaskOperation].
type TaskOperationConverter interface {
	AsTaskOperation() *TaskOperation
}

// GetAccountID возвращает ID учётной записи.
func (task Task) GetAccountID() uuid.UUID {
	return Deref(task.AccountID)
}

// GetAgent возвращает Метаданные Контрагента или юрлица, связанного с задачей.
//
// Задача может быть привязана либо к контрагенту, либо к юрлицу, либо к документу.
func (task Task) GetAgent() Agent {
	return Deref(task.Agent)
}

// GetAssignee возвращает Метаданные ответственного за выполнение задачи.
func (task Task) GetAssignee() Employee {
	return Deref(task.Assignee)
}

// GetAuthor возвращает Метаданные Сотрудника, создавшего задачу (администратор аккаунта, если автор - Приложение).
func (task Task) GetAuthor() Employee {
	return Deref(task.Author)
}

// GetAuthorApplication возвращает Метаданные Приложения, создавшего задачу.
func (task Task) GetAuthorApplication() Application {
	return Deref(task.AuthorApplication)
}

// GetCompleted возвращает Время выполнения задачи.
func (task Task) GetCompleted() Timestamp {
	return Deref(task.Completed)
}

// GetCreated возвращает Момент создания.
func (task Task) GetCreated() Timestamp {
	return Deref(task.Created)
}

// GetDescription возвращает Текст задачи.
func (task Task) GetDescription() string {
	return Deref(task.Description)
}

// GetDone возвращает Отметку о выполнении задачи.
func (task Task) GetDone() bool {
	return Deref(task.Done)
}

// GetDueToDate возвращает Срок задачи.
func (task Task) GetDueToDate() Timestamp {
	return Deref(task.DueToDate)
}

// GetFiles возвращает Метаданные массива Файлов.
func (task Task) GetFiles() MetaArray[File] {
	return Deref(task.Files)
}

// GetID возвращает ID Задачи.
func (task Task) GetID() uuid.UUID {
	return Deref(task.ID)
}

// GetImplementer возвращает Метаданные Сотрудника, выполнившего задачу.
func (task Task) GetImplementer() Employee {
	return Deref(task.Implementer)
}

// GetMeta возвращает Метаданные Задачи.
func (task Task) GetMeta() Meta {
	return Deref(task.Meta)
}

// GetState возвращает Метаданные Типа задачи.
func (task Task) GetState() State {
	return Deref(task.State)
}

// GetNotes возвращает Массив комментариев к задаче.
func (task Task) GetNotes() MetaArray[TaskNote] {
	return Deref(task.Notes)
}

// GetOperation возвращает Метаданные Документа, связанного с задачей.
//
// Задача может быть привязана либо к контрагенту, либо к юрлицу, либо к документу.
func (task Task) GetOperation() TaskOperation {
	return Deref(task.Operation)
}

// GetUpdated возвращает Момент последнего обновления Задачи.
func (task Task) GetUpdated() Timestamp {
	return Deref(task.Updated)
}

// SetAgent устанавливает Метаданные Контрагента или юрлица, связанного с задачей.
//
// Задача может быть привязана либо к контрагенту, либо к юрлицу, либо к документу.
//
// Принимает [Counterparty] или [Organization].
func (task *Task) SetAgent(agent AgentOrganizationConverter) *Task {
	if agent != nil {
		task.Agent = agent.AsOrganizationAgent()
	}
	return task
}

// SetAssignee устанавливает Метаданные ответственного за выполнение задачи.
func (task *Task) SetAssignee(assignee *Employee) *Task {
	if assignee != nil {
		task.Assignee = assignee.Clean()
	}
	return task
}

// SetDescription устанавливает Текст задачи.
func (task *Task) SetDescription(description string) *Task {
	task.Description = &description
	return task
}

// SetDone устанавливает Отметку о выполнении задачи.
func (task *Task) SetDone(done bool) *Task {
	task.Done = &done
	return task
}

// SetDueToDate устанавливает Срок задачи.
func (task *Task) SetDueToDate(dueToDate time.Time) *Task {
	task.DueToDate = NewTimestamp(dueToDate)
	return task
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (task *Task) SetFiles(files ...*File) *Task {
	task.Files = NewMetaArrayFrom(files)
	return task
}

// SetMeta устанавливает Метаданные Задачи.
func (task *Task) SetMeta(meta *Meta) *Task {
	task.Meta = meta
	return task
}

// SetState устанавливает Метаданные Типа задачи.
func (task *Task) SetState(state *State) *Task {
	if state != nil {
		task.State = state
	}
	return task
}

// SetNotes устанавливает комментарии к задаче.
//
// Принимает множество комментариев к задаче (тип string).
func (task *Task) SetNotes(notes ...string) *Task {
	var taskNotes Slice[TaskNote]
	for _, note := range notes {
		taskNotes.Push(&TaskNote{Description: String(note)})
	}
	task.Notes = NewMetaArrayFrom(taskNotes)
	return task
}

// SetOperation устанавливает Метаданные Документа, связанного с задачей.
//
// Задача может быть привязана либо к контрагенту, либо к юрлицу, либо к документу.
func (task *Task) SetOperation(operation TaskOperationConverter) *Task {
	if operation != nil {
		task.Operation = operation.AsTaskOperation()
	}
	return task
}

// String реализует интерфейс [fmt.Stringer].
func (task Task) String() string {
	return Stringify(task)
}

// MetaType возвращает код сущности.
func (Task) MetaType() MetaType {
	return MetaTypeTask
}

// Update shortcut
func (task Task) Update(ctx context.Context, client *Client, params ...*Params) (*Task, *resty.Response, error) {
	return NewTaskService(client).Update(ctx, task.GetID(), &task, params...)
}

// Create shortcut
func (task Task) Create(ctx context.Context, client *Client, params ...*Params) (*Task, *resty.Response, error) {
	return NewTaskService(client).Create(ctx, &task, params...)
}

// Delete shortcut
func (task Task) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewTaskService(client).Delete(ctx, task.GetID())
}

// TaskNote Комментарий задачи.
//
// Код сущности: tasknote
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-zadachi-kommentarii-zadachi
type TaskNote struct {
	Author            *Meta      `json:"author,omitempty"`            // Метаданные Сотрудника, создавшего комментарий (администратор аккаунта, если автор - приложение)
	AuthorApplication *Meta      `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего комментарий
	Moment            *Timestamp `json:"moment,omitempty"`            // Момент создания комментария
	Description       *string    `json:"description,omitempty"`       // Текст комментария
}

// GetAuthor возвращает Метаданные Сотрудника, создавшего комментарий (администратор аккаунта, если автор - приложение).
func (taskNote TaskNote) GetAuthor() Meta {
	return Deref(taskNote.Author)
}

// GetAuthorApplication возвращает Метаданные Приложения, создавшего комментарий.
func (taskNote TaskNote) GetAuthorApplication() Meta {
	return Deref(taskNote.AuthorApplication)
}

// GetMoment возвращает Момент создания комментария.
func (taskNote TaskNote) GetMoment() Timestamp {
	return Deref(taskNote.Moment)
}

// GetDescription возвращает Текст комментария.
func (taskNote TaskNote) GetDescription() string {
	return Deref(taskNote.Description)
}

// SetDescription устанавливает Текст комментария.
func (taskNote *TaskNote) SetDescription(description string) *TaskNote {
	taskNote.Description = &description
	return taskNote
}

// String реализует интерфейс [fmt.Stringer].
func (taskNote TaskNote) String() string {
	return Stringify(taskNote)
}

// MetaType возвращает код сущности.
func (TaskNote) MetaType() MetaType {
	return MetaTypeTaskNote
}

// TaskOperation Метаданные Документа, связанного с задачей. Задача может быть привязана либо к контрагенту, либо к юрлицу, либо к документу
type TaskOperation struct {
	Meta *Meta     `json:"meta,omitempty"` // Метаданные документа
	Name *string   `json:"name,omitempty"` // Наименование документа
	raw  []byte    // сырые данные для последующей конвертации в нужный тип
	ID   uuid.UUID `json:"id"` // ID документа
}

// MetaType возвращает код сущности.
func (taskOperation TaskOperation) MetaType() MetaType {
	return taskOperation.Meta.GetType()
}

// GetMeta возвращает Метаданные документа.
func (taskOperation TaskOperation) GetMeta() Meta {
	return Deref(taskOperation.Meta)
}

// GetName возвращает Наименование документа.
func (taskOperation TaskOperation) GetName() string {
	return Deref(taskOperation.Name)
}

// Raw реализует интерфейс [RawMetaTyper].
func (taskOperation TaskOperation) Raw() []byte {
	return taskOperation.raw
}

// String реализует интерфейс [fmt.Stringer].
func (taskOperation TaskOperation) String() string {
	return Stringify(taskOperation.Meta)
}

// IsPurchaseOrder возвращает true, если объект является типом [PurchaseOrder].
func (taskOperation TaskOperation) IsPurchaseOrder() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePurchaseOrder
}

// IsInvoiceIn возвращает true, если объект является типом [InvoiceIn].
func (taskOperation TaskOperation) IsInvoiceIn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeInvoiceIn
}

// IsSupply возвращает true, если объект является типом [Supply].
func (taskOperation TaskOperation) IsSupply() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeSupply
}

// IsPurchaseReturn возвращает true, если объект является типом [PurchaseReturn].
func (taskOperation TaskOperation) IsPurchaseReturn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePurchaseReturn
}

// IsFactureIn возвращает true, если объект является типом [FactureIn].
func (taskOperation TaskOperation) IsFactureIn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeFactureIn
}

// IsCustomerOrder возвращает true, если объект является типом [CustomerOrder].
func (taskOperation TaskOperation) IsCustomerOrder() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeCustomerOrder
}

// IsInvoiceOut возвращает true, если объект является типом [InvoiceOut].
func (taskOperation TaskOperation) IsInvoiceOut() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeInvoiceOut
}

// IsDemand возвращает true, если объект является типом [Demand].
func (taskOperation TaskOperation) IsDemand() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeDemand
}

// IsSalesReturn возвращает true, если объект является типом [SalesReturn].
func (taskOperation TaskOperation) IsSalesReturn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeSalesReturn
}

// IsFactureOut возвращает true, если объект является типом [FactureOut].
func (taskOperation TaskOperation) IsFactureOut() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeFactureOut
}

// IsPriceList возвращает true, если объект является типом [PriceList].
func (taskOperation TaskOperation) IsPriceList() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePriceList
}

// IsLoss возвращает true, если объект является типом [Loss].
func (taskOperation TaskOperation) IsLoss() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeLoss
}

// IsEnter возвращает true, если объект является типом [Enter].
func (taskOperation TaskOperation) IsEnter() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeEnter
}

// IsMove возвращает true, если объект является типом [Move].
func (taskOperation TaskOperation) IsMove() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeMove
}

// IsInventory возвращает true, если объект является типом [Inventory].
func (taskOperation TaskOperation) IsInventory() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeInventory
}

// IsProcessing возвращает true, если объект является типом [Processing].
func (taskOperation TaskOperation) IsProcessing() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeProcessing
}

// IsProcessingOrder возвращает true, если объект является типом [ProcessingOrder].
func (taskOperation TaskOperation) IsProcessingOrder() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeProcessingOrder
}

// IsInternalOrder возвращает true, если объект является типом [InternalOrder].
func (taskOperation TaskOperation) IsInternalOrder() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeInternalOrder
}

// IsPaymentIn возвращает true, если объект является типом [PaymentIn].
func (taskOperation TaskOperation) IsPaymentIn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePaymentIn
}

// IsCashIn возвращает true, если объект является типом [CashIn].
func (taskOperation TaskOperation) IsCashIn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeCashIn
}

// IsPaymentOut возвращает true, если объект является типом [PaymentOut].
func (taskOperation TaskOperation) IsPaymentOut() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePaymentOut
}

// IsCashOut возвращает true, если объект является типом [CashOut].
func (taskOperation TaskOperation) IsCashOut() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeCashOut
}

// IsCounterpartyAdjustment возвращает true, если объект является типом [CounterpartyAdjustment].
func (taskOperation TaskOperation) IsCounterpartyAdjustment() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeCounterpartyAdjustment
}

// IsRetailDemand возвращает true, если объект является типом [RetailDemand].
func (taskOperation TaskOperation) IsRetailDemand() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeRetailDemand
}

// IsRetailSalesReturn возвращает true, если объект является типом [RetailSalesReturn].
func (taskOperation TaskOperation) IsRetailSalesReturn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeRetailSalesReturn
}

// IsCommissionReportIn возвращает true, если объект является типом [CommissionReportIn].
func (taskOperation TaskOperation) IsCommissionReportIn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeCommissionReportIn
}

// IsCommissionReportOut возвращает true, если объект является типом [CommissionReportOut].
func (taskOperation TaskOperation) IsCommissionReportOut() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeCommissionReportOut
}

// IsRetailShift возвращает true, если объект является типом [RetailShift].
func (taskOperation TaskOperation) IsRetailShift() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeRetailShift
}

// IsRetailDrawerCashIn возвращает true, если объект является типом [RetailDrawerCashIn].
func (taskOperation TaskOperation) IsRetailDrawerCashIn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeRetailDrawerCashIn
}

// IsRetailDrawerCashOut возвращает true, если объект является типом [RetailDrawerCashOut].
func (taskOperation TaskOperation) IsRetailDrawerCashOut() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeRetailDrawerCashOut
}

// IsBonusTransaction возвращает true, если объект является типом [BonusTransaction].
func (taskOperation TaskOperation) IsBonusTransaction() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeBonusTransaction
}

// IsPrepayment возвращает true, если объект является типом [Prepayment].
func (taskOperation TaskOperation) IsPrepayment() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePrepayment
}

// IsPrepaymentReturn возвращает true, если объект является типом [PrepaymentReturn].
func (taskOperation TaskOperation) IsPrepaymentReturn() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePrepaymentReturn
}

// IsProductionTask возвращает true, если объект является типом [ProductionTask].
func (taskOperation TaskOperation) IsProductionTask() bool {
	return taskOperation.GetMeta().GetType() == MetaTypeProductionTask
}

// IsPayroll возвращает true, если объект является типом [Payroll].
func (taskOperation TaskOperation) IsPayroll() bool {
	return taskOperation.GetMeta().GetType() == MetaTypePayroll
}

// AsPurchaseOrder пытается привести объект к типу [PurchaseOrder].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PurchaseOrder] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPurchaseOrder() *PurchaseOrder {
	return UnmarshalAsType[PurchaseOrder](taskOperation)
}

// AsInvoiceIn пытается привести объект к типу [InvoiceIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [InvoiceIn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsInvoiceIn() *InvoiceIn {
	return UnmarshalAsType[InvoiceIn](taskOperation)
}

// AsSupply пытается привести объект к типу [Supply].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Supply] или nil в случае неудачи.
func (taskOperation TaskOperation) AsSupply() *Supply {
	return UnmarshalAsType[Supply](taskOperation)
}

// AsPurchaseReturn пытается привести объект к типу [PurchaseReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PurchaseReturn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPurchaseReturn() *PurchaseReturn {
	return UnmarshalAsType[PurchaseReturn](taskOperation)
}

// AsFactureIn пытается привести объект к типу [FactureIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [FactureIn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsFactureIn() *FactureIn {
	return UnmarshalAsType[FactureIn](taskOperation)
}

// AsCustomerOrder пытается привести объект к типу [CustomerOrder].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CustomerOrder] или nil в случае неудачи.
func (taskOperation TaskOperation) AsCustomerOrder() *CustomerOrder {
	return UnmarshalAsType[CustomerOrder](taskOperation)
}

// AsInvoiceOut пытается привести объект к типу [InvoiceOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [InvoiceOut] или nil в случае неудачи.
func (taskOperation TaskOperation) AsInvoiceOut() *InvoiceOut {
	return UnmarshalAsType[InvoiceOut](taskOperation)
}

// AsDemand пытается привести объект к типу [Demand].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Demand] или nil в случае неудачи.
func (taskOperation TaskOperation) AsDemand() *Demand {
	return UnmarshalAsType[Demand](taskOperation)
}

// AsSalesReturn пытается привести объект к типу [SalesReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [SalesReturn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsSalesReturn() *SalesReturn {
	return UnmarshalAsType[SalesReturn](taskOperation)
}

// AsFactureOut пытается привести объект к типу [FactureOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [FactureOut] или nil в случае неудачи.
func (taskOperation TaskOperation) AsFactureOut() *FactureOut {
	return UnmarshalAsType[FactureOut](taskOperation)
}

// AsPriceList пытается привести объект к типу [PriceList].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PriceList] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPriceList() *PriceList {
	return UnmarshalAsType[PriceList](taskOperation)
}

// AsLoss пытается привести объект к типу [Loss].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Loss] или nil в случае неудачи.
func (taskOperation TaskOperation) AsLoss() *Loss {
	return UnmarshalAsType[Loss](taskOperation)
}

// AsEnter пытается привести объект к типу [Enter].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Enter] или nil в случае неудачи.
func (taskOperation TaskOperation) AsEnter() *Enter {
	return UnmarshalAsType[Enter](taskOperation)
}

// AsMove пытается привести объект к типу [Move].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Move] или nil в случае неудачи.
func (taskOperation TaskOperation) AsMove() *Move {
	return UnmarshalAsType[Move](taskOperation)
}

// AsInventory пытается привести объект к типу [Inventory].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Inventory] или nil в случае неудачи.
func (taskOperation TaskOperation) AsInventory() *Inventory {
	return UnmarshalAsType[Inventory](taskOperation)
}

// AsProcessing пытается привести объект к типу [Processing].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Processing] или nil в случае неудачи.
func (taskOperation TaskOperation) AsProcessing() *Processing {
	return UnmarshalAsType[Processing](taskOperation)
}

// AsProcessingOrder пытается привести объект к типу [ProcessingOrder].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [ProcessingOrder] или nil в случае неудачи.
func (taskOperation TaskOperation) AsProcessingOrder() *ProcessingOrder {
	return UnmarshalAsType[ProcessingOrder](taskOperation)
}

// AsInternalOrder пытается привести объект к типу [InternalOrder].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [InternalOrder] или nil в случае неудачи.
func (taskOperation TaskOperation) AsInternalOrder() *InternalOrder {
	return UnmarshalAsType[InternalOrder](taskOperation)
}

// AsPaymentIn пытается привести объект к типу [PaymentIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PaymentIn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPaymentIn() *PaymentIn {
	return UnmarshalAsType[PaymentIn](taskOperation)
}

// AsCashIn пытается привести объект к типу [CashIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CashIn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsCashIn() *CashIn {
	return UnmarshalAsType[CashIn](taskOperation)
}

// AsPaymentOut пытается привести объект к типу [PaymentOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PaymentOut] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPaymentOut() *PaymentOut {
	return UnmarshalAsType[PaymentOut](taskOperation)
}

// AsCashOut пытается привести объект к типу [CashOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CashOut] или nil в случае неудачи.
func (taskOperation TaskOperation) AsCashOut() *CashOut {
	return UnmarshalAsType[CashOut](taskOperation)
}

// AsCounterpartyAdjustment пытается привести объект к типу [CounterpartyAdjustment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CounterpartyAdjustment] или nil в случае неудачи.
func (taskOperation TaskOperation) AsCounterpartyAdjustment() *CounterpartyAdjustment {
	return UnmarshalAsType[CounterpartyAdjustment](taskOperation)
}

// AsRetailDemand пытается привести объект к типу [RetailDemand].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailDemand] или nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailDemand() *RetailDemand {
	return UnmarshalAsType[RetailDemand](taskOperation)
}

// AsRetailSalesReturn пытается привести объект к типу [RetailSalesReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailSalesReturn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailSalesReturn() *RetailSalesReturn {
	return UnmarshalAsType[RetailSalesReturn](taskOperation)
}

// AsCommissionReportIn пытается привести объект к типу [CommissionReportIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CommissionReportIn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsCommissionReportIn() *CommissionReportIn {
	return UnmarshalAsType[CommissionReportIn](taskOperation)
}

// AsCommissionReportOut пытается привести объект к типу [CommissionReportOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CommissionReportOut] или nil в случае неудачи.
func (taskOperation TaskOperation) AsCommissionReportOut() *CommissionReportOut {
	return UnmarshalAsType[CommissionReportOut](taskOperation)
}

// AsRetailShift пытается привести объект к типу [RetailShift].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailShift] или nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailShift() *RetailShift {
	return UnmarshalAsType[RetailShift](taskOperation)
}

// AsRetailDrawerCashIn пытается привести объект к типу [RetailDrawerCashIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailDrawerCashIn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailDrawerCashIn() *RetailDrawerCashIn {
	return UnmarshalAsType[RetailDrawerCashIn](taskOperation)
}

// AsRetailDrawerCashOut пытается привести объект к типу [RetailDrawerCashOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailDrawerCashOut] или nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailDrawerCashOut() *RetailDrawerCashOut {
	return UnmarshalAsType[RetailDrawerCashOut](taskOperation)
}

// AsBonusTransaction пытается привести объект к типу [BonusTransaction].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [BonusTransaction] или nil в случае неудачи.
func (taskOperation TaskOperation) AsBonusTransaction() *BonusTransaction {
	return UnmarshalAsType[BonusTransaction](taskOperation)
}

// AsPrepayment пытается привести объект к типу [Prepayment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Prepayment] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPrepayment() *Prepayment {
	return UnmarshalAsType[Prepayment](taskOperation)
}

// AsPrepaymentReturn пытается привести объект к типу [PrepaymentReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PrepaymentReturn] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPrepaymentReturn() *PrepaymentReturn {
	return UnmarshalAsType[PrepaymentReturn](taskOperation)
}

// AsProductionTask пытается привести объект к типу [ProductionTask].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [ProductionTask] или nil в случае неудачи.
func (taskOperation TaskOperation) AsProductionTask() *ProductionTask {
	return UnmarshalAsType[ProductionTask](taskOperation)
}

// AsPayroll пытается привести объект к типу [Payroll].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Payroll] или nil в случае неудачи.
func (taskOperation TaskOperation) AsPayroll() *Payroll {
	return UnmarshalAsType[Payroll](taskOperation)
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (taskOperation *TaskOperation) UnmarshalJSON(data []byte) error {
	type alias TaskOperation
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.raw = data
	*taskOperation = TaskOperation(t)
	return nil
}

// TaskService описывает методы сервиса для работы с задачами.
type TaskService interface {
	// GetList выполняет запрос на получение списка задач.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Task], *resty.Response, error)

	// Create выполняет запрос на создание задачи.
	// Создать новую задачу. Для создания новых задач необходима активная тарифная опция CRM.
	// Обязательные поля для заполнения:
	//	- description (Текст задачи)
	//	- assignee (Метаданные Сотрудника, ответственного за выполнение задачи)
	// Принимает контекст, задачу и опционально объект параметров запроса Params.
	// Возвращает созданную задачу.
	Create(ctx context.Context, task *Task, params ...*Params) (*Task, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение задач.
	// Изменяемые задачи должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список задач и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых задач.
	CreateUpdateMany(ctx context.Context, taskList Slice[Task], params ...*Params) (*Slice[Task], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление задач.
	// Принимает контекст и множество задач.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Task) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление задачи.
	// Принимает контекст и ID задачи.
	// Возвращает «true» в случае успешного удаления задачи.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной задачи по ID.
	// Принимает контекст, ID задачи и опционально объект параметров запроса Params.
	// Возвращает задачу.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Task, *resty.Response, error)

	// Update выполняет запрос на изменение задачи.
	// Принимает контекст, задачу и опционально объект параметров запроса Params.
	// Возвращает изменённую задачу.
	Update(ctx context.Context, id uuid.UUID, task *Task, params ...*Params) (*Task, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// GetNoteList выполняет запрос на получение списка всех комментариев данной задачи.
	// Принимает контекст, ID задачи и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNoteList(ctx context.Context, taskID uuid.UUID, params ...*Params) (*List[TaskNote], *resty.Response, error)

	// CreateNote выполняет запрос на создание нового комментария к задаче.
	// Обязательные поля для заполнения:
	//	- text (Текст комментария)
	// Принимает контекст, ID задачи и комментарий к задаче (тип string).
	// Возвращает созданный комментарий к задаче.
	CreateNote(ctx context.Context, taskID uuid.UUID, taskNoteText string) (*TaskNote, *resty.Response, error)

	// CreateNoteMany выполняет запрос на массовое создание комментариев к задаче.
	// Принимает контекст, ID задачи и список комментариев к задаче (тип string).
	// Возвращает список созданных комментариев к задаче.
	CreateNoteMany(ctx context.Context, taskID uuid.UUID, taskNotesText ...string) (*Slice[TaskNote], *resty.Response, error)

	// GetNoteByID выполняет запрос на получение отдельного комментария к задаче по ID.
	// Принимает контекст, ID задачи и ID комментария к задаче.
	// Возвращает комментарий к задаче.
	GetNoteByID(ctx context.Context, taskID, taskNoteID uuid.UUID) (*TaskNote, *resty.Response, error)

	// UpdateNote выполняет запрос на изменение комментария к задаче.
	// Принимает контекст, ID задачи, ID комментария к задаче и комментарий к задаче (тип string).
	// Возвращает изменённый комментарий к задаче.
	UpdateNote(ctx context.Context, taskID, taskNoteID uuid.UUID, taskNoteText string) (*TaskNote, *resty.Response, error)

	// DeleteNote выполняет запрос на удаление комментария к задаче.
	// Принимает контекст, ID задачи и ID комментария к задаче.
	// Возвращает «true» в случае успешного удаления комментария к задаче.
	DeleteNote(ctx context.Context, taskID, taskNoteID uuid.UUID) (bool, *resty.Response, error)

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
	// Возвращает «true» в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
}

type taskService struct {
	Endpoint
	endpointGetList[Task]
	endpointCreate[Task]
	endpointCreateUpdateMany[Task]
	endpointDeleteMany[Task]
	endpointDelete
	endpointGetByID[Task]
	endpointUpdate[Task]
	endpointNamedFilter
	endpointFiles
}

func (service *taskService) GetNoteList(ctx context.Context, taskID uuid.UUID, params ...*Params) (*List[TaskNote], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[List[TaskNote]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *taskService) CreateNote(ctx context.Context, taskID uuid.UUID, taskNoteText string) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[TaskNote](service.client, path).Post(ctx, &TaskNote{Description: String(taskNoteText)})
}

func (service *taskService) CreateNoteMany(ctx context.Context, taskID uuid.UUID, taskNoteText ...string) (*Slice[TaskNote], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	var taskNotes Slice[TaskNote]
	for _, text := range taskNoteText {
		taskNotes.Push(&TaskNote{Description: String(text)})
	}
	return NewRequestBuilder[Slice[TaskNote]](service.client, path).Post(ctx, taskNotes)
}

func (service *taskService) GetNoteByID(ctx context.Context, taskID, taskNoteID uuid.UUID) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, taskID, taskNoteID)
	return NewRequestBuilder[TaskNote](service.client, path).Get(ctx)
}

func (service *taskService) UpdateNote(ctx context.Context, taskID, taskNoteID uuid.UUID, taskNoteText string) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, taskID, taskNoteID)
	return NewRequestBuilder[TaskNote](service.client, path).Put(ctx, &TaskNote{Description: String(taskNoteText)})
}

func (service *taskService) DeleteNote(ctx context.Context, taskID, taskNoteID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/notes/%s", taskID, taskNoteID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// NewTaskService принимает [Client] и возвращает сервис для работы с задачами.
func NewTaskService(client *Client) TaskService {
	e := NewEndpoint(client, "entity/task")
	return &taskService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Task]{e},
		endpointCreate:           endpointCreate[Task]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Task]{e},
		endpointDeleteMany:       endpointDeleteMany[Task]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[Task]{e},
		endpointUpdate:           endpointUpdate[Task]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointFiles:            endpointFiles{e},
	}
}
