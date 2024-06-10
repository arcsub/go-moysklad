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
	Acquire                             *Counterparty             `json:"acquire,omitempty"`
	OnlyInStock                         *bool                     `json:"onlyInStock,omitempty"`
	Active                              *bool                     `json:"active,omitempty"`
	AccountID                           *uuid.UUID                `json:"accountId,omitempty"`
	AddressFull                         *Address                  `json:"addressFull,omitempty"`
	Meta                                *Meta                     `json:"meta,omitempty"`
	AllowCustomPrice                    *bool                     `json:"allowCustomPrice,omitempty"`
	AllowDeleteReceiptPositions         *bool                     `json:"allowDeleteReceiptPositions,omitempty"`
	AllowSellTobaccoWithoutMRC          *bool                     `json:"allowSellTobaccoWithoutMRC,omitempty"`
	Archived                            *bool                     `json:"archived,omitempty"`
	AuthTokenAttached                   *bool                     `json:"authTokenAttached,omitempty"`
	BankPercent                         *float64                  `json:"bankPercent,omitempty"`
	Cashiers                            *MetaArray[Cashier]       `json:"cashiers,omitempty"`
	ControlCashierChoice                *bool                     `json:"controlCashierChoice,omitempty"`
	ControlShippingStock                *bool                     `json:"controlShippingStock,omitempty"`
	CreateCashInOnRetailShiftClosing    *bool                     `json:"createCashInOnRetailShiftClosing,omitempty"`
	CreateOrderWithState                *State                    `json:"createOrderWithState,omitempty"`
	CreatePaymentInOnRetailShiftClosing *bool                     `json:"createPaymentInOnRetailShiftClosing,omitempty"`
	DemandPrefix                        *string                   `json:"demandPrefix,omitempty"`
	Description                         *string                   `json:"description,omitempty"`
	DiscountEnable                      *bool                     `json:"discountEnable,omitempty"`
	DiscountMaxPercent                  *float64                  `json:"discountMaxPercent,omitempty"`
	EnableReturnsWithNoReason           *bool                     `json:"enableReturnsWithNoReason,omitempty"`
	Environment                         *Environment              `json:"environment,omitempty"`
	ExternalCode                        *string                   `json:"externalCode,omitempty"`
	Group                               *Group                    `json:"group,omitempty"`
	ID                                  *uuid.UUID                `json:"id,omitempty"`
	Store                               *Store                    `json:"store,omitempty"`
	IDQR                                *string                   `json:"idQR,omitempty"`
	IssueOrders                         *bool                     `json:"issueOrders,omitempty"`
	ShowBeerOnTap                       *bool                     `json:"showBeerOnTap,omitempty"`
	SellReserves                        *bool                     `json:"sellReserves,omitempty"`
	ReturnFromClosedShiftEnabled        *bool                     `json:"returnFromClosedShiftEnabled,omitempty"`
	ReservePrepaidGoods                 *bool                     `json:"reservePrepaidGoods,omitempty"`
	Address                             *string                   `json:"address,omitempty"`
	SyncAgents                          *bool                     `json:"syncAgents,omitempty"`
	Updated                             *Timestamp                `json:"updated,omitempty"`
	RequiredDiscountCardNumber          *bool                     `json:"requiredDiscountCardNumber,omitempty"`
	AllowCreateProducts                 *bool                     `json:"allowCreateProducts,omitempty"`
	Name                                *string                   `json:"name,omitempty"`
	OFDEnabled                          *bool                     `json:"ofdEnabled,omitempty"`
	State                               *RetailStoreState         `json:"state,omitempty"`
	Shared                              *bool                     `json:"shared,omitempty"`
	OrderToState                        *State                    `json:"orderToState,omitempty"`
	Organization                        *Organization             `json:"organization,omitempty"`
	Owner                               *Employee                 `json:"owner,omitempty"`
	PriceType                           *PriceType                `json:"priceType,omitempty"`
	PrintAlways                         *bool                     `json:"printAlways,omitempty"`
	SendMarksForCheck                   *bool                     `json:"sendMarksForCheck,omitempty"`
	ProductFolders                      *MetaArray[ProductFolder] `json:"productFolders,omitempty"`
	QRAcquire                           *Counterparty             `json:"qrAcquire,omitempty"`
	QRBankPercent                       *float64                  `json:"qrBankPercent,omitempty"`
	QRPayEnabled                        *bool                     `json:"qrPayEnabled,omitempty"`
	QRTerminalId                        *string                   `json:"qrTerminalId,omitempty"`
	ReceiptTemplate                     *Meta                     `json:"receiptTemplate,omitempty"`
	RequiredFio                         *bool                     `json:"requiredFio,omitempty"`
	RequiredPhone                       *bool                     `json:"requiredPhone,omitempty"`
	RequiredEmail                       *bool                     `json:"requiredEmail,omitempty"`
	RequiredBirthdate                   *bool                     `json:"RequiredBirthdate,omitempty"`
	RequiredSex                         *bool                     `json:"requiredSex,omitempty"`
	MinionToMasterType                  MinionToMaster            `json:"minionToMasterType,omitempty"`
	FiscalType                          FiscalType                `json:"fiscalType,omitempty"`
	PriorityOFDSend                     PriorityOfdSend           `json:"priorityOfdSend,omitempty"`
	MarkingSellingMode                  MarkingSellingMode        `json:"markingSellingMode,omitempty"`
	OrderTaxSystem                      TaxSystem                 `json:"orderTaxSystem,omitempty"`
	TobaccoMRCControlType               TobaccoMRCControlType     `json:"tobaccoMrcControlType,omitempty"`
	DefaultTaxSystem                    TaxSystem                 `json:"defaulTaxSystem,omitempty"`
	MarksCheckMode                      MarksCheckMode            `json:"marksCheckMode,omitempty"`
	MasterRetailStores                  Slice[RetailStore]        `json:"masterRetailStores,omitempty"`
	LastOperationNames                  Slice[LastOperation]      `json:"lastOperationNames,omitempty"`
	FilterAgentsTags                    Slice[string]             `json:"filterAgentsTags,omitempty"`
	CustomerOrderStates                 Slice[State]              `json:"customerOrderStates,omitempty"`
	CreateAgentsTags                    Slice[string]             `json:"createAgentsTags,omitempty"`
}

