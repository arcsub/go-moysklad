package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailStore Точка продаж.
// Ключевое слово: retailstore
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh
type RetailStore struct {
	Meta                                *Meta                 `json:"meta,omitempty"`
	AllowCreateProducts                 *bool                 `json:"allowCreateProducts,omitempty"`
	Active                              *bool                 `json:"active,omitempty"`
	AccountID                           *uuid.UUID            `json:"accountId,omitempty"`
	AddressFull                         *Address              `json:"addressFull,omitempty"`
	Updated                             *Timestamp            `json:"updated,omitempty"`
	AllowCustomPrice                    *bool                 `json:"allowCustomPrice,omitempty"`
	AllowDeleteReceiptPositions         *bool                 `json:"allowDeleteReceiptPositions,omitempty"`
	AllowSellTobaccoWithoutMRC          *bool                 `json:"allowSellTobaccoWithoutMRC,omitempty"`
	Archived                            *bool                 `json:"archived,omitempty"`
	AuthTokenAttached                   *bool                 `json:"authTokenAttached,omitempty"`
	BankPercent                         *Decimal              `json:"bankPercent,omitempty"`
	Cashiers                            *Cashiers             `json:"cashiers,omitempty"`
	ControlCashierChoice                *bool                 `json:"controlCashierChoice,omitempty"`
	ControlShippingStock                *bool                 `json:"controlShippingStock,omitempty"`
	CreateAgentsTags                    *Tags                 `json:"createAgentsTags,omitempty"`
	CreateCashInOnRetailShiftClosing    *bool                 `json:"createCashInOnRetailShiftClosing,omitempty"`
	CreateOrderWithState                *State                `json:"createOrderWithState,omitempty"`
	CreatePaymentInOnRetailShiftClosing *bool                 `json:"createPaymentInOnRetailShiftClosing,omitempty"`
	CustomerOrderStates                 *States               `json:"customerOrderStates,omitempty"`
	DemandPrefix                        *string               `json:"demandPrefix,omitempty"`
	Description                         *string               `json:"description,omitempty"`
	DiscountEnable                      *bool                 `json:"discountEnable,omitempty"`
	DiscountMaxPercent                  *Decimal              `json:"discountMaxPercent,omitempty"`
	EnableReturnsWithNoReason           *bool                 `json:"enableReturnsWithNoReason,omitempty"`
	Environment                         *Environment          `json:"environment,omitempty"`
	ExternalCode                        *string               `json:"externalCode,omitempty"`
	FilterAgentsTags                    *Tags                 `json:"filterAgentsTags,omitempty"`
	Group                               *Group                `json:"group,omitempty"`
	ID                                  *uuid.UUID            `json:"id,omitempty"`
	IdQR                                *string               `json:"idQR,omitempty"`
	IssueOrders                         *bool                 `json:"issueOrders,omitempty"`
	LastOperationNames                  *LastOperations       `json:"lastOperationNames,omitempty"`
	MasterRetailStores                  *RetailStores         `json:"masterRetailStores,omitempty"`
	Address                             *string               `json:"address,omitempty"`
	SyncAgents                          *bool                 `json:"syncAgents,omitempty"`
	Acquire                             *Counterparty         `json:"acquire,omitempty"`
	Store                               *Store                `json:"store,omitempty"`
	State                               *RetailStoreState     `json:"state,omitempty"`
	Name                                *string               `json:"name,omitempty"`
	OFDEnabled                          *bool                 `json:"ofdEnabled,omitempty"`
	OnlyInStock                         *bool                 `json:"onlyInStock,omitempty"`
	Shared                              *bool                 `json:"shared,omitempty"`
	OrderToState                        *State                `json:"orderToState,omitempty"`
	Organization                        *Organization         `json:"organization,omitempty"`
	Owner                               *Employee             `json:"owner,omitempty"`
	PriceType                           *PriceType            `json:"priceType,omitempty"`
	PrintAlways                         *bool                 `json:"printAlways,omitempty"`
	SendMarksForCheck                   *bool                 `json:"sendMarksForCheck,omitempty"`
	ProductFolders                      *ProductFolders       `json:"productFolders,omitempty"`
	QRAcquire                           *Counterparty         `json:"qrAcquire,omitempty"`
	QRBankPercent                       *Decimal              `json:"qrBankPercent,omitempty"`
	QRPayEnabled                        *bool                 `json:"qrPayEnabled,omitempty"`
	QRTerminalId                        *string               `json:"qrTerminalId,omitempty"`
	ReceiptTemplate                     *Meta                 `json:"receiptTemplate,omitempty"`
	RequiredFio                         *bool                 `json:"requiredFio,omitempty"`
	RequiredPhone                       *bool                 `json:"requiredPhone,omitempty"`
	RequiredEmail                       *bool                 `json:"requiredEmail,omitempty"`
	RequiredBirthdate                   *bool                 `json:"RequiredBirthdate,omitempty"`
	RequiredSex                         *bool                 `json:"requiredSex,omitempty"`
	RequiredDiscountCardNumber          *bool                 `json:"requiredDiscountCardNumber,omitempty"`
	ReservePrepaidGoods                 *bool                 `json:"reservePrepaidGoods,omitempty"`
	ReturnFromClosedShiftEnabled        *bool                 `json:"returnFromClosedShiftEnabled,omitempty"`
	SellReserves                        *bool                 `json:"sellReserves,omitempty"`
	ShowBeerOnTap                       *bool                 `json:"showBeerOnTap,omitempty"` // [28-11-2023]
	PriorityOFDSend                     PriorityOfdSend       `json:"priorityOfdSend,omitempty"`
	OrderTaxSystem                      TaxSystem             `json:"orderTaxSystem,omitempty"`
	MinionToMasterType                  MinionToMaster        `json:"minionToMasterType,omitempty"`
	MarkingSellingMode                  MarkingSellingMode    `json:"markingSellingMode,omitempty"`
	FiscalType                          FiscalType            `json:"fiscalType,omitempty"`
	TobaccoMRCControlType               TobaccoMRCControlType `json:"tobaccoMrcControlType,omitempty"`
	DefaultTaxSystem                    TaxSystem             `json:"defaulTaxSystem,omitempty"`
	MarksCheckMode                      MarksCheckMode        `json:"marksCheckMode,omitempty"` // [07-02-2024]
}

