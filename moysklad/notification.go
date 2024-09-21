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
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-tipy-uwedomlenij-formaty-polej
type NotificationFieldValue struct {
	OldValue string `json:"oldValue"` // Значение атрибута до удаления
	NewValue string `json:"newValue"` // Значение атрибута после обновления
}

// String реализует интерфейс [fmt.Stringer].
func (notificationFieldValue NotificationFieldValue) String() string {
	return Stringify(notificationFieldValue)
}

// NotificationInvoice Метаданные счета.
type NotificationInvoice struct {
	Meta                 Meta      `json:"meta"` // Метаданные объекта. Содержит тип конкретного уведомления
	PaymentPlannedMoment Timestamp `json:"paymentPlannedMoment"`
	Name                 string    `json:"name"`
	CustomerName         string    `json:"customerName"`
	Sum                  float64   `json:"sum"`
	ID                   uuid.UUID `json:"id"` // ID Уведомления
}

// String реализует интерфейс [fmt.Stringer].
func (invoice NotificationInvoice) String() string {
	return Stringify(invoice)
}

// Order Метаданные заказа.
type Order struct {
	Meta      Meta      `json:"meta"` // Метаданные объекта. Содержит тип конкретного уведомления
	Name      string    `json:"name"`
	AgentName string    `json:"agentName"` // Имя контрагента
	Sum       float64   `json:"sum"`
	ID        uuid.UUID `json:"id"` // ID Уведомления
}

// String реализует интерфейс [fmt.Stringer].
func (order Order) String() string {
	return Stringify(order)
}

// NotificationTaskState Статус завершения.
//
// Возможные значения:
//   - NotificationTaskStateCompleted            – выполнено
//   - NotificationTaskStateInterrupted          – прервано
//   - NotificationTaskStateInterruptedByUser    – прервано пользователем
//   - NotificationTaskStateInterruptedByTimeout – прервано по таймауту
//   - NotificationTaskStateInterruptedBySystem  – прервано системой
type NotificationTaskState string

const (
	NotificationTaskStateCompleted            NotificationTaskState = "completed"              // выполнено
	NotificationTaskStateInterrupted          NotificationTaskState = "interrupted"            // прервано
	NotificationTaskStateInterruptedByUser    NotificationTaskState = "interrupted_by_user"    // прервано пользователем
	NotificationTaskStateInterruptedByTimeout NotificationTaskState = "interrupted_by_timeout" // прервано по таймауту
	NotificationTaskStateInterruptedBySystem  NotificationTaskState = "interrupted_by_system"  // прервано системой
)

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskState NotificationTaskState) String() string {
	return string(notificationTaskState)
}

// NotificationTaskType Тип задачи.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-tipy-uwedomlenij-formaty-polej
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

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskType NotificationTaskType) String() string {
	return string(notificationTaskType)
}

// NotificationExportCompleted Завершение экспорта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zawershenie-axporta
type NotificationExportCompleted struct {
	Meta                Meta                  `json:"meta"`                // Метаданные объекта. Содержит тип конкретного уведомления
	Created             Timestamp             `json:"created"`             // Дата и время формирования Уведомления
	CreatedDocumentName string                `json:"createdDocumentName"` // Имя экспортированного документа
	Description         string                `json:"description"`         // Описание уведомления
	ErrorMessage        string                `json:"errorMessage"`        // Сообщение об ошибке
	Message             string                `json:"message"`             // Сообщение о завершении экспорта
	TaskState           NotificationTaskState `json:"taskState"`           // Статус завершения
	TaskType            NotificationTaskType  `json:"taskType"`            // Тип экспорта
	Title               string                `json:"title"`               // Краткий текст уведомления
	AccountID           uuid.UUID             `json:"accountId"`           // ID учетной записи
	ID                  uuid.UUID             `json:"id"`                  // ID Уведомления
	Read                bool                  `json:"read"`                // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationExportCompleted NotificationExportCompleted) String() string {
	return Stringify(notificationExportCompleted)
}

// MetaType возвращает код сущности.
func (NotificationExportCompleted) MetaType() MetaType {
	return MetaTypeNotificationExportCompleted
}

// NotificationImportCompleted Завершение импорта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zawershenie-importa
type NotificationImportCompleted struct {
	Meta                Meta                  `json:"meta"`                // Метаданные объекта. Содержит тип конкретного уведомления
	Created             Timestamp             `json:"created"`             // Дата и время формирования Уведомления
	CreatedDocumentName string                `json:"createdDocumentName"` // Имя импортированного документа
	Description         string                `json:"description"`         // Описание уведомления
	ErrorMessage        string                `json:"errorMessage"`        // Сообщение об ошибке
	Message             string                `json:"message"`             // Сообщение о завершении импорта
	TaskState           NotificationTaskState `json:"taskState"`           // Статус завершения
	TaskType            NotificationTaskType  `json:"taskType"`            // Тип экспорта
	Title               string                `json:"title"`               // Краткий текст уведомления
	AccountID           uuid.UUID             `json:"accountId"`           // ID учетной записи
	ID                  uuid.UUID             `json:"id"`                  // ID Уведомления
	Read                bool                  `json:"read"`                // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationImportCompleted NotificationImportCompleted) String() string {
	return Stringify(notificationImportCompleted)
}

