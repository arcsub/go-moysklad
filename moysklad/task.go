package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Task Задача.
// Ключевое слово: task
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha
type Task struct {
	AccountID         *uuid.UUID           `json:"accountId,omitempty"`
	Agent             *Agent               `json:"agent,omitempty"`
	Assignee          *Employee            `json:"assignee,omitempty"`
	Author            *Employee            `json:"author,omitempty"`
	AuthorApplication *Application         `json:"authorApplication,omitempty"`
	Completed         *Timestamp           `json:"completed,omitempty"`
	Created           *Timestamp           `json:"created,omitempty"`
	Description       *string              `json:"description,omitempty"`
	Done              *bool                `json:"done,omitempty"`
	DueToDate         *Timestamp           `json:"dueToDate,omitempty"`
	Files             *MetaArray[File]     `json:"files,omitempty"`
	ID                *uuid.UUID           `json:"id,omitempty"`
	Implementer       *Employee            `json:"implementer,omitempty"`
	Meta              *Meta                `json:"meta,omitempty"`
	State             *State               `json:"state,omitempty"`
	Notes             *MetaArray[TaskNote] `json:"notes,omitempty"`
	Operation         *TaskOperation       `json:"operation,omitempty"`
	Updated           *Timestamp           `json:"updated,omitempty"`
}

// AsAgentInterface описывает метод, который возвращает *Agent
type AsAgentInterface interface {
	AsAgent() *Agent
}

// AsTaskOperationInterface описывает метод, который возвращает *TaskOperation
type AsTaskOperationInterface interface {
	AsTaskOperation() *TaskOperation
}

func (task Task) GetAccountID() uuid.UUID {
	return Deref(task.AccountID)
}

func (task Task) GetAgent() Agent {
	return Deref(task.Agent)
}

func (task Task) GetAssignee() Employee {
	return Deref(task.Assignee)
}

func (task Task) GetAuthor() Employee {
	return Deref(task.Author)
}

func (task Task) GetAuthorApplication() Application {
	return Deref(task.AuthorApplication)
}

func (task Task) GetCompleted() Timestamp {
	return Deref(task.Completed)
}

func (task Task) GetCreated() Timestamp {
	return Deref(task.Created)
}

func (task Task) GetDescription() string {
	return Deref(task.Description)
}

func (task Task) GetDone() bool {
	return Deref(task.Done)
}

func (task Task) GetDueToDate() Timestamp {
	return Deref(task.DueToDate)
}

func (task Task) GetFiles() MetaArray[File] {
	return Deref(task.Files)
}

func (task Task) GetID() uuid.UUID {
	return Deref(task.ID)
}

func (task Task) GetImplementer() Employee {
	return Deref(task.Implementer)
}

func (task Task) GetMeta() Meta {
	return Deref(task.Meta)
}

func (task Task) GetState() State {
	return Deref(task.State)
}

func (task Task) GetNotes() MetaArray[TaskNote] {
	return Deref(task.Notes)
}

func (task Task) GetOperation() TaskOperation {
	return Deref(task.Operation)
}

func (task Task) GetUpdated() Timestamp {
	return Deref(task.Updated)
}

func (task *Task) SetAgent(agent AsAgentInterface) *Task {
	task.Agent = agent.AsAgent()
	return task
}

func (task *Task) SetAssignee(assignee *Employee) *Task {
	task.Assignee = assignee
	return task
}

func (task *Task) SetDescription(description string) *Task {
	task.Description = &description
	return task
}

func (task *Task) SetDone(done bool) *Task {
	task.Done = &done
	return task
}

func (task *Task) SetDueToDate(dueToDate *Timestamp) *Task {
	task.DueToDate = dueToDate
	return task
}

func (task *Task) SetFiles(files Slice[File]) *Task {
	task.Files = NewMetaArrayRows(files)
	return task
}

func (task *Task) SetMeta(meta *Meta) *Task {
	task.Meta = meta
	return task
}

func (task *Task) SetState(state *State) *Task {
	task.State = state
	return task
}