func (retailStore RetailStore) Clean() *RetailStore {
	return &RetailStore{Meta: retailStore.Meta}
}

func (retailStore RetailStore) GetAcquire() Counterparty {
	return Deref(retailStore.Acquire)
}

func (retailStore RetailStore) GetOnlyInStock() bool {
	return Deref(retailStore.OnlyInStock)
}

func (retailStore RetailStore) GetActive() bool {
	return Deref(retailStore.Active)
}

func (retailStore RetailStore) GetAccountID() uuid.UUID {
	return Deref(retailStore.AccountID)
}

func (retailStore RetailStore) GetAddressFull() Address {
	return Deref(retailStore.AddressFull)
}

func (retailStore RetailStore) GetMeta() Meta {
	return Deref(retailStore.Meta)
}

func (retailStore RetailStore) GetAllowCustomPrice() bool {
	return Deref(retailStore.AllowCustomPrice)
}

func (retailStore RetailStore) GetAllowDeleteReceiptPositions() bool {
	return Deref(retailStore.AllowDeleteReceiptPositions)
}

func (retailStore RetailStore) GetAllowSellTobaccoWithoutMRC() bool {
	return Deref(retailStore.AllowSellTobaccoWithoutMRC)
}

func (retailStore RetailStore) GetArchived() bool {
	return Deref(retailStore.Archived)
}

func (retailStore RetailStore) GetAuthTokenAttached() bool {
	return Deref(retailStore.AuthTokenAttached)
}

func (retailStore RetailStore) GetBankPercent() float64 {
	return Deref(retailStore.BankPercent)
}

func (retailStore RetailStore) GetCashiers() MetaArray[Cashier] {
	return Deref(retailStore.Cashiers)
}

func (retailStore RetailStore) GetControlCashierChoice() bool {
	return Deref(retailStore.ControlCashierChoice)
}