// MetaType возвращает код сущности.
func (NotificationImportCompleted) MetaType() MetaType {
	return MetaTypeNotificationImportCompleted
}

// NotificationGoodCountTooLow Снижение остатка товара ниже неснижаемого.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-snizhenie-ostatka-towara-nizhe-nesnizhaemogo
type NotificationGoodCountTooLow struct {
	Meta           Meta       `json:"meta"`           // Метаданные объекта. Содержит тип конкретного уведомления
	Created        Timestamp  `json:"created"`        // Дата и время формирования Уведомления
	Description    string     `json:"description"`    // Описание уведомления
	Title          string     `json:"title"`          // Краткий текст уведомления
	Good           MetaNameID `json:"good"`           // Метаданные товара
	ActualBalance  int        `json:"actualBalance"`  // Остаток товара
	MinimumBalance int        `json:"minimumBalance"` // Неснижаемый остаток товара
	AccountID      uuid.UUID  `json:"accountId"`      // ID учетной записи
	ID             uuid.UUID  `json:"id"`             // ID Уведомления
	Read           bool       `json:"read"`           // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationGoodCountTooLow NotificationGoodCountTooLow) String() string {
	return Stringify(notificationGoodCountTooLow)
}

// MetaType возвращает код сущности.
func (NotificationGoodCountTooLow) MetaType() MetaType {
	return MetaTypeNotificationGoodCountTooLow
}

// NotificationInvoiceOutOverdue Просрочен счет покупателя.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-prosrochen-schet-pokupatelq
type NotificationInvoiceOutOverdue struct {
	Meta                 Meta                `json:"meta"`                 // Метаданные объекта. Содержит тип конкретного уведомления
	Created              Timestamp           `json:"created"`              // Дата и время формирования Уведомления
	PaymentPlannedMoment Timestamp           `json:"paymentPlannedMoment"` // Запланированная дата оплаты
	AgentName            string              `json:"agentName"`            // Имя контрагента
	Description          string              `json:"description"`          // Описание уведомления
	Invoice              NotificationInvoice `json:"invoice"`              // Метаданные счета
	Sum                  float64             `json:"sum"`                  // Сумма счета
	AccountID            uuid.UUID           `json:"accountId"`            // ID учетной записи
	ID                   uuid.UUID           `json:"id"`                   // ID Уведомления
	Read                 bool                `json:"read"`                 // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationInvoiceOutOverdue NotificationInvoiceOutOverdue) String() string {
	return Stringify(notificationInvoiceOutOverdue)
}

// MetaType возвращает код сущности.
func (NotificationInvoiceOutOverdue) MetaType() MetaType {
	return MetaTypeNotificationInvoiceOutOverdue
}

// NotificationOrderNew Новый заказ.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowyj-zakaz
type NotificationOrderNew struct {
	Meta                  Meta      `json:"meta"`                  // Метаданные объекта. Содержит тип конкретного уведомления
	Created               Timestamp `json:"created"`               // Дата и время формирования Уведомления
	DeliveryPlannedMoment Timestamp `json:"deliveryPlannedMoment"` // Планируемое время отгрузки
	AgentName             string    `json:"agentName"`             // Имя контрагента
	Description           string    `json:"description"`           // Описание уведомления
	Title                 string    `json:"title"`                 // Краткий текст уведомления
	Order                 Order     `json:"order"`                 // Метаданные заказа
	Sum                   float64   `json:"sum"`                   // Сумма
	AccountID             uuid.UUID `json:"accountId"`             // ID учетной записи
	ID                    uuid.UUID `json:"id"`                    // ID Уведомления
	Read                  bool      `json:"read"`                  // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationOrderNew NotificationOrderNew) String() string {
	return Stringify(notificationOrderNew)
}

// MetaType возвращает код сущности.
func (NotificationOrderNew) MetaType() MetaType {
	return MetaTypeNotificationOrderNew
}

