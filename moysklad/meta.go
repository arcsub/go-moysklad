package moysklad

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

type HasMeta interface {
	MetaTyper
	MetaOwner
}

type MetaTyper interface {
	MetaType() MetaType
}

type MetaOwner interface {
	GetMeta() Meta
}

// Meta Метаданные объекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-ob-ekta
type Meta struct {
	Href         *string   `json:"href,omitempty"`
	MetadataHref *string   `json:"metadataHref,omitempty"`
	MediaType    *string   `json:"mediaType,omitempty"`
	UUIDHref     *string   `json:"uuidHref,omitempty"`
	DownloadHref *string   `json:"downloadHref,omitempty"`
	Type         *MetaType `json:"type,omitempty"`
}

func (meta Meta) GetHref() string {
	return Deref(meta.Href)
}

func (meta Meta) GetMetadataHref() string {
	return Deref(meta.MetadataHref)
}

func (meta Meta) GetMediaType() string {
	return Deref(meta.MediaType)
}

func (meta Meta) GetUUIDHref() string {
	return Deref(meta.UUIDHref)
}

func (meta Meta) GetDownloadHref() string {
	return Deref(meta.DownloadHref)
}

func (meta Meta) GetType() MetaType {
	return Deref(meta.Type)
}

func (meta *Meta) SetHref(href string) *Meta {
	meta.Href = &href
	return meta
}

func (meta *Meta) SetType(metaType MetaType) *Meta {
	meta.Type = &metaType
	return meta
}

// Wrap оборачивает текущий объект Meta в MetaWrapper
func (meta Meta) Wrap() MetaWrapper {
	return MetaWrapper{meta}
}

func (meta Meta) String() string {
	return Stringify(meta)
}

func (meta *Meta) IsEqual(other *Meta) bool {
	return IsEqualPtr(meta.Href, other.Href)
}

// GetIDFromHref возвращает UUID из поля Href
// Возвращает nil, если поле Href пусто или не содержит id
func (meta Meta) GetIDFromHref() *uuid.UUID {
	href := Deref(meta.Href)
	if href == "" {
		return nil
	}

	sep := strings.Split(href, "/")
	if len(sep) == 0 {
		return nil
	}

	if id, err := uuid.Parse(sep[len(sep)-1]); err == nil {
		return &id
	}
	return nil
}

// MetaWrapper объект-обёртка для Meta
type MetaWrapper struct {
	Meta Meta `json:"meta"`
}

func (metaWrapper MetaWrapper) String() string {
	return Stringify(metaWrapper)
}

// MetaNameWrapper объект-обёртка для Meta и поля Name
type MetaNameWrapper struct {
	Meta Meta   `json:"meta"`
	Name string `json:"name"`
}

func (metaNameWrapper MetaNameWrapper) String() string {
	return Stringify(metaNameWrapper)
}

// MetaCollection Метаданные коллекции.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-kollekcii
type MetaCollection struct {
	Href         string `json:"href,omitempty"`
	Type         string `json:"type,omitempty"`
	MediaType    string `json:"mediaType,omitempty"`
	NextHref     string `json:"nextHref,omitempty"`
	PreviousHref string `json:"previousHref,omitempty"`
	Size         int    `json:"size,omitempty"`
	Limit        int    `json:"Limit,omitempty"`
	Offset       int    `json:"Offset,omitempty"`
}

func (metaCollection MetaCollection) String() string {
	return Stringify(metaCollection)
}

// MetaArray Объект с полями meta и rows, где rows - массив объектов
type MetaArray[T any] struct {
	Rows Slice[T]       `json:"rows,omitempty"`
	Meta MetaCollection `json:"meta,omitempty"`
}

func NewMetaArrayFrom[T any](rows []*T) *MetaArray[T] {
	return &MetaArray[T]{Rows: rows}
}

// Len возвращает количество элементов Rows
func (metaArray MetaArray[T]) Len() int {
	return len(metaArray.Rows)
}