func (retailStore RetailStore) GetControlShippingStock() bool {
	return Deref(retailStore.ControlShippingStock)
}

func (retailStore RetailStore) GetCreateCashInOnRetailShiftClosing() bool {
	return Deref(retailStore.CreateCashInOnRetailShiftClosing)
}

func (retailStore RetailStore) GetCreateOrderWithState() State {
	return Deref(retailStore.CreateOrderWithState)
}

func (retailStore RetailStore) GetCreatePaymentInOnRetailShiftClosing() bool {
	return Deref(retailStore.CreatePaymentInOnRetailShiftClosing)
}

func (retailStore RetailStore) GetDemandPrefix() string {
	return Deref(retailStore.DemandPrefix)
}

func (retailStore RetailStore) GetDescription() string {
	return Deref(retailStore.Description)
}

func (retailStore RetailStore) GetDiscountEnable() bool {
	return Deref(retailStore.DiscountEnable)
}

func (retailStore RetailStore) GetDiscountMaxPercent() float64 {
	return Deref(retailStore.DiscountMaxPercent)
}

func (retailStore RetailStore) GetEnableReturnsWithNoReason() bool {
	return Deref(retailStore.EnableReturnsWithNoReason)
}

func (retailStore RetailStore) GetEnvironment() Environment {
	return Deref(retailStore.Environment)
}

func (retailStore RetailStore) GetExternalCode() string {
	return Deref(retailStore.ExternalCode)
}

func (retailStore RetailStore) GetGroup() Group {
	return Deref(retailStore.Group)
}

func (retailStore RetailStore) GetID() uuid.UUID {
	return Deref(retailStore.ID)
}

func (retailStore RetailStore) GetStore() Store {
	return Deref(retailStore.Store)
}

func (retailStore RetailStore) GetIdQR() string {
	return Deref(retailStore.IDQR)
}

func (retailStore RetailStore) GetIssueOrders() bool {
	return Deref(retailStore.IssueOrders)
}

func (retailStore RetailStore) GetShowBeerOnTap() bool {
	return Deref(retailStore.ShowBeerOnTap)
}

func (retailStore RetailStore) GetSellReserves() bool {
	return Deref(retailStore.SellReserves)
}

func (retailStore RetailStore) GetReturnFromClosedShiftEnabled() bool {
	return Deref(retailStore.ReturnFromClosedShiftEnabled)
}

func (retailStore RetailStore) GetReservePrepaidGoods() bool {
	return Deref(retailStore.ReservePrepaidGoods)
}

func (retailStore RetailStore) GetAddress() string {
	return Deref(retailStore.Address)
}

func (retailStore RetailStore) GetSyncAgents() bool {
	return Deref(retailStore.SyncAgents)
}

func (retailStore RetailStore) GetUpdated() Timestamp {
	return Deref(retailStore.Updated)
}

func (retailStore RetailStore) GetRequiredDiscountCardNumber() bool {
	return Deref(retailStore.RequiredDiscountCardNumber)
}

func (retailStore RetailStore) GetAllowCreateProducts() bool {
	return Deref(retailStore.AllowCreateProducts)
}

func (retailStore RetailStore) GetName() string {
	return Deref(retailStore.Name)
}

func (retailStore RetailStore) GetOFDEnabled() bool {
	return Deref(retailStore.OFDEnabled)
}

func (retailStore RetailStore) GetState() RetailStoreState {
	return Deref(retailStore.State)
}

func (retailStore RetailStore) GetShared() bool {
	return Deref(retailStore.Shared)
}

func (retailStore RetailStore) GetOrderToState() State {
	return Deref(retailStore.OrderToState)
}

func (retailStore RetailStore) GetOrganization() Organization {
	return Deref(retailStore.Organization)
}

func (retailStore RetailStore) GetOwner() Employee {
	return Deref(retailStore.Owner)
}

func (retailStore RetailStore) GetPriceType() PriceType {
	return Deref(retailStore.PriceType)
}

func (retailStore RetailStore) GetPrintAlways() bool {
	return Deref(retailStore.PrintAlways)
}