// NotificationOrderOverdue Просроченный заказ.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-prosrochennyj-zakaz
type NotificationOrderOverdue struct {
	Meta                  Meta      `json:"meta"`                  // Метаданные объекта. Содержит тип конкретного уведомления
	Created               Timestamp `json:"created"`               // Дата и время формирования Уведомления
	DeliveryPlannedMoment Timestamp `json:"deliveryPlannedMoment"` // Планируемое время отгрузки
	AgentName             string    `json:"agentName"`             // Имя контрагента
	Description           string    `json:"description"`           // Описание уведомления
	Title                 string    `json:"title"`                 // Краткий текст уведомления
	Order                 Order     `json:"order"`                 // Метаданные заказа
	Sum                   float64   `json:"sum"`                   // Сумма
	AccountID             uuid.UUID `json:"accountId"`             // ID учетной записи
	ID                    uuid.UUID `json:"id"`                    // ID Уведомления
	Read                  bool      `json:"read"`                  // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationOrderOverdue NotificationOrderOverdue) String() string {
	return Stringify(notificationOrderOverdue)
}

// MetaType возвращает код сущности.
func (NotificationOrderOverdue) MetaType() MetaType {
	return MetaTypeNotificationOrderOverdue
}

// NotificationSubscribeExpired Окончание подписки.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-okonchanie-podpiski
type NotificationSubscribeExpired struct {
	Meta        Meta      `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp `json:"created"`     // Дата и время формирования Уведомления
	Description string    `json:"description"` // Описание уведомления
	Title       string    `json:"title"`       // Краткий текст уведомления
	AccountID   uuid.UUID `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID `json:"id"`          // ID Уведомления
	Read        bool      `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationSubscribeExpired NotificationSubscribeExpired) String() string {
	return Stringify(notificationSubscribeExpired)
}

// MetaType возвращает код сущности.
func (NotificationSubscribeExpired) MetaType() MetaType {
	return MetaTypeNotificationSubscribeExpired
}

// NotificationSubscribeTermsExpired Условия подписки истекают.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uslowiq-podpiski-istekaut
type NotificationSubscribeTermsExpired struct {
	Meta        Meta      `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp `json:"created"`     // Дата и время формирования Уведомления
	Description string    `json:"description"` // Описание уведомления
	Title       string    `json:"title"`       // Краткий текст уведомления
	DaysLeft    int       `json:"daysLeft"`    // Количество оставшихся дней подписки
	AccountID   uuid.UUID `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID `json:"id"`          // ID Уведомления
	Read        bool      `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationSubscribeTermsExpired NotificationSubscribeTermsExpired) String() string {
	return Stringify(notificationSubscribeTermsExpired)
}

// MetaType возвращает код сущности.
func (NotificationSubscribeTermsExpired) MetaType() MetaType {
	return MetaTypeNotificationSubscribeTermsExpired
}

// NotificationTaskAssigned Задача назначена.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-naznachena
type NotificationTaskAssigned struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	PerformedBy MetaNameID       `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskAssigned NotificationTaskAssigned) String() string {
	return Stringify(notificationTaskAssigned)
}

// MetaType возвращает код сущности.
func (NotificationTaskAssigned) MetaType() MetaType {
	return MetaTypeNotificationTaskAssigned
}

// NotificationTaskUnassigned Задача снята.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-snqta
type NotificationTaskUnassigned struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	PerformedBy MetaNameID       `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskUnassigned NotificationTaskUnassigned) String() string {
	return Stringify(notificationTaskUnassigned)
}

// MetaType возвращает код сущности.
func (NotificationTaskUnassigned) MetaType() MetaType {
	return MetaTypeNotificationTaskUnassigned
}

// NotificationTask Задача.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-naznachena-atributy-wlozhennyh-suschnostej-zadacha
type NotificationTask struct {
	Meta     Meta      `json:"meta"`     // Метаданные объекта. Содержит тип конкретного уведомления
	Deadline Timestamp `json:"deadline"` // Планируемая дата завершения задачи
	Name     string    `json:"name"`     // Наименование Контрагента
	ID       uuid.UUID `json:"id"`       // ID Уведомления
}

// NotificationTaskChanged Задача изменена.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-izmenena
type NotificationTaskChanged struct {
	Meta        Meta                        `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp                   `json:"created"`     // Дата и время формирования Уведомления
	Diff        NotificationTaskChangedDiff `json:"diff"`        // Измененные поля
	Description string                      `json:"description"` // Описание уведомления
	Title       string                      `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask            `json:"task"`        // Задача
	PerformedBy MetaNameID                  `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID                   `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID                   `json:"id"`          // ID Уведомления
	Read        bool                        `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskChanged NotificationTaskChanged) String() string {
	return Stringify(notificationTaskChanged)
}

// MetaType возвращает код сущности.
func (NotificationTaskChanged) MetaType() MetaType {
	return MetaTypeNotificationTaskChanged
}

// NotificationTaskChangedDiff Измененные поля.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-izmenena-atributy-wlozhennyh-suschnostej-izmenennye-polq
type NotificationTaskChangedDiff struct {
	Description  string `json:"description"`  // Изменение описания задачи
	Deadline     string `json:"deadline"`     // Изменение планируемой даты завершения задачи
	AgentLink    string `json:"agentLink"`    // Изменение контрагента
	DocumentLink string `json:"documentLink"` // Изменение связанного документа
	Assignee     string `json:"assignee"`     // Изменение исполнителя
}

// NotificationTaskCompleted Задача выполнена.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-wypolnena
type NotificationTaskCompleted struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	PerformedBy MetaNameID       `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskCompleted NotificationTaskCompleted) String() string {
	return Stringify(notificationTaskCompleted)
}

