package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"net/http"
)

// NotificationFieldValue Формат измененного поля.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-tipy-uwedomlenij-formaty-polej
type NotificationFieldValue struct {
	OldValue string `json:"oldValue"` // Значение атрибута до удаления
	NewValue string `json:"newValue"` // Значение атрибута после обновления
}

func (notificationFieldValue NotificationFieldValue) String() string {
	return Stringify(notificationFieldValue)
}

type MetaNameID struct {
	Meta Meta      `json:"meta"`
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func (metaNameID MetaNameID) String() string {
	return Stringify(metaNameID)
}

// NotificationInvoice Метаданные счета.
type NotificationInvoice struct {
	Meta                 Meta      `json:"meta"`
	PaymentPlannedMoment Timestamp `json:"paymentPlannedMoment"`
	Name                 string    `json:"name"`
	CustomerName         string    `json:"customerName"`
	Sum                  float64   `json:"sum"`
	ID                   uuid.UUID `json:"id"`
}

func (invoice NotificationInvoice) String() string {
	return Stringify(invoice)
}

// Order Метаданные заказа.
type Order struct {
	Meta      Meta      `json:"meta"`
	Name      string    `json:"name"`
	AgentName string    `json:"agentName"`
	Sum       float64   `json:"sum"`
	ID        uuid.UUID `json:"id"`
}

func (order Order) String() string {
	return Stringify(order)
}

// NotificationTaskState Статус завершения.
type NotificationTaskState string

const (
	NotificationTaskStateCompleted            NotificationTaskState = "completed"
	NotificationTaskStateInterrupted          NotificationTaskState = "interrupted"
	NotificationTaskStateInterruptedByUser    NotificationTaskState = "interrupted_by_user"
	NotificationTaskStateInterruptedByTimeout NotificationTaskState = "interrupted_by_timeout"
	NotificationTaskStateInterruptedBySystem  NotificationTaskState = "interrupted_by_system"
)

func (notificationTaskState NotificationTaskState) String() string {
	return string(notificationTaskState)
}

// NotificationTaskType Тип задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-tipy-uwedomlenij-formaty-polej
type NotificationTaskType string

const (
	NotificationTaskTypeExportCSVGood                 NotificationTaskType = "export_csv_good"
	NotificationTaskTypeExportCSVAgent                NotificationTaskType = "export_csv_agent"
	NotificationTaskTypeExportMSXML                   NotificationTaskType = "export_ms_xml"
	NotificationTaskTypeExport1CV2XML                 NotificationTaskType = "export_1c_v2_xml"
	NotificationTaskTypeExportUnisender               NotificationTaskType = "export_unisender"
	NotificationTaskTypeExport1CV3XML                 NotificationTaskType = "export_1c_v3_xml"
	NotificationTaskTypeExportSubscribePro            NotificationTaskType = "export_subscribepro"
	NotificationTaskTypeExport1CClientBank            NotificationTaskType = "export_1c_client_bank"
	NotificationTaskTypeExportAlfaPayments            NotificationTaskType = "export_alfa_payments"
	NotificationTaskTypeExportTochkaPayments          NotificationTaskType = "export_tochka_payments"
	NotificationTaskTypeExportModulBankPayments       NotificationTaskType = "export_modulbank_payments"
	NotificationTaskTypeExport1CEnterpriseData        NotificationTaskType = "export_1c_enterprise_data"
	NotificationTaskTypeExportTinkoffPayments         NotificationTaskType = "export_tinkoff_payments"
	NotificationTaskTypeExportGood                    NotificationTaskType = "export_good"
	NotificationTaskTypeExportCustomEntity            NotificationTaskType = "export_custom_entity"
	NotificationTaskTypeImportCVS                     NotificationTaskType = "importer_csv"
	NotificationTaskTypeImportYML                     NotificationTaskType = "importer_yml"
	NotificationTaskTypeImportCSVAgent                NotificationTaskType = "importer_csv_agent"
	NotificationTaskTypeImportCSVCustomerOrder        NotificationTaskType = "importer_csv_customerorder"
	NotificationTaskTypeImportCSVPurchaseOrder        NotificationTaskType = "importer_csv_purchaseorder"
	NotificationTaskTypeImportCSVPriceList            NotificationTaskType = "importer_csv_pricelist"
	NotificationTaskTypeImportMSXML                   NotificationTaskType = "importer_ms_xml"
	NotificationTaskTypeImport1CClientBank            NotificationTaskType = "importer_1c_client_bank"
	NotificationTaskTypeImportAlfaPayments            NotificationTaskType = "import_alfa_payments"
	NotificationTaskTypeImportAlfaPaymentsRequest     NotificationTaskType = "import_alfa_payments_request"
	NotificationTaskTypeImportAlfaPaymentsSave        NotificationTaskType = "import_alfa_payments_save"
	NotificationTaskTypeImportTochkaPayments          NotificationTaskType = "import_tochka_payments"
	NotificationTaskTypeImportModulBankPayments       NotificationTaskType = "import_modulbank_payments"
	NotificationTaskTypeImportTochkaPaymentsSave      NotificationTaskType = "import_tochka_payments_save"
	NotificationTaskTypeImportModulBankPaymentsSave   NotificationTaskType = "import_modulbank_payments_save"
	NotificationTaskTypeImportTinkoffPayments         NotificationTaskType = "import_tinkoff_payments"
	NotificationTaskTypeImportTinkoffPaymentsSave     NotificationTaskType = "import_tinkoff_payments_save"
	NotificationTaskTypeImportGood                    NotificationTaskType = "importer_good"
	NotificationTaskTypeImportGoodInDoc               NotificationTaskType = "importer_good_in_doc"
	NotificationTaskTypeImportEDOSupply               NotificationTaskType = "import_edo_supply"
	NotificationTaskTypeImportUnionCompany            NotificationTaskType = "import_union_company"
	NotificationTaskTypeImportSberbankPaymentsRequest NotificationTaskType = "import_sberbank_payments_request"
	NotificationTaskTypeImportSberbankPaymentsSave    NotificationTaskType = "import_sberbank_payments_save"
	NotificationTaskTypeImportUpdateVatTo20Percents   NotificationTaskType = "import_update_vat_to_20_percents"
	NotificationTaskTypeImportCustomEntity            NotificationTaskType = "import_custom_entity"
)

func (notificationTaskType NotificationTaskType) String() string {
	return string(notificationTaskType)
}

// NotificationExportCompleted Завершение экспорта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zawershenie-axporta
type NotificationExportCompleted struct {
	Meta                Meta                  `json:"meta"`
	Created             Timestamp             `json:"created"`
	CreatedDocumentName string                `json:"createdDocumentName"`
	Description         string                `json:"description"`
	ErrorMessage        string                `json:"errorMessage"`
	Message             string                `json:"message"`
	TaskState           NotificationTaskState `json:"taskState"`
	TaskType            NotificationTaskType  `json:"taskType"`
	Title               string                `json:"title"`
	AccountID           uuid.UUID             `json:"accountId"`
	ID                  uuid.UUID             `json:"id"`
	Read                bool                  `json:"read"`
}

func (notificationExportCompleted NotificationExportCompleted) String() string {
	return Stringify(notificationExportCompleted)
}

// MetaType возвращает тип сущности.
func (NotificationExportCompleted) MetaType() MetaType {
	return MetaTypeNotificationExportCompleted
}

// NotificationImportCompleted Завершение импорта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zawershenie-importa
type NotificationImportCompleted struct {
	Meta                Meta                  `json:"meta"`
	Created             Timestamp             `json:"created"`
	CreatedDocumentName string                `json:"createdDocumentName"`
	Description         string                `json:"description"`
	ErrorMessage        string                `json:"errorMessage"`
	Message             string                `json:"message"`
	TaskState           NotificationTaskState `json:"taskState"`
	TaskType            NotificationTaskType  `json:"taskType"`
	Title               string                `json:"title"`
	AccountID           uuid.UUID             `json:"accountId"`
	ID                  uuid.UUID             `json:"id"`
	Read                bool                  `json:"read"`
}

func (notificationImportCompleted NotificationImportCompleted) String() string {
	return Stringify(notificationImportCompleted)
}

// MetaType возвращает тип сущности.
func (NotificationImportCompleted) MetaType() MetaType {
	return MetaTypeNotificationImportCompleted
}

// NotificationGoodCountTooLow Снижение остатка товара ниже неснижаемого.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-snizhenie-ostatka-towara-nizhe-nesnizhaemogo
type NotificationGoodCountTooLow struct {
	Meta           Meta       `json:"meta"`
	Created        Timestamp  `json:"created"`
	Description    string     `json:"description"`
	Title          string     `json:"title"`
	Good           MetaNameID `json:"good"`
	ActualBalance  int        `json:"actualBalance"`
	MinimumBalance int        `json:"minimumBalance"`
	AccountID      uuid.UUID  `json:"accountId"`
	ID             uuid.UUID  `json:"id"`
	Read           bool       `json:"read"`
}

func (notificationGoodCountTooLow NotificationGoodCountTooLow) String() string {
	return Stringify(notificationGoodCountTooLow)
}

// MetaType возвращает тип сущности.
func (NotificationGoodCountTooLow) MetaType() MetaType {
	return MetaTypeNotificationGoodCountTooLow
}

// NotificationInvoiceOutOverdue Просрочен счет покупателя.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-prosrochen-schet-pokupatelq
type NotificationInvoiceOutOverdue struct {
	Meta                 Meta                `json:"meta"`
	Created              Timestamp           `json:"created"`
	PaymentPlannedMoment Timestamp           `json:"paymentPlannedMoment"`
	AgentName            string              `json:"agentName"`
	Description          string              `json:"description"`
	Invoice              NotificationInvoice `json:"invoice"`
	Sum                  float64             `json:"sum"`
	AccountID            uuid.UUID           `json:"accountId"`
	ID                   uuid.UUID           `json:"id"`
	Read                 bool                `json:"read"`
}

func (notificationInvoiceOutOverdue NotificationInvoiceOutOverdue) String() string {
	return Stringify(notificationInvoiceOutOverdue)
}

// MetaType возвращает тип сущности.
func (NotificationInvoiceOutOverdue) MetaType() MetaType {
	return MetaTypeNotificationInvoiceOutOverdue
}

// NotificationOrderNew Новый заказ.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowyj-zakaz
type NotificationOrderNew struct {
	Meta                  Meta      `json:"meta"`
	Created               Timestamp `json:"created"`
	DeliveryPlannedMoment Timestamp `json:"deliveryPlannedMoment"`
	AgentName             string    `json:"agentName"`
	Description           string    `json:"description"`
	Title                 string    `json:"title"`
	Order                 Order     `json:"order"`
	Sum                   float64   `json:"sum"`
	AccountID             uuid.UUID `json:"accountId"`
	ID                    uuid.UUID `json:"id"`
	Read                  bool      `json:"read"`
}

func (notificationOrderNew NotificationOrderNew) String() string {
	return Stringify(notificationOrderNew)
}

// MetaType возвращает тип сущности.
func (NotificationOrderNew) MetaType() MetaType {
	return MetaTypeNotificationOrderNew
}

// NotificationOrderOverdue Просроченный заказ.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-prosrochennyj-zakaz
type NotificationOrderOverdue struct {
	Meta                  Meta      `json:"meta"`
	Created               Timestamp `json:"created"`
	DeliveryPlannedMoment Timestamp `json:"deliveryPlannedMoment"`
	AgentName             string    `json:"agentName"`
	Description           string    `json:"description"`
	Title                 string    `json:"title"`
	Order                 Order     `json:"order"`
	Sum                   float64   `json:"sum"`
	AccountID             uuid.UUID `json:"accountId"`
	ID                    uuid.UUID `json:"id"`
	Read                  bool      `json:"read"`
}

func (notificationOrderOverdue NotificationOrderOverdue) String() string {
	return Stringify(notificationOrderOverdue)
}

// MetaType возвращает тип сущности.
func (NotificationOrderOverdue) MetaType() MetaType {
	return MetaTypeNotificationOrderOverdue
}

// NotificationSubscribeExpired Окончание подписки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-okonchanie-podpiski
type NotificationSubscribeExpired struct {
	Meta        Meta      `json:"meta"`
	Created     Timestamp `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	AccountID   uuid.UUID `json:"accountId"`
	ID          uuid.UUID `json:"id"`
	Read        bool      `json:"read"`
}

func (notificationSubscribeExpired NotificationSubscribeExpired) String() string {
	return Stringify(notificationSubscribeExpired)
}

// MetaType возвращает тип сущности.
func (NotificationSubscribeExpired) MetaType() MetaType {
	return MetaTypeNotificationSubscribeExpired
}

// NotificationSubscribeTermsExpired Условия подписки истекают.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uslowiq-podpiski-istekaut
type NotificationSubscribeTermsExpired struct {
	Meta        Meta      `json:"meta"`
	Created     Timestamp `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	DaysLeft    int       `json:"daysLeft"`
	AccountID   uuid.UUID `json:"accountId"`
	ID          uuid.UUID `json:"id"`
	Read        bool      `json:"read"`
}

func (notificationSubscribeTermsExpired NotificationSubscribeTermsExpired) String() string {
	return Stringify(notificationSubscribeTermsExpired)
}

// MetaType возвращает тип сущности.
func (NotificationSubscribeTermsExpired) MetaType() MetaType {
	return MetaTypeNotificationSubscribeTermsExpired
}

// NotificationTaskAssigned Задача назначена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-naznachena
type NotificationTaskAssigned struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	PerformedBy MetaNameID       `json:"performedBy"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskAssigned NotificationTaskAssigned) String() string {
	return Stringify(notificationTaskAssigned)
}

// MetaType возвращает тип сущности.
func (NotificationTaskAssigned) MetaType() MetaType {
	return MetaTypeNotificationTaskAssigned
}

// NotificationTaskUnassigned Задача снята.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-snqta
type NotificationTaskUnassigned struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	PerformedBy MetaNameID       `json:"performedBy"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskUnassigned NotificationTaskUnassigned) String() string {
	return Stringify(notificationTaskUnassigned)
}

