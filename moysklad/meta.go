package moysklad

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// MetaTyper описывает метод, возвращающий код сущности.
type MetaTyper interface {
	MetaType() MetaType
}

// MetaOwner описывает метод, возвращающий [Meta].
type MetaOwner interface {
	GetMeta() Meta
}

// Meta Метаданные объекта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-ob-ekta
type Meta struct {
	Href         *string   `json:"href,omitempty"`         // ссылка на объект
	MetadataHref *string   `json:"metadataHref,omitempty"` // ссылка на метаданные сущности
	MediaType    *string   `json:"mediaType,omitempty"`    // тип данных, который приходят в ответ от сервиса, либо отправляется в теле запроса
	UUIDHref     *string   `json:"uuidHref,omitempty"`     // ссылка на объект в веб-версии МоегоСклада. Присутствует не во всех сущностях
	DownloadHref *string   `json:"downloadHref,omitempty"` // ссылка на скачивание
	Type         *MetaType `json:"type,omitempty"`         // тип объекта (код сущности)
}

func newMeta(metaType MetaType, id string) *Meta {
	href := fmt.Sprintf("%s%s%s/%s", baseApiURL, EndpointEntity, metaType, id)

	meta := new(Meta).SetHref(href).SetMediaType(ApplicationJson).SetType(metaType)

	return meta
}

func NewMetaProduct(id string) *Meta {
	return newMeta(MetaTypeProduct, id)
}

// GetHref возвращает ссылку на объект.
func (meta Meta) GetHref() string {
	return Deref(meta.Href)
}

// GetMetadataHref возвращает ссылку на метаданные сущности.
func (meta Meta) GetMetadataHref() string {
	return Deref(meta.MetadataHref)
}

// GetMediaType возвращает тип данных, который приходят в ответ от сервиса, либо отправляется в теле запроса.
func (meta Meta) GetMediaType() string {
	return Deref(meta.MediaType)
}

// GetUUIDHref возвращает ссылка на объект в веб-версии МоегоСклада. Присутствует не во всех сущностях.
func (meta Meta) GetUUIDHref() string {
	return Deref(meta.UUIDHref)
}

// GetDownloadHref возвращает ссылку на скачивание.
func (meta Meta) GetDownloadHref() string {
	return Deref(meta.DownloadHref)
}

// GetType возвращает тип объекта (код сущности).
func (meta Meta) GetType() MetaType {
	return Deref(meta.Type)
}

// SetHref устанавливает ссылку на объект.
func (meta *Meta) SetHref(href string) *Meta {
	meta.Href = &href
	return meta
}

// SetType устанавливает тип объекта (код сущности).
func (meta *Meta) SetType(metaType MetaType) *Meta {
	meta.Type = &metaType
	return meta
}

// SetMediaType устанавливает тип данных, который приходят в ответ от сервиса, либо отправляется в теле запроса.
func (meta *Meta) SetMediaType(mediaType string) *Meta {
	meta.MediaType = &mediaType
	return meta
}

// Wrap оборачивает текущий объект [Meta] в [MetaWrapper].
func (meta Meta) Wrap() MetaWrapper {
	return MetaWrapper{meta}
}

// String реализует интерфейс [fmt.Stringer].
func (meta Meta) String() string {
	return Stringify(meta)
}

// IsEqual сравнивает текущий объект [Meta], с объектом [Meta], переданным в качестве аргумента по полю Href.
func (meta *Meta) IsEqual(other *Meta) bool {
	return IsEqualPtr(meta.Href, other.Href)
}

// GetUUIDFromHref возвращает ID из поля Href.
//
// Возвращает "<empty id>", если поле Href пустое или не содержит идентификатора.
func (meta Meta) GetUUIDFromHref() string {
	href := Deref(meta.Href)
	if href == "" {
		return "<empty id>"
	}

	sep := strings.Split(href, "/")
	if len(sep) == 0 {
		return "<empty id>"
	}

	id := sep[len(sep)-1]

	if strings.Contains(id, "?") {
		id = strings.Split(id, "?")[0]
	}

	return id
}