// MetaType возвращает код сущности.
func (NotificationTaskCompleted) MetaType() MetaType {
	return MetaTypeNotificationTaskCompleted
}

// NotificationTaskDeleted Задача удалена.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-udalena
type NotificationTaskDeleted struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	PerformedBy MetaNameID       `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskDeleted NotificationTaskDeleted) String() string {
	return Stringify(notificationTaskDeleted)
}

// MetaType возвращает код сущности.
func (NotificationTaskDeleted) MetaType() MetaType {
	return MetaTypeNotificationTaskDeleted
}

// NotificationTaskOverdue Задача просрочена.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-prosrochena
type NotificationTaskOverdue struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskOverdue NotificationTaskOverdue) String() string {
	return Stringify(notificationTaskOverdue)
}

// MetaType возвращает код сущности.
func (NotificationTaskOverdue) MetaType() MetaType {
	return MetaTypeNotificationTaskOverdue
}

// NotificationTaskReopened Задача переоткрыта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-pereotkryta
type NotificationTaskReopened struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	PerformedBy MetaNameID       `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskReopened NotificationTaskReopened) String() string {
	return Stringify(notificationTaskReopened)
}

// MetaType возвращает код сущности.
func (NotificationTaskReopened) MetaType() MetaType {
	return MetaTypeNotificationTaskReopened
}

// NotificationTaskNewComment Новый комментарий к задаче.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowyj-kommentarij-k-zadache
type NotificationTaskNewComment struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	NoteContent string           `json:"noteContent"` // Содержимое комментария
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	PerformedBy MetaNameID       `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskNewComment NotificationTaskNewComment) String() string {
	return Stringify(notificationTaskNewComment)
}

// MetaType возвращает код сущности.
func (NotificationTaskNewComment) MetaType() MetaType {
	return MetaTypeNotificationTaskNewComment
}

// NotificationTaskCommentChanged Изменен комментарий к задаче.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-izmenen-kommentarij-k-zadache
type NotificationTaskCommentChanged struct {
	Meta        Meta                        `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp                   `json:"created"`     // Дата и время формирования Уведомления
	Diff        NotificationTaskChangedDiff `json:"diff"`        // Изменения комментария
	Description string                      `json:"description"` // Описание уведомления
	NoteContent string                      `json:"noteContent"` // Содержимое комментария
	Title       string                      `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask            `json:"task"`        // Задача
	PerformedBy MetaNameID                  `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID                   `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID                   `json:"id"`          // ID Уведомления
	Read        bool                        `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskCommentChanged NotificationTaskCommentChanged) String() string {
	return Stringify(notificationTaskCommentChanged)
}

// MetaType возвращает код сущности.
func (NotificationTaskCommentChanged) MetaType() MetaType {
	return MetaTypeNotificationTaskCommentChanged
}

// NotificationTaskCommentDeleted Удален комментарий к задаче.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-udalen-kommentarij-k-zadache
type NotificationTaskCommentDeleted struct {
	Meta        Meta             `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp        `json:"created"`     // Дата и время формирования Уведомления
	Description string           `json:"description"` // Описание уведомления
	NoteContent string           `json:"noteContent"` // Содержимое комментария
	Title       string           `json:"title"`       // Краткий текст уведомления
	Task        NotificationTask `json:"task"`        // Задача
	PerformedBy MetaNameID       `json:"performedBy"` // Сотрудник, выполнивший изменение
	AccountID   uuid.UUID        `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID        `json:"id"`          // ID Уведомления
	Read        bool             `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationTaskCommentDeleted NotificationTaskCommentDeleted) String() string {
	return Stringify(notificationTaskCommentDeleted)
}

// MetaType возвращает код сущности.
func (NotificationTaskCommentDeleted) MetaType() MetaType {
	return MetaTypeNotificationTaskCommentDeleted
}

type NotificationRetailShift struct {
	Meta    Meta      `json:"meta"`    // Метаданные смены
	Open    Timestamp `json:"open"`    // Дата открытия смены
	Name    string    `json:"name"`    // Номер смены
	Proceed float64   `json:"proceed"` // Выручка
	ID      uuid.UUID `json:"id"`      // ID смены
}

// NotificationRetailShiftOpened Смена открыта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-smena-otkryta
type NotificationRetailShiftOpened struct {
	Meta        Meta                    `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp               `json:"created"`     // Дата и время формирования Уведомления
	Description string                  `json:"description"` // Описание уведомления
	Title       string                  `json:"title"`       // Краткий текст уведомления
	RetailStore MetaNameID              `json:"retailStore"` // Точка продаж
	User        MetaNameID              `json:"user"`        // Сотрудник
	RetailShift NotificationRetailShift `json:"retailShift"` // Описание смены
	AccountID   uuid.UUID               `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID               `json:"id"`          // ID Уведомления
	Read        bool                    `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationRetailShiftOpened NotificationRetailShiftOpened) String() string {
	return Stringify(notificationRetailShiftOpened)
}

