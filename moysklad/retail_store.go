package moysklad

import (
	"github.com/google/uuid"
)

// RetailStore Точка продаж.
// Ключевое слово: retailstore
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh
type RetailStore struct {
	AccountID                           *uuid.UUID            `json:"accountId,omitempty"`                           // ID учетной записи
	Acquire                             *Counterparty         `json:"acquire,omitempty"`                             // Метаданные Банка-эквайера по операциям по карте
	Active                              *bool                 `json:"active,omitempty"`                              // Состояние точки продаж (Включена/Отключена)
	Address                             *string               `json:"address,omitempty"`                             // Адрес Точки продаж
	AddressFull                         *Address              `json:"addressFull,omitempty"`                         // Адрес с детализацией по отдельным полям
	AllowCreateProducts                 *bool                 `json:"allowCreateProducts,omitempty"`                 // Контроль остатков. Не может быть true, если controlShippingStock имеет значение true
	AllowCustomPrice                    *bool                 `json:"allowCustomPrice,omitempty"`                    // Разрешить продажу по свободной цене
	AllowDeleteReceiptPositions         *bool                 `json:"allowDeleteReceiptPositions,omitempty"`         // Разрешить удалять позиции в чеке
	AllowSellTobaccoWithoutMRC          *bool                 `json:"allowSellTobaccoWithoutMRC,omitempty"`          // Разрешить продавать табачную продукцию не по МРЦ
	Archived                            *bool                 `json:"archived,omitempty"`                            // Добавлена ли Точка продаж в архив
	AuthTokenAttached                   *bool                 `json:"authTokenAttached,omitempty"`                   // Создан ли токен для точки продаж
	BankPercent                         *float64              `json:"bankPercent,omitempty"`                         // Комиссия банка-эквайера по операциям по карте (в процентах)
	Cashiers                            *Cashiers             `json:"cashiers,omitempty"`                            // Метаданные Кассиров
	ControlCashierChoice                *bool                 `json:"controlCashierChoice,omitempty"`                // Выбор продавца
	ControlShippingStock                *bool                 `json:"controlShippingStock,omitempty"`                // Контроль остатков. Не может быть true, если AllowCreateProducts имеет значение true
	CreateAgentsTags                    *Tags                 `json:"createAgentsTags,omitempty"`                    // Коллекция групп покупателей, представленных в формате строк. Определяет группы, в которые добавляются новые покупатели. Значения null игнорируются
	CreateCashInOnRetailShiftClosing    *bool                 `json:"createCashInOnRetailShiftClosing,omitempty"`    // Создавать ПКО при закрытии смены
	CreateOrderWithState                *State                `json:"createOrderWithState,omitempty"`                // Метаданные статуса, который будет указан при создании заказа
	CreatePaymentInOnRetailShiftClosing *bool                 `json:"createPaymentInOnRetailShiftClosing,omitempty"` // Создавать входящий платеж при закрытии смены
	CustomerOrderStates                 *States               `json:"customerOrderStates,omitempty"`                 // Метаданные статусов, в которых выгружаются заказы в точку продаж (если указано)
	DefaultTaxSystem                    TaxSystem             `json:"defaulTaxSystem,omitempty"`                     // Код системы налогообложения по умолчанию
	DemandPrefix                        *string               `json:"demandPrefix,omitempty"`                        // Префикс номера продаж
	Description                         *string               `json:"description,omitempty"`                         // Комментарий к Точке продаж
	DiscountEnable                      *bool                 `json:"discountEnable,omitempty"`                      // Разрешить скидки
	DiscountMaxPercent                  *float64              `json:"discountMaxPercent,omitempty"`                  // Максимальная скидка (в процентах)
	EnableReturnsWithNoReason           *bool                 `json:"enableReturnsWithNoReason,omitempty"`           // Разрешить возвраты без основания
	Environment                         *Environment          `json:"environment,omitempty"`                         // Информация об окружении
	ExternalCode                        *string               `json:"externalCode,omitempty"`                        // Внешний код Точки продаж
	FilterAgentsTags                    *Tags                 `json:"filterAgentsTags,omitempty"`                    // Коллекция групп покупателей, представленных в формате строк. Определяет группы, из которых выгружаются покупатели. Значения null игнорируются
	FiscalType                          FiscalType            `json:"fiscalType,omitempty"`                          // Тип формирования чеков
	Group                               *Group                `json:"group,omitempty"`                               // Отдел сотрудника
	ID                                  *uuid.UUID            `json:"id,omitempty"`                                  // ID сущности
	IdQR                                *string               `json:"idQR,omitempty"`                                // Идентификатор устройства QR (IdQR) для приложения оплаты по QR
	IssueOrders                         *bool                 `json:"issueOrders,omitempty"`                         // Выдача заказов
	LastOperationNames                  *LastOperations       `json:"lastOperationNames,omitempty"`                  // Последние операции
	MarkingSellingMode                  MarkingSellingMode    `json:"markingSellingMode,omitempty"`                  // Режим продажи маркированной продукции, если используется формат фискальных документов версии 1.2
	MasterRetailStores                  *RetailStores         `json:"masterRetailStores,omitempty"`                  // Ссылка на точки продаж, которые могут фискализировать операции с текущей точки продаж, если minionToMaster = CHOSEN
	Meta                                *Meta                 `json:"meta,omitempty"`                                // Метаданные
	MinionToMasterType                  MinionToMaster        `json:"minionToMasterType,omitempty"`                  // Стратегия выбора кассы для фискализации облачных чеков
	Name                                *string               `json:"name,omitempty"`                                // Название
	OFDEnabled                          *bool                 `json:"ofdEnabled,omitempty"`                          // Отправлять электронный чек через ОФД
	OnlyInStock                         *bool                 `json:"onlyInStock,omitempty"`                         // Выгружать только товары в наличии. Доступно только при активном контроле остатков. Влияет только на выгрузку остатков в POS API
	OrderTaxSystem                      TaxSystem             `json:"orderTaxSystem,omitempty"`                      // Код системы налогообложения для заказов
	OrderToState                        *State                `json:"orderToState,omitempty"`                        // Метаданные статуса, который проставится заказу после проведения продажи на его основании (если указано)
	Organization                        *Organization         `json:"organization,omitempty"`                        // Метаданные Юрлица
	Owner                               *Employee             `json:"owner,omitempty"`                               // Владелец (Сотрудник)
	PriceType                           *PriceType            `json:"priceType,omitempty"`                           // Тип цен, с которыми будут продаваться товары в рознице
	PrintAlways                         *bool                 `json:"printAlways,omitempty"`                         // Всегда печатать кассовые чеки
	PriorityOFDSend                     PriorityOfdSend       `json:"priorityOfdSend,omitempty"`                     // Приоритет отправки электронного чека. Активен только, когда отправка электронных чеков через ОФД включена
	ProductFolders                      *ProductFolders       `json:"productFolders,omitempty"`                      // Коллекция Метаданных групп товаров, из которых можно выгружать товары
	QRAcquire                           *Counterparty         `json:"qrAcquire,omitempty"`                           // Метаданные Банка-эквайера по операциям по QR-коду
	QRBankPercent                       *float64              `json:"qrBankPercent,omitempty"`                       // Комиссия банка-эквайера по операция по QR-коду (в процентах)
	QRPayEnabled                        *bool                 `json:"qrPayEnabled,omitempty"`                        // Возможность оплаты по QR-коду на точке продаж
	QRTerminalId                        *string               `json:"qrTerminalId,omitempty"`                        // Идентификатор терминала (TerminalID) для приложения оплаты по QR
	ReceiptTemplate                     *Meta                 `json:"receiptTemplate,omitempty"`                     // Метаданные шаблона печати кассовых чеков TODO: expand не работает
	RequiredFio                         *bool                 `json:"requiredFio,omitempty"`                         // Обязательность поля ФИО при создании контрагента
	RequiredPhone                       *bool                 `json:"requiredPhone,omitempty"`                       // Обязательность поля телефон при создании контрагента
	RequiredEmail                       *bool                 `json:"requiredEmail,omitempty"`                       // Обязательность поля эл. почта при создании контрагента
	RequiredBirthdate                   *bool                 `json:"RequiredBirthdate,omitempty"`                   // Обязательность поля дата рождения при создании контрагента
	RequiredSex                         *bool                 `json:"requiredSex,omitempty"`                         // Обязательность поля пол при создании контрагента
	RequiredDiscountCardNumber          *bool                 `json:"requiredDiscountCardNumber,omitempty"`          // Обязательность поля номер бонусной карты при создании контрагента
	ReservePrepaidGoods                 *bool                 `json:"reservePrepaidGoods,omitempty"`                 // Резервировать товары, за которые внесена предоплата
	ReturnFromClosedShiftEnabled        *bool                 `json:"returnFromClosedShiftEnabled,omitempty"`        // Разрешить возвраты в закрытых сменах
	SellReserves                        *bool                 `json:"sellReserves,omitempty"`                        // Учёт резервов
	SendMarksForCheck                   *bool                 `json:"sendMarksForCheck,omitempty"`                   // Для облачных точек — до продажи отправлять коды маркировки на проверку на точку с ККТ
	Shared                              *bool                 `json:"shared,omitempty"`                              // Общий доступ
	State                               *RetailStoreState     `json:"state,omitempty"`                               // Информация статусе точки продаж
	Store                               *Store                `json:"store,omitempty"`                               // Метаданные Склада
	TobaccoMRCControlType               TobaccoMRCControlType `json:"tobaccoMrcControlType,omitempty"`               // Контроль МРЦ для табачной продукции
	Updated                             *Timestamp            `json:"updated,omitempty"`                             // Момент последнего обновления Точки продаж
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
	Message           string    `json:"message,omitempty"`          // Состояние синхронизации
	LastAttemptMoment Timestamp `json:"lastAttempMoment,omitempty"` // Дата последней сихронизации (не обязательно успешной)
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
	FiscalDataVersion  string    `json:"fiscalDataVersion,omitempty"`  // Версия фискальной памяти
	FiscalValidityDate Timestamp `json:"fiscalValidityDate,omitempty"` // Версия фискальной памяти
}

func (f FiscalMemory) String() string {
	return Stringify(f)
}

// FiscalMemoryState Информация о фискальном накопителе.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-fiskal-naq-pamqt
type FiscalMemoryState struct {
	Error                 FiscalMemoryStateError `json:"error,omitempty"`                 // Информация об ошибке ФН
	NotSendDocCount       int                    `json:"notSendDocCount,omitempty"`       // Количество неотправленных документов в ОФД
	NotSendFirstDocMoment Timestamp              `json:"notSendFirstDocMoment,omitempty"` // Дата последнего неотправленного документа (?)
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