// MetaWrapper объект-обёртка для [Meta]
type MetaWrapper struct {
	Meta Meta `json:"meta"` // метаданные
}

// GetMeta возвращает метаданные.
func (metaWrapper MetaWrapper) GetMeta() Meta {
	return metaWrapper.Meta
}

// String реализует интерфейс [fmt.Stringer].
func (metaWrapper MetaWrapper) String() string {
	return Stringify(metaWrapper)
}

// MetaNameWrapper объект-обёртка для [Meta] и поля Name
type MetaNameWrapper struct {
	Meta Meta   `json:"meta"`
	Name string `json:"name"`
}

// String реализует интерфейс [fmt.Stringer].
func (metaNameWrapper MetaNameWrapper) String() string {
	return Stringify(metaNameWrapper)
}

// MetaCollection Метаданные коллекции.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-kollekcii
type MetaCollection struct {
	Href         string `json:"href,omitempty"`         // ссылка на объект
	Type         string `json:"type,omitempty"`         // тип объекта (код сущности)
	MediaType    string `json:"mediaType,omitempty"`    // тип данных, который приходят в ответ от сервиса, либо отправляется в теле запроса
	NextHref     string `json:"nextHref,omitempty"`     // ссылка на следующую страницу коллекции
	PreviousHref string `json:"previousHref,omitempty"` // ссылка на предыдущую страницу коллекции
	Size         int    `json:"size,omitempty"`         // количество элементов в коллекции
	Limit        int    `json:"limit,omitempty"`        // максимальное число элементов в коллекции, возвращаемых за один запрос
	Offset       int    `json:"offset,omitempty"`       // смещение выборки коллекции от первого элемента
}

// String реализует интерфейс [fmt.Stringer].
func (metaCollection MetaCollection) String() string {
	return Stringify(metaCollection)
}

// MetaArray объект с полями meta и rows, где rows - массив объектов T.
type MetaArray[T any] struct {
	Rows Slice[T]       `json:"rows,omitempty"`
	Meta MetaCollection `json:"meta,omitempty"`
}

// NewMetaArrayFrom принимает срез объектов T и возвращает [MetaArray] с переданными объектами в поле Rows.
//
// nil значения игнорируются.
func NewMetaArrayFrom[T any](rows Slice[T]) *MetaArray[T] {
	return &MetaArray[T]{Rows: rows.Filter(func(t *T) bool { return t != nil })}
}

// Len возвращает количество элементов Rows.
func (metaArray MetaArray[T]) Len() int {
	return len(metaArray.Rows)
}

// Size возвращает размер выданного списка.
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

// String реализует интерфейс [fmt.Stringer].
func (metaArray MetaArray[T]) String() string {
	return Stringify(metaArray)
}

// MarshalJSON реализует интерфейс [json.Marshaler].
func (metaArray MetaArray[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(metaArray.Rows)
}

type MetaNameID struct {
	Meta Meta   `json:"meta"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (metaNameID MetaNameID) String() string {
	return Stringify(metaNameID)
}

// MetaType код сущности.
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
	MetaTypeCounterpartyAdjustment              MetaType = "counterpartyadjustment"
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
	MetaTypeRetailDemand                        MetaType = "retaildemand"
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
	MetaTypeEntitySettings                      MetaType = "entitysettings"
	MetaTypeStateSettings                       MetaType = "statesettings"
	MetaTypeTemplateSettings                    MetaType = "templatesettings"
	MetaTypeUnknown                             MetaType = ""
)

// MetaTypeFromEntity принимает объект v и определяет его код сущности.
//
// Функция вернёт [MetaTypeUnknown], если код сущности определить невозможно.
func MetaTypeFromEntity(v any) MetaType {
	var metaType MetaType

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
		metaType = MetaTypeCounterpartyAdjustment
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
		metaType = MetaTypeUnknown
	}
	return metaType
}