// MetaType возвращает код сущности.
func (NotificationRetailShiftOpened) MetaType() MetaType {
	return MetaTypeNotificationRetailShiftOpened
}

// NotificationRetailShiftClosed Смена закрыта.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-smena-zakryta
type NotificationRetailShiftClosed struct {
	Meta        Meta                    `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp               `json:"created"`     // Дата и время формирования Уведомления
	Description string                  `json:"description"` // Описание уведомления
	Title       string                  `json:"title"`       // Краткий текст уведомления
	RetailStore MetaNameID              `json:"retailStore"` // Точка продаж
	User        MetaNameID              `json:"user"`        // Сотрудник
	RetailShift NotificationRetailShift `json:"retailShift"` // Описание смены
	Returns     int                     `json:"returns"`     // Количество возвратов
	Sales       int                     `json:"sales"`       // Количество продаж
	AccountID   uuid.UUID               `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID               `json:"id"`          // ID Уведомления
	Read        bool                    `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationRetailShiftClosed NotificationRetailShiftClosed) String() string {
	return Stringify(notificationRetailShiftClosed)
}

// MetaType возвращает код сущности.
func (NotificationRetailShiftClosed) MetaType() MetaType {
	return MetaTypeNotificationRetailShiftClosed
}

// NotificationScript Уведомление из сценария.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uwedomlenie-iz-scenariq
type NotificationScript struct {
	Meta        Meta       `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp  `json:"created"`     // Дата и время формирования Уведомления
	Description string     `json:"description"` // Описание уведомления
	EventType   EventType  `json:"eventType"`   // Тип события сценария
	Title       string     `json:"title"`       // Краткий текст уведомления
	Entity      MetaNameID `json:"entity"`      // Ссылка на объект сценария
	AccountID   uuid.UUID  `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID  `json:"id"`          // ID Уведомления
	Read        bool       `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationScript NotificationScript) String() string {
	return Stringify(notificationScript)
}

// MetaType возвращает код сущности.
func (NotificationScript) MetaType() MetaType {
	return MetaTypeNotificationScript
}

// EventType Тип события сценария.
//
// Возможные значения:
//   - EventTypeAdd             – создан
//   - EventTypeModify          – изменен
//   - EventTypeAddChangeStatus – изменен статус
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uwedomlenie-iz-scenariq-atributy-wlozhennyh-suschnostej-sobytie
type EventType string

const (
	EventTypeAdd             EventType = "ADD"           // создан
	EventTypeModify          EventType = "MODIFY"        // изменен
	EventTypeAddChangeStatus EventType = "CHANGE_STATUS" // изменен статус
)

// String реализует интерфейс [fmt.Stringer].
func (eventType EventType) String() string {
	return string(eventType)
}

// FacebookTokenExpirationNotification Предупреждение о скором окончании действия доступа к аккаунту Facebook.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-preduprezhdenie-o-skorom-okonchanii-dejstwiq-dostupa-k-akkauntu-facebook
type FacebookTokenExpirationNotification struct {
	Meta                 Meta      `json:"meta"`                 // Метаданные объекта. Содержит тип конкретного уведомления
	Created              Timestamp `json:"created"`              // Дата и время формирования Уведомления
	ConnectorName        string    `json:"connectorName"`        // Название коннектора "Instagram and Facebook"
	Description          string    `json:"description"`          // Описание уведомления
	Title                string    `json:"title"`                // Краткий текст уведомления
	DaysLeftToExpiration int       `json:"daysLeftToExpiration"` // Количество дней, оставшихся до окончания действия доступа к аккаунту Facebook
	AccountID            uuid.UUID `json:"accountId"`            // ID учетной записи
	ID                   uuid.UUID `json:"id"`                   // ID Уведомления
	Read                 bool      `json:"read"`                 // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (facebookTokenExpirationNotification FacebookTokenExpirationNotification) String() string {
	return Stringify(facebookTokenExpirationNotification)
}

// MetaType возвращает код сущности.
func (FacebookTokenExpirationNotification) MetaType() MetaType {
	return MetaTypeFacebookTokenExpirationNotification
}