// MetaType возвращает тип сущности.
func (NotificationTaskUnassigned) MetaType() MetaType {
	return MetaTypeNotificationTaskUnassigned
}

// NotificationTask Задача.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-naznachena-atributy-wlozhennyh-suschnostej-zadacha
type NotificationTask struct {
	Meta     Meta      `json:"meta"`
	Deadline Timestamp `json:"deadline"`
	Name     string    `json:"name"`
	ID       uuid.UUID `json:"id"`
}

// NotificationTaskChanged Задача изменена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-izmenena
type NotificationTaskChanged struct {
	Meta        Meta                        `json:"meta"`
	Created     Timestamp                   `json:"created"`
	Diff        NotificationTaskChangedDiff `json:"diff"`
	Description string                      `json:"description"`
	Title       string                      `json:"title"`
	Task        NotificationTask            `json:"task"`
	PerformedBy MetaNameID                  `json:"performedBy"`
	AccountID   uuid.UUID                   `json:"accountId"`
	ID          uuid.UUID                   `json:"id"`
	Read        bool                        `json:"read"`
}

func (notificationTaskChanged NotificationTaskChanged) String() string {
	return Stringify(notificationTaskChanged)
}

// MetaType возвращает тип сущности.
func (NotificationTaskChanged) MetaType() MetaType {
	return MetaTypeNotificationTaskChanged
}