func (r RetailStore) String() string {
	return Stringify(r)
}

func (r RetailStore) MetaType() MetaType {
	return MetaTypeRetailStore
}

type RetailStores = Slice[RetailStore]

// RetailStoreState Информация статусе точки продаж
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status
type RetailStoreState struct {
	Sync            RetailStoreStateSync `json:"sync,omitempty"`         // Состояние синхронизации
	LastCheckMoment Timestamp            `json:"lastCheckMoment,"`       // Дата и время последней синхронизации
	FiscalMemory    FiscalMemoryState    `json:"fiscalMemory,omitempty"` // Информация о фискальном накопителе
}

func (r RetailStoreState) String() string {
	return Stringify(r)
}

// RetailStoreStateSync Состояние синхронизации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-sinhronizaciq
type RetailStoreStateSync struct {
	LastAttemptMoment Timestamp `json:"lastAttempMoment,omitempty"`
	Message           string    `json:"message,omitempty"`
}

func (r RetailStoreStateSync) String() string {
	return Stringify(r)
}

// Environment Информация об окружении.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie
type Environment struct {
	Device          string          `json:"device,omitempty"`          // Информация об устройстве
	OS              string          `json:"os,omitempty"`              // Информация об операционной системе
	Software        Software        `json:"software,omitempty"`        // Информация о ПО
	ChequePrinter   ChequePrinter   `json:"chequePrinter,omitempty"`   // Данные о ККТ
	PaymentTerminal PaymentTerminal `json:"paymentTerminal,omitempty"` // Информация о платежном терминале
}

func (e Environment) String() string {
	return Stringify(e)
}

// Software Аттрибуты сущности ПО
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-po
type Software struct {
	Name    string `json:"name,omitempty"`    // Наименование ПО
	Vendor  string `json:"vendor,omitempty"`  // Производитель
	Version string `json:"version,omitempty"` // Версия ПО
}

func (s Software) String() string {
	return Stringify(s)
}

// PaymentTerminal Информация о платежном терминале.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-platezhnyj-terminal
type PaymentTerminal struct {
	AcquiringType string `json:"acquiringType,omitempty"` // Информация о типе эквайера (например: inpas/payme)
}