// NotificationBonusMoney На счет зачислены бонусные деньги.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-na-schet-zachisleny-bonusnye-den-gi
type NotificationBonusMoney struct {
	Meta        Meta      `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp `json:"created"`     // Дата и время формирования Уведомления
	Description string    `json:"description"` // Описание уведомления
	Title       string    `json:"title"`       // Краткий текст уведомления
	AccountID   uuid.UUID `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID `json:"id"`          // ID Уведомления
	Read        bool      `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notificationBonusMoney NotificationBonusMoney) String() string {
	return Stringify(notificationBonusMoney)
}

// MetaType возвращает код сущности.
func (NotificationBonusMoney) MetaType() MetaType {
	return MetaTypeNotificationBonusMoney
}

// NewMentionInEvent Новое упоминание в ленте событий.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowoe-upominanie-w-lente-sobytij
type NewMentionInEvent struct {
	Meta        Meta          `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp     `json:"created"`     // Дата и время формирования Уведомления
	Description string        `json:"description"` // Описание уведомления
	Title       string        `json:"title"`       // Краткий текст уведомления
	Operation   TaskOperation `json:"operation"`   // Объект, в ленте которого было добавлено событие с упоминанием
	AccountID   uuid.UUID     `json:"accountId"`   // ID учетной записи
	ID          uuid.UUID     `json:"id"`          // ID Уведомления
	Read        bool          `json:"read"`        // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (newMentionInEvent NewMentionInEvent) String() string {
	return Stringify(newMentionInEvent)
}

// MetaType возвращает код сущности.
func (NewMentionInEvent) MetaType() MetaType {
	return MetaTypeNewMentionInEvent
}

// SubscriptionGroup Значения кода группы уведомлений.
//
// Возможные значения:
//   - SubscriptionGroupCustomerOrder – Заказы покупателей
//   - SubscriptionGroupDataExchange  – Обмен данными
//   - SubscriptionGroupInvoice       – Счета покупателей
//   - SubscriptionGroupRetail        – Розничная торговля
//   - SubscriptionGroupScripts       – Сценарии
//   - SubscriptionGroupStock         – Складские остатки
//   - SubscriptionGroupTask          – Задачи
//   - SubscriptionGroupMentions      – Упоминания сотрудников
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

// String реализует интерфейс [fmt.Stringer].
func (subscriptionGroup SubscriptionGroup) String() string {
	return string(subscriptionGroup)
}

// SubscriptionChannel каналы уведомлений.
//
// Возможные значения:
//   - SubscriptionChannelEmail – Email-уведомления
//   - SubscriptionChannelPush  – Push-уведомления
type SubscriptionChannel string

const (
	SubscriptionChannelEmail SubscriptionChannel = "email" // Email-уведомления
	SubscriptionChannelPush  SubscriptionChannel = "push"  // Push-уведомления
)

// String реализует интерфейс [fmt.Stringer].
func (subscriptionChannel SubscriptionChannel) String() string {
	return string(subscriptionChannel)
}

// SubscriptionElement группу уведомлений.
type SubscriptionElement struct {
	Channels []SubscriptionChannel `json:"channels"` // Массив каналов
	Enabled  bool                  `json:"enabled"`  // Признак "активна" для подписки на уведомления данной группы
}

// NotificationSubscription Настройки уведомлений.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-nastrojki-uwedomlenij-atributy-suschnosti
type NotificationSubscription struct {
	Groups map[SubscriptionGroup]SubscriptionElement // Подписка на уведомления по группам
}