func (retailStore RetailStore) GetSendMarksForCheck() bool {
	return Deref(retailStore.SendMarksForCheck)
}

func (retailStore RetailStore) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(retailStore.ProductFolders)
}

func (retailStore RetailStore) GetQRAcquire() Counterparty {
	return Deref(retailStore.QRAcquire)
}

func (retailStore RetailStore) GetQRBankPercent() float64 {
	return Deref(retailStore.QRBankPercent)
}

func (retailStore RetailStore) GetQRPayEnabled() bool {
	return Deref(retailStore.QRPayEnabled)
}

func (retailStore RetailStore) GetQRTerminalId() string {
	return Deref(retailStore.QRTerminalId)
}

func (retailStore RetailStore) GetReceiptTemplate() Meta {
	return Deref(retailStore.ReceiptTemplate)
}

func (retailStore RetailStore) GetRequiredFio() bool {
	return Deref(retailStore.RequiredFio)
}

func (retailStore RetailStore) GetRequiredPhone() bool {
	return Deref(retailStore.RequiredPhone)
}

func (retailStore RetailStore) GetRequiredEmail() bool {
	return Deref(retailStore.RequiredEmail)
}

func (retailStore RetailStore) GetRequiredBirthdate() bool {
	return Deref(retailStore.RequiredBirthdate)
}

func (retailStore RetailStore) GetRequiredSex() bool {
	return Deref(retailStore.RequiredSex)
}

func (retailStore RetailStore) GetMinionToMasterType() MinionToMaster {
	return retailStore.MinionToMasterType
}

func (retailStore RetailStore) GetFiscalType() FiscalType {
	return retailStore.FiscalType
}

func (retailStore RetailStore) GetPriorityOFDSend() PriorityOfdSend {
	return retailStore.PriorityOFDSend
}

func (retailStore RetailStore) GetMarkingSellingMode() MarkingSellingMode {
	return retailStore.MarkingSellingMode
}

func (retailStore RetailStore) GetOrderTaxSystem() TaxSystem {
	return retailStore.OrderTaxSystem
}

func (retailStore RetailStore) GetTobaccoMRCControlType() TobaccoMRCControlType {
	return retailStore.TobaccoMRCControlType
}

func (retailStore RetailStore) GetDefaultTaxSystem() TaxSystem {
	return retailStore.DefaultTaxSystem
}

func (retailStore RetailStore) GetMarksCheckMode() MarksCheckMode {
	return retailStore.MarksCheckMode
}

func (retailStore RetailStore) GetMasterRetailStores() Slice[RetailStore] {
	return retailStore.MasterRetailStores
}

func (retailStore RetailStore) GetLastOperationNames() Slice[LastOperation] {
	return retailStore.LastOperationNames
}

func (retailStore RetailStore) GetFilterAgentsTags() Slice[string] {
	return retailStore.FilterAgentsTags
}

func (retailStore RetailStore) GetCustomerOrderStates() Slice[State] {
	return retailStore.CustomerOrderStates
}

func (retailStore RetailStore) GetCreateAgentsTags() Slice[string] {
	return retailStore.CreateAgentsTags
}

func (retailStore *RetailStore) SetAcquire(acquire *Counterparty) *RetailStore {
	retailStore.Acquire = acquire
	return retailStore
}

func (retailStore *RetailStore) SetOnlyInStock(onlyInStock bool) *RetailStore {
	retailStore.OnlyInStock = &onlyInStock
	return retailStore
}

func (retailStore *RetailStore) SetActive(active bool) *RetailStore {
	retailStore.Active = &active
	return retailStore
}

func (retailStore *RetailStore) SetAddressFull(addressFull *Address) *RetailStore {
	retailStore.AddressFull = addressFull
	return retailStore
}

func (retailStore *RetailStore) SetMeta(meta *Meta) *RetailStore {
	retailStore.Meta = meta
	return retailStore
}

func (retailStore *RetailStore) SetAllowDeleteReceiptPositions(allowDeleteReceiptPositions bool) *RetailStore {
	retailStore.AllowDeleteReceiptPositions = &allowDeleteReceiptPositions
	return retailStore
}