// NotificationTaskChangedDiff Измененные поля.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-izmenena-atributy-wlozhennyh-suschnostej-izmenennye-polq
type NotificationTaskChangedDiff struct {
	Description  string `json:"description"`
	Deadline     string `json:"deadline"`
	AgentLink    string `json:"agentLink"`
	DocumentLink string `json:"documentLink"`
	Assignee     string `json:"assignee"`
}

// NotificationTaskCompleted Задача выполнена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-wypolnena
type NotificationTaskCompleted struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	PerformedBy MetaNameID       `json:"performedBy"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskCompleted NotificationTaskCompleted) String() string {
	return Stringify(notificationTaskCompleted)
}

// MetaType возвращает тип сущности.
func (NotificationTaskCompleted) MetaType() MetaType {
	return MetaTypeNotificationTaskCompleted
}

// NotificationTaskDeleted Задача удалена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-udalena
type NotificationTaskDeleted struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	PerformedBy MetaNameID       `json:"performedBy"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskDeleted NotificationTaskDeleted) String() string {
	return Stringify(notificationTaskDeleted)
}

// MetaType возвращает тип сущности.
func (NotificationTaskDeleted) MetaType() MetaType {
	return MetaTypeNotificationTaskDeleted
}

// NotificationTaskOverdue Задача просрочена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-prosrochena
type NotificationTaskOverdue struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskOverdue NotificationTaskOverdue) String() string {
	return Stringify(notificationTaskOverdue)
}

