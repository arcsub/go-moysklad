package moysklad

import (
	"github.com/google/uuid"
)

// RetailShift Розничная смена.
// Ключевое слово: retailshift
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena
type RetailShift struct {
	AccountID           *uuid.UUID        `json:"accountId,omitempty"`           // ID учетной записи
	Acquire             *Counterparty     `json:"acquire,omitempty"`             // Метаданные Банка-эквайера по операциям по карте
	AgentAccount        *AgentAccount     `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Attributes          *Attributes       `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	BankCommission      *Decimal          `json:"bankComission,omitempty"`       // Сумма комиссии эквайера за проведение безналичных платежей по банковской карте. Не может превышать общую сумму безналичных платежей по карте. Если не указано, заполняется 0 автоматически.
	BankPercent         *Decimal          `json:"bankPercent,omitempty"`         // Комиссия банка-эквайера по операциям по карте (в процентах)
	Cheque              *Cheque           `json:"cheque,omitempty"`              // Информация о смене ККТ
	CloseDate           *Timestamp        `json:"closeDate,omitempty"`           // Дата закрытия смены
	Contract            *Contract         `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp        `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp        `json:"deleted,omitempty"`             // Момент последнего удаления Розничной смены
	Description         *string           `json:"description,omitempty"`         // Комментарий Розничной смены
	ExternalCode        *string           `json:"externalCode,omitempty"`        // Внешний код Розничной смены
	Files               *Files            `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group            `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID        `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta             `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp        `json:"moment,omitempty"`              // Дата документа
	Name                *string           `json:"name,omitempty"`                // Наименование
	Operations          *RetailOperations `json:"operations,omitempty"`          // Коллекция метаданных связанных операций
	Organization        *Organization     `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount     `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee         `json:"owner,omitempty"`               // Владелец (Сотрудник)
	Payments            *Payments         `json:"paymentOperations,omitempty"`   // Коллекция метаданных платежных операций
	Printed             *bool             `json:"printed,omitempty"`             // Напечатан ли документ
	ProceedsCash        *Decimal          `json:"proceedsCash,omitempty"`        // Выручка наличными
	ProceedsNoCash      *Decimal          `json:"proceedsNoCash,omitempty"`      // Выручка безнал
	Published           *bool             `json:"published,omitempty"`           // Опубликован ли документ
	QRAcquire           *Counterparty     `json:"qrAcquire,omitempty"`           // Метаданные Банка-эквайера по операциям по QR-коду
	QRBankCommission    *Decimal          `json:"qrBankComission,omitempty"`     // Сумма комиссии эквайера за проведение безналичных платежей по QR-коду. Не может превышать общую сумму безналичных платежей по QR-коду. Если не указано, заполняется 0 автоматически.
	QRBankPercent       *Decimal          `json:"qrBankPercent,omitempty"`       // Комиссия банка-эквайера по операция по QR-коду (в процентах)
	ReceivedCash        *Decimal          `json:"receivedCash,omitempty"`        // Получено наличными
	ReceivedNoCash      *Decimal          `json:"receivedNoCash,omitempty"`      // Получено безнал
	RetailStore         *RetailStore      `json:"retailStore,omitempty"`         // Метаданные точки продаж
	Shared              *bool             `json:"shared,omitempty"`              // Общий доступ
	Store               *Store            `json:"store,omitempty"`               // Метаданные склада. Если не указано, заполняется с точки продаж автоматически
	SyncID              *uuid.UUID        `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp        `json:"updated,omitempty"`             // Момент последнего обновления
	VatEnabled          *bool             `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool             `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
}

func (r RetailShift) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailShift) GetMeta() *Meta {
	return r.Meta
}

func (r RetailShift) MetaType() MetaType {
	return MetaTypeRetailShift
}

// Cheque Информация о смене ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-smene-kkt
type Cheque struct {
	Start ChequeStart `json:"start,omitempty"` // Информация об открытии смены
	End   ChequeEnd   `json:"end,omitempty"`   // Информация о закрытии смены
}

func (c Cheque) String() string {
	return Stringify(c)
}

// ChequeStart Информация об открытии смены ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-ob-otkrytii-smeny-kkt
type ChequeStart struct {
	Time            Timestamp `json:"time,omitempty"`
	FnNumber        string    `json:"fnNumber,omitempty"`
	KktRegNumber    string    `json:"kktRegNumber,omitempty"`
	FiscalDocSign   string    `json:"fiscalDocSign,omitempty"`
	ShiftNumber     string    `json:"shiftNumber,omitempty"`
	FiscalDocNumber string    `json:"fiscalDocNumber,omitempty"`
}

func (c ChequeStart) String() string {
	return Stringify(c)
}

// ChequeEnd Информация о закрытии смены ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-zakrytii-smeny-kkt
type ChequeEnd struct {
	ChequeStart
	ChequesTotal    float64 `json:"chequesTotal,omitempty"`    // Количество чеков за смену
	FiscalDocsTotal float64 `json:"fiscalDocsTotal,omitempty"` // Количество фискальных документов за смену
}

func (c ChequeEnd) String() string {
	return Stringify(c)
}