// Size возвращает размер выданного списка
func (metaArray MetaArray[T]) Size() int {
	return metaArray.Meta.Size
}

// NextHref возвращает ссылку на следующую страницу сущностей.
func (metaArray MetaArray[T]) NextHref() string {
	return metaArray.Meta.NextHref
}

// PreviousHref возвращает ссылку на предыдущую страницу сущностей.
func (metaArray MetaArray[T]) PreviousHref() string {
	return metaArray.Meta.PreviousHref
}

// Push добавляет элементы в срез.
func (metaArray *MetaArray[T]) Push(elements ...*T) *MetaArray[T] {
	metaArray.Rows.Push(elements...)
	return metaArray
}

func (metaArray MetaArray[T]) String() string {
	return Stringify(metaArray)
}

// MarshalJSON реализует интерфейс json.Marshaler
func (metaArray MetaArray[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(metaArray.Rows)
}

type MetaType string

func (metaType MetaType) String() string {
	return string(metaType)
}

const (
	MetaTypeAccount                             MetaType = "account"
	MetaTypeAccumulationDiscount                MetaType = "accumulationdiscount"
	MetaTypeApplication                         MetaType = "application"
	MetaTypeAssortment                          MetaType = "assortment"
	MetaTypeAssortmentSettings                  MetaType = "assortmentsettings"
	MetaTypeAsync                               MetaType = "async"
	MetaTypeAttribute                           MetaType = "attributemetadata"
	MetaTypeAudit                               MetaType = "audit"
	MetaTypeAuditEvent                          MetaType = "auditevent"
	MetaTypeBonusProgram                        MetaType = "bonusprogram"
	MetaTypeBonusTransaction                    MetaType = "bonustransaction"
	MetaTypeBundle                              MetaType = "bundle"
	MetaTypeBundleComponent                     MetaType = "bundlecomponent"
	MetaTypeCashier                             MetaType = "cashier"
	MetaTypeCashIn                              MetaType = "cashin"
	MetaTypeCashOut                             MetaType = "cashout"
	MetaTypeCharacteristic                      MetaType = "attributemetadata"
	MetaTypeCommissionReportIn                  MetaType = "commissionreportin"
	MetaTypeCommissionReportInPosition          MetaType = "commissionreportinposition"
	MetaTypeCommissionReportInReturnPosition    MetaType = "commissionreportinreturnedposition"
	MetaTypeCommissionReportOut                 MetaType = "commissionreportout"
	MetaTypeCommissionReportOutPosition         MetaType = "commissionreportoutposition"
	MetaTypeCompanySettings                     MetaType = "companysettings"
	MetaTypeConsignment                         MetaType = "consignment"
	MetaTypeContactPerson                       MetaType = "contactperson"
	MetaTypeContract                            MetaType = "contract"
	MetaTypeCounterparty                        MetaType = "counterparty"
	MetaTypeCounterPartyAdjustment              MetaType = "counterpartyadjustment"
	MetaTypeCountry                             MetaType = "country"
	MetaTypeCurrency                            MetaType = "currency"
	MetaTypeCustomerOrder                       MetaType = "customerorder"
	MetaTypeCustomerOrderPosition               MetaType = "customerorderposition"
	MetaTypeCustomEntity                        MetaType = "customentity"
	MetaTypeCustomTemplate                      MetaType = "customtemplate"
	MetaTypeDemand                              MetaType = "demand"
	MetaTypeDemandPosition                      MetaType = "demandposition"
	MetaTypeDiscount                            MetaType = "discount"
	MetaTypeEmbeddedTemplate                    MetaType = "embeddedtemplate"
	MetaTypeEmployee                            MetaType = "employee"
	MetaTypeEmployeeContext                     MetaType = "employee"
	MetaTypeEnter                               MetaType = "enter"
	MetaTypeEnterPosition                       MetaType = "enterposition"
	MetaTypeExpenseItem                         MetaType = "expenseitem"
	MetaTypeFactureIn                           MetaType = "facturein"
	MetaTypeFactureOut                          MetaType = "factureout"
	MetaTypeFiles                               MetaType = "files"
	MetaTypeGroup                               MetaType = "group"
	MetaTypeImage                               MetaType = "image"
	MetaTypeInternalOrder                       MetaType = "internalorder"
	MetaTypeInternalOrderPosition               MetaType = "internalorderposition"
	MetaTypeInventory                           MetaType = "inventory"
	MetaTypeInventoryPosition                   MetaType = "inventoryposition"
	MetaTypeInvoiceIn                           MetaType = "invoicein"
	MetaTypeInvoiceOut                          MetaType = "invoiceout"
	MetaTypeInvoicePosition                     MetaType = "invoiceposition"
	MetaTypeLoss                                MetaType = "loss"
	MetaTypeLossPosition                        MetaType = "lossposition"
	MetaTypeNamedFilter                         MetaType = "namedfilter"
	MetaTypeMove                                MetaType = "move"
	MetaTypeMovePosition                        MetaType = "moveposition"
	MetaTypeNote                                MetaType = "note"
	MetaTypeNotification                        MetaType = "notification"
	MetaTypeNotificationExportCompleted         MetaType = "NotificationExportCompleted"
	MetaTypeNotificationImportCompleted         MetaType = "NotificationImportCompleted"
	MetaTypeNotificationGoodCountTooLow         MetaType = "NotificationGoodCountTooLow"
	MetaTypeNotificationInvoiceOutOverdue       MetaType = "NotificationInvoiceOutOverdue"
	MetaTypeNotificationOrderNew                MetaType = "NotificationOrderNew"
	MetaTypeNotificationOrderOverdue            MetaType = "NotificationOrderOverdue"
	MetaTypeNotificationSubscribeExpired        MetaType = "NotificationSubscribeExpired"
	MetaTypeNotificationSubscribeTermsExpired   MetaType = "NotificationSubscribeTermsExpired"
	MetaTypeNotificationTaskAssigned            MetaType = "NotificationTaskAssigned"
	MetaTypeNotificationTaskUnassigned          MetaType = "NotificationTaskUnassigned"
	MetaTypeNotificationTaskChanged             MetaType = "NotificationTaskChanged"
	MetaTypeNotificationTaskCompleted           MetaType = "NotificationTaskCompleted"
	MetaTypeNotificationTaskDeleted             MetaType = "NotificationTaskDeleted"
	MetaTypeNotificationTaskOverdue             MetaType = "NotificationTaskOverdue"
	MetaTypeNotificationTaskReopened            MetaType = "NotificationTaskReopened"
	MetaTypeNotificationTaskNewComment          MetaType = "NotificationTaskNewComment"
	MetaTypeNotificationTaskCommentChanged      MetaType = "NotificationTaskCommentChanged"
	MetaTypeNotificationTaskCommentDeleted      MetaType = "NotificationTaskCommentDeleted"
	MetaTypeNotificationRetailShiftOpened       MetaType = "NotificationRetailShiftOpened"
	MetaTypeNotificationRetailShiftClosed       MetaType = "NotificationRetailShiftClosed"
	MetaTypeNotificationScript                  MetaType = "NotificationScript"
	MetaTypeFacebookTokenExpirationNotification MetaType = "FacebookTokenExpirationNotification"
	MetaTypeNotificationBonusMoney              MetaType = "NotificationBonusMoney"
	MetaTypeNewMentionInEvent                   MetaType = "NewMentionInEvent"
	MetaTypePublication                         MetaType = "operationpublication"
	MetaTypeOrganization                        MetaType = "organization"
	MetaTypePaymentIn                           MetaType = "paymentin"
	MetaTypePaymentOut                          MetaType = "paymentout"
	MetaTypePersonalDiscount                    MetaType = "personaldiscount"
	MetaTypePrepayment                          MetaType = "prepayment"
	MetaTypePrepaymentPosition                  MetaType = "prepaymentposition"
	MetaTypePrepaymentReturn                    MetaType = "prepaymentreturn"
	MetaTypePrepaymentReturnPosition            MetaType = "prepaymentreturnposition"
	MetaTypePriceList                           MetaType = "pricelist"
	MetaTypePriceListPosition                   MetaType = "pricelistrow"
	MetaTypePriceType                           MetaType = "pricetype"
	MetaTypeProcessing                          MetaType = "processing"
	MetaTypeProcessingOrder                     MetaType = "processingorder"
	MetaTypeProcessingOrderPosition             MetaType = "processingorderposition"
	MetaTypeProcessingPlan                      MetaType = "processingplan"
	MetaTypeProcessingPlanMaterial              MetaType = "processingplanmaterial"
	MetaTypeProcessingPlanProduct               MetaType = "processingplanresult"
	MetaTypeProcessingPositionMaterial          MetaType = "processingpositionmaterial"
	MetaTypeProcessingPositionProduct           MetaType = "processingpositionresult"
	MetaTypeProcessingProcess                   MetaType = "processingprocess"
	MetaTypeProcessingProcessPosition           MetaType = "processingprocessposition"
	MetaTypeProcessingStage                     MetaType = "processingstage"
	MetaTypeProduct                             MetaType = "product"
	MetaTypeProductFolder                       MetaType = "productfolder"
	MetaTypeProject                             MetaType = "project"
	MetaTypePurchaseOrder                       MetaType = "purchaseorder"
	MetaTypePurchaseOrderPosition               MetaType = "purchaseorderposition"
	MetaTypePurchaseReturn                      MetaType = "purchasereturn"
	MetaTypePurchaseReturnPosition              MetaType = "purchasereturnposition"
	MetaTypeReceiptTemplate                     MetaType = "receipttemplate"
	MetaTypeRegion                              MetaType = "region"
	MetaTypeRetailDemand                        MetaType = "demand"
	MetaTypeRetailDemandPosition                MetaType = "demandposition"
	MetaTypeRetailDrawerCashIn                  MetaType = "retaildrawercashin"
	MetaTypeRetailDrawerCashOut                 MetaType = "retaildrawercashout"
	MetaTypeRetailSalesReturn                   MetaType = "retailsalesreturn"
	MetaTypeRetailSalesReturnPosition           MetaType = "salesreturnposition"
	MetaTypeRetailShift                         MetaType = "retailshift"
	MetaTypeRetailStore                         MetaType = "retailstore"
	MetaTypeSalesReturn                         MetaType = "salesreturn"
	MetaTypeSalesReturnPosition                 MetaType = "salesreturnposition"
	MetaTypeService                             MetaType = "service"
	MetaTypeSlot                                MetaType = "slot"
	MetaTypeSpecialPriceDiscount                MetaType = "specialpricediscount"
	MetaTypeState                               MetaType = "state"
	MetaTypeStore                               MetaType = "store"
	MetaTypeStoreZone                           MetaType = "storezone"
	MetaTypeSupply                              MetaType = "supply"
	MetaTypeSupplyPosition                      MetaType = "supplyposition"
	MetaTypeTask                                MetaType = "task"
	MetaTypeTaskNote                            MetaType = "tasknote"
	MetaTypeUom                                 MetaType = "uom"
	MetaTypeVariant                             MetaType = "variant"
	MetaTypeWebhook                             MetaType = "webhook"
	MetaTypeCounterpartySettings                MetaType = "counterpartysettings"
	MetaTypeRole                                MetaType = "role"
	MetaTypeSystemRole                          MetaType = "systemrole"
	MetaTypeIndividualRole                      MetaType = "individualrole"
	MetaTypeCustomRole                          MetaType = "customrole"
	MetaTypeUserSettings                        MetaType = "usersettings"
	MetaTypeSubscription                        MetaType = "subscription"
	MetaTypeSalesChannel                        MetaType = "saleschannel"
	MetaTypeMetadata                            MetaType = "metadata"
	MetaTypeTaxRate                             MetaType = "taxrate"
	MetaTypeThing                               MetaType = "thing"
	MetaTypeToken                               MetaType = "token"
	MetaTypeReportStock                         MetaType = "stock"
	MetaTypeReportStockByOperation              MetaType = "stockbyoperation"
	MetaTypeReportStockByStore                  MetaType = "stockbystore"
	MetaTypeReportMoney                         MetaType = "money"
	MetaTypeReportMoneyPlotSeries               MetaType = "moneyplotseries"
	MetaTypeReportProfitByCounterparty          MetaType = "salesbyCounterparty"
	MetaTypeReportProfitByEmployee              MetaType = "salesbyemployee"
	MetaTypeReportProfitByProduct               MetaType = "salesbyproduct"
	MetaTypeReportProfitBySalesChannel          MetaType = "salesbysaleschannel"
	MetaTypeReportProfitByVariant               MetaType = "salesbyvariant"
	MetaTypeReportOrders                        MetaType = "ordersplotseries"
	MetaTypeReportSales                         MetaType = "salesplotseries"
	MetaTypeReportTurnover                      MetaType = "turnover"
	MetaTypeReportDashboard                     MetaType = "dashboard"
	MetaTypeReportCounterparty                  MetaType = "counterparty"
	MetaTypeWebhookStock                        MetaType = "webhookstock"
	MetaTypeProcessingPlanFolder                MetaType = "processingplanfolder"
	MetaTypeProductionTask                      MetaType = "productiontask"
	MetaTypeProductionTaskMaterial              MetaType = "productiontaskmaterial"
	MetaTypeProductionRow                       MetaType = "productionrow"
	MetaTypeProductionTaskResult                MetaType = "productiontaskresult"
	MetaTypeProductionStage                     MetaType = "productionstage"
	MetaTypeProductionStageCompletion           MetaType = "productionstagecompletion"
	MetaTypeProductionStageCompletionMaterial   MetaType = "productionstagecompletionmaterial"
	MetaTypeProductionStageCompletionResult     MetaType = "productionstagecompletionresult"
	MetaTypeProcessingPlanStages                MetaType = "processingplanstages"
	MetaTypePayroll                             MetaType = "payroll"
	MetaTypeUnknown                             MetaType = ""
)

func MetaTypeFromEntity(v any) (MetaType, error) {
	var metaType MetaType
	var err error

	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		v = val.Elem().Interface()
		val = reflect.ValueOf(v)
	}

	switch v.(type) {
	case ContextEmployee:
		metaType = MetaTypeEmployeeContext
	case AccumulationDiscount:
		metaType = MetaTypeAccumulationDiscount
	case Assortment:
		metaType = MetaTypeAssortment
	case Application:
		metaType = MetaTypeApplication
	case BonusProgram:
		metaType = MetaTypeBonusProgram
	case BonusTransaction:
		metaType = MetaTypeBonusTransaction
	case Bundle:
		metaType = MetaTypeBundle
	case BundleComponent:
		metaType = MetaTypeBundleComponent
	case Cashier:
		metaType = MetaTypeCashier
	case CashIn:
		metaType = MetaTypeCashIn
	case CashOut:
		metaType = MetaTypeCashOut
	case CommissionReportIn:
		metaType = MetaTypeCommissionReportIn
	case CommissionReportInPosition:
		metaType = MetaTypeCommissionReportInPosition
	case CommissionReportOut:
		metaType = MetaTypeCommissionReportOut
	case CommissionReportOutPosition:
		metaType = MetaTypeCommissionReportOutPosition
	case CompanySettings:
		metaType = MetaTypeCompanySettings
	case Consignment:
		metaType = MetaTypeConsignment
	case ContactPerson:
		metaType = MetaTypeContactPerson
	case Contract:
		metaType = MetaTypeContract
	case Counterparty:
		metaType = MetaTypeCounterparty
	case Country:
		metaType = MetaTypeCountry
	case Currency:
		metaType = MetaTypeCurrency
	case CustomerOrder:
		metaType = MetaTypeCustomerOrder
	case CustomerOrderPosition:
		metaType = MetaTypeCustomerOrderPosition
	case CustomEntity:
	case CustomEntityElement:
		metaType = MetaTypeCustomEntity
	case Demand:
		metaType = MetaTypeDemand
	case DemandPosition:
		metaType = MetaTypeDemandPosition
	case Discount:
		metaType = MetaTypeDiscount
	case Employee:
		metaType = MetaTypeEmployee
	case Enter:
		metaType = MetaTypeEnter
	case EnterPosition:
		metaType = MetaTypeEnterPosition
	case ExpenseItem:
		metaType = MetaTypeExpenseItem
	case FactureIn:
		metaType = MetaTypeFactureIn
	case FactureOut:
		metaType = MetaTypeFactureOut
	case Group:
		metaType = MetaTypeGroup
	case Image:
		metaType = MetaTypeImage
	case InternalOrder:
		metaType = MetaTypeInternalOrder
	case InternalOrderPosition:
		metaType = MetaTypeInternalOrderPosition
	case Inventory:
		metaType = MetaTypeInventory
	case InventoryPosition:
		metaType = MetaTypeInventoryPosition
	case InvoiceIn:
		metaType = MetaTypeInvoiceIn
	case InvoiceOut:
		metaType = MetaTypeInvoiceOut
	case InvoiceInPosition:
		metaType = MetaTypeInvoicePosition
	case InvoiceOutPosition:
		metaType = MetaTypeInvoicePosition
	case Loss:
		metaType = MetaTypeLoss
	case LossPosition:
		metaType = MetaTypeLossPosition
	case Move:
		metaType = MetaTypeMove
	case MovePosition:
		metaType = MetaTypeMovePosition
	case Note:
		metaType = MetaTypeNote
	case Organization:
		metaType = MetaTypeOrganization
	case PaymentIn:
		metaType = MetaTypePaymentIn
	case PaymentOut:
		metaType = MetaTypePaymentOut
	case PersonalDiscount:
		metaType = MetaTypePersonalDiscount
	case Prepayment:
		metaType = MetaTypePrepayment
	case PrepaymentPosition:
		metaType = MetaTypePrepaymentPosition
	case PrepaymentReturn:
		metaType = MetaTypePrepaymentReturn
	case PrepaymentReturnPosition:
		metaType = MetaTypePrepaymentReturnPosition
	case PriceList:
		metaType = MetaTypePriceList
	case PriceListPosition:
		metaType = MetaTypePriceListPosition
	case Processing:
		metaType = MetaTypeProcessing
	case ProcessingOrder:
		metaType = MetaTypeProcessingOrder
	case ProcessingOrderPosition:
		metaType = MetaTypeProcessingOrderPosition
	case ProcessingPlan:
		metaType = MetaTypeProcessingPlan
	case ProcessingPlanMaterial:
		metaType = MetaTypeProcessingPlanMaterial
	case ProcessingPlanProduct:
		metaType = MetaTypeProcessingPlanProduct
	case ProcessingPositionMaterial:
		metaType = MetaTypeProcessingPositionMaterial
	case ProcessingPositionProduct:
		metaType = MetaTypeProcessingPositionProduct
	case ProcessingProcess:
		metaType = MetaTypeProcessingProcess
	case ProcessingStage:
		metaType = MetaTypeProcessingStage
	case Product:
		metaType = MetaTypeProduct
	case ProductFolder:
		metaType = MetaTypeProductFolder
	case Project:
		metaType = MetaTypeProject
	case PurchaseOrder:
		metaType = MetaTypePurchaseOrder
	case PurchaseOrderPosition:
		metaType = MetaTypePurchaseOrderPosition
	case PurchaseReturn:
		metaType = MetaTypePurchaseReturn
	case PurchaseReturnPosition:
		metaType = MetaTypePurchaseReturnPosition
	case ReceiptTemplate:
		metaType = MetaTypeReceiptTemplate
	case Region:
		metaType = MetaTypeRegion
	case RetailDemand:
		metaType = MetaTypeRetailDemand
	case RetailDrawerCashIn:
		metaType = MetaTypeRetailDrawerCashIn
	case RetailDrawerCashOut:
		metaType = MetaTypeRetailDrawerCashOut
	case RetailSalesReturn:
		metaType = MetaTypeRetailSalesReturn
	case RetailShift:
		metaType = MetaTypeRetailShift
	case RetailStore:
		metaType = MetaTypeRetailStore
	case SalesReturn:
		metaType = MetaTypeSalesReturn
	case SalesReturnPosition:
		metaType = MetaTypeSalesReturnPosition
	case Service:
		metaType = MetaTypeService
	case SpecialPriceDiscount:
		metaType = MetaTypeSpecialPriceDiscount
	case State:
		metaType = MetaTypeState
	case Store:
		metaType = MetaTypeStore
	case Supply:
		metaType = MetaTypeSupply
	case SupplyPosition:
		metaType = MetaTypeSupplyPosition
	case Task:
		metaType = MetaTypeTask
	case TaskNote:
		metaType = MetaTypeTaskNote
	case Uom:
		metaType = MetaTypeUom
	case Variant:
		metaType = MetaTypeVariant
	case Webhook:
		metaType = MetaTypeWebhook
	case CounterpartySettings:
		metaType = MetaTypeCounterpartySettings
	case UserSettings:
		metaType = MetaTypeUserSettings
	case Subscription:
		metaType = MetaTypeSubscription
	case Role:
		metaType = MetaTypeRole
	case PriceType:
		metaType = MetaTypePriceType
	case AssortmentSettings:
		metaType = MetaTypeAssortmentSettings
	case SalesChannel:
		metaType = MetaTypeSalesChannel
	case CounterpartyAdjustment:
		metaType = MetaTypeCounterPartyAdjustment
	case Metadata:
		metaType = MetaTypeMetadata
	case Token:
		metaType = MetaTypeToken
	case TaxRate:
		metaType = MetaTypeTaxRate
	case Thing:
		metaType = MetaTypeThing
	case WebhookStock:
		metaType = MetaTypeWebhookStock
	case CustomTemplate:
		metaType = MetaTypeCustomTemplate
	case EmbeddedTemplate:
		metaType = MetaTypeEmbeddedTemplate
	case ProcessingPlanFolder:
		metaType = MetaTypeProcessingPlanFolder
	case Publication:
		metaType = MetaTypePublication
	case Attribute:
		metaType = MetaTypeAttribute
	case Async:
		metaType = MetaTypeAsync

	case ProductionTask:
		metaType = MetaTypeProductionTask
	case ProductionRow:
		metaType = MetaTypeProductionRow
	case ProductionTaskResult:
		metaType = MetaTypeProductionTaskResult
	case ProductionStage:
		metaType = MetaTypeProductionStage

	case ProductionStageCompletion:
		metaType = MetaTypeProductionStageCompletion
	case ProductionStageCompletionMaterial:
		metaType = MetaTypeProductionStageCompletionMaterial
	case ProductionStageCompletionResult:
		metaType = MetaTypeProductionStageCompletionResult

	default:
		err = fmt.Errorf("unrecognized entity: %v", metaType)
	}
	return metaType, err
}