// MetaType возвращает тип сущности.
func (NotificationTaskOverdue) MetaType() MetaType {
	return MetaTypeNotificationTaskOverdue
}

// NotificationTaskReopened Задача переоткрыта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-pereotkryta
type NotificationTaskReopened struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	PerformedBy MetaNameID       `json:"performedBy"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskReopened NotificationTaskReopened) String() string {
	return Stringify(notificationTaskReopened)
}

// MetaType возвращает тип сущности.
func (NotificationTaskReopened) MetaType() MetaType {
	return MetaTypeNotificationTaskReopened
}

// NotificationTaskNewComment Новый комментарий к задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowyj-kommentarij-k-zadache
type NotificationTaskNewComment struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	NoteContent string           `json:"noteContent"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	PerformedBy MetaNameID       `json:"performedBy"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskNewComment NotificationTaskNewComment) String() string {
	return Stringify(notificationTaskNewComment)
}

// MetaType возвращает тип сущности.
func (NotificationTaskNewComment) MetaType() MetaType {
	return MetaTypeNotificationTaskNewComment
}

// NotificationTaskCommentChanged Изменен комментарий к задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-izmenen-kommentarij-k-zadache
type NotificationTaskCommentChanged struct {
	Meta        Meta                        `json:"meta"`
	Created     Timestamp                   `json:"created"`
	Diff        NotificationTaskChangedDiff `json:"diff"`
	Description string                      `json:"description"`
	NoteContent string                      `json:"noteContent"`
	Title       string                      `json:"title"`
	Task        NotificationTask            `json:"task"`
	PerformedBy MetaNameID                  `json:"performedBy"`
	AccountID   uuid.UUID                   `json:"accountId"`
	ID          uuid.UUID                   `json:"id"`
	Read        bool                        `json:"read"`
}

func (notificationTaskCommentChanged NotificationTaskCommentChanged) String() string {
	return Stringify(notificationTaskCommentChanged)
}

// MetaType возвращает тип сущности.
func (NotificationTaskCommentChanged) MetaType() MetaType {
	return MetaTypeNotificationTaskCommentChanged
}

// NotificationTaskCommentDeleted Удален комментарий к задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-udalen-kommentarij-k-zadache
type NotificationTaskCommentDeleted struct {
	Meta        Meta             `json:"meta"`
	Created     Timestamp        `json:"created"`
	Description string           `json:"description"`
	NoteContent string           `json:"noteContent"`
	Title       string           `json:"title"`
	Task        NotificationTask `json:"task"`
	PerformedBy MetaNameID       `json:"performedBy"`
	AccountID   uuid.UUID        `json:"accountId"`
	ID          uuid.UUID        `json:"id"`
	Read        bool             `json:"read"`
}

func (notificationTaskCommentDeleted NotificationTaskCommentDeleted) String() string {
	return Stringify(notificationTaskCommentDeleted)
}

// MetaType возвращает тип сущности.
func (NotificationTaskCommentDeleted) MetaType() MetaType {
	return MetaTypeNotificationTaskCommentDeleted
}

type NotificationRetailShift struct {
	Meta    Meta      `json:"meta"`
	Open    Timestamp `json:"open"`
	Name    string    `json:"name"`
	Proceed float64   `json:"proceed"`
	ID      uuid.UUID `json:"id"`
}

// NotificationRetailShiftOpened Смена открыта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-smena-otkryta
type NotificationRetailShiftOpened struct {
	Meta        Meta                    `json:"meta"`
	Created     Timestamp               `json:"created"`
	Description string                  `json:"description"`
	Title       string                  `json:"title"`
	RetailStore MetaNameID              `json:"retailStore"`
	User        MetaNameID              `json:"user"`
	RetailShift NotificationRetailShift `json:"retailShift"`
	AccountID   uuid.UUID               `json:"accountId"`
	ID          uuid.UUID               `json:"id"`
	Read        bool                    `json:"read"`
}

func (notificationRetailShiftOpened NotificationRetailShiftOpened) String() string {
	return Stringify(notificationRetailShiftOpened)
}

// MetaType возвращает тип сущности.
func (NotificationRetailShiftOpened) MetaType() MetaType {
	return MetaTypeNotificationRetailShiftOpened
}

// NotificationRetailShiftClosed Смена закрыта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-smena-zakryta
type NotificationRetailShiftClosed struct {
	Meta        Meta                    `json:"meta"`
	Created     Timestamp               `json:"created"`
	Description string                  `json:"description"`
	Title       string                  `json:"title"`
	RetailStore MetaNameID              `json:"retailStore"`
	User        MetaNameID              `json:"user"`
	RetailShift NotificationRetailShift `json:"retailShift"`
	Returns     int                     `json:"returns"`
	Sales       int                     `json:"sales"`
	AccountID   uuid.UUID               `json:"accountId"`
	ID          uuid.UUID               `json:"id"`
	Read        bool                    `json:"read"`
}

func (notificationRetailShiftClosed NotificationRetailShiftClosed) String() string {
	return Stringify(notificationRetailShiftClosed)
}

// MetaType возвращает тип сущности.
func (NotificationRetailShiftClosed) MetaType() MetaType {
	return MetaTypeNotificationRetailShiftClosed
}

// NotificationScript Уведомление из сценария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uwedomlenie-iz-scenariq
type NotificationScript struct {
	Meta        Meta       `json:"meta"`
	Created     Timestamp  `json:"created"`
	Description string     `json:"description"`
	EventType   EventType  `json:"eventType"`
	Title       string     `json:"title"`
	Entity      MetaNameID `json:"entity"`
	AccountID   uuid.UUID  `json:"accountId"`
	ID          uuid.UUID  `json:"id"`
	Read        bool       `json:"read"`
}

func (notificationScript NotificationScript) String() string {
	return Stringify(notificationScript)
}

// MetaType возвращает тип сущности.
func (NotificationScript) MetaType() MetaType {
	return MetaTypeNotificationScript
}

// EventType Тип события сценария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uwedomlenie-iz-scenariq-atributy-wlozhennyh-suschnostej-sobytie
type EventType string

const (
	EventTypeAdd             EventType = "ADD"           // создан
	EventTypeModify          EventType = "MODIFY"        // изменен
	EventTypeAddChangeStatus EventType = "CHANGE_STATUS" // изменен статус
)

func (eventType EventType) String() string {
	return string(eventType)
}

// FacebookTokenExpirationNotification Предупреждение о скором окончании действия доступа к аккаунту Facebook.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-preduprezhdenie-o-skorom-okonchanii-dejstwiq-dostupa-k-akkauntu-facebook
type FacebookTokenExpirationNotification struct {
	Meta                 Meta      `json:"meta"`
	Created              Timestamp `json:"created"`
	ConnectorName        string    `json:"connectorName"`
	Description          string    `json:"description"`
	Title                string    `json:"title"`
	DaysLeftToExpiration int       `json:"daysLeftToExpiration"`
	AccountID            uuid.UUID `json:"accountId"`
	ID                   uuid.UUID `json:"id"`
	Read                 bool      `json:"read"`
}

func (facebookTokenExpirationNotification FacebookTokenExpirationNotification) String() string {
	return Stringify(facebookTokenExpirationNotification)
}

// MetaType возвращает тип сущности.
func (FacebookTokenExpirationNotification) MetaType() MetaType {
	return MetaTypeFacebookTokenExpirationNotification
}

// NotificationBonusMoney На счет зачислены бонусные деньги.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-na-schet-zachisleny-bonusnye-den-gi
type NotificationBonusMoney struct {
	Meta        Meta      `json:"meta"`
	Created     Timestamp `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	AccountID   uuid.UUID `json:"accountId"`
	ID          uuid.UUID `json:"id"`
	Read        bool      `json:"read"`
}