func (retailStore *RetailStore) SetAllowSellTobaccoWithoutMRC(allowSellTobaccoWithoutMRC bool) *RetailStore {
	retailStore.AllowSellTobaccoWithoutMRC = &allowSellTobaccoWithoutMRC
	return retailStore
}

func (retailStore *RetailStore) SetArchived(archived bool) *RetailStore {
	retailStore.Archived = &archived
	return retailStore
}

func (retailStore *RetailStore) SetBankPercent(bankPercent float64) *RetailStore {
	retailStore.BankPercent = &bankPercent
	return retailStore
}

func (retailStore *RetailStore) SetCashiers(cashiers *MetaArray[Cashier]) *RetailStore {
	retailStore.Cashiers = cashiers
	return retailStore
}

func (retailStore *RetailStore) SetControlCashierChoice(controlCashierChoice bool) *RetailStore {
	retailStore.ControlCashierChoice = &controlCashierChoice
	return retailStore
}

func (retailStore *RetailStore) SetControlShippingStock(controlShippingStock bool) *RetailStore {
	retailStore.ControlShippingStock = &controlShippingStock
	return retailStore
}

func (retailStore *RetailStore) SetCreateCashInOnRetailShiftClosing(createCashInOnRetailShiftClosing bool) *RetailStore {
	retailStore.CreateCashInOnRetailShiftClosing = &createCashInOnRetailShiftClosing
	return retailStore
}

func (retailStore *RetailStore) SetCreateOrderWithState(createOrderWithState *State) *RetailStore {
	retailStore.CreateOrderWithState = createOrderWithState
	return retailStore
}

func (retailStore *RetailStore) SetCreatePaymentInOnRetailShiftClosing(createPaymentInOnRetailShiftClosing bool) *RetailStore {
	retailStore.CreatePaymentInOnRetailShiftClosing = &createPaymentInOnRetailShiftClosing
	return retailStore
}

func (retailStore *RetailStore) SetDemandPrefix(demandPrefix string) *RetailStore {
	retailStore.DemandPrefix = &demandPrefix
	return retailStore
}

func (retailStore *RetailStore) SetDescription(description string) *RetailStore {
	retailStore.Description = &description
	return retailStore
}

func (retailStore *RetailStore) SetDiscountEnable(discountEnable bool) *RetailStore {
	retailStore.DiscountEnable = &discountEnable
	return retailStore
}

func (retailStore *RetailStore) SetDiscountMaxPercent(discountMaxPercent float64) *RetailStore {
	retailStore.DiscountMaxPercent = &discountMaxPercent
	return retailStore
}

func (retailStore *RetailStore) SetEnableReturnsWithNoReason(enableReturnsWithNoReason bool) *RetailStore {
	retailStore.EnableReturnsWithNoReason = &enableReturnsWithNoReason
	return retailStore
}

func (retailStore *RetailStore) SetGroup(group *Group) *RetailStore {
	retailStore.Group = group
	return retailStore
}

func (retailStore *RetailStore) SetStore(store *Store) *RetailStore {
	retailStore.Store = store
	return retailStore
}

func (retailStore *RetailStore) SetIDQR(idqr string) *RetailStore {
	retailStore.IDQR = &idqr
	return retailStore
}

func (retailStore *RetailStore) SetIssueOrders(issueOrders bool) *RetailStore {
	retailStore.IssueOrders = &issueOrders
	return retailStore
}

func (retailStore *RetailStore) SetShowBeerOnTap(showBeerOnTap bool) *RetailStore {
	retailStore.ShowBeerOnTap = &showBeerOnTap
	return retailStore
}

func (retailStore *RetailStore) SetSellReserves(sellReserves bool) *RetailStore {
	retailStore.SellReserves = &sellReserves
	return retailStore
}

func (retailStore *RetailStore) SetReturnFromClosedShiftEnabled(returnFromClosedShiftEnabled bool) *RetailStore {
	retailStore.ReturnFromClosedShiftEnabled = &returnFromClosedShiftEnabled
	return retailStore
}