// Notification Уведомление (общие поля).
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-lenta-uwedomlenij-obschie-atributy-uwedomlenij
type Notification struct {
	Meta        Meta      `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Created     Timestamp `json:"created"`     // Дата и время формирования Уведомления
	Description string    `json:"description"` // Описание уведомления
	Title       string    `json:"title"`       // Краткий текст уведомления
	raw         []byte
	AccountID   uuid.UUID `json:"accountId"` // ID учетной записи
	ID          uuid.UUID `json:"id"`        // ID Уведомления
	Read        bool      `json:"read"`      // Признак того, было ли Уведомление прочитано
}

// String реализует интерфейс [fmt.Stringer].
func (notification Notification) String() string {
	return Stringify(notification)
}

// MetaType возвращает код сущности.
func (notification Notification) MetaType() MetaType {
	return notification.Meta.GetType()
}

// Raw реализует интерфейс [RawMetaTyper].
func (notification Notification) Raw() []byte {
	return notification.raw
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
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

// AsFacebookTokenExpirationNotification пытается привести объект к типу [FacebookTokenExpirationNotification].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [FacebookTokenExpirationNotification] или nil в случае неудачи.
func (notification Notification) AsFacebookTokenExpirationNotification() *FacebookTokenExpirationNotification {
	return UnmarshalAsType[FacebookTokenExpirationNotification](notification)
}

// AsNotificationExportCompleted пытается привести объект к типу [NotificationExportCompleted].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationExportCompleted] или nil в случае неудачи.
func (notification Notification) AsNotificationExportCompleted() *NotificationExportCompleted {
	return UnmarshalAsType[NotificationExportCompleted](notification)
}

// AsNotificationGoodCountTooLow пытается привести объект к типу [NotificationGoodCountTooLow].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationGoodCountTooLow] или nil в случае неудачи.
func (notification Notification) AsNotificationGoodCountTooLow() *NotificationGoodCountTooLow {
	return UnmarshalAsType[NotificationGoodCountTooLow](notification)
}

// AsNotificationImportCompleted пытается привести объект к типу [NotificationImportCompleted].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationImportCompleted] или nil в случае неудачи.
func (notification Notification) AsNotificationImportCompleted() *NotificationImportCompleted {
	return UnmarshalAsType[NotificationImportCompleted](notification)
}

// AsNotificationInvoiceOutOverdue пытается привести объект к типу [NotificationInvoiceOutOverdue].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationInvoiceOutOverdue] или nil в случае неудачи.
func (notification Notification) AsNotificationInvoiceOutOverdue() *NotificationInvoiceOutOverdue {
	return UnmarshalAsType[NotificationInvoiceOutOverdue](notification)
}

// AsNotificationOrderNew пытается привести объект к типу [NotificationOrderNew].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationOrderNew] или nil в случае неудачи.
func (notification Notification) AsNotificationOrderNew() *NotificationOrderNew {
	return UnmarshalAsType[NotificationOrderNew](notification)
}

// AsNotificationOrderOverdue пытается привести объект к типу [NotificationOrderOverdue].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationOrderOverdue] или nil в случае неудачи.
func (notification Notification) AsNotificationOrderOverdue() *NotificationOrderOverdue {
	return UnmarshalAsType[NotificationOrderOverdue](notification)
}

// AsNotificationRetailShiftClosed пытается привести объект к типу [NotificationRetailShiftClosed].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationRetailShiftClosed] или nil в случае неудачи.
func (notification Notification) AsNotificationRetailShiftClosed() *NotificationRetailShiftClosed {
	return UnmarshalAsType[NotificationRetailShiftClosed](notification)
}

// AsNotificationRetailShiftOpened пытается привести объект к типу [NotificationRetailShiftOpened].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationRetailShiftOpened] или nil в случае неудачи.
func (notification Notification) AsNotificationRetailShiftOpened() *NotificationRetailShiftOpened {
	return UnmarshalAsType[NotificationRetailShiftOpened](notification)
}

// AsNotificationScript пытается привести объект к типу [NotificationScript].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationScript] или nil в случае неудачи.
func (notification Notification) AsNotificationScript() *NotificationScript {
	return UnmarshalAsType[NotificationScript](notification)
}

// AsNotificationSubscribeExpired пытается привести объект к типу [NotificationSubscribeExpired].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationSubscribeExpired] или nil в случае неудачи.
func (notification Notification) AsNotificationSubscribeExpired() *NotificationSubscribeExpired {
	return UnmarshalAsType[NotificationSubscribeExpired](notification)
}

// AsNotificationSubscribeTermsExpired пытается привести объект к типу [NotificationSubscribeTermsExpired].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationSubscribeTermsExpired] или nil в случае неудачи.
func (notification Notification) AsNotificationSubscribeTermsExpired() *NotificationSubscribeTermsExpired {
	return UnmarshalAsType[NotificationSubscribeTermsExpired](notification)
}

// AsNotificationTaskAssigned пытается привести объект к типу [NotificationTaskAssigned].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskAssigned] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskAssigned() *NotificationTaskAssigned {
	return UnmarshalAsType[NotificationTaskAssigned](notification)
}

// AsNotificationTaskChanged пытается привести объект к типу [NotificationTaskChanged].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskChanged] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskChanged() *NotificationTaskChanged {
	return UnmarshalAsType[NotificationTaskChanged](notification)
}

// AsNotificationTaskCommentChanged пытается привести объект к типу [NotificationTaskCommentChanged].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskCommentChanged] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskCommentChanged() *NotificationTaskCommentChanged {
	return UnmarshalAsType[NotificationTaskCommentChanged](notification)
}

// AsNotificationTaskCommentDeleted пытается привести объект к типу [NotificationTaskCommentDeleted].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskCommentDeleted] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskCommentDeleted() *NotificationTaskCommentDeleted {
	return UnmarshalAsType[NotificationTaskCommentDeleted](notification)
}

// AsNotificationTaskCompleted пытается привести объект к типу [NotificationTaskCompleted].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskCompleted] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskCompleted() *NotificationTaskCompleted {
	return UnmarshalAsType[NotificationTaskCompleted](notification)
}

// AsNotificationTaskDeleted пытается привести объект к типу [NotificationTaskCompleted].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskCompleted] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskDeleted() *NotificationTaskDeleted {
	return UnmarshalAsType[NotificationTaskDeleted](notification)
}

// AsNotificationTaskNewComment пытается привести объект к типу [NotificationTaskNewComment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskNewComment] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskNewComment() *NotificationTaskNewComment {
	return UnmarshalAsType[NotificationTaskNewComment](notification)
}

// AsNotificationTaskOverdue пытается привести объект к типу [NotificationTaskOverdue].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskOverdue] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskOverdue() *NotificationTaskOverdue {
	return UnmarshalAsType[NotificationTaskOverdue](notification)
}

// AsNotificationTaskReopened пытается привести объект к типу [NotificationTaskReopened].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskReopened] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskReopened() *NotificationTaskReopened {
	return UnmarshalAsType[NotificationTaskReopened](notification)
}

// AsNotificationTaskUnassigned пытается привести объект к типу [NotificationTaskUnassigned].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationTaskUnassigned] или nil в случае неудачи.
func (notification Notification) AsNotificationTaskUnassigned() *NotificationTaskUnassigned {
	return UnmarshalAsType[NotificationTaskUnassigned](notification)
}

// AsNotificationBonusMoney пытается привести объект к типу [NotificationBonusMoney].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NotificationBonusMoney] или nil в случае неудачи.
func (notification Notification) AsNotificationBonusMoney() *NotificationBonusMoney {
	return UnmarshalAsType[NotificationBonusMoney](notification)
}

// AsNewMentionInEvent пытается привести объект к типу [NewMentionInEvent].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [NewMentionInEvent] или nil в случае неудачи.
func (notification Notification) AsNewMentionInEvent() *NewMentionInEvent {
	return UnmarshalAsType[NewMentionInEvent](notification)
}

// NotificationService описывает методы сервиса для работы с уведомлениями.
type NotificationService interface {
	// GetList выполняет запрос на получение ленты уведомлений.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Notification], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех уведомлений в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Notification], *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного уведомления по ID.
	// Принимает контекст, ID уведомления и опционально объект параметров запроса Params.
	// Возвращает найденное уведомление.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Notification, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление уведомления.
	// Принимает контекст и ID уведомления.
	// Возвращает «true» в случае успешного удаления уведомления.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// MarkAsRead выполняет запрос на отметку о прочтении Уведомления.
	// Принимает контекст и ID уведомления.
	// Возвращает «true» в случае успешного запроса.
	MarkAsRead(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// MarkAsReadAll выполняет запрос на отметку о прочтении всех уведомлений.
	// Принимает контекст.
	// Возвращает «true» в случае успешного запроса.
	MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error)

	// GetSubscription выполняет запрос настроек Уведомлений текущего пользователя.
	// Принимает контекст.
	// Возвращает Настройки уведомлений.
	GetSubscription(ctx context.Context) (*NotificationSubscription, *resty.Response, error)

	// UpdateSubscription выполняет запрос на изменение настроек Уведомлений текущего пользователя.
	// Принимает контекст и настройки уведомлений.
	// Возвращает «true» в случае успешного запроса.
	UpdateSubscription(ctx context.Context, notificationSubscription *NotificationSubscription) (bool, *resty.Response, error)
}

const (
	EndpointNotification              = string(MetaTypeNotification)
	EndpointNotificationSubscription  = EndpointNotification + "/subscription"
	EndpointNotificationMarkAsRead    = EndpointNotification + "/%s/markasread"
	EndpointNotificationMarkAsReadAll = EndpointNotification + "/markasreadall"
)

type notificationService struct {
	Endpoint
	endpointGetList[Notification]
	endpointGetByID[Notification]
	endpointDeleteByID
}

func (service *notificationService) MarkAsRead(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointNotificationMarkAsRead, id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}

func (service *notificationService) MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error) {
	_, resp, err := NewRequestBuilder[any](service.client, EndpointNotificationMarkAsReadAll).Put(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}

func (service *notificationService) GetSubscription(ctx context.Context) (*NotificationSubscription, *resty.Response, error) {
	return NewRequestBuilder[NotificationSubscription](service.client, EndpointNotificationSubscription).Get(ctx)
}

func (service *notificationService) UpdateSubscription(ctx context.Context, notificationSubscription *NotificationSubscription) (bool, *resty.Response, error) {
	_, resp, err := NewRequestBuilder[any](service.client, EndpointNotificationSubscription).Put(ctx, notificationSubscription)
	return resp.StatusCode() == http.StatusOK, resp, err
}

// NewNotificationService принимает [Client] и возвращает сервис для работы с уведомлениями.
func NewNotificationService(client *Client) NotificationService {
	e := NewEndpoint(client, EndpointNotification)
	return &notificationService{
		Endpoint:           e,
		endpointGetList:    endpointGetList[Notification]{e},
		endpointGetByID:    endpointGetByID[Notification]{e},
		endpointDeleteByID: endpointDeleteByID{e},
	}
}
