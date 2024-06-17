package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// NotificationFieldValue Формат измененного поля.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-tipy-uwedomlenij-formaty-polej
type NotificationFieldValue struct {
	OldValue string `json:"oldValue"` // Значение атрибута до удаления
	NewValue string `json:"newValue"` // Значение атрибута после обновления
}

// NotificationTaskState Статус завершения.
type NotificationTaskState string

const (
	NotificationTaskStateCompleted_            NotificationTaskState = "completed"
	NotificationTaskStateInterrupted_          NotificationTaskState = "interrupted"
	NotificationTaskStateInterruptedByUser_    NotificationTaskState = "interrupted_by_user"
	NotificationTaskStateInterruptedByTimeout_ NotificationTaskState = "interrupted_by_timeout"
	NotificationTaskStateInterruptedBySystem_  NotificationTaskState = "interrupted_by_system"
)

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

// NotificationExportCompleted Завершение экспорта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zawershenie-axporta
type NotificationExportCompleted struct {
	Neta                Meta                  `json:"neta"`
	Created             Timestamp             `json:"created"`
	CreatedDocumentName string                `json:"createdDocumentName"`
	Description         string                `json:"description"`
	ErrorMessage        string                `json:"errorMessage"`
	Message             string                `json:"message"`
	TaskState           NotificationTaskState `json:"taskState"`
	TaskType            NotificationTaskType  `json:"taskType"`
	Title               string                `json:"title"`
	AccountId           uuid.UUID             `json:"accountId"`
	ID                  uuid.UUID             `json:"id"`
	Read                bool                  `json:"read"`
}