func (retailStore *RetailStore) SetReservePrepaidGoods(reservePrepaidGoods bool) *RetailStore {
	retailStore.ReservePrepaidGoods = &reservePrepaidGoods
	return retailStore
}

func (retailStore *RetailStore) SetAddress(address string) *RetailStore {
	retailStore.Address = &address
	return retailStore
}

func (retailStore *RetailStore) SetSyncAgents(syncAgents bool) *RetailStore {
	retailStore.SyncAgents = &syncAgents
	return retailStore
}

func (retailStore *RetailStore) SetRequiredDiscountCardNumber(requiredDiscountCardNumber bool) *RetailStore {
	retailStore.RequiredDiscountCardNumber = &requiredDiscountCardNumber
	return retailStore
}

func (retailStore *RetailStore) SetAllowCreateProducts(allowCreateProducts bool) *RetailStore {
	retailStore.AllowCreateProducts = &allowCreateProducts
	return retailStore
}

func (retailStore *RetailStore) SetName(name string) *RetailStore {
	retailStore.Name = &name
	return retailStore
}

func (retailStore *RetailStore) SetShared(shared bool) *RetailStore {
	retailStore.Shared = &shared
	return retailStore
}

func (retailStore *RetailStore) SetOrderToState(orderToState *State) *RetailStore {
	retailStore.OrderToState = orderToState
	return retailStore
}

func (retailStore *RetailStore) SetOrganization(organization *Organization) *RetailStore {
	retailStore.Organization = organization
	return retailStore
}

func (retailStore *RetailStore) SetOwner(owner *Employee) *RetailStore {
	retailStore.Owner = owner
	return retailStore
}

func (retailStore *RetailStore) SetPriceType(priceType *PriceType) *RetailStore {
	retailStore.PriceType = priceType
	return retailStore
}

func (retailStore *RetailStore) SetPrintAlways(printAlways bool) *RetailStore {
	retailStore.PrintAlways = &printAlways
	return retailStore
}

func (retailStore *RetailStore) SetSendMarksForCheck(sendMarksForCheck bool) *RetailStore {
	retailStore.SendMarksForCheck = &sendMarksForCheck
	return retailStore
}

func (retailStore *RetailStore) SetProductFolders(productFolders Slice[ProductFolder]) *RetailStore {
	retailStore.ProductFolders = NewMetaArrayRows(productFolders)
	return retailStore
}

func (retailStore *RetailStore) SetQRAcquire(qrAcquire *Counterparty) *RetailStore {
	retailStore.QRAcquire = qrAcquire
	return retailStore
}

func (retailStore *RetailStore) SetQRBankPercent(qrBankPercent float64) *RetailStore {
	retailStore.QRBankPercent = &qrBankPercent
	return retailStore
}

func (retailStore *RetailStore) SetQRPayEnabled(qrPayEnabled bool) *RetailStore {
	retailStore.QRPayEnabled = &qrPayEnabled
	return retailStore
}

func (retailStore *RetailStore) SetQRTerminalId(qrTerminalId string) *RetailStore {
	retailStore.QRTerminalId = &qrTerminalId
	return retailStore
}

func (retailStore *RetailStore) SetReceiptTemplate(receiptTemplate *Meta) *RetailStore {
	retailStore.ReceiptTemplate = receiptTemplate
	return retailStore
}

func (retailStore *RetailStore) SetRequiredFio(requiredFio bool) *RetailStore {
	retailStore.RequiredFio = &requiredFio
	return retailStore
}

func (retailStore *RetailStore) SetRequiredPhone(requiredPhone bool) *RetailStore {
	retailStore.RequiredPhone = &requiredPhone
	return retailStore
}

func (retailStore *RetailStore) SetRequiredEmail(requiredEmail bool) *RetailStore {
	retailStore.RequiredEmail = &requiredEmail
	return retailStore
}

func (retailStore *RetailStore) SetRequiredBirthdate(requiredBirthdate bool) *RetailStore {
	retailStore.RequiredBirthdate = &requiredBirthdate
	return retailStore
}

func (retailStore *RetailStore) SetRequiredSex(requiredSex bool) *RetailStore {
	retailStore.RequiredSex = &requiredSex
	return retailStore
}

