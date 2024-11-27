package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"

	"time"
)

// RetailStore Точка продаж.
//
// Код сущности: retailstore
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh
type RetailStore struct {
	Acquire                             *NullValue[Agent]           `json:"acquire,omitempty"`                             // Метаданные Банка-эквайера по операциям по карте
	OnlyInStock                         *bool                       `json:"onlyInStock,omitempty"`                         // Выгружать только товары в наличии. Доступно только при активном контроле остатков. Влияет только на выгрузку остатков в POS API
	Active                              *bool                       `json:"active,omitempty"`                              // Состояние точки продаж (Включена/Отключена)
	AccountID                           *string                     `json:"accountId,omitempty"`                           // ID учётной записи
	AddressFull                         *Address                    `json:"addressFull,omitempty"`                         // Адрес с детализацией по отдельным полям
	Meta                                *Meta                       `json:"meta,omitempty"`                                // Метаданные Точки продаж
	AllowCustomPrice                    *bool                       `json:"allowCustomPrice,omitempty"`                    // Разрешить продажу по свободной цене
	AllowDeleteReceiptPositions         *bool                       `json:"allowDeleteReceiptPositions,omitempty"`         // Разрешить удалять позиции в чеке
	AllowSellTobaccoWithoutMRC          *bool                       `json:"allowSellTobaccoWithoutMRC,omitempty"`          // Разрешить продавать табачную продукцию не по МРЦ
	Archived                            *bool                       `json:"archived,omitempty"`                            // Добавлена ли Точка продаж в архив
	AuthTokenAttached                   *bool                       `json:"authTokenAttached,omitempty"`                   // Создан ли токен для точки продаж
	BankPercent                         *float64                    `json:"bankPercent,omitempty"`                         // Комиссия банка-эквайера по операциям по карте (в процентах)
	Cashiers                            *MetaArray[Cashier]         `json:"cashiers,omitempty"`                            // Метаданные Кассиров
	ControlCashierChoice                *bool                       `json:"controlCashierChoice,omitempty"`                // Выбор продавца
	ControlShippingStock                *bool                       `json:"controlShippingStock,omitempty"`                // Контроль остатков. Не может быть true, если AllowCreateProducts имеет значение true
	CreateCashInOnRetailShiftClosing    *bool                       `json:"createCashInOnRetailShiftClosing,omitempty"`    // Создавать ПКО при закрытии смены
	CreateOrderWithState                *NullValue[State]           `json:"createOrderWithState,omitempty"`                // Метаданные статуса, который будет указан при создании заказа
	CreatePaymentInOnRetailShiftClosing *bool                       `json:"createPaymentInOnRetailShiftClosing,omitempty"` // Создавать входящий платеж при закрытии смены
	DemandPrefix                        *string                     `json:"demandPrefix,omitempty"`                        // Префикс номера продаж
	Description                         *string                     `json:"description,omitempty"`                         // Комментарий к Точке продаж
	DiscountEnable                      *bool                       `json:"discountEnable,omitempty"`                      // Разрешить скидки
	DiscountMaxPercent                  *float64                    `json:"discountMaxPercent,omitempty"`                  // Максимальная скидка (в процентах)
	EnableReturnsWithNoReason           *bool                       `json:"enableReturnsWithNoReason,omitempty"`           // Разрешить возвраты без основания
	Environment                         *Environment                `json:"environment,omitempty"`                         // Информация об окружении
	ExternalCode                        *string                     `json:"externalCode,omitempty"`                        // Внешний код Точки продаж
	Group                               *Group                      `json:"group,omitempty"`                               // Отдел сотрудника
	ID                                  *string                     `json:"id,omitempty"`                                  // ID Точки продаж
	Store                               *Store                      `json:"store,omitempty"`                               // Метаданные Склада
	IDQR                                *string                     `json:"idQR,omitempty"`                                // Идентификатор устройства QR (IdQR) для приложения оплаты по QR
	IssueOrders                         *bool                       `json:"issueOrders,omitempty"`                         // Выдача заказов
	ShowBeerOnTap                       *bool                       `json:"showBeerOnTap,omitempty"`                       // Отображать или нет вскрытые кеги на кассе
	SellReserves                        *bool                       `json:"sellReserves,omitempty"`                        // Учет резервов
	ReturnFromClosedShiftEnabled        *bool                       `json:"returnFromClosedShiftEnabled,omitempty"`        // Разрешить возвраты в закрытых сменах
	ReservePrepaidGoods                 *bool                       `json:"reservePrepaidGoods,omitempty"`                 // Резервировать товары, за которые внесена предоплата
	Address                             *string                     `json:"address,omitempty"`                             // Адрес Точки продаж
	SyncAgents                          *bool                       `json:"syncAgents,omitempty"`                          // Выгружать покупателей для работы оффлайн
	Updated                             *Timestamp                  `json:"updated,omitempty"`                             // Момент последнего обновления Точки продаж
	RequiredDiscountCardNumber          *bool                       `json:"requiredDiscountCardNumber,omitempty"`          // Обязательность поля номер бонусной карты при создании контрагента
	AllowCreateProducts                 *bool                       `json:"allowCreateProducts,omitempty"`                 // Контроль остатков. Не может быть true, если controlShippingStock имеет значение true
	Name                                *string                     `json:"name,omitempty"`                                // Наименование Точки продаж
	OFDEnabled                          *bool                       `json:"ofdEnabled,omitempty"`                          // Отправлять электронный чек через ОФД
	State                               *RetailStoreState           `json:"state,omitempty"`                               // Информация о статусе точки продаж
	Shared                              *bool                       `json:"shared,omitempty"`                              // Общий доступ
	OrderToState                        *NullValue[State]           `json:"orderToState,omitempty"`                        // Метаданные статуса, который проставится заказу после проведения продажи на его основании (если указано)
	Organization                        *Organization               `json:"organization,omitempty"`                        // Метаданные Юрлица
	Owner                               *Employee                   `json:"owner,omitempty"`                               // Метаданные владельца (Сотрудника)
	PriceType                           *PriceType                  `json:"priceType,omitempty"`                           // Тип цен, с которыми будут продаваться товары в рознице
	PrintAlways                         *bool                       `json:"printAlways,omitempty"`                         // Всегда печатать кассовые чеки
	SendMarksForCheck                   *bool                       `json:"sendMarksForCheck,omitempty"`                   // Для облачных точек — до продажи отправлять коды маркировки на проверку на точку с ККТ
	ProductFolders                      *MetaArray[ProductFolder]   `json:"productFolders,omitempty"`                      // Коллекция Метаданных групп товаров, из которых можно выгружать товары
	QRAcquire                           *NullValue[Agent]           `json:"qrAcquire,omitempty"`                           // Метаданные Банка-эквайера по операциям по QR-коду
	QRBankPercent                       *float64                    `json:"qrBankPercent,omitempty"`                       // Комиссия банка-эквайера по операция по QR-коду (в процентах)
	QRPayEnabled                        *bool                       `json:"qrPayEnabled,omitempty"`                        // Возможность оплаты по QR-коду на точке продаж
	QRTerminalID                        *string                     `json:"qrTerminalId,omitempty"`                        // Идентификатор терминала (TerminalID) для приложения оплаты по QR
	ReceiptTemplate                     *NullValue[ReceiptTemplate] `json:"receiptTemplate,omitempty"`                     // Метаданные шаблона печати кассовых чеков
	RequiredFio                         *bool                       `json:"requiredFio,omitempty"`                         // Обязательность поля ФИО при создании контрагента
	RequiredPhone                       *bool                       `json:"requiredPhone,omitempty"`                       // Обязательность поля телефон при создании контрагента
	RequiredEmail                       *bool                       `json:"requiredEmail,omitempty"`                       // Обязательность поля эл. почта при создании контрагента
	RequiredBirthdate                   *bool                       `json:"requiredBirthdate,omitempty"`                   // Обязательность поля дата рождения при создании контрагента
	RequiredSex                         *bool                       `json:"requiredSex,omitempty"`                         // Обязательность поля пол при создании контрагента
	SendMarksToChestnyZnakOnCloud       *bool                       `json:"sendMarksToChestnyZnakOnCloud,omitempty"`       // Для облачных точек — отправлять коды маркировки на проверку в Честный Знак
	MinionToMasterType                  MinionToMaster              `json:"minionToMasterType,omitempty"`                  // Стратегия выбора кассы для фискализации облачных чеков
	FiscalType                          FiscalType                  `json:"fiscalType,omitempty"`                          // Тип формирования чеков
	PriorityOFDSend                     PriorityOFDSend             `json:"priorityOfdSend,omitempty"`                     // Приоритет отправки электронного чека. Активен только, когда отправка электронных чеков через ОФД включена
	MarkingSellingMode                  MarkingSellingMode          `json:"markingSellingMode,omitempty"`                  // Режим продажи маркированной продукции, если используется формат фискальных документов версии 1.2
	OrderTaxSystem                      TaxSystem                   `json:"orderTaxSystem,omitempty"`                      // Код системы налогообложения для заказов
	TobaccoMRCControlType               MRCType                     `json:"tobaccoMrcControlType,omitempty"`               // Контроль МРЦ для табачной продукции
	DefaultTaxSystem                    TaxSystem                   `json:"defaultTaxSystem,omitempty"`                    // Код системы налогообложения по умолчанию
	MarksCheckMode                      MarksCheckMode              `json:"marksCheckMode,omitempty"`                      // Настройка проверки КМ перед продажей в ГИС МТ (по умолчанию CORRECT_MARKS_ONLY)
	MasterRetailStores                  Slice[RetailStore]          `json:"masterRetailStores,omitempty"`                  // Ссылка на точки продаж, которые могут фискализировать операции с текущей точки продаж, если minionToMaster = CHOSEN
	LastOperationNames                  Slice[LastOperation]        `json:"lastOperationNames,omitempty"`                  // Последние операции
	FilterAgentsTags                    Slice[string]               `json:"filterAgentsTags,omitempty"`                    // Коллекция групп покупателей, представленных в формате строк. Определяет группы, из которых выгружаются покупатели. Значения null игнорируются
	CustomerOrderStates                 Slice[State]                `json:"customerOrderStates,omitempty"`                 // Метаданные статусов, в которых выгружаются заказы в точку продаж (если указано)
	CreateAgentsTags                    Slice[string]               `json:"createAgentsTags,omitempty"`                    // Коллекция групп покупателей, представленных в формате строк. Определяет группы, в которые добавляются новые покупатели. Значения null игнорируются
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (retailStore RetailStore) Clean() *RetailStore {
	if retailStore.Meta == nil {
		return nil
	}
	return &RetailStore{Meta: retailStore.Meta}
}

// GetAcquire возвращает Метаданные Банка-эквайера по операциям по карте.
func (retailStore RetailStore) GetAcquire() Agent {
	return Deref(retailStore.Acquire).getValue()
}

// GetOnlyInStock возвращает признак выгрузки товаров только в наличии.
//
// Доступно только при активном контроле остатков.
//
// Влияет только на выгрузку остатков в POS API.
func (retailStore RetailStore) GetOnlyInStock() bool {
	return Deref(retailStore.OnlyInStock)
}

// GetActive возвращает состояние точки продаж (Включена/Отключена).
func (retailStore RetailStore) GetActive() bool {
	return Deref(retailStore.Active)
}

// GetAccountID возвращает ID учётной записи.
func (retailStore RetailStore) GetAccountID() string {
	return Deref(retailStore.AccountID)
}

// GetAddressFull возвращает Адрес точки продаж с детализацией по отдельным полям.
func (retailStore RetailStore) GetAddressFull() Address {
	return Deref(retailStore.AddressFull)
}

// GetMeta возвращает Метаданные Точки продаж.
func (retailStore RetailStore) GetMeta() Meta {
	return Deref(retailStore.Meta)
}

// GetAllowCustomPrice возвращает true, если разрешена продажа по свободной цене.
func (retailStore RetailStore) GetAllowCustomPrice() bool {
	return Deref(retailStore.AllowCustomPrice)
}

// GetAllowDeleteReceiptPositions возвращает true, если разрешено удаление позиции в чеке.
func (retailStore RetailStore) GetAllowDeleteReceiptPositions() bool {
	return Deref(retailStore.AllowDeleteReceiptPositions)
}

// GetAllowSellTobaccoWithoutMRC возвращает true, если разрешена продажа табачной продукции не по МРЦ.
func (retailStore RetailStore) GetAllowSellTobaccoWithoutMRC() bool {
	return Deref(retailStore.AllowSellTobaccoWithoutMRC)
}

// GetArchived возвращает true, если Точка продаж добавлена в архив.
func (retailStore RetailStore) GetArchived() bool {
	return Deref(retailStore.Archived)
}

// GetAuthTokenAttached возвращает true, если создан токен для точки продаж.
func (retailStore RetailStore) GetAuthTokenAttached() bool {
	return Deref(retailStore.AuthTokenAttached)
}

// GetBankPercent возвращает Комиссию банка-эквайера по операциям по карте (в процентах).
func (retailStore RetailStore) GetBankPercent() float64 {
	return Deref(retailStore.BankPercent)
}

// GetCashiers возвращает Метаданные кассиров.
func (retailStore RetailStore) GetCashiers() MetaArray[Cashier] {
	return Deref(retailStore.Cashiers)
}

// GetControlCashierChoice возвращает true, если разрешено выбирать продавца.
func (retailStore RetailStore) GetControlCashierChoice() bool {
	return Deref(retailStore.ControlCashierChoice)
}

// GetControlShippingStock возвращает Контроль остатков.
//
// Не может быть true, если AllowCreateProducts имеет значение true.
func (retailStore RetailStore) GetControlShippingStock() bool {
	return Deref(retailStore.ControlShippingStock)
}

// GetCreateCashInOnRetailShiftClosing возвращает true, если при закрытии смены создаётся ПКО.
func (retailStore RetailStore) GetCreateCashInOnRetailShiftClosing() bool {
	return Deref(retailStore.CreateCashInOnRetailShiftClosing)
}

// GetCreateOrderWithState возвращает Метаданные статуса, который будет указан при создании заказ.
func (retailStore RetailStore) GetCreateOrderWithState() State {
	return Deref(retailStore.CreateOrderWithState).getValue()
}

// GetCreatePaymentInOnRetailShiftClosing возвращает true, если при закрытии смены создаётся входящий платеж.
func (retailStore RetailStore) GetCreatePaymentInOnRetailShiftClosing() bool {
	return Deref(retailStore.CreatePaymentInOnRetailShiftClosing)
}

// GetDemandPrefix возвращает Префикс номера продаж.
func (retailStore RetailStore) GetDemandPrefix() string {
	return Deref(retailStore.DemandPrefix)
}

// GetDescription возвращает Комментарий к Точке продаж.
func (retailStore RetailStore) GetDescription() string {
	return Deref(retailStore.Description)
}

// GetDiscountEnable возвращает true, если скидки разрешены.
func (retailStore RetailStore) GetDiscountEnable() bool {
	return Deref(retailStore.DiscountEnable)
}

// GetDiscountMaxPercent возвращает Максимальную скидку (в процентах).
func (retailStore RetailStore) GetDiscountMaxPercent() float64 {
	return Deref(retailStore.DiscountMaxPercent)
}

// GetEnableReturnsWithNoReason возвращает true, если разрешены возвраты без основания.
func (retailStore RetailStore) GetEnableReturnsWithNoReason() bool {
	return Deref(retailStore.EnableReturnsWithNoReason)
}

// GetEnvironment возвращает Информация об окружении.
func (retailStore RetailStore) GetEnvironment() Environment {
	return Deref(retailStore.Environment)
}

// GetExternalCode возвращает Внешний код Точки продаж.
func (retailStore RetailStore) GetExternalCode() string {
	return Deref(retailStore.ExternalCode)
}

// GetGroup возвращает Отдел сотрудника.
func (retailStore RetailStore) GetGroup() Group {
	return Deref(retailStore.Group)
}

// GetID возвращает ID Точки продаж.
func (retailStore RetailStore) GetID() string {
	return Deref(retailStore.ID)
}

// GetStore возвращает Метаданные склада.
func (retailStore RetailStore) GetStore() Store {
	return Deref(retailStore.Store)
}

// GetIDQR возвращает Идентификатор устройства QR (IdQR) для приложения оплаты по QR.
func (retailStore RetailStore) GetIDQR() string {
	return Deref(retailStore.IDQR)
}

// GetIssueOrders возвращает true, если включена выдача заказов.
func (retailStore RetailStore) GetIssueOrders() bool {
	return Deref(retailStore.IssueOrders)
}

// GetShowBeerOnTap возвращает true, если вскрытые кеги отображаются на кассе.
func (retailStore RetailStore) GetShowBeerOnTap() bool {
	return Deref(retailStore.ShowBeerOnTap)
}

// GetSellReserves возвращает Учет резервов.
func (retailStore RetailStore) GetSellReserves() bool {
	return Deref(retailStore.SellReserves)
}

// GetReturnFromClosedShiftEnabled возвращает true, если разрешены возвраты в закрытых сменах.
func (retailStore RetailStore) GetReturnFromClosedShiftEnabled() bool {
	return Deref(retailStore.ReturnFromClosedShiftEnabled)
}

// GetReservePrepaidGoods возвращает true, если товары, за которые внесена предоплата, резервируются.
func (retailStore RetailStore) GetReservePrepaidGoods() bool {
	return Deref(retailStore.ReservePrepaidGoods)
}

// GetAddress возвращает Адрес точки продаж.
func (retailStore RetailStore) GetAddress() string {
	return Deref(retailStore.Address)
}

// GetSyncAgents возвращает true, если покупатели выгружаются для работы оффлайн.
func (retailStore RetailStore) GetSyncAgents() bool {
	return Deref(retailStore.SyncAgents)
}

// GetUpdated возвращает Момент последнего обновления точки продаж.
func (retailStore RetailStore) GetUpdated() time.Time {
	return Deref(retailStore.Updated).Time()
}

// GetRequiredDiscountCardNumber возвращает Обязательность поля номер бонусной карты при создании контрагента.
func (retailStore RetailStore) GetRequiredDiscountCardNumber() bool {
	return Deref(retailStore.RequiredDiscountCardNumber)
}

// GetAllowCreateProducts возвращает Контроль остатков.
//
// Не может быть true, если controlShippingStock имеет значение true.
func (retailStore RetailStore) GetAllowCreateProducts() bool {
	return Deref(retailStore.AllowCreateProducts)
}

// GetName возвращает Наименование Точки продаж.
func (retailStore RetailStore) GetName() string {
	return Deref(retailStore.Name)
}

// GetOFDEnabled возвращает true, если электронные чеки отправляются через ОФД.
func (retailStore RetailStore) GetOFDEnabled() bool {
	return Deref(retailStore.OFDEnabled)
}

// GetState возвращает Информацию о статусе точки продаж.
func (retailStore RetailStore) GetState() RetailStoreState {
	return Deref(retailStore.State)
}

// GetShared возвращает флаг Общего доступа.
func (retailStore RetailStore) GetShared() bool {
	return Deref(retailStore.Shared)
}

// GetOrderToState возвращает Метаданные статуса, который проставится заказу после проведения продажи на его основании (если указано).
func (retailStore RetailStore) GetOrderToState() State {
	return Deref(retailStore.OrderToState).getValue()
}

// GetOrganization возвращает Метаданные юрлица.
func (retailStore RetailStore) GetOrganization() Organization {
	return Deref(retailStore.Organization)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (retailStore RetailStore) GetOwner() Employee {
	return Deref(retailStore.Owner)
}

// GetPriceType возвращает Тип цены, с которыми будут продаваться товары в рознице.
func (retailStore RetailStore) GetPriceType() PriceType {
	return Deref(retailStore.PriceType)
}

// GetPrintAlways возвращает true, если кассовые чеки печатаются всегда.
func (retailStore RetailStore) GetPrintAlways() bool {
	return Deref(retailStore.PrintAlways)
}

// GetSendMarksForCheck возвращает true, если до продажи коды маркировки отправляются на проверку на точку с ККТ (для облачных точек).
func (retailStore RetailStore) GetSendMarksForCheck() bool {
	return Deref(retailStore.SendMarksForCheck)
}

// GetProductFolders возвращает Коллекцию Метаданных групп товаров, из которых можно выгружать товар.
func (retailStore RetailStore) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(retailStore.ProductFolders)
}

// GetQRAcquire возвращает Метаданные Банка-эквайера по операциям по QR-коду.
func (retailStore RetailStore) GetQRAcquire() Agent {
	return Deref(retailStore.QRAcquire).getValue()
}

// GetQRBankPercent возвращает Комиссию банка-эквайера по операция по QR-коду (в процентах).
func (retailStore RetailStore) GetQRBankPercent() float64 {
	return Deref(retailStore.QRBankPercent)
}

// GetQRPayEnabled возвращает Возможность оплаты по QR-коду на точке продаж.
func (retailStore RetailStore) GetQRPayEnabled() bool {
	return Deref(retailStore.QRPayEnabled)
}

// GetQRTerminalID возвращает Идентификатор терминала (TerminalID) для приложения оплаты по QR.
func (retailStore RetailStore) GetQRTerminalID() string {
	return Deref(retailStore.QRTerminalID)
}

// GetReceiptTemplate возвращает Метаданные шаблона печати кассовых чеков.
func (retailStore RetailStore) GetReceiptTemplate() ReceiptTemplate {
	return Deref(retailStore.ReceiptTemplate).getValue()
}

// GetRequiredFio возвращает Обязательность поля ФИО при создании контрагента.
func (retailStore RetailStore) GetRequiredFio() bool {
	return Deref(retailStore.RequiredFio)
}

// GetRequiredPhone возвращает Обязательность поля телефон при создании контрагента.
func (retailStore RetailStore) GetRequiredPhone() bool {
	return Deref(retailStore.RequiredPhone)
}

// GetRequiredEmail возвращает Обязательность поля эл. почта при создании контрагента.
func (retailStore RetailStore) GetRequiredEmail() bool {
	return Deref(retailStore.RequiredEmail)
}

// GetRequiredBirthdate возвращает Обязательность поля дата рождения при создании контрагента.
func (retailStore RetailStore) GetRequiredBirthdate() bool {
	return Deref(retailStore.RequiredBirthdate)
}

// GetRequiredSex возвращает Обязательность поля пол при создании контрагента.
func (retailStore RetailStore) GetRequiredSex() bool {
	return Deref(retailStore.RequiredSex)
}

// GetMinionToMasterType возвращает Стратегию выбора кассы для фискализации облачных чеков.
func (retailStore RetailStore) GetMinionToMasterType() MinionToMaster {
	return retailStore.MinionToMasterType
}

// GetFiscalType возвращает Тип формирования чеков.
func (retailStore RetailStore) GetFiscalType() FiscalType {
	return retailStore.FiscalType
}

// GetPriorityOFDSend возвращает Приоритет отправки электронного чека.
//
// Активен только, когда отправка электронных чеков через ОФД включена.
func (retailStore RetailStore) GetPriorityOFDSend() PriorityOFDSend {
	return retailStore.PriorityOFDSend
}

// GetMarkingSellingMode возвращает Режим продажи маркированной продукции, если используется формат фискальных документов версии 1.2.
func (retailStore RetailStore) GetMarkingSellingMode() MarkingSellingMode {
	return retailStore.MarkingSellingMode
}

// GetOrderTaxSystem возвращает Код системы налогообложения для заказов.
func (retailStore RetailStore) GetOrderTaxSystem() TaxSystem {
	return retailStore.OrderTaxSystem
}

// GetTobaccoMRCControlType возвращает Контроль МРЦ для табачной продукции.
func (retailStore RetailStore) GetTobaccoMRCControlType() MRCType {
	return retailStore.TobaccoMRCControlType
}

// GetDefaultTaxSystem возвращает Код системы налогообложения по умолчанию.
func (retailStore RetailStore) GetDefaultTaxSystem() TaxSystem {
	return retailStore.DefaultTaxSystem
}

// GetMarksCheckMode возвращает Настройку проверки КМ перед продажей в ГИС МТ (по умолчанию CORRECT_MARKS_ONLY).
func (retailStore RetailStore) GetMarksCheckMode() MarksCheckMode {
	return retailStore.MarksCheckMode
}

// GetMasterRetailStores возвращает Ссылку на точки продаж, которые могут фискализировать операции с текущей точки продаж, если minionToMaster = CHOSEN.
func (retailStore RetailStore) GetMasterRetailStores() Slice[RetailStore] {
	return retailStore.MasterRetailStores
}

// GetLastOperationNames возвращает Последние операции.
func (retailStore RetailStore) GetLastOperationNames() Slice[LastOperation] {
	return retailStore.LastOperationNames
}

// GetFilterAgentsTags возвращает Коллекцию групп покупателей, представленных в формате строк.
//
// Определяет группы, из которых выгружаются покупатели.
func (retailStore RetailStore) GetFilterAgentsTags() Slice[string] {
	return retailStore.FilterAgentsTags
}

// GetCustomerOrderStates возвращает Метаданные статусов, в которых выгружаются заказы в точку продаж (если указано).
func (retailStore RetailStore) GetCustomerOrderStates() Slice[State] {
	return retailStore.CustomerOrderStates
}

// GetCreateAgentsTags возвращает Коллекцию групп покупателей, представленных в формате строк.
//
// Определяет группы, в которые добавляются новые покупатели.
func (retailStore RetailStore) GetCreateAgentsTags() Slice[string] {
	return retailStore.CreateAgentsTags
}

// SetAcquire устанавливает Метаданные Банка-эквайера по операциям по карте.
func (retailStore *RetailStore) SetAcquire(acquire AgentOrganizationConverter) *RetailStore {
	if acquire == nil {
		retailStore.Acquire = &NullValue[Agent]{null: true}
	} else {
		retailStore.Acquire = NewNullValue(acquire.AsOrganizationAgent())
	}
	return retailStore
}

// SetOnlyInStock устанавливает правило выгрузки товаров только в наличии.
func (retailStore *RetailStore) SetOnlyInStock(onlyInStock bool) *RetailStore {
	retailStore.OnlyInStock = &onlyInStock
	return retailStore
}

// SetActive устанавливает Состояние точки продаж (Включена/Отключена).
func (retailStore *RetailStore) SetActive(active bool) *RetailStore {
	retailStore.Active = &active
	return retailStore
}

// SetAddressFull устанавливает Адрес точки продаж с детализацией по отдельным полям
func (retailStore *RetailStore) SetAddressFull(addressFull *Address) *RetailStore {
	if addressFull == nil {
		retailStore.SetAddress("")
	} else {
		retailStore.AddressFull = addressFull
	}
	return retailStore
}

// SetMeta устанавливает Метаданные Точки продаж.
func (retailStore *RetailStore) SetMeta(meta *Meta) *RetailStore {
	retailStore.Meta = meta
	return retailStore
}

// SetAllowDeleteReceiptPositions устанавливает разрешение удалять позиции в чеке.
func (retailStore *RetailStore) SetAllowDeleteReceiptPositions(allowDeleteReceiptPositions bool) *RetailStore {
	retailStore.AllowDeleteReceiptPositions = &allowDeleteReceiptPositions
	return retailStore
}

// SetAllowSellTobaccoWithoutMRC устанавливает разрешение продавать табачную продукцию не по МРЦ.
func (retailStore *RetailStore) SetAllowSellTobaccoWithoutMRC(allowSellTobaccoWithoutMRC bool) *RetailStore {
	retailStore.AllowSellTobaccoWithoutMRC = &allowSellTobaccoWithoutMRC
	return retailStore
}

// SetArchived устанавливает флаг нахождения в архиве.
func (retailStore *RetailStore) SetArchived(archived bool) *RetailStore {
	retailStore.Archived = &archived
	return retailStore
}

// SetBankPercent устанавливает Комиссию банка-эквайера по операциям по карте (в процентах).
func (retailStore *RetailStore) SetBankPercent(bankPercent float64) *RetailStore {
	retailStore.BankPercent = &bankPercent
	return retailStore
}

// SetCashiers устанавливает Метаданные Кассиров.
//
// Принимает множество объектов [Cashier].
func (retailStore *RetailStore) SetCashiers(cashiers ...*Cashier) *RetailStore {
	retailStore.Cashiers = NewMetaArrayFrom(cashiers)
	return retailStore
}

// SetControlCashierChoice устанавливает разрешения выбирать продавца.
func (retailStore *RetailStore) SetControlCashierChoice(controlCashierChoice bool) *RetailStore {
	retailStore.ControlCashierChoice = &controlCashierChoice
	return retailStore
}

// SetControlShippingStock устанавливает Контроль остатков.
//
// Не может быть true, если AllowCreateProducts имеет значение true.
func (retailStore *RetailStore) SetControlShippingStock(controlShippingStock bool) *RetailStore {
	retailStore.ControlShippingStock = &controlShippingStock
	return retailStore
}

// SetCreateCashInOnRetailShiftClosing устанавливает правило создания ПКО при закрытии смены.
func (retailStore *RetailStore) SetCreateCashInOnRetailShiftClosing(createCashInOnRetailShiftClosing bool) *RetailStore {
	retailStore.CreateCashInOnRetailShiftClosing = &createCashInOnRetailShiftClosing
	return retailStore
}

// SetCreateOrderWithState устанавливает Метаданные статуса, который будет указан при создании заказа.
//
// Передача nil передаёт сброс значения (null).
func (retailStore *RetailStore) SetCreateOrderWithState(createOrderWithState *State) *RetailStore {
	retailStore.CreateOrderWithState = NewNullValue(createOrderWithState)
	return retailStore
}

// SetCreatePaymentInOnRetailShiftClosing устанавливает правило создания входящего платежа при закрытии смены.
func (retailStore *RetailStore) SetCreatePaymentInOnRetailShiftClosing(createPaymentInOnRetailShiftClosing bool) *RetailStore {
	retailStore.CreatePaymentInOnRetailShiftClosing = &createPaymentInOnRetailShiftClosing
	return retailStore
}

// SetDemandPrefix устанавливает Префикс номера продаж.
func (retailStore *RetailStore) SetDemandPrefix(demandPrefix string) *RetailStore {
	retailStore.DemandPrefix = &demandPrefix
	return retailStore
}

// SetDescription устанавливает Комментарий к Точке продаж.
func (retailStore *RetailStore) SetDescription(description string) *RetailStore {
	retailStore.Description = &description
	return retailStore
}

// SetDiscountEnable устанавливает разрешение делать скидки.
func (retailStore *RetailStore) SetDiscountEnable(discountEnable bool) *RetailStore {
	retailStore.DiscountEnable = &discountEnable
	return retailStore
}

// SetDiscountMaxPercent устанавливает Максимальную скидку (в процентах).
func (retailStore *RetailStore) SetDiscountMaxPercent(discountMaxPercent float64) *RetailStore {
	retailStore.DiscountMaxPercent = &discountMaxPercent
	return retailStore
}

// SetEnableReturnsWithNoReason устанавливает разрешение создавать возвраты без основания.
func (retailStore *RetailStore) SetEnableReturnsWithNoReason(enableReturnsWithNoReason bool) *RetailStore {
	retailStore.EnableReturnsWithNoReason = &enableReturnsWithNoReason
	return retailStore
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (retailStore *RetailStore) SetGroup(group *Group) *RetailStore {
	if group != nil {
		retailStore.Group = group.Clean()
	}
	return retailStore
}

// SetStore устанавливает Метаданные склада.
func (retailStore *RetailStore) SetStore(store *Store) *RetailStore {
	if store != nil {
		retailStore.Store = store.Clean()
	}
	return retailStore
}

// SetIDQR устанавливает Идентификатор устройства QR (IdQR) для приложения оплаты по QR.
func (retailStore *RetailStore) SetIDQR(idqr string) *RetailStore {
	retailStore.IDQR = &idqr
	return retailStore
}

// SetIssueOrders устанавливает разрешение выдавать заказы.
func (retailStore *RetailStore) SetIssueOrders(issueOrders bool) *RetailStore {
	retailStore.IssueOrders = &issueOrders
	return retailStore
}

// SetShowBeerOnTap устанавливает отображение вскрытых кег на кассе.
func (retailStore *RetailStore) SetShowBeerOnTap(showBeerOnTap bool) *RetailStore {
	retailStore.ShowBeerOnTap = &showBeerOnTap
	return retailStore
}

// SetSellReserves устанавливает Учет резервов.
func (retailStore *RetailStore) SetSellReserves(sellReserves bool) *RetailStore {
	retailStore.SellReserves = &sellReserves
	return retailStore
}

// SetReturnFromClosedShiftEnabled устанавливает разрешение создавать возвраты в закрытых сменах.
func (retailStore *RetailStore) SetReturnFromClosedShiftEnabled(returnFromClosedShiftEnabled bool) *RetailStore {
	retailStore.ReturnFromClosedShiftEnabled = &returnFromClosedShiftEnabled
	return retailStore
}

// SetReservePrepaidGoods устанавливает правило резерва товаров, за которые внесена предоплата.
func (retailStore *RetailStore) SetReservePrepaidGoods(reservePrepaidGoods bool) *RetailStore {
	retailStore.ReservePrepaidGoods = &reservePrepaidGoods
	return retailStore
}

// SetAddress устанавливает Адрес Точки продаж.
func (retailStore *RetailStore) SetAddress(address string) *RetailStore {
	retailStore.Address = &address
	return retailStore
}

// SetSyncAgents устанавливает правило выгрузки покупателей для работы оффлайн.
func (retailStore *RetailStore) SetSyncAgents(syncAgents bool) *RetailStore {
	retailStore.SyncAgents = &syncAgents
	return retailStore
}

// SetRequiredDiscountCardNumber устанавливает Обязательность поля номер бонусной карты при создании контрагента.
func (retailStore *RetailStore) SetRequiredDiscountCardNumber(requiredDiscountCardNumber bool) *RetailStore {
	retailStore.RequiredDiscountCardNumber = &requiredDiscountCardNumber
	return retailStore
}

// SetAllowCreateProducts устанавливает Контроль остатков.
//
// Не может быть true, если controlShippingStock имеет значение true.
func (retailStore *RetailStore) SetAllowCreateProducts(allowCreateProducts bool) *RetailStore {
	retailStore.AllowCreateProducts = &allowCreateProducts
	return retailStore
}

// SetName устанавливает Наименование Точки продаж.
func (retailStore *RetailStore) SetName(name string) *RetailStore {
	retailStore.Name = &name
	return retailStore
}

// SetShared устанавливает флаг общего доступа.
func (retailStore *RetailStore) SetShared(shared bool) *RetailStore {
	retailStore.Shared = &shared
	return retailStore
}

// SetOrderToState устанавливает Метаданные статуса, который проставится заказу после проведения продажи на его основании (если указано).
//
// Передача nil передаёт сброс значения (null).
func (retailStore *RetailStore) SetOrderToState(orderToState *State) *RetailStore {
	retailStore.OrderToState = NewNullValue(orderToState)
	return retailStore
}

// SetOrganization устанавливает Метаданные юрлица.
func (retailStore *RetailStore) SetOrganization(organization *Organization) *RetailStore {
	if organization != nil {
		retailStore.Organization = organization.Clean()
	}
	return retailStore
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (retailStore *RetailStore) SetOwner(owner *Employee) *RetailStore {
	if owner != nil {
		retailStore.Owner = owner.Clean()
	}
	return retailStore
}

// SetPriceType устанавливает Тип цен, с которыми будут продаваться товары в рознице.
func (retailStore *RetailStore) SetPriceType(priceType *PriceType) *RetailStore {
	if priceType != nil {
		retailStore.PriceType = priceType.Clean()
	}
	return retailStore
}

// SetPrintAlways устанавливает правило всегда печатать кассовые чеки.
func (retailStore *RetailStore) SetPrintAlways(printAlways bool) *RetailStore {
	retailStore.PrintAlways = &printAlways
	return retailStore
}

// SetSendMarksForCheck устанавливает правило до продажи отправлять коды маркировки на проверку на точку с ККТ (для облачных точек).
func (retailStore *RetailStore) SetSendMarksForCheck(sendMarksForCheck bool) *RetailStore {
	retailStore.SendMarksForCheck = &sendMarksForCheck
	return retailStore
}

// SetProductFolders устанавливает Коллекцию Метаданных групп товаров, из которых можно выгружать товары.
//
// Принимает множество объектов [ProductFolder].
func (retailStore *RetailStore) SetProductFolders(productFolders ...*ProductFolder) *RetailStore {
	retailStore.ProductFolders = NewMetaArrayFrom(productFolders)
	return retailStore
}

// SetQRAcquire устанавливает Метаданные Банка-эквайера по операциям по QR-коду.
func (retailStore *RetailStore) SetQRAcquire(qrAcquire AgentOrganizationConverter) *RetailStore {
	if qrAcquire == nil {
		retailStore.QRAcquire = &NullValue[Agent]{null: true}
	} else {
		retailStore.QRAcquire = NewNullValue(qrAcquire.AsOrganizationAgent())
	}
	return retailStore
}

// SetQRBankPercent устанавливает Комиссию банка-эквайера по операция по QR-коду (в процентах).
func (retailStore *RetailStore) SetQRBankPercent(qrBankPercent float64) *RetailStore {
	retailStore.QRBankPercent = &qrBankPercent
	return retailStore
}

// SetQRPayEnabled устанавливает Возможность оплаты по QR-коду на точке продаж.
func (retailStore *RetailStore) SetQRPayEnabled(qrPayEnabled bool) *RetailStore {
	retailStore.QRPayEnabled = &qrPayEnabled
	return retailStore
}

// SetQRTerminalID устанавливает Идентификатор терминала (TerminalID) для приложения оплаты по QR.
func (retailStore *RetailStore) SetQRTerminalID(qrTerminalID string) *RetailStore {
	retailStore.QRTerminalID = &qrTerminalID
	return retailStore
}

// SetReceiptTemplate устанавливает Метаданные шаблона печати кассовых чеков.
func (retailStore *RetailStore) SetReceiptTemplate(receiptTemplate *ReceiptTemplate) *RetailStore {
	retailStore.ReceiptTemplate = NewNullValue(receiptTemplate)
	return retailStore
}

// SetRequiredFio устанавливает Обязательность поля ФИО при создании контрагента.
func (retailStore *RetailStore) SetRequiredFio(requiredFio bool) *RetailStore {
	retailStore.RequiredFio = &requiredFio
	return retailStore
}

// SetRequiredPhone устанавливает Обязательность поля телефон при создании контрагента.
func (retailStore *RetailStore) SetRequiredPhone(requiredPhone bool) *RetailStore {
	retailStore.RequiredPhone = &requiredPhone
	return retailStore
}

// SetRequiredEmail устанавливает Обязательность поля эл. почта при создании контрагента.
func (retailStore *RetailStore) SetRequiredEmail(requiredEmail bool) *RetailStore {
	retailStore.RequiredEmail = &requiredEmail
	return retailStore
}

// SetRequiredBirthdate устанавливает Обязательность поля дата рождения при создании контрагента.
func (retailStore *RetailStore) SetRequiredBirthdate(requiredBirthdate bool) *RetailStore {
	retailStore.RequiredBirthdate = &requiredBirthdate
	return retailStore
}

// SetRequiredSex устанавливает Обязательность поля пол при создании контрагента.
func (retailStore *RetailStore) SetRequiredSex(requiredSex bool) *RetailStore {
	retailStore.RequiredSex = &requiredSex
	return retailStore
}

// SetMinionToMasterType устанавливает Стратегию выбора кассы для фискализации облачных чеков.
func (retailStore *RetailStore) SetMinionToMasterType(minionToMasterType MinionToMaster) *RetailStore {
	retailStore.MinionToMasterType = minionToMasterType
	return retailStore
}

// SetFiscalType устанавливает Тип формирования чеков.
func (retailStore *RetailStore) SetFiscalType(fiscalType FiscalType) *RetailStore {
	retailStore.FiscalType = fiscalType
	return retailStore
}

// SetPriorityOFDSend устанавливает Приоритет отправки электронного чека.
//
// Активен только, когда отправка электронных чеков через ОФД включена.
func (retailStore *RetailStore) SetPriorityOFDSend(priorityOFDSend PriorityOFDSend) *RetailStore {
	retailStore.PriorityOFDSend = priorityOFDSend
	return retailStore
}

// SetMarkingSellingMode устанавливает Режим продажи маркированной продукции, если используется формат фискальных документов версии 1.2.
func (retailStore *RetailStore) SetMarkingSellingMode(markingSellingMode MarkingSellingMode) *RetailStore {
	retailStore.MarkingSellingMode = markingSellingMode
	return retailStore
}

// SetOrderTaxSystem устанавливает Код системы налогообложения для заказов.
func (retailStore *RetailStore) SetOrderTaxSystem(orderTaxSystem TaxSystem) *RetailStore {
	retailStore.OrderTaxSystem = orderTaxSystem
	return retailStore
}

// SetTobaccoMRCControlType устанавливает Контроль МРЦ для табачной продукции.
func (retailStore *RetailStore) SetTobaccoMRCControlType(tobaccoMRCControlType MRCType) *RetailStore {
	retailStore.TobaccoMRCControlType = tobaccoMRCControlType
	return retailStore
}

// SetDefaultTaxSystem устанавливает Код системы налогообложения по умолчанию.
func (retailStore *RetailStore) SetDefaultTaxSystem(defaultTaxSystem TaxSystem) *RetailStore {
	retailStore.DefaultTaxSystem = defaultTaxSystem
	return retailStore
}

// SetMarksCheckMode устанавливает Настройка проверки КМ перед продажей в ГИС МТ (по умолчанию CORRECT_MARKS_ONLY).
func (retailStore *RetailStore) SetMarksCheckMode(marksCheckMode MarksCheckMode) *RetailStore {
	retailStore.MarksCheckMode = marksCheckMode
	return retailStore
}

// SetMasterRetailStores устанавливает Ссылки на точки продаж, которые могут фискализировать операции с текущей точки продаж, если minionToMaster = CHOSEN.
//
// Принимает множество объектов [RetailStore].
func (retailStore *RetailStore) SetMasterRetailStores(masterRetailStores ...*RetailStore) *RetailStore {
	retailStore.MasterRetailStores.Push(masterRetailStores...)
	return retailStore
}

// SetFilterAgentsTags устанавливает Коллекцию групп покупателей, представленных в формате строк.
//
// Определяет группы, из которых выгружаются покупатели.
//
// Принимает множество string.
func (retailStore *RetailStore) SetFilterAgentsTags(filterAgentsTags ...string) *RetailStore {
	retailStore.FilterAgentsTags = NewSliceFrom(filterAgentsTags)
	return retailStore
}

// SetCustomerOrderStates устанавливает Метаданные статусов, в которых выгружаются заказы в точку продаж (если указано).
//
// Принимает множество объектов [State].
func (retailStore *RetailStore) SetCustomerOrderStates(customerOrderStates ...*State) *RetailStore {
	for _, state := range customerOrderStates {
		if state != nil {
			retailStore.CustomerOrderStates.Push(state.Clean())
		}
	}
	return retailStore
}

// SetCreateAgentsTags устанавливает Коллекцию групп покупателей, представленных в формате строк.
//
// # Определяет группы, в которые добавляются новые покупатели
//
// Принимает множество string.
func (retailStore *RetailStore) SetCreateAgentsTags(createAgentsTags ...string) *RetailStore {
	retailStore.CreateAgentsTags = NewSliceFrom(createAgentsTags)
	return retailStore
}

// String реализует интерфейс [fmt.Stringer].
func (retailStore RetailStore) String() string {
	return Stringify(retailStore)
}

// MetaType возвращает код сущности.
func (RetailStore) MetaType() MetaType {
	return MetaTypeRetailStore
}

// Update shortcut
func (retailStore *RetailStore) Update(ctx context.Context, client *Client, params ...*Params) (*RetailStore, *resty.Response, error) {
	return NewRetailStoreService(client).Update(ctx, retailStore.GetID(), retailStore, params...)
}

// Create shortcut
func (retailStore *RetailStore) Create(ctx context.Context, client *Client, params ...*Params) (*RetailStore, *resty.Response, error) {
	return NewRetailStoreService(client).Create(ctx, retailStore, params...)
}

// Delete shortcut
func (retailStore *RetailStore) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewRetailStoreService(client).Delete(ctx, retailStore)
}

// Cashier Кассир.
//
// Код сущности: cashier
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir
type Cashier struct {
	AccountID   *string      `json:"accountId,omitempty"`   // ID учётной записи
	Employee    *Employee    `json:"employee,omitempty"`    // Метаданные сотрудника, которого представляет собой кассир
	ID          *string      `json:"id,omitempty"`          // ID кассира
	Meta        *Meta        `json:"meta,omitempty"`        // Метаданные кассира
	RetailStore *RetailStore `json:"retailStore,omitempty"` // Метаданные точки продаж, к которой прикреплен кассир
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (cashier Cashier) Clean() *Cashier {
	if cashier.Meta == nil {
		return nil
	}
	return &Cashier{Meta: cashier.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (cashier Cashier) GetAccountID() string {
	return Deref(cashier.AccountID)
}

// GetEmployee возвращает Метаданные сотрудника, которого представляет собой кассир.
func (cashier Cashier) GetEmployee() Employee {
	return Deref(cashier.Employee)
}

// GetID возвращает ID кассира.
func (cashier Cashier) GetID() string {
	return Deref(cashier.ID)
}

// GetMeta возвращает Метаданные кассира.
func (cashier Cashier) GetMeta() Meta {
	return Deref(cashier.Meta)
}

// GetRetailStore возвращает Метаданные точки продаж, к которой прикреплен кассир.
func (cashier Cashier) GetRetailStore() RetailStore {
	return Deref(cashier.RetailStore)
}

// SetEmployee устанавливает Метаданные сотрудника, которого представляет собой кассир.
func (cashier *Cashier) SetEmployee(employee *Employee) *Cashier {
	if employee != nil {
		cashier.Employee = employee.Clean()
	}
	return cashier
}

// SetMeta устанавливает Метаданные кассира.
func (cashier *Cashier) SetMeta(meta *Meta) *Cashier {
	cashier.Meta = meta
	return cashier
}

// SetRetailStore устанавливает Метаданные сотрудника, которого представляет собой кассир.
func (cashier *Cashier) SetRetailStore(retailStore *RetailStore) *Cashier {
	if retailStore != nil {
		cashier.RetailStore = retailStore.Clean()
	}
	return cashier
}

// String реализует интерфейс [fmt.Stringer].
func (cashier Cashier) String() string {
	return Stringify(cashier)
}

// MetaType возвращает код сущности.
func (Cashier) MetaType() MetaType {
	return MetaTypeCashier
}

// RetailStoreState Информация статусе точки продаж
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status
type RetailStoreState struct {
	Sync            RetailStoreStateSync `json:"sync,omitempty"`         // Состояние синхронизации
	LastCheckMoment Timestamp            `json:"lastCheckMoment,"`       // Дата и время последней синхронизации
	FiscalMemory    FiscalMemoryState    `json:"fiscalMemory,omitempty"` // Информация о фискальном накопителе
}

// String реализует интерфейс [fmt.Stringer].
func (retailStoreState RetailStoreState) String() string {
	return Stringify(retailStoreState)
}

// RetailStoreStateSync Состояние синхронизации.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-sinhronizaciq
type RetailStoreStateSync struct {
	LastAttemptMoment Timestamp `json:"lastAttempMoment,omitempty"`
	Message           string    `json:"message,omitempty"`
}

// String реализует интерфейс [fmt.Stringer].
func (retailStoreStateSync RetailStoreStateSync) String() string {
	return Stringify(retailStoreStateSync)
}

// Environment Информация об окружении.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie
type Environment struct {
	Device          string          `json:"device,omitempty"`          // Информация об устройстве
	OS              string          `json:"os,omitempty"`              // Информация об операционной системе
	Software        Software        `json:"software,omitempty"`        // Информация о ПО
	ChequePrinter   ChequePrinter   `json:"chequePrinter,omitempty"`   // Данные о ККТ
	PaymentTerminal PaymentTerminal `json:"paymentTerminal,omitempty"` // Информация о платежном терминале
}

// String реализует интерфейс [fmt.Stringer].
func (environment Environment) String() string {
	return Stringify(environment)
}

// Software Аттрибуты сущности ПО.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-po
type Software struct {
	Name    string `json:"name,omitempty"`    // Наименование ПО
	Vendor  string `json:"vendor,omitempty"`  // Производитель
	Version string `json:"version,omitempty"` // Версия ПО
}

// String реализует интерфейс [fmt.Stringer].
func (software Software) String() string {
	return Stringify(software)
}

// PaymentTerminal Информация о платежном терминале.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-platezhnyj-terminal
type PaymentTerminal struct {
	AcquiringType string `json:"acquiringType,omitempty"` // Информация о типе эквайера (например: inpas/payme)
}

// String реализует интерфейс [fmt.Stringer].
func (paymentTerminal PaymentTerminal) String() string {
	return Stringify(paymentTerminal)
}

// LastOperation Последняя операция.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-poslednie-operacii
type LastOperation struct {
	Entity MetaType `json:"entity,omitempty"` // Код сущности, обозначающее тип последней операции
	Name   string   `json:"name,omitempty"`   // Наименование (номер) последней операции
}

// String реализует интерфейс [fmt.Stringer].
func (lastOperation LastOperation) String() string {
	return Stringify(lastOperation)
}

// ChequePrinter Данные о ККТ.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-kkt
type ChequePrinter struct {
	Driver            Driver       `json:"driver,omitempty"`            // Информация об используемом драйвере
	FirmwareVersion   string       `json:"firmwareVersion,omitempty"`   // Версия прошивки ККТ
	FiscalDataVersion string       `json:"fiscalDataVersion,omitempty"` // Формат фискальных данных
	FiscalMemory      FiscalMemory `json:"fiscalMemory,omitempty"`      // Информация о фискальном накопителе
	Name              string       `json:"name,omitempty"`              // Наименование ПО
	Serial            string       `json:"serial,omitempty"`            // Серийный номер
	Vendor            string       `json:"vendor,omitempty"`            // Производитель
}

// String реализует интерфейс [fmt.Stringer].
func (chequePrinter ChequePrinter) String() string {
	return Stringify(chequePrinter)
}

// Driver Атрибуты сущности Драйвер.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-drajwer
type Driver struct {
	Name    string `json:"name,omitempty"`    // Наименование драйвера
	Version string `json:"version,omitempty"` // Версия драйвера
}

// String реализует интерфейс [fmt.Stringer].
func (driver Driver) String() string {
	return Stringify(driver)
}

// FiscalMemory Атрибуты сущности Фискальный накопитель.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-fiskal-nyj-nakopitel
type FiscalMemory struct {
	FiscalValidityDate Timestamp `json:"fiscalValidityDate,omitempty"`
	FiscalDataVersion  string    `json:"fiscalDataVersion,omitempty"`
}

// String реализует интерфейс [fmt.Stringer].
func (fiscalMemory FiscalMemory) String() string {
	return Stringify(fiscalMemory)
}

// FiscalMemoryState Информация о фискальном накопителе.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-fiskal-naq-pamqt
type FiscalMemoryState struct {
	NotSendFirstDocMoment Timestamp              `json:"notSendFirstDocMoment,omitempty"` // Дата последнего неотправленного документа (?)
	Error                 FiscalMemoryStateError `json:"error,omitempty"`                 // Информация об ошибке ФН
	NotSendDocCount       int                    `json:"notSendDocCount,omitempty"`       // Количество неотправленных документов в ОФД
}

// String реализует интерфейс [fmt.Stringer].
func (fiscalMemoryState FiscalMemoryState) String() string {
	return Stringify(fiscalMemoryState)
}

// FiscalMemoryStateError Информация об ошибке ФН.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-oshibka
type FiscalMemoryStateError struct {
	Code    string `json:"code,omitempty"`    // Код ошибки ФН
	Message string `json:"message,omitempty"` // Описание ошибки
}

// String реализует интерфейс [fmt.Stringer].
func (fiscalMemoryStateError FiscalMemoryStateError) String() string {
	return Stringify(fiscalMemoryStateError)
}

// ReceiptTemplate Шаблон печати кассовых чеков.
//
// TODO: МойСклад не работает с кодом сущности 'receipttemplate'
type ReceiptTemplate struct {
	//AccountID *string `json:"accountId,omitempty"` // ID учётной записи
	//ID        *string `json:"id,omitempty"`        // ID сущности
	//Name   *string `json:"name,omitempty"`   // Наименование
	Meta   *Meta   `json:"meta,omitempty"`   // Метаданные
	Header *string `json:"header,omitempty"` // Верхний блок
	Footer *string `json:"footer,omitempty"` // Нижний блок
}

// String реализует интерфейс [fmt.Stringer].
func (receiptTemplate ReceiptTemplate) String() string {
	return Stringify(receiptTemplate)
}

// FiscalType Тип формирования чеков.
//
// Возможные значения:
//   - FiscalTypeStandard – Стандартное
//   - FiscalTypeMaster   – Стандартное с обработкой облачных операций
//   - FiscalTypeCloud    – Облачное
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-tip-formirowaniq-chekow
type FiscalType string

const (
	FiscalTypeStandard FiscalType = "STANDARD" // Стандартное
	FiscalTypeMaster   FiscalType = "MASTER"   // Стандартное с обработкой облачных операций
	FiscalTypeCloud    FiscalType = "CLOUD"    // Облачное
)

// MarksCheckMode Продажа маркированных товаров.
//
// Возможные значения:
//   - MarksCheckModeCorrectMarksOnly – Только проверенные и правильные коды маркировки
//   - MarksCheckModeWithoutErrors    – Правильные коды и те, которые не удалось проверить
//   - MarksCheckModeAll              – Все — независимо от результатов проверки кодов
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-prodazha-markirowannyh-towarow
type MarksCheckMode string

const (
	MarksCheckModeCorrectMarksOnly MarksCheckMode = "CORRECT_MARKS_ONLY" // Только проверенные и правильные коды маркировки
	MarksCheckModeWithoutErrors    MarksCheckMode = "WITHOUT_ERRORS"     // Правильные коды и те, которые не удалось проверить
	MarksCheckModeAll              MarksCheckMode = "ALL"                // Все — независимо от результатов проверки кодов
)

// MarkingSellingMode Продажа маркированных товаров.
//
// Возможные значения:
//   - MarkingSellingModeCorrectMarksOnly – Только с правильными кодами маркировки
//   - MarkingSellingModeWithoutErrors    – С правильными кодами и те, которые не удалось проверить
//   - MarkingSellingModeAll              – Все – независимо от результатов проверки кодов маркировки
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-prodazha-markirowannyh-towarow
type MarkingSellingMode string

const (
	MarkingSellingModeCorrectMarksOnly MarkingSellingMode = "CORRECT_MARKS_ONLY" // Только с правильными кодами маркировки
	MarkingSellingModeWithoutErrors    MarkingSellingMode = "WITHOUT_ERRORS"     // С правильными кодами и те, которые не удалось проверить
	MarkingSellingModeAll              MarkingSellingMode = "ALL"                // Все – независимо от результатов проверки кодов маркировки
)

// MinionToMaster Стратегия выбора кассы для фискализации облачных чеков.
//
// Возможные значения:
//   - MinionToMasterAny       – Любая мастер касса
//   - MinionToMasterSameGroup – Только кассы из того же отдела
//   - MinionToMasterChosen    – Выбранные кассы из списка в поле masterRetailStores
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-strategiq-wybora-kassy-dlq-fiskalizacii-oblachnyh-chekow
type MinionToMaster string

const (
	MinionToMasterAny       MinionToMaster = "ANY"        // Любая мастер касса
	MinionToMasterSameGroup MinionToMaster = "SAME_GROUP" // Только кассы из того же отдела
	MinionToMasterChosen    MinionToMaster = "CHOSEN"     // Выбранные кассы из списка в поле masterRetailStores
)

// PriorityOFDSend Приоритет отправки электронного чека.
//
// Возможные значения:
//   - PriorityOFDSendPhone – Приоритет отправки на телефон
//   - PriorityOFDSendEmail – Приоритет отправки на e-mail
//   - PriorityOFDSendNone  – Отсутствие отправки чека
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-prioritet-otprawki-alektronnogo-cheka
type PriorityOFDSend string

const (
	PriorityOFDSendPhone PriorityOFDSend = "phone" // Приоритет отправки на телефон
	PriorityOFDSendEmail PriorityOFDSend = "email" // Приоритет отправки на e-mail
	PriorityOFDSendNone  PriorityOFDSend = "none"  // Отсутствие отправки чека
)

// MRCType Тип контроля МРЦ для табачной продукции.
//
// Возможные значения:
//   - MRCTypeUserPrice – Не контролировать МРЦ
//   - MRCTypeMRCPrice  – Продавать по МРЦ указанной на пачке
//   - MRCTypeSamePrice – Запрещать продажу, если цена продажи не совпадает с МРЦ
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-tip-kontrolq-mrc-dlq-tabachnoj-produkcii
type MRCType string

const (
	MRCTypeUserPrice MRCType = "USER_PRICE" // Не контролировать МРЦ
	MRCTypeMRCPrice  MRCType = "MRC_PRICE"  // Продавать по МРЦ указанной на пачке
	MRCTypeSamePrice MRCType = "SAME_PRICE" // Запрещать продажу, если цена продажи не совпадает с МРЦ
)

// RetailStoreService описывает методы сервиса для работы с точками продаж.
type RetailStoreService interface {
	// GetList выполняет запрос на получение списка точек продаж.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[RetailStore], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех точек продаж в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[RetailStore], *resty.Response, error)

	// Create выполняет запрос на создание точи продаж.
	// Обязательные поля для заполнения:
	//	- name (Наименование точки продаж)
	//	- organization (Ссылка на ваше юрлицо)
	//	- store (Ссылка на склад)
	//	- priceType (Тип цен, с которыми будут продаваться товары в рознице)
	// Принимает контекст, точку продаж и опционально объект параметров запроса Params.
	// Возвращает созданную точку продаж.
	Create(ctx context.Context, retailStore *RetailStore, params ...*Params) (*RetailStore, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение точек продаж.
	// Изменяемые точки продаж должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список точек продаж и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых точек продаж.
	CreateUpdateMany(ctx context.Context, retailStore Slice[RetailStore], params ...*Params) (*Slice[RetailStore], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление точек продаж.
	// Принимает контекст и множество точек продаж.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*RetailStore) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление точки продаж по ID.
	// Принимает контекст и ID точки продаж.
	// Возвращает «true» в случае успешного удаления точки продаж.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление точки продаж.
	// Принимает контекст и точку продаж.
	// Возвращает «true» в случае успешного удаления точки продаж.
	Delete(ctx context.Context, entity *RetailStore) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной точки продаж по ID.
	// Принимает контекст, ID точки продаж и опционально объект параметров запроса Params.
	// Возвращает точку продаж.
	GetByID(ctx context.Context, id string, params ...*Params) (*RetailStore, *resty.Response, error)

	// Update выполняет запрос на изменение точки продаж.
	// Принимает контекст, точку продаж и опционально объект параметров запроса Params.
	// Возвращает изменённую точку продаж.
	Update(ctx context.Context, id string, entity *RetailStore, params ...*Params) (*RetailStore, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id string) (*NamedFilter, *resty.Response, error)

	// GetCashiers выполняет запрос на получение списка кассиров.
	// Принимает контекст, ID точки продаж и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetCashiers(ctx context.Context, id string, params ...*Params) (*MetaArray[Cashier], *resty.Response, error)

	// GetCashierByID выполняет запрос на получение отдельного кассира по ID.
	// Принимает контекст, ID точки продаж, ID кассира и опционально объект параметров запроса Params.
	// Возвращает кассира.
	GetCashierByID(ctx context.Context, id, cashierID string, params ...*Params) (*Cashier, *resty.Response, error)
}

const (
	EndpointRetailStore           = EndpointEntity + string(MetaTypeRetailStore)
	EndpointRetailStoreCashiers   = EndpointRetailStore + "/%s/cashiers"
	EndpointRetailStoreCashiersID = EndpointRetailStoreCashiers + "/%s"
)

type retailStoreService struct {
	Endpoint
	endpointGetList[RetailStore]
	endpointCreate[RetailStore]
	endpointCreateUpdateMany[RetailStore]
	endpointDeleteMany[RetailStore]
	endpointDeleteByID
	endpointDelete[RetailStore]
	endpointGetByID[RetailStore]
	endpointUpdate[RetailStore]
	endpointNamedFilter
}

func (service *retailStoreService) GetCashiers(ctx context.Context, id string, params ...*Params) (*MetaArray[Cashier], *resty.Response, error) {
	path := fmt.Sprintf(EndpointRetailStoreCashiers, id)
	return NewRequestBuilder[MetaArray[Cashier]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *retailStoreService) GetCashierByID(ctx context.Context, id, cashierID string, params ...*Params) (*Cashier, *resty.Response, error) {
	path := fmt.Sprintf(EndpointRetailStoreCashiersID, id, cashierID)
	return NewRequestBuilder[Cashier](service.client, path).SetParams(params...).Get(ctx)
}

// NewRetailStoreService принимает [Client] и возвращает сервис для работы с точками продаж.
func NewRetailStoreService(client *Client) RetailStoreService {
	e := NewEndpoint(client, EndpointRetailStore)
	return &retailStoreService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[RetailStore]{e},
		endpointCreate:           endpointCreate[RetailStore]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[RetailStore]{e},
		endpointDeleteMany:       endpointDeleteMany[RetailStore]{e},
		endpointDeleteByID:       endpointDeleteByID{e},
		endpointDelete:           endpointDelete[RetailStore]{e},
		endpointGetByID:          endpointGetByID[RetailStore]{e},
		endpointUpdate:           endpointUpdate[RetailStore]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}