func (notificationBonusMoney NotificationBonusMoney) String() string {
	return Stringify(notificationBonusMoney)
}

// MetaType возвращает тип сущности.
func (NotificationBonusMoney) MetaType() MetaType {
	return MetaTypeNotificationBonusMoney
}

// NewMentionInEvent Новое упоминание в ленте событий.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowoe-upominanie-w-lente-sobytij
type NewMentionInEvent struct {
	Meta        Meta          `json:"meta"`
	Created     Timestamp     `json:"created"`
	Description string        `json:"description"`
	Title       string        `json:"title"`
	Operation   TaskOperation `json:"operation"`
	AccountID   uuid.UUID     `json:"accountId"`
	ID          uuid.UUID     `json:"id"`
	Read        bool          `json:"read"`
}

func (newMentionInEvent NewMentionInEvent) String() string {
	return Stringify(newMentionInEvent)
}

// MetaType возвращает тип сущности.
func (NewMentionInEvent) MetaType() MetaType {
	return MetaTypeNewMentionInEvent
}

type SubscriptionGroup string

const (
	SubscriptionGroupCustomerOrder SubscriptionGroup = "customer_order" // Заказы покупателей
	SubscriptionGroupDataExchange  SubscriptionGroup = "data_exchange"  // Обмен данными
	SubscriptionGroupInvoice       SubscriptionGroup = "invoice"        // Счета покупателей
	SubscriptionGroupRetail        SubscriptionGroup = "retail"         // Розничная торговля
	SubscriptionGroupScripts       SubscriptionGroup = "scripts"        // Сценарии
	SubscriptionGroupStock         SubscriptionGroup = "stock"          // Складские остатки
	SubscriptionGroupTask          SubscriptionGroup = "task"           // Задачи
	SubscriptionGroupMentions      SubscriptionGroup = "mentions"       // Упоминания сотрудников
)