// NotificationImportCompleted Завершение импорта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zawershenie-importa
type NotificationImportCompleted struct {
	// accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//createdDocumentName	String(255)	Имя экспортированного документа
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//errorMessage	String(255)	Сообщение об ошибке
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//message	String(255)	Сообщение о завершении экспорта
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//taskState	Object	Статус завершения. Может принимать значения completed, interrupted, interrupted_by_user, interrupted_by_timeout, interrupted_by_system
	//Обязательное при ответе Необходимо при создании
	//taskType	Object	Тип экспорта
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationGoodCountTooLow Снижение остатка товара ниже неснижаемого.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-snizhenie-ostatka-towara-nizhe-nesnizhaemogo
type NotificationGoodCountTooLow struct {
	// accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//actualBalance	Int	Остаток товара
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//good	Meta	Метаданные товара
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//minimumBalance	Int	Неснижаемый остаток товара
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationInvoiceOutOverdue Просрочен счет покупателя.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-prosrochen-schet-pokupatelq
type NotificationInvoiceOutOverdue struct {
	// accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//agentName	String(255)	Имя контрагента
	//Обязательное при ответе
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//invoice	Meta	Метаданные счета
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//paymentPlannedMoment	DateTime	Запланированная дата оплаты
	//Обязательное при ответе
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//sum	Int	Сумма счета
	//Обязательное при ответе
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationOrderNew Новый заказ.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowyj-zakaz
type NotificationOrderNew struct {
	// accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//*agentName	String(255)	Имя контрагента
	//Обязательное при ответе
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//deliveryPlannedMoment	DateTime	Планируемое время отгрузки
	//Обязательное при ответе
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//order	Meta	Метаданные заказа
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//sum	Int	Сумма счета
	//Обязательное при ответе
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationOrderOverdue Просроченный заказ.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-prosrochennyj-zakaz
type NotificationOrderOverdue struct {
	// accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//*agentName	String(255)	Имя контрагента
	//Обязательное при ответе
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//deliveryPlannedMoment	DateTime	Планируемое время отгрузки
	//Обязательное при ответе
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//order	Meta	Метаданные заказа
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//sum	Int	Сумма счета
	//Обязательное при ответе
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationSubscribeExpired Окончание подписки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-okonchanie-podpiski
type NotificationSubscribeExpired struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationSubscribeTermsExpired Условия подписки истекают.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uslowiq-podpiski-istekaut
type NotificationSubscribeTermsExpired struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//daysLeft	Int	Количество оставшихся дней подписки
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskAssigned Задача назначена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-naznachena
type NotificationTaskAssigned struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// PerformedBy Сотрудник, выполнивший изменение.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-naznachena-atributy-wlozhennyh-suschnostej-sotrudnik-wypolniwshij-izmenenie
type PerformedBy struct {
	//id	UUID	ID Уведомления
	//Обязательное при ответе
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе
	//name	String(255)	Наименование Контрагента
	//Обязательное при ответе
}

// NotificationTask Задача.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-naznachena-atributy-wlozhennyh-suschnostej-zadacha
type NotificationTask struct {
	// deadline	DateTime	Планируемая дата завершения задачи
	//Обязательное при ответе
	//id	UUID	ID Уведомления
	//Обязательное при ответе
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе
	//name	String(255)	Наименование Контрагента
	//Обязательное при ответе
}

// NotificationTaskChanged Задача изменена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-izmenena
type NotificationTaskChanged struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//diff	Object	Измененные поля Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskChangedDiff Измененные поля.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-izmenena-atributy-wlozhennyh-suschnostej-izmenennye-polq
type NotificationTaskChangedDiff struct {
	//description	String(255)	Изменение описания задачи в формате изменения поля
	//Обязательное при ответе
	//deadline	String(255)	Изменение планируемой даты завершения задачи в формате изменения поля
	//Обязательное при ответе
	//agentLink	String(255)	Изменение контрагента в формате изменения поля
	//Обязательное при ответе
	//documentLink	String(255)	Изменение связанного документа в формате изменения поля
	//Обязательное при ответе
	//assignee	String(255)	Изменение исполнителя в формате изменения поля
	//Обязательное при ответе
}

// NotificationTaskCompleted Задача выполнена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-wypolnena
type NotificationTaskCompleted struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskDeleted Задача удалена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-udalena
type NotificationTaskDeleted struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskOverdue Задача просрочена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-prosrochena
type NotificationTaskOverdue struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskReopened Задача переоткрыта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-zadacha-pereotkryta
type NotificationTaskReopened struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskNewComment Новый комментарий к задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowyj-kommentarij-k-zadache
type NotificationTaskNewComment struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//noteContent	String(4096)	Содержимое комментария
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskCommentChanged Изменен комментарий к задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-izmenen-kommentarij-k-zadache
type NotificationTaskCommentChanged struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//diff	String(255)	Изменения комментария в формате изменения поля
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//noteContent	String(4096)	Содержимое комментария
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationTaskCommentDeleted Удален комментарий к задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-udalen-kommentarij-k-zadache
type NotificationTaskCommentDeleted struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//noteContent	String(4096)	Содержимое комментария
	//Обязательное при ответе Необходимо при создании
	//performedBy	Object	Сотрудник, выполнивший изменение. Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//task	Object	Задача Подробнее тут
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationRetailShiftOpened Смена открыта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-smena-otkryta
type NotificationRetailShiftOpened struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//retailShift	Object	Описание смены
	//Обязательное при ответе Необходимо при создании
	//retailStore	Object	Точка продаж
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
	//user	Object	Сотрудник
	//Обязательное при ответе Необходимо при создании
}

// NotificationRetailShiftClosed Смена закрыта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-smena-zakryta
type NotificationRetailShiftClosed struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//retailShift	Object	Описание смены
	//Обязательное при ответе Необходимо при создании
	//retailStore	Object	Точка продаж
	//Обязательное при ответе Необходимо при создании
	//returns	Int	Количество возвратов
	//Обязательное при ответе Необходимо при создании
	//sales	Int	Количество продаж
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
	//user	Object	Сотрудник
	//Обязательное при ответе Необходимо при создании
}

// NotificationScript Уведомление из сценария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uwedomlenie-iz-scenariq
type NotificationScript struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Только для чтения
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Только для чтения
	//description	String(255)	Описание уведомления
	//Обязательное при ответе Только для чтения
	//entity	Object	Ссылка на объект сценария
	//Обязательное при ответе Только для чтения
	//eventType	Событие	Тип события сценария
	//Обязательное при ответе Только для чтения
	//id	UUID	ID Уведомления
	//Обязательное при ответе Только для чтения
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Только для чтения
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Только для чтения
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Только для чтения
}

// EventType Тип события сценария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uwedomlenie-iz-scenariq-atributy-wlozhennyh-suschnostej-sobytie
type EventType string

const (
	EventTypeAdd             EventType = "ADD"           // создан
	EventTypeModify          EventType = "MODIFY"        // изменен
	EventTypeAddChangeStatus EventType = "CHANGE_STATUS" // изменен статус
)

// NotificationScriptEntity Событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-uwedomlenie-iz-scenariq-atributy-wlozhennyh-suschnostej-ob-ekt
type NotificationScriptEntity struct {
	//id	UUID	ID объекта
	//Обязательное при ответе Только для чтения
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Только для чтения
	//name	String(255)	Наименование объекта
	//Обязательное при ответе Только для чтения
}

// FacebookTokenExpirationNotification Предупреждение о скором окончании действия доступа к аккаунту Facebook.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-preduprezhdenie-o-skorom-okonchanii-dejstwiq-dostupa-k-akkauntu-facebook
type FacebookTokenExpirationNotification struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//connectorName	String(4096)	Название коннектора "Instagram and Facebook"
	//Обязательное при ответе
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//daysLeftToExpiration	Int	Количество дней, оставшихся до окончания действия доступа к аккаунту Facebook
	//Обязательное при ответе
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NotificationBonusMoney На счет зачислены бонусные деньги.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-na-schet-zachisleny-bonusnye-den-gi
type NotificationBonusMoney struct {
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Необходимо при создании
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Необходимо при создании
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Необходимо при создании
	//id	UUID	ID Уведомления
	//Обязательное при ответе Необходимо при создании
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Необходимо при создании
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Необходимо при создании
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Необходимо при создании
}

// NewMentionInEvent Новое упоминание в ленте событий.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowoe-upominanie-w-lente-sobytij
type NewMentionInEvent struct {
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Только для чтения
	//id	UUID	ID Уведомления
	//Обязательное при ответе Только для чтения
	//accountId	UUID	ID учетной записи
	//Обязательное при ответе Только для чтения
	//created	DateTime	Дата и время формирования Уведомления
	//Обязательное при ответе Только для чтения
	//read	Boolean	Признак того, было ли Уведомление прочитано
	//Обязательное при ответе Только для чтения
	//title	String(255)	Краткий текст уведомления
	//Обязательное при ответе Только для чтения
	//description	String(4096)	Описание уведомления
	//Обязательное при ответе Только для чтения
	//operation	Object	Объект, в ленте которого было добавлено событие с упоминанием
	//Обязательное при ответе Только для чтения
}

// NewMentionInEventOperation Объект, в ленте которого добавили событие с упоминанием.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-podrobnoe-opisanie-tipow-uwedomlenij-nowoe-upominanie-w-lente-sobytij-atributy-wlozhennyh-suschnostej-ob-ekt
type NewMentionInEventOperation struct {
	//meta	Meta	Метаданные объекта
	//Обязательное при ответе Только для чтения
	//id	UUID	ID объекта
	//Обязательное при ответе Только для чтения
	//name	String(255)	Наименование объекта
	//Обязательное при ответе Только для чтения
}

// NotificationSubscription Настройки уведомлений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-nastrojki-uwedomlenij-atributy-suschnosti
type NotificationSubscription struct {
}

// Notification TODO: Общие атрибуты уведомлений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-lenta-uwedomlenij-obschie-atributy-uwedomlenij
type Notification struct {
	Meta        Meta      `json:"meta"`
	Created     Timestamp `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	AccountID   uuid.UUID `json:"accountId"`
	ID          uuid.UUID `json:"id"`
	Read        bool      `json:"read"`
}

func (notification Notification) String() string {
	return Stringify(notification)
}

func (notification Notification) MetaType() MetaType {
	return MetaTypeNotification
}

// NotificationService
// Сервис для работы с уведомлениями.
type NotificationService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Notification], *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Notification, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	MarkAsRead(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error)
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