func (retailStore *RetailStore) SetMinionToMasterType(minionToMasterType MinionToMaster) *RetailStore {
	retailStore.MinionToMasterType = minionToMasterType
	return retailStore
}

func (retailStore *RetailStore) SetFiscalType(fiscalType FiscalType) *RetailStore {
	retailStore.FiscalType = fiscalType
	return retailStore
}

func (retailStore *RetailStore) SetPriorityOFDSend(priorityOFDSend PriorityOfdSend) *RetailStore {
	retailStore.PriorityOFDSend = priorityOFDSend
	return retailStore
}

func (retailStore *RetailStore) SetMarkingSellingMode(markingSellingMode MarkingSellingMode) *RetailStore {
	retailStore.MarkingSellingMode = markingSellingMode
	return retailStore
}

func (retailStore *RetailStore) SetOrderTaxSystem(orderTaxSystem TaxSystem) *RetailStore {
	retailStore.OrderTaxSystem = orderTaxSystem
	return retailStore
}

func (retailStore *RetailStore) SetTobaccoMRCControlType(tobaccoMRCControlType TobaccoMRCControlType) *RetailStore {
	retailStore.TobaccoMRCControlType = tobaccoMRCControlType
	return retailStore
}

func (retailStore *RetailStore) SetDefaultTaxSystem(defaultTaxSystem TaxSystem) *RetailStore {
	retailStore.DefaultTaxSystem = defaultTaxSystem
	return retailStore
}

func (retailStore *RetailStore) SetMarksCheckMode(marksCheckMode MarksCheckMode) *RetailStore {
	retailStore.MarksCheckMode = marksCheckMode
	return retailStore
}

func (retailStore *RetailStore) SetMasterRetailStores(masterRetailStores Slice[RetailStore]) *RetailStore {
	retailStore.MasterRetailStores = masterRetailStores
	return retailStore
}

func (retailStore *RetailStore) SetFilterAgentsTags(filterAgentsTags Slice[string]) *RetailStore {
	retailStore.FilterAgentsTags = filterAgentsTags
	return retailStore
}

func (retailStore *RetailStore) SetCustomerOrderStates(customerOrderStates Slice[State]) *RetailStore {
	retailStore.CustomerOrderStates = customerOrderStates
	return retailStore
}

func (retailStore *RetailStore) SetCreateAgentsTags(createAgentsTags Slice[string]) *RetailStore {
	retailStore.CreateAgentsTags = createAgentsTags
	return retailStore
}

func (retailStore RetailStore) String() string {
	return Stringify(retailStore)
}

func (retailStore RetailStore) MetaType() MetaType {
	return MetaTypeRetailStore
}

// RetailStoreState Информация статусе точки продаж
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status
type RetailStoreState struct {
	Sync            RetailStoreStateSync `json:"sync,omitempty"`         // Состояние синхронизации
	LastCheckMoment Timestamp            `json:"lastCheckMoment,"`       // Дата и время последней синхронизации
	FiscalMemory    FiscalMemoryState    `json:"fiscalMemory,omitempty"` // Информация о фискальном накопителе
}

func (retailStoreState RetailStoreState) String() string {
	return Stringify(retailStoreState)
}

// RetailStoreStateSync Состояние синхронизации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-sinhronizaciq
type RetailStoreStateSync struct {
	LastAttemptMoment Timestamp `json:"lastAttempMoment,omitempty"`
	Message           string    `json:"message,omitempty"`
}

