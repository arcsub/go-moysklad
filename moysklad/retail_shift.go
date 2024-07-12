package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// RetailShift Розничная смена.
//
// Код сущности: retailshift
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena
type RetailShift struct {
	Organization        *Organization          `json:"organization,omitempty"`        // Метаданные юрлица
	Shared              *bool                  `json:"shared,omitempty"`              // Общий доступ
	AgentAccount        *AgentAccount          `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	VatIncluded         *bool                  `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	BankCommission      *float64               `json:"bankComission,omitempty"`       // Сумма комиссии эквайера за проведение безналичных платежей по банковской карте. Не может превышать общую сумму безналичных платежей по карте. Если не указано, заполняется 0 автоматически.
	BankPercent         *float64               `json:"bankPercent,omitempty"`         // Комиссия банка-эквайера по операциям по карте (в процентах)
	Name                *string                `json:"name,omitempty"`                // Наименование Розничной смены
	CloseDate           *Timestamp             `json:"closeDate,omitempty"`           // Дата закрытия смены
	Contract            *NullValue[Contract]   `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp             `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp             `json:"deleted,omitempty"`             // Момент последнего удаления Розничной смены
	Description         *string                `json:"description,omitempty"`         // Комментарий Розничной смены
	ExternalCode        *string                `json:"externalCode,omitempty"`        // Внешний код Розничной смены
	Files               *MetaArray[File]       `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                 `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID             `json:"id,omitempty"`                  // ID Розничной смены
	Meta                *Meta                  `json:"meta,omitempty"`                // Метаданные Розничной смены
	AccountID           *uuid.UUID             `json:"accountId,omitempty"`           // ID учётной записи
	Cheque              *Cheque                `json:"cheque,omitempty"`              // Информация о смене ККТ
	Acquire             *Agent                 `json:"acquire,omitempty"`             // Метаданные Банка-эквайера по операциям по карте
	Moment              *Timestamp             `json:"moment,omitempty"`              // Дата смены
	OrganizationAccount *AgentAccount          `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee              `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Printed             *bool                  `json:"printed,omitempty"`             // Напечатан ли документ
	ProceedsCash        *float64               `json:"proceedsCash,omitempty"`        // Выручка наличными
	ProceedsNoCash      *float64               `json:"proceedsNoCash,omitempty"`      // Выручка безнал
	Published           *bool                  `json:"published,omitempty"`           // Опубликован ли документ
	QRAcquire           *Agent                 `json:"qrAcquire,omitempty"`           // Метаданные Банка-эквайера по операциям по QR-коду
	QRBankCommission    *float64               `json:"qrBankComission,omitempty"`     // Сумма комиссии эквайера за проведение безналичных платежей по QR-коду. Не может превышать общую сумму безналичных платежей по QR-коду. Если не указано, заполняется 0 автоматически.
	QRBankPercent       *float64               `json:"qrBankPercent,omitempty"`       // Комиссия банка-эквайера по операция по QR-коду (в процентах)
	ReceivedCash        *float64               `json:"receivedCash,omitempty"`        // Получено наличными
	ReceivedNoCash      *float64               `json:"receivedNoCash,omitempty"`      // Получено безнал
	RetailStore         *RetailStore           `json:"retailStore,omitempty"`         // Метаданные точки продаж
	Store               *Store                 `json:"store,omitempty"`               // Метаданные склада. Если не указано, заполняется с точки продаж автоматически
	SyncID              *uuid.UUID             `json:"syncId,omitempty"`              // ID синхронизации
	Updated             *Timestamp             `json:"updated,omitempty"`             // Момент последнего обновления Розничной смены
	VatEnabled          *bool                  `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	PaymentOperations   Slice[Payment]         `json:"paymentOperations,omitempty"`   // Коллекция метаданных платежных операций
	Operations          Slice[RetailOperation] `json:"operations,omitempty"`          // Коллекция метаданных связанных операций
	Attributes          Slice[Attribute]       `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (retailShift RetailShift) Clean() *RetailShift {
	if retailShift.Meta == nil {
		return nil
	}
	return &RetailShift{Meta: retailShift.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (retailShift RetailShift) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: retailShift.Meta}
}

// AsOperation реализует интерфейс [OperationConverter].
func (retailShift RetailShift) AsOperation() *Operation {
	return newOperation(retailShift)
}

// AsOperationIn реализует интерфейс [OperationInConverter].
func (retailShift RetailShift) AsOperationIn() *Operation {
	return retailShift.AsOperation()
}

// GetProceedsSum возвращает сумму выручки.
func (retailShift RetailShift) GetProceedsSum() float64 {
	return retailShift.GetProceedsCash() + retailShift.GetProceedsNoCash()
}

// GetReceivedSum возвращает сумму полученного.
func (retailShift RetailShift) GetReceivedSum() float64 {
	return retailShift.GetReceivedCash() + retailShift.GetReceivedNoCash()
}

// GetOrganization возвращает Метаданные юрлица.
func (retailShift RetailShift) GetOrganization() Organization {
	return Deref(retailShift.Organization)
}

// GetShared возвращает флаг Общего доступа.
func (retailShift RetailShift) GetShared() bool {
	return Deref(retailShift.Shared)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (retailShift RetailShift) GetAgentAccount() AgentAccount {
	return Deref(retailShift.AgentAccount)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (retailShift RetailShift) GetVatIncluded() bool {
	return Deref(retailShift.VatIncluded)
}

// GetBankCommission возвращает Сумма комиссии эквайера за проведение безналичных платежей по банковской карте.
//
// Не может превышать общую сумму безналичных платежей по карте.
//
// Если не указано, заполняется 0 автоматически.
func (retailShift RetailShift) GetBankCommission() float64 {
	return Deref(retailShift.BankCommission)
}

// GetBankPercent возвращает Комиссия банка-эквайера по операциям по карте (в процентах).
func (retailShift RetailShift) GetBankPercent() float64 {
	return Deref(retailShift.BankPercent)
}

// GetName возвращает Наименование Розничной смены.
func (retailShift RetailShift) GetName() string {
	return Deref(retailShift.Name)
}

// GetCloseDate возвращает Дату закрытия смены.
func (retailShift RetailShift) GetCloseDate() Timestamp {
	return Deref(retailShift.CloseDate)
}

// GetContract возвращает Метаданные договора.
func (retailShift RetailShift) GetContract() Contract {
	return Deref(retailShift.Contract).getValue()
}

// GetCreated возвращает Дату создания.
func (retailShift RetailShift) GetCreated() Timestamp {
	return Deref(retailShift.Created)
}

// GetDeleted возвращает Момент последнего удаления Розничной смены.
func (retailShift RetailShift) GetDeleted() Timestamp {
	return Deref(retailShift.Deleted)
}

// GetDescription возвращает Комментарий Розничной смены.
func (retailShift RetailShift) GetDescription() string {
	return Deref(retailShift.Description)
}

// GetExternalCode возвращает Внешний код Розничной смены.
func (retailShift RetailShift) GetExternalCode() string {
	return Deref(retailShift.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (retailShift RetailShift) GetFiles() MetaArray[File] {
	return Deref(retailShift.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (retailShift RetailShift) GetGroup() Group {
	return Deref(retailShift.Group)
}

// GetID возвращает ID Розничной смены.
func (retailShift RetailShift) GetID() uuid.UUID {
	return Deref(retailShift.ID)
}

// GetMeta возвращает Метаданные Розничной смены.
func (retailShift RetailShift) GetMeta() Meta {
	return Deref(retailShift.Meta)
}

// GetAccountID возвращает ID учётной записи.
func (retailShift RetailShift) GetAccountID() uuid.UUID {
	return Deref(retailShift.AccountID)
}

// GetCheque возвращает Информацию о смене ККТ.
func (retailShift RetailShift) GetCheque() Cheque {
	return Deref(retailShift.Cheque)
}

// GetAcquire возвращает Метаданные Банка-эквайера по операциям по карте.
func (retailShift RetailShift) GetAcquire() Agent {
	return Deref(retailShift.Acquire)
}

// GetMoment возвращает Дату смены.
func (retailShift RetailShift) GetMoment() Timestamp {
	return Deref(retailShift.Moment)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (retailShift RetailShift) GetOrganizationAccount() AgentAccount {
	return Deref(retailShift.OrganizationAccount)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (retailShift RetailShift) GetOwner() Employee {
	return Deref(retailShift.Owner)
}

// GetPrinted возвращает true, если документ напечатан.
func (retailShift RetailShift) GetPrinted() bool {
	return Deref(retailShift.Printed)
}

// GetProceedsCash возвращает Выручку наличными.
func (retailShift RetailShift) GetProceedsCash() float64 {
	return Deref(retailShift.ProceedsCash)
}

// GetProceedsNoCash возвращает Выручку безнала.
func (retailShift RetailShift) GetProceedsNoCash() float64 {
	return Deref(retailShift.ProceedsNoCash)
}

// GetPublished возвращает true, если документ опубликован.
func (retailShift RetailShift) GetPublished() bool {
	return Deref(retailShift.Published)
}

// GetQRAcquire возвращает Метаданные Банка-эквайера по операциям по QR-коду.
func (retailShift RetailShift) GetQRAcquire() Agent {
	return Deref(retailShift.QRAcquire)
}

// GetQRBankCommission возвращает Сумма комиссии эквайера за проведение безналичных платежей по QR-коду.
//
// Не может превышать общую сумму безналичных платежей по QR-коду.
//
// Если не указано, заполняется 0 автоматически.
func (retailShift RetailShift) GetQRBankCommission() float64 {
	return Deref(retailShift.QRBankCommission)
}

// GetQRBankPercent возвращает Комиссию банка-эквайера по операция по QR-коду (в процентах).
func (retailShift RetailShift) GetQRBankPercent() float64 {
	return Deref(retailShift.QRBankPercent)
}

// GetReceivedCash возвращает Получено наличными.
func (retailShift RetailShift) GetReceivedCash() float64 {
	return Deref(retailShift.ReceivedCash)
}

// GetReceivedNoCash возвращает Получено безнала.
func (retailShift RetailShift) GetReceivedNoCash() float64 {
	return Deref(retailShift.ReceivedNoCash)
}

// GetRetailStore возвращает Метаданные Точки продаж.
func (retailShift RetailShift) GetRetailStore() RetailStore {
	return Deref(retailShift.RetailStore)
}

// GetStore возвращает Метаданные склада.
func (retailShift RetailShift) GetStore() Store {
	return Deref(retailShift.Store)
}

// GetSyncID возвращает ID синхронизации.
func (retailShift RetailShift) GetSyncID() uuid.UUID {
	return Deref(retailShift.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Списания.
func (retailShift RetailShift) GetUpdated() Timestamp {
	return Deref(retailShift.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (retailShift RetailShift) GetVatEnabled() bool {
	return Deref(retailShift.VatEnabled)
}

// GetPayments возвращает Коллекцию метаданных платежных операций.
func (retailShift RetailShift) GetPayments() Slice[Payment] {
	return retailShift.PaymentOperations
}

// GetOperations возвращает Коллекцию метаданных связанных операций.
func (retailShift RetailShift) GetOperations() Slice[RetailOperation] {
	return retailShift.Operations
}

// GetAttributes возвращает Список метаданных доп. полей.
func (retailShift RetailShift) GetAttributes() Slice[Attribute] {
	return retailShift.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (retailShift *RetailShift) SetOrganization(organization *Organization) *RetailShift {
	if organization != nil {
		retailShift.Organization = organization.Clean()
	}
	return retailShift
}

// SetShared устанавливает флаг общего доступа.
func (retailShift *RetailShift) SetShared(shared bool) *RetailShift {
	retailShift.Shared = &shared
	return retailShift
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (retailShift *RetailShift) SetVatIncluded(vatIncluded bool) *RetailShift {
	retailShift.VatIncluded = &vatIncluded
	return retailShift
}

// SetBankCommission устанавливает Сумму комиссии эквайера за проведение безналичных платежей по банковской карте.
//
// Не может превышать общую сумму безналичных платежей по карте.
//
// Если не указано, заполняется 0 автоматически.
func (retailShift *RetailShift) SetBankCommission(bankCommission float64) *RetailShift {
	retailShift.BankCommission = &bankCommission
	return retailShift
}

// SetBankPercent устанавливает Комиссию банка-эквайера по операциям по карте (в процентах).
func (retailShift *RetailShift) SetBankPercent(bankPercent float64) *RetailShift {
	retailShift.BankPercent = &bankPercent
	return retailShift
}

// SetName устанавливает Наименование Розничной смены.
func (retailShift *RetailShift) SetName(name string) *RetailShift {
	retailShift.Name = &name
	return retailShift
}

// SetCloseDate устанавливает Дату закрытия смены.
func (retailShift *RetailShift) SetCloseDate(closeDate time.Time) *RetailShift {
	retailShift.CloseDate = NewTimestamp(closeDate)
	return retailShift
}

// SetDescription устанавливает Комментарий Розничной смены.
func (retailShift *RetailShift) SetDescription(description string) *RetailShift {
	retailShift.Description = &description
	return retailShift
}

// SetExternalCode устанавливает Внешний код Розничной смены.
func (retailShift *RetailShift) SetExternalCode(externalCode string) *RetailShift {
	retailShift.ExternalCode = &externalCode
	return retailShift
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (retailShift *RetailShift) SetFiles(files ...*File) *RetailShift {
	retailShift.Files = NewMetaArrayFrom(files)
	return retailShift
}

// SetMeta устанавливает Метаданные Розничной смены.
func (retailShift *RetailShift) SetMeta(meta *Meta) *RetailShift {
	retailShift.Meta = meta
	return retailShift
}

// SetAcquire устанавливает Метаданные Банка-эквайера по операциям по карте.
func (retailShift *RetailShift) SetAcquire(acquire AgentOrganizationConverter) *RetailShift {
	if acquire != nil {
		retailShift.Acquire = acquire.AsOrganizationAgent()
	}
	return retailShift
}

// SetMoment устанавливает Дату смены.
func (retailShift *RetailShift) SetMoment(moment time.Time) *RetailShift {
	retailShift.Moment = NewTimestamp(moment)
	return retailShift
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (retailShift *RetailShift) SetOrganizationAccount(organizationAccount *AgentAccount) *RetailShift {
	if organizationAccount != nil {
		retailShift.OrganizationAccount = organizationAccount.Clean()
	}
	return retailShift
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (retailShift *RetailShift) SetOwner(owner *Employee) *RetailShift {
	if owner != nil {
		retailShift.Owner = owner.Clean()
	}
	return retailShift
}

// SetQRAcquire устанавливает Метаданные Банка-эквайера по операциям по QR-коду.
func (retailShift *RetailShift) SetQRAcquire(qrAcquire AgentOrganizationConverter) *RetailShift {
	if qrAcquire != nil {
		retailShift.QRAcquire = qrAcquire.AsOrganizationAgent()
	}
	return retailShift
}

// SetQRBankCommission устанавливает Сумму комиссии эквайера за проведение безналичных платежей по QR-коду.
//
// Не может превышать общую сумму безналичных платежей по QR-коду.
//
// Если не указано, заполняется 0 автоматически.
func (retailShift *RetailShift) SetQRBankCommission(qrBankCommission float64) *RetailShift {
	retailShift.QRBankCommission = &qrBankCommission
	return retailShift
}

// SetQRBankPercent устанавливает Комиссию банка-эквайера по операция по QR-коду (в процентах).
func (retailShift *RetailShift) SetQRBankPercent(qrBankPercent float64) *RetailShift {
	retailShift.QRBankPercent = &qrBankPercent
	return retailShift
}

// SetRetailStore устанавливает Метаданные Точки продаж.
func (retailShift *RetailShift) SetRetailStore(retailStore *RetailStore) *RetailShift {
	if retailStore != nil {
		retailShift.RetailStore = retailStore.Clean()
	}
	return retailShift
}

// SetStore устанавливает Метаданные склада.
func (retailShift *RetailShift) SetStore(store *Store) *RetailShift {
	if store != nil {
		retailShift.Store = store.Clean()
	}
	return retailShift
}

// SetSyncID устанавливает ID синхронизации.
func (retailShift *RetailShift) SetSyncID(syncID uuid.UUID) *RetailShift {
	retailShift.SyncID = &syncID
	return retailShift
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (retailShift *RetailShift) SetAttributes(attributes ...*Attribute) *RetailShift {
	retailShift.Attributes.Push(attributes...)
	return retailShift
}

// String реализует интерфейс [fmt.Stringer].
func (retailShift RetailShift) String() string {
	return Stringify(retailShift)
}

// MetaType возвращает код сущности.
func (RetailShift) MetaType() MetaType {
	return MetaTypeRetailShift
}

// Update shortcut
func (retailShift RetailShift) Update(ctx context.Context, client *Client, params ...*Params) (*RetailShift, *resty.Response, error) {
	return NewRetailShiftService(client).Update(ctx, retailShift.GetID(), &retailShift, params...)
}

// Create shortcut
func (retailShift RetailShift) Create(ctx context.Context, client *Client, params ...*Params) (*RetailShift, *resty.Response, error) {
	return NewRetailShiftService(client).Create(ctx, &retailShift, params...)
}

// Delete shortcut
func (retailShift RetailShift) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewRetailShiftService(client).Delete(ctx, retailShift.GetID())
}

// Cheque Информация о смене ККТ.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-smene-kkt
type Cheque struct {
	Start ChequeStart `json:"start,omitempty"` // Информация об открытии смены
	End   ChequeEnd   `json:"end,omitempty"`   // Информация о закрытии смены
}

// String реализует интерфейс [fmt.Stringer].
func (cheque Cheque) String() string {
	return Stringify(cheque)
}

// ChequeStart Информация об открытии смены ККТ.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-ob-otkrytii-smeny-kkt
type ChequeStart struct {
	Time            Timestamp `json:"time,omitempty"`            // Дата и время открытия смены
	FnNumber        string    `json:"fnNumber,omitempty"`        // Номер фискального накопителя
	KKTRegNumber    string    `json:"kktRegNumber,omitempty"`    // Регистрационный номер ККТ
	FiscalDocSign   string    `json:"fiscalDocSign,omitempty"`   // Фискальный признак документа
	ShiftNumber     string    `json:"shiftNumber,omitempty"`     // Номер смены ККТ
	FiscalDocNumber string    `json:"fiscalDocNumber,omitempty"` // Номер фискального документа
}

// String реализует интерфейс [fmt.Stringer].
func (ChequeStart ChequeStart) String() string {
	return Stringify(ChequeStart)
}

// ChequeEnd Информация о закрытии смены ККТ.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-zakrytii-smeny-kkt
type ChequeEnd struct {
	ChequeStart
	ChequesTotal    float64 `json:"chequesTotal,omitempty"`    // Количество чеков за смену
	FiscalDocsTotal float64 `json:"fiscalDocsTotal,omitempty"` // Количество фискальных документов за смену
}

// String реализует интерфейс [fmt.Stringer].
func (chequeEnd ChequeEnd) String() string {
	return Stringify(chequeEnd)
}

// RetailShiftService описывает методы сервиса для работы с розничными сменами.
type RetailShiftService interface {
	// GetList выполняет запрос на получение списка розничных смен.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[RetailShift], *resty.Response, error)

	// Create выполняет запрос на создание розничной смены.
	// Обязательные поля для заполнения:
	//	- organization (Метаданные юрлица)
	//	- retailStore (Метаданные точки продаж)
	// Принимает контекст, розничную смены и опционально объект параметров запроса Params.
	// Возвращает созданную розничную смены.
	Create(ctx context.Context, retailShift *RetailShift, params ...*Params) (*RetailShift, *resty.Response, error)

	// Delete выполняет запрос на удаление розничной смены.
	// Принимает контекст и ID розничной смены.
	// Возвращает «true» в случае успешного удаления розничной смены.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение розничной смены по ID.
	// Принимает контекст, ID розничной смены и опционально объект параметров запроса Params.
	// Возвращает розничную смену.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*RetailShift, *resty.Response, error)

	// Update выполняет запрос на изменение розничной смены.
	// Принимает контекст, розничную смену и опционально объект параметров запроса Params.
	// Возвращает изменённую розничную смену.
	Update(ctx context.Context, id uuid.UUID, retailShift *RetailShift, params ...*Params) (*RetailShift, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных розничных смен.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

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
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailShift, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

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

const (
	EndpointRetailShift = EndpointEntity + string(MetaTypeRetailShift)
)

// NewRetailShiftService принимает [Client] и возвращает сервис для работы с розничными сменами.
func NewRetailShiftService(client *Client) RetailShiftService {
	return newMainService[RetailShift, any, MetaAttributesStatesSharedWrapper, any](client, EndpointRetailShift)
}