func (subscriptionGroup SubscriptionGroup) String() string {
	return string(subscriptionGroup)
}

type SubscriptionChannel string

const (
	SubscriptionChannelEmail SubscriptionChannel = "email"
	SubscriptionChannelPush  SubscriptionChannel = "push"
)

func (subscriptionChannel SubscriptionChannel) String() string {
	return string(subscriptionChannel)
}

type SubscriptionElement struct {
	Channels []SubscriptionChannel `json:"channels"`
	Enabled  bool                  `json:"enabled"`
}

// NotificationSubscription Настройки уведомлений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-nastrojki-uwedomlenij-atributy-suschnosti
type NotificationSubscription struct {
	Groups map[SubscriptionGroup]SubscriptionElement
}

// Notification
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-lenta-uwedomlenij-obschie-atributy-uwedomlenij
type Notification struct {
	Meta        Meta      `json:"meta"`
	Created     Timestamp `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	raw         []byte
	AccountID   uuid.UUID `json:"accountId"`
	ID          uuid.UUID `json:"id"`
	Read        bool      `json:"read"`
}

func (notification Notification) String() string {
	return Stringify(notification)
}

// MetaType возвращает тип сущности.
func (notification Notification) MetaType() MetaType {
	return notification.Meta.GetType()
}

// Raw реализует интерфейс RawMetaTyper
func (notification Notification) Raw() []byte {
	return notification.raw
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (notification *Notification) UnmarshalJSON(data []byte) error {
	type alias Notification
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.raw = data
	*notification = Notification(t)
	return nil
}

// AsFacebookTokenExpirationNotification десериализует объект в тип *FacebookTokenExpirationNotification
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsFacebookTokenExpirationNotification() *FacebookTokenExpirationNotification {
	return UnmarshalAsType[FacebookTokenExpirationNotification](notification)
}

// AsNotificationExportCompleted десериализует объект в тип *NotificationExportCompleted
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationExportCompleted() *NotificationExportCompleted {
	return UnmarshalAsType[NotificationExportCompleted](notification)
}

// AsNotificationGoodCountTooLow десериализует объект в тип *NotificationGoodCountTooLow
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationGoodCountTooLow() *NotificationGoodCountTooLow {
	return UnmarshalAsType[NotificationGoodCountTooLow](notification)
}

// AsNotificationImportCompleted десериализует объект в тип *NotificationImportCompleted
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationImportCompleted() *NotificationImportCompleted {
	return UnmarshalAsType[NotificationImportCompleted](notification)
}

// AsNotificationInvoiceOutOverdue десериализует объект в тип *NotificationInvoiceOutOverdue
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationInvoiceOutOverdue() *NotificationInvoiceOutOverdue {
	return UnmarshalAsType[NotificationInvoiceOutOverdue](notification)
}

// AsNotificationOrderNew десериализует объект в тип *NotificationOrderNew
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationOrderNew() *NotificationOrderNew {
	return UnmarshalAsType[NotificationOrderNew](notification)
}

// AsNotificationOrderOverdue десериализует объект в тип *NotificationOrderOverdue
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationOrderOverdue() *NotificationOrderOverdue {
	return UnmarshalAsType[NotificationOrderOverdue](notification)
}

// AsNotificationRetailShiftClosed десериализует объект в тип *NotificationRetailShiftClosed
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationRetailShiftClosed() *NotificationRetailShiftClosed {
	return UnmarshalAsType[NotificationRetailShiftClosed](notification)
}

// AsNotificationRetailShiftOpened десериализует объект в тип *NotificationRetailShiftOpened
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationRetailShiftOpened() *NotificationRetailShiftOpened {
	return UnmarshalAsType[NotificationRetailShiftOpened](notification)
}

// AsNotificationScript десериализует объект в тип *NotificationScript
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationScript() *NotificationScript {
	return UnmarshalAsType[NotificationScript](notification)
}

// AsNotificationSubscribeExpired десериализует объект в тип *NotificationSubscribeExpired
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationSubscribeExpired() *NotificationSubscribeExpired {
	return UnmarshalAsType[NotificationSubscribeExpired](notification)
}

// AsNotificationSubscribeTermsExpired десериализует объект в тип *NotificationSubscribeTermsExpired
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationSubscribeTermsExpired() *NotificationSubscribeTermsExpired {
	return UnmarshalAsType[NotificationSubscribeTermsExpired](notification)
}

// AsNotificationTaskAssigned десериализует объект в тип *NotificationTaskAssigned
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskAssigned() *NotificationTaskAssigned {
	return UnmarshalAsType[NotificationTaskAssigned](notification)
}

// AsNotificationTaskChanged десериализует объект в тип *NotificationTaskChanged
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskChanged() *NotificationTaskChanged {
	return UnmarshalAsType[NotificationTaskChanged](notification)
}

// AsNotificationTaskCommentChanged десериализует объект в тип *NotificationTaskCommentChanged
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskCommentChanged() *NotificationTaskCommentChanged {
	return UnmarshalAsType[NotificationTaskCommentChanged](notification)
}

// AsNotificationTaskCommentDeleted десериализует объект в тип *NotificationTaskCommentDeleted
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskCommentDeleted() *NotificationTaskCommentDeleted {
	return UnmarshalAsType[NotificationTaskCommentDeleted](notification)
}

// AsNotificationTaskCompleted десериализует объект в тип *NotificationTaskCompleted
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskCompleted() *NotificationTaskCompleted {
	return UnmarshalAsType[NotificationTaskCompleted](notification)
}

// AsNotificationTaskDeleted десериализует объект в тип *NotificationTaskCompleted
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskDeleted() *NotificationTaskDeleted {
	return UnmarshalAsType[NotificationTaskDeleted](notification)
}

// AsNotificationTaskNewComment десериализует объект в тип *NotificationTaskNewComment
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskNewComment() *NotificationTaskNewComment {
	return UnmarshalAsType[NotificationTaskNewComment](notification)
}

// AsNotificationTaskOverdue десериализует объект в тип *NotificationTaskOverdue
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskOverdue() *NotificationTaskOverdue {
	return UnmarshalAsType[NotificationTaskOverdue](notification)
}

// AsNotificationTaskReopened десериализует объект в тип *NotificationTaskReopened
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskReopened() *NotificationTaskReopened {
	return UnmarshalAsType[NotificationTaskReopened](notification)
}

// AsNotificationTaskUnassigned десериализует объект в тип *NotificationTaskUnassigned
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationTaskUnassigned() *NotificationTaskUnassigned {
	return UnmarshalAsType[NotificationTaskUnassigned](notification)
}

// AsNotificationBonusMoney десериализует объект в тип *NotificationBonusMoney
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNotificationBonusMoney() *NotificationBonusMoney {
	return UnmarshalAsType[NotificationBonusMoney](notification)
}

// AsNewMentionInEvent десериализует объект в тип *NewMentionInEvent
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (notification Notification) AsNewMentionInEvent() *NewMentionInEvent {
	return UnmarshalAsType[NewMentionInEvent](notification)
}

// NotificationService Сервис для работы с уведомлениями.
type NotificationService interface {
	// GetList Получить ленту Уведомлений.
	GetList(ctx context.Context, params ...*Params) (*List[Notification], *resty.Response, error)

	// GetByID Запрос на получение Уведомления с указанным id.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Notification, *resty.Response, error)

	// Delete Запрос на удаление Уведомления с указанным id.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// MarkAsRead Отметить Уведомление как прочитанное.
	MarkAsRead(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// MarkAsReadAll Отметить все Уведомления как прочитанные.
	MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error)

	// GetSubscription Запрос настроек Уведомлений текущего пользователя.
	GetSubscription(ctx context.Context) (*NotificationSubscription, *resty.Response, error)

	// UpdateSubscription Изменение настроек Уведомлений текущего пользователя.
	UpdateSubscription(ctx context.Context, notificationSubscription *NotificationSubscription) (bool, *resty.Response, error)
}

type notificationService struct {
	Endpoint
	endpointGetList[Notification]
	endpointGetByID[Notification]
	endpointDelete
}

func NewNotificationService(client *Client) NotificationService {
	e := NewEndpoint(client, "notification")
	return &notificationService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Notification]{e},
		endpointGetByID: endpointGetByID[Notification]{e},
		endpointDelete:  endpointDelete{e},
	}
}

// MarkAsRead Отметить Уведомление как прочитанное.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-uwedomlenie-otmetit-uwedomlenie-kak-prochitannoe
func (service *notificationService) MarkAsRead(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/markasread", service.uri, id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}

// MarkAsReadAll Отметить все Уведомления как прочитанные.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-uwedomlenie-otmetit-wse-uwedomleniq-kak-prochitannye
func (service *notificationService) MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/markasreadall", service.uri)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}

func (service *notificationService) GetSubscription(ctx context.Context) (*NotificationSubscription, *resty.Response, error) {
	path := "subscription"
	return NewRequestBuilder[NotificationSubscription](service.client, path).Get(ctx)
}

func (service *notificationService) UpdateSubscription(ctx context.Context, notificationSubscription *NotificationSubscription) (bool, *resty.Response, error) {
	path := "subscription"
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, notificationSubscription)
	return resp.StatusCode() == http.StatusOK, resp, err
}