func (task *Task) SetNotes(notes Slice[TaskNote]) *Task {
	task.Notes = NewMetaArrayRows(notes)
	return task
}

func (task *Task) SetOperation(operation AsTaskOperationInterface) *Task {
	task.Operation = operation.AsTaskOperation()
	return task
}

func (task Task) String() string {
	return Stringify(task)
}

func (task Task) MetaType() MetaType {
	return MetaTypeTask
}

// Update shortcut
func (task Task) Update(ctx context.Context, client *Client, params ...*Params) (*Task, *resty.Response, error) {
	return client.Entity().Task().Update(ctx, task.GetID(), &task, params...)
}

// Create shortcut
func (task Task) Create(ctx context.Context, client *Client, params ...*Params) (*Task, *resty.Response, error) {
	return client.Entity().Task().Create(ctx, &task, params...)
}

// Delete shortcut
func (task Task) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Task().Delete(ctx, task.GetID())
}

// TaskNote Комментарии задачи.
// Ключевое слово: tasknote
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-zadachi-kommentarii-zadachi
type TaskNote struct {
	Author            *Meta      `json:"author,omitempty"`            // Метаданные Сотрудника, создавшего комментарий (администратор аккаунта, если автор - приложение)
	AuthorApplication *Meta      `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего комментарий
	Moment            *Timestamp `json:"moment,omitempty"`            // Момент создания комментария
	Description       *string    `json:"description,omitempty"`       // Текст комментария
}

func (taskNote TaskNote) GetAuthor() Meta {
	return Deref(taskNote.Author)
}

func (taskNote TaskNote) GetAuthorApplication() Meta {
	return Deref(taskNote.AuthorApplication)
}

func (taskNote TaskNote) GetMoment() Timestamp {
	return Deref(taskNote.Moment)
}

func (taskNote TaskNote) GetDescription() string {
	return Deref(taskNote.Description)
}

func (taskNote *TaskNote) SetDescription(description string) *TaskNote {
	taskNote.Description = &description
	return taskNote
}

func (taskNote TaskNote) String() string {
	return Stringify(taskNote)
}

func (taskNote TaskNote) MetaType() MetaType {
	return MetaTypeTaskNote
}

// TaskOperation Метаданные Документа, связанного с задачей. Задача может быть привязана либо к конрагенту, либо к юрлицу, либо к документу
type TaskOperation struct {
	Meta *Meta   `json:"meta,omitempty"`
	Name *string `json:"name,omitempty"`
	raw  []byte
	ID   uuid.UUID `json:"id"`
}

// MetaType реализует интерфейс MetaTyper
func (taskOperation TaskOperation) MetaType() MetaType {
	return taskOperation.Meta.GetType()
}

func (taskOperation TaskOperation) GetMeta() Meta {
	return Deref(taskOperation.Meta)
}

func (taskOperation TaskOperation) GetName() string {
	return Deref(taskOperation.Name)
}

// Raw реализует интерфейс RawMetaTyper
func (taskOperation TaskOperation) Raw() []byte {
	return taskOperation.raw
}

func (taskOperation TaskOperation) String() string {
	return Stringify(taskOperation.Meta)
}

// AsPurchaseOrder десериализует объект в тип *PurchaseOrder
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPurchaseOrder() *PurchaseOrder {
	return UnmarshalAsType[PurchaseOrder](taskOperation)
}

// AsInvoiceIn десериализует объект в тип *InvoiceIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsInvoiceIn() *InvoiceIn {
	return UnmarshalAsType[InvoiceIn](taskOperation)
}

// AsSupply десериализует объект в тип *Supply
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsSupply() *Supply {
	return UnmarshalAsType[Supply](taskOperation)
}

// AsPurchaseReturn десериализует объект в тип *PurchaseReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPurchaseReturn() *PurchaseReturn {
	return UnmarshalAsType[PurchaseReturn](taskOperation)
}

// AsFactureIn десериализует объект в тип *FactureIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsFactureIn() *FactureIn {
	return UnmarshalAsType[FactureIn](taskOperation)
}

// AsCustomerOrder десериализует объект в тип *CustomerOrder
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsCustomerOrder() *CustomerOrder {
	return UnmarshalAsType[CustomerOrder](taskOperation)
}

// AsInvoiceOut десериализует объект в тип *InvoiceOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsInvoiceOut() *InvoiceOut {
	return UnmarshalAsType[InvoiceOut](taskOperation)
}

// AsDemand десериализует объект в тип *Demand
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsDemand() *Demand {
	return UnmarshalAsType[Demand](taskOperation)
}

// AsSalesReturn десериализует объект в тип *SalesReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsSalesReturn() *SalesReturn {
	return UnmarshalAsType[SalesReturn](taskOperation)
}

// AsFactureOut десериализует объект в тип *FactureOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsFactureOut() *FactureOut {
	return UnmarshalAsType[FactureOut](taskOperation)
}

// AsPriceList десериализует объект в тип *PriceList
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPriceList() *PriceList {
	return UnmarshalAsType[PriceList](taskOperation)
}

// AsLoss десериализует объект в тип *Loss
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsLoss() *Loss {
	return UnmarshalAsType[Loss](taskOperation)
}

// AsEnter десериализует объект в тип *Enter
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsEnter() *Enter {
	return UnmarshalAsType[Enter](taskOperation)
}

// AsMove десериализует объект в тип *Move
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsMove() *Move {
	return UnmarshalAsType[Move](taskOperation)
}

// AsInventory десериализует объект в тип *Inventory
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsInventory() *Inventory {
	return UnmarshalAsType[Inventory](taskOperation)
}

// AsProcessing десериализует объект в тип *Processing
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsProcessing() *Processing {
	return UnmarshalAsType[Processing](taskOperation)
}

// AsProcessingOrder десериализует объект в тип *ProcessingOrder
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsProcessingOrder() *ProcessingOrder {
	return UnmarshalAsType[ProcessingOrder](taskOperation)
}

// AsInternalOrder десериализует объект в тип *InternalOrder
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsInternalOrder() *InternalOrder {
	return UnmarshalAsType[InternalOrder](taskOperation)
}

// AsPaymentIn десериализует объект в тип *PaymentIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPaymentIn() *PaymentIn {
	return UnmarshalAsType[PaymentIn](taskOperation)
}

// AsCashIn десериализует объект в тип *CashIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsCashIn() *CashIn {
	return UnmarshalAsType[CashIn](taskOperation)
}

// AsPaymentOut десериализует объект в тип *PaymentOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPaymentOut() *PaymentOut {
	return UnmarshalAsType[PaymentOut](taskOperation)
}

// AsCashOut десериализует объект в тип *CashOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsCashOut() *CashOut {
	return UnmarshalAsType[CashOut](taskOperation)
}

// AsCounterpartyAdjustment десериализует объект в тип *CounterpartyAdjustment
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsCounterpartyAdjustment() *CounterpartyAdjustment {
	return UnmarshalAsType[CounterpartyAdjustment](taskOperation)
}

// AsRetailDemand десериализует объект в тип *RetailDemand
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailDemand() *RetailDemand {
	return UnmarshalAsType[RetailDemand](taskOperation)
}

// AsRetailSalesReturn десериализует объект в тип *RetailSalesReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailSalesReturn() *RetailSalesReturn {
	return UnmarshalAsType[RetailSalesReturn](taskOperation)
}

// AsCommissionReportIn десериализует объект в тип *CommissionReportIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsCommissionReportIn() *CommissionReportIn {
	return UnmarshalAsType[CommissionReportIn](taskOperation)
}

// AsCommissionReportOut десериализует объект в тип *CommissionReportOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsCommissionReportOut() *CommissionReportOut {
	return UnmarshalAsType[CommissionReportOut](taskOperation)
}

// AsRetailShift десериализует объект в тип *RetailShift
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailShift() *RetailShift {
	return UnmarshalAsType[RetailShift](taskOperation)
}

// AsRetailDrawerCashIn десериализует объект в тип *RetailDrawerCashIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailDrawerCashIn() *RetailDrawerCashIn {
	return UnmarshalAsType[RetailDrawerCashIn](taskOperation)
}

// AsRetailDrawerCashOut десериализует объект в тип *RetailDrawerCashOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsRetailDrawerCashOut() *RetailDrawerCashOut {
	return UnmarshalAsType[RetailDrawerCashOut](taskOperation)
}

// AsBonusTransaction десериализует объект в тип *BonusTransaction
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsBonusTransaction() *BonusTransaction {
	return UnmarshalAsType[BonusTransaction](taskOperation)
}

// AsPrepayment десериализует объект в тип *Prepayment
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPrepayment() *Prepayment {
	return UnmarshalAsType[Prepayment](taskOperation)
}

// AsPrepaymentReturn десериализует объект в тип *PrepaymentReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPrepaymentReturn() *PrepaymentReturn {
	return UnmarshalAsType[PrepaymentReturn](taskOperation)
}

// AsProductionTask десериализует объект в тип *ProductionTask
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsProductionTask() *ProductionTask {
	return UnmarshalAsType[ProductionTask](taskOperation)
}

// AsPayroll десериализует объект в тип *Payroll
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (taskOperation TaskOperation) AsPayroll() *Payroll {
	return UnmarshalAsType[Payroll](taskOperation)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
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

// TaskService
// Сервис для работы с задачами.
type TaskService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Task], *resty.Response, error)
	Create(ctx context.Context, task *Task, params ...*Params) (*Task, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, taskList Slice[Task], params ...*Params) (*Slice[Task], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...Task) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Task, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, task *Task, params ...*Params) (*Task, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetNotes(ctx context.Context, taskID uuid.UUID, params ...*Params) (*List[TaskNote], *resty.Response, error)
	CreateNote(ctx context.Context, taskID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error)
	CreateNotes(ctx context.Context, taskID uuid.UUID, taskNotes Slice[TaskNote]) (*Slice[TaskNote], *resty.Response, error)
	GetNoteByID(ctx context.Context, taskID, taskNoteID uuid.UUID) (*TaskNote, *resty.Response, error)
	UpdateNote(ctx context.Context, taskID, taskNoteID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error)
	DeleteNote(ctx context.Context, taskID, taskNoteID uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
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

// GetNotes Запрос на получение списка всех комментариев данной Задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarii-zadachi
func (service *taskService) GetNotes(ctx context.Context, taskID uuid.UUID, params ...*Params) (*List[TaskNote], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[List[TaskNote]](service.client, path).SetParams(params...).Get(ctx)
}

// CreateNote Запрос на создание нового комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-sozdat-kommentarij-zadachi
func (service *taskService) CreateNote(ctx context.Context, taskID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[TaskNote](service.client, path).Post(ctx, taskNote)
}

// CreateNotes Запрос на создание нескольких комментариев к Задаче.
func (service *taskService) CreateNotes(ctx context.Context, taskID uuid.UUID, taskNotes Slice[TaskNote]) (*Slice[TaskNote], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[Slice[TaskNote]](service.client, path).Post(ctx, taskNotes)
}

// GetNoteByID Отдельный комментарий к Задаче с указанным id комментария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarij-k-zadache
func (service *taskService) GetNoteByID(ctx context.Context, taskID, taskNoteID uuid.UUID) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, taskID, taskNoteID)
	return NewRequestBuilder[TaskNote](service.client, path).Get(ctx)
}

// UpdateNote Запрос на обновление отдельного комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-izmenit-kommentarij-k-zadache
func (service *taskService) UpdateNote(ctx context.Context, taskID, taskNoteID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, taskID, taskNoteID)
	return NewRequestBuilder[TaskNote](service.client, path).Put(ctx, taskNote)
}

// DeleteNote Запрос на удаление отдельного комментария к Задаче с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-udalit-kommentarij
func (service *taskService) DeleteNote(ctx context.Context, taskID, taskNoteID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/notes/%s", taskID, taskNoteID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