func (retailStoreStateSync RetailStoreStateSync) String() string {
	return Stringify(retailStoreStateSync)
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

func (environment Environment) String() string {
	return Stringify(environment)
}

// Software Аттрибуты сущности ПО
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-po
type Software struct {
	Name    string `json:"name,omitempty"`    // Наименование ПО
	Vendor  string `json:"vendor,omitempty"`  // Производитель
	Version string `json:"version,omitempty"` // Версия ПО
}

func (software Software) String() string {
	return Stringify(software)
}

// PaymentTerminal Информация о платежном терминале.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-platezhnyj-terminal
type PaymentTerminal struct {
	AcquiringType string `json:"acquiringType,omitempty"` // Информация о типе эквайера (например: inpas/payme)
}

func (paymentTerminal PaymentTerminal) String() string {
	return Stringify(paymentTerminal)
}

// LastOperation Последняя операция.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-poslednie-operacii
type LastOperation struct {
	Entity MetaType `json:"entity,omitempty"` // Ключевое слово, обозначающее тип последней операции
	Name   string   `json:"name,omitempty"`   // Наименование (номер) последней операции
}

func (lastOperation LastOperation) String() string {
	return Stringify(lastOperation)
}

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

func (chequePrinter ChequePrinter) String() string {
	return Stringify(chequePrinter)
}

// Driver Атрибуты сущности Драйвер.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-drajwer
type Driver struct {
	Name    string `json:"name,omitempty"`    // Наименование драйвера
	Version string `json:"version,omitempty"` // Версия драйвера
}

func (driver Driver) String() string {
	return Stringify(driver)
}

// FiscalMemory Атрибуты сущности Фискальный накопитель.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-okruzhenie-attributy-suschnosti-fiskal-nyj-nakopitel
type FiscalMemory struct {
	FiscalValidityDate Timestamp `json:"fiscalValidityDate,omitempty"`
	FiscalDataVersion  string    `json:"fiscalDataVersion,omitempty"`
}

func (fiscalMemory FiscalMemory) String() string {
	return Stringify(fiscalMemory)
}

// FiscalMemoryState Информация о фискальном накопителе.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-fiskal-naq-pamqt
type FiscalMemoryState struct {
	NotSendFirstDocMoment Timestamp              `json:"notSendFirstDocMoment,omitempty"` // Дата последнего неотправленного документа (?)
	Error                 FiscalMemoryStateError `json:"error,omitempty"`                 // Информация об ошибке ФН
	NotSendDocCount       int                    `json:"notSendDocCount,omitempty"`       // Количество неотправленных документов в ОФД
}

func (fiscalMemoryState FiscalMemoryState) String() string {
	return Stringify(fiscalMemoryState)
}

// FiscalMemoryStateError Информация об ошибке ФН.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-attributy-suschnosti-status-attributy-suschnosti-oshibka
type FiscalMemoryStateError struct {
	Code    string `json:"code,omitempty"`    // Код ошибки ФН
	Message string `json:"message,omitempty"` // Описание ошибки
}

func (fiscalMemoryStateError FiscalMemoryStateError) String() string {
	return Stringify(fiscalMemoryStateError)
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
	DeleteMany(ctx context.Context, retailStore *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*RetailStore, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, entity *RetailStore, params *Params) (*RetailStore, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetCashiers(ctx context.Context, id uuid.UUID) (*MetaArray[Cashier], *resty.Response, error)
	GetCashierByID(ctx context.Context, id, cashierID uuid.UUID) (*Cashier, *resty.Response, error)
}

type retailStoreService struct {
	Endpoint
	endpointGetList[RetailStore]
	endpointCreate[RetailStore]
	endpointCreateUpdateMany[RetailStore]
	endpointDeleteMany[RetailStore]
	endpointDelete
	endpointGetByID[RetailStore]
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
		endpointGetByID:          endpointGetByID[RetailStore]{e},
		endpointUpdate:           endpointUpdate[RetailStore]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// GetCashiers Получить Кассиров.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassirow
func (s *retailStoreService) GetCashiers(ctx context.Context, id uuid.UUID) (*MetaArray[Cashier], *resty.Response, error) {
	path := fmt.Sprintf("entity/retailstore/%s/cashiers", id)
	return NewRequestBuilder[MetaArray[Cashier]](s.client, path).Get(ctx)
}

// GetCashierByID Получить Кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassira
func (s *retailStoreService) GetCashierByID(ctx context.Context, id, cashierID uuid.UUID) (*Cashier, *resty.Response, error) {
	path := fmt.Sprintf("entity/retailstore/%s/cashiers/%s", id, cashierID)
	return NewRequestBuilder[Cashier](s.client, path).Get(ctx)
}