func (p PaymentTerminal) String() string {
	return Stringify(p)
}

// LastOperation Последняя операция.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-poslednie-operacii
type LastOperation struct {
	Entity MetaType `json:"entity,omitempty"` // Ключевое слово, обозначающее тип последней операции
	Name   string   `json:"name,omitempty"`   // Наименование (номер) последней операции
}

func (l LastOperation) String() string {
	return Stringify(l)
}

type LastOperations = Slice[LastOperation]

// ChequePrinter Данные о ККТ.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-kkt
type ChequePrinter struct {
	Driver            Driver       `json:"driver,omitempty"`            // Информация об используемом драйвере
	FirmwareVersion   string       `json:"firmwareVersion,omitempty"`   // Версия прошивки ККТ
	FiscalDataVersion string       `json:"fiscalDataVersion,omitempty"` // Формат фискальных данных
	FiscalMemory      FiscalMemory `json:"fiscalMemory,omitempty"`      // Информация о фискальном накопителе
	Name              string       `json:"name,omitempty"`              // Наименование ПО
	Serial            string       `json:"serial,omitempty"`            // Серийный номер
	Vendor            string       `json:"vendor,omitempty"`            // Производитель
}

func (c ChequePrinter) String() string {
	return Stringify(c)
}

// Driver Атрибуты сущности Драйвер.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-drajwer
type Driver struct {
	Name    string `json:"name,omitempty"`    // Наименование драйвера
	Version string `json:"version,omitempty"` // Версия драйвера
}

func (d Driver) String() string {
	return Stringify(d)
}

// FiscalMemory Атрибуты сущности Фискальный накопитель.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-fiskal-nyj-nakopitel
type FiscalMemory struct {
	FiscalValidityDate Timestamp `json:"fiscalValidityDate,omitempty"`
	FiscalDataVersion  string    `json:"fiscalDataVersion,omitempty"`
}

func (f FiscalMemory) String() string {
	return Stringify(f)
}

// FiscalMemoryState Информация о фискальном накопителе.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-fiskal-naq-pamqt
type FiscalMemoryState struct {
	NotSendFirstDocMoment Timestamp              `json:"notSendFirstDocMoment,omitempty"` // Дата последнего неотправленного документа (?)
	Error                 FiscalMemoryStateError `json:"error,omitempty"`                 // Информация об ошибке ФН
	NotSendDocCount       int                    `json:"notSendDocCount,omitempty"`       // Количество неотправленных документов в ОФД
}

func (f FiscalMemoryState) String() string {
	return Stringify(f)
}

// FiscalMemoryStateError Информация об ошибке ФН.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-oshibka
type FiscalMemoryStateError struct {
	Code    string `json:"code,omitempty"`    // Код ошибки ФН
	Message string `json:"message,omitempty"` // Описание ошибки
}

func (f FiscalMemoryStateError) String() string {
	return Stringify(f)
}

// FiscalType Тип формирования чеков.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-tip-formirowaniq-chekow
type FiscalType string

const (
	FiscalTypeStandart FiscalType = "STANDARD" // Стандартное
	FiscalTypeMaster   FiscalType = "MASTER"   // Стандартное с обработкой облачных операций
	FiscalTypeCloud    FiscalType = "CLOUD"    // Облачное
)

// MarksCheckMode Продажа маркированных товаров
// https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-prodazha-markirowannyh-towarow
type MarksCheckMode string

const (
	MarksCheckModeCorrectMarksOnly MarksCheckMode = "CORRECT_MARKS_ONLY" // Только проверенные и правильные коды маркировки
	MarksCheckModeWithoutErrors    MarksCheckMode = "WITHOUT_ERRORS"     // Правильные коды и те, которые не удалось проверить
	MarksCheckModeAll              MarksCheckMode = "ALL"                // Все — независимо от результатов проверки кодов
)

// MarkingSellingMode Продажа маркированных товаров.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-prodazha-markirowannyh-towarow
type MarkingSellingMode string

const (
	MarkingSellingModeCorrectMarksOnly MarkingSellingMode = "CORRECT_MARKS_ONLY" // Только с правильными кодами маркировки
	MarkingSellingModeWithoutErrors    MarkingSellingMode = "WITHOUT_ERRORS"     // С правильными кодами и те, которые не удалось проверить
	MarkingSellingModeAll              MarkingSellingMode = "ALL"                // Все – независимо от результатов проверки кодов маркировки
)

// MinionToMaster Стратегия выбора кассы для фискализации облачных чеков.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-strategiq-wybora-kassy-dlq-fiskalizacii-oblachnyh-chekow
type MinionToMaster string

const (
	MinionToMasterAny       MinionToMaster = "ANY"        // Любая мастер касса
	MinionToMasterSameGroup MinionToMaster = "SAME_GROUP" // Только кассы из того же отдела
	MinionToMasterChosen    MinionToMaster = "CHOSEN"     // Выбранные кассы из списка в поле masterRetailStores
)

// PriorityOfdSend Приоритет отправки электронного чека.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-prioritet-otprawki-alektronnogo-cheka
type PriorityOfdSend string

const (
	PriorityOfdSendPhone PriorityOfdSend = "phone" // Приоритет отправки на телефон
	PriorityOfdSendEmail PriorityOfdSend = "email" // Приоритет отправки на e-mail
	PriorityOfdSendNone  PriorityOfdSend = "none"  // Отсутствие отправки чека
)

// TobaccoMRCControlType Тип контроля МРЦ для табачной продукции.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-tip-kontrolq-mrc-dlq-tabachnoj-produkcii
type TobaccoMRCControlType string

const (
	TobaccoMRCControlTypeUserPrice TobaccoMRCControlType = "USER_PRICE" // Не контролировать МРЦ
	TobaccoMRCControlTypeMrcPrice  TobaccoMRCControlType = "MRC_PRICE"  // Продавать по МРЦ указанной на пачке
	TobaccoMRCControlTypeSamePrice TobaccoMRCControlType = "SAME_PRICE" // Запрещать продажу, если цена продажи не совпадает с МРЦ
)

// RetailStoreService
// Сервис для работы с точками продаж.
type RetailStoreService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailStore], *resty.Response, error)
	Create(ctx context.Context, retailStore *RetailStore, params *Params) (*RetailStore, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailStore []*RetailStore, params *Params) (*[]RetailStore, *resty.Response, error)
	DeleteMany(ctx context.Context, retailStore []*RetailStore) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailStore, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, entity *RetailStore, params *Params) (*RetailStore, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetCashiers(ctx context.Context, id *uuid.UUID) (*MetaArray[Cashier], *resty.Response, error)
	GetCashierById(ctx context.Context, id, cashierId *uuid.UUID) (*Cashier, *resty.Response, error)
}

type retailStoreService struct {
	Endpoint
	endpointGetList[RetailStore]
	endpointCreate[RetailStore]
	endpointCreateUpdateMany[RetailStore]
	endpointDeleteMany[RetailStore]
	endpointDelete
	endpointGetById[RetailStore]
	endpointUpdate[RetailStore]
	endpointNamedFilter
}

func NewRetailStoreService(client *Client) RetailStoreService {
	e := NewEndpoint(client, "entity/retailstore")
	return &retailStoreService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[RetailStore]{e},
		endpointCreate:           endpointCreate[RetailStore]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[RetailStore]{e},
		endpointDeleteMany:       endpointDeleteMany[RetailStore]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[RetailStore]{e},
		endpointUpdate:           endpointUpdate[RetailStore]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// GetCashiers Получить Кассиров.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassirow
func (s *retailStoreService) GetCashiers(ctx context.Context, id *uuid.UUID) (*MetaArray[Cashier], *resty.Response, error) {
	path := fmt.Sprintf("entity/retailstore/%s/cashiers", id)
	return NewRequestBuilder[MetaArray[Cashier]](s.client, path).Get(ctx)
}

// GetCashierById Получить Кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassira
func (s *retailStoreService) GetCashierById(ctx context.Context, id, cashierId *uuid.UUID) (*Cashier, *resty.Response, error) {
	path := fmt.Sprintf("entity/retailstore/%s/cashiers/%s", id, cashierId)
	return NewRequestBuilder[Cashier](s.client, path).Get(ctx)
}
